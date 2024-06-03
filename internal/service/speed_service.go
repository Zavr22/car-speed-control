package service

import (
	models "github.com/Zavr22/car-speed-control/model"
	"time"
)

type SpeedRepo interface {
	GetByDate(date time.Time) ([]*models.SpeedRecord, error)
	Save(record models.SpeedRecord) error
}

type SpeedService struct {
	repo SpeedRepo
}

func NewSpeedService(repo SpeedRepo) *SpeedService {
	return &SpeedService{repo: repo}
}

func (s *SpeedService) AddRecord(record models.SpeedRecord) error {
	return s.repo.Save(record)
}

func (s *SpeedService) GetRecordsExceedingSpeed(date time.Time, threshold float64) ([]*models.SpeedRecord, error) {
	records, err := s.repo.GetByDate(date)
	if err != nil {
		return nil, err
	}

	var filteredRecords []*models.SpeedRecord
	for _, record := range records {
		if record.Speed > threshold {
			filteredRecords = append(filteredRecords, record)
		}
	}

	return filteredRecords, nil
}

func (s *SpeedService) GetSpeedStats(date time.Time) (minRecord, maxRecord models.SpeedRecord, err error) {
	records, err := s.repo.GetByDate(date)
	if err != nil {
		return models.SpeedRecord{}, models.SpeedRecord{}, err
	}

	if len(records) == 0 {
		return models.SpeedRecord{}, models.SpeedRecord{}, nil
	}

	minRecord = *records[0]
	maxRecord = *records[0]

	for _, record := range records {
		if record.Speed < minRecord.Speed {
			minRecord = *record
		}
		if record.Speed > maxRecord.Speed {
			maxRecord = *record
		}
	}

	return minRecord, maxRecord, nil
}
