package service

import (
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/models"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service/mappers"
)

type LogService struct {
	repo repository.LogRepository
}

func NewLogService(repository repository.LogRepository) *LogService {
	return &LogService{repo: repository}
}

func (p *LogService) Register(
	log models.LogAttributes,
) (newLog models.Log, err error) {
	logEntity := mappers.LogAttributesToLogEntity(&log)

	newLogEntity, err := p.repo.Register(*logEntity)
	if err != nil {
		err = error_management.HandleServiceError(
			error_management.HandleLogServiceError(err),
			err,
		)
		return
	}

	newLog = *mappers.LogEntityToLog(&newLogEntity)

	return
}
