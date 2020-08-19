package controllers

import (
	"fmt"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{"data": record})
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

	c.JSON(http.StatusOK, gin.H{"data": record})
}

//CreateRecordForm renders create record form
func CreateRecordForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create", gin.H{})
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
	c.HTML(http.StatusOK, "edit", gin.H{"data": record})
}

//UpdateRecord updates record at the given ID
func UpdateRecord(c *gin.Context) {
	var record models.Record
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var updatedRecord models.Record
	c.ShouldBindJSON(&updatedRecord)
	models.DB.Model(&record).Updates(updatedRecord)
}

//DeleteRecord deletes record at the given ID
func DeleteRecord(c *gin.Context) {
	var record models.Record
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&record)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
