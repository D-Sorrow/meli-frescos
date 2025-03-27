package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/mappers"
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
