package repository

import (
	"encoding/csv"
	"fmt"
	models "github.com/Zavr22/car-speed-control/model"
	"os"
	"strconv"
	"time"
)

type SpeedRepository struct {
	filePath string
}

func NewSpeedRepository(filePath string) *SpeedRepository {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write([]string{"Timestamp", "VehicleID", "Speed"})
	}

	return &SpeedRepository{filePath: filePath}
}

func (r *SpeedRepository) Save(record models.SpeedRecord) error {
	file, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write([]string{
		record.Timestamp.Format(time.RFC3339),
		record.VehicleID,
		fmt.Sprintf("%.2f", record.Speed),
	})
}

func (r *SpeedRepository) GetByDate(date time.Time) ([]*models.SpeedRecord, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var results []*models.SpeedRecord

	for _, record := range records[1:] {
		timestamp, _ := time.Parse(time.RFC3339, record[0])
		speed, _ := strconv.ParseFloat(record[2], 64)

		if timestamp.Format("2006.01.02") == date.Format("2006.01.02") {

			results = append(results, &models.SpeedRecord{Timestamp: timestamp, VehicleID: record[1], Speed: speed})

		}
	}
	return results, nil
}
