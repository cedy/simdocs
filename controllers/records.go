package controllers

import (
	"net/http"

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

// GetAllRecords returns all records
func GetAllRecords(c *gin.Context) {
	var records []models.Record
	models.DB.Find(&records)

	c.JSON(http.StatusOK, gin.H{"data": records})
}

// CreateRecord creates new record
func CreateRecord(c *gin.Context) {
	var record models.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: handle files after POC
	models.DB.Create(&record)
	c.JSON(http.StatusOK, gin.H{"data": record})
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
