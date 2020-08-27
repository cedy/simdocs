package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cedy/simdocs/models"
	"github.com/gin-gonic/gin"
)

//GetRecord returns idividual record fetched by ID
func GetRecord(c *gin.Context) {
	var record models.Record
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var files []models.File
	err := models.DB.Where("record_id = ?", record.ID).Find(&files).Error
	if err != nil {
		files = append(files, models.File{Name: "error getting files, please try reloading page."})
	}
	c.HTML(http.StatusOK, "record", gin.H{"data": record, "files": files, "title": fmt.Sprintf("ID: %d %s %s", record.ID, record.Lastname, record.Firstname)})
}

// GetAllRecords renders page with all records
func GetAllRecords(c *gin.Context) {
	var records []models.Record
	var files []models.File
	models.DB.Find(&records)
	//TODO: can we do better
	models.DB.Find(&files)
	recordIDToFile := make(map[uint][]models.File)
	for _, file := range files {
		recordIDToFile[file.RecordID] = append(recordIDToFile[file.RecordID], file)
	}
	c.HTML(http.StatusOK, "index", gin.H{"title": "All Records", "data": records, "files": recordIDToFile})
}

// CreateRecord creates new record
func CreateRecord(c *gin.Context) {
	var record models.Record
	c.ShouldBind(&record)

	if err := c.ShouldBind(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&record)
	// Write attached files to disk and save info about in in DB
	files := form.File["Files"]
	for _, file := range files {
		var f models.File
		localPath := fmt.Sprintf("docs/%d_%s", time.Now().Unix(), file.Filename)
		if err := c.SaveUploadedFile(file, localPath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		f.Name = file.Filename
		f.Path = localPath
		f.RecordID = record.ID
		models.DB.Create(&f)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "id": record.ID})
}

//CreateRecordForm renders create record form
func CreateRecordForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create", gin.H{"data": models.OrderTypes, "title": "Create Record"})
}

//GetRecordsSearch returns record(s) filtered by querying params*
func GetRecordsSearch(c *gin.Context) {
	var records []models.Record
	lastname := c.Query("lastname") + "%"
	orderType := c.Query("orderType") + "%"
	address := c.Query("address") + "%"
	phone := c.Query("phone") + "%"
	orderTime := c.Query("orderTime") + "%"
	if err := models.DB.Where("lastname like ? and order_type like ? and address like ? and phone like ? and order_time like ?", lastname, orderType, address, phone, orderTime).Find(&records).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": records})
}

//EditRecordForm renders edit record form
func EditRecordForm(c *gin.Context) {
	var record models.Record
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var files []models.File
	err := models.DB.Where("record_id = ?", record.ID).Find(&files).Error
	if err != nil {
		files = append(files, models.File{Name: "error getting files, please try reloading page."})
	}
	c.HTML(http.StatusOK, "edit", gin.H{"data": record, "files": files, "orderTypes": models.OrderTypes, "title": fmt.Sprintf("Edit record ID: %d %s %s", record.ID, record.Lastname, record.Firstname)})
}

//UpdateRecord updates record at the given ID
func UpdateRecord(c *gin.Context) {
	var record models.Record

	if err := models.DB.Where("id = ?", c.Request.FormValue("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var updatedRecord models.Record
	c.ShouldBind(&updatedRecord)

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if models.DB.Model(&record).Updates(updatedRecord).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	files := form.File["Files"]
	for _, file := range files {
		var f models.File
		localPath := fmt.Sprintf("docs/%d_%s", time.Now().Unix(), file.Filename)
		if err := c.SaveUploadedFile(file, localPath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		f.Name = file.Filename
		f.Path = localPath
		f.RecordID = record.ID
		models.DB.Create(&f)
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

//DeleteRecord deletes record at the given ID and all files attached to that record
func DeleteRecord(c *gin.Context) {
	var record models.Record
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var files []models.File
	models.DB.Where("record_id = ?", record.ID).Find(&files)
	if models.DB.Delete(&record).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error!"})
		return
	}
	for _, file := range files {
		models.DB.Delete(&file)
		os.Remove(file.Path)
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

//DeleteFile deletes file and DB file record for given FileID
func DeleteFile(c *gin.Context) {
	var file models.File
	if err := models.DB.Where("id = ?", c.Param("id")).First(&file).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File doesn't exist!"})
		return
	}
	// delete file record from DB and physical record from disk
	if models.DB.Delete(&file).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error!"})
		return
	}
	// TODO: rethink this behavior
	// ignore an error if file can't be deleted for any reason.
	_ = os.Remove(file.Path)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
