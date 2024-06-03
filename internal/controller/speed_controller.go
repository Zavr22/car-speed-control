package controller

import (
	"github.com/Zavr22/car-speed-control/config"
	models "github.com/Zavr22/car-speed-control/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SpeedService interface {
	GetSpeedStats(date time.Time) (minRecord, maxRecord models.SpeedRecord, err error)
	GetRecordsExceedingSpeed(date time.Time, speed float64) ([]*models.SpeedRecord, error)
	AddRecord(record models.SpeedRecord) error
}

type SpeedController struct {
	service SpeedService
	config  *config.Config
}

func NewSpeedController(service SpeedService, config *config.Config) *SpeedController {
	return &SpeedController{service: service, config: config}
}

func (s *SpeedController) AddSpeedRecord(c *gin.Context) {
	var record models.SpeedRecord

	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.service.AddRecord(record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "record added"})
}

func (s *SpeedController) GetSpeedRecords(c *gin.Context) {
	dateStr := c.Query("date")
	speedStr := c.Query("speed")

	date, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	speed, err := strconv.ParseFloat(speedStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid speed format"})
		return
	}

	records, err := s.service.GetRecordsExceedingSpeed(date, speed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

func (s *SpeedController) GetSpeedStats(c *gin.Context) {
	dateStr := c.Query("date")

	date, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	minRecord, maxRecord, err := s.service.GetSpeedStats(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"min_speed": minRecord,
		"max_speed": maxRecord,
	})
}
