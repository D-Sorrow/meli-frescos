package mappers

import (
	"time"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/entities"
)

func LogAttributesToLogEntity(
	v *models.LogAttributes,
) *entities.LogEntity {
	return &entities.LogEntity{
		Level:  string(v.Level),
		Source: v.Source,
		Detail: v.Detail,
	}
}

func LogEntityToLog(
	v *entities.LogEntity,
) *models.Log {
	dt, err := time.Parse("2006-01-02 15:04:05", v.Timestamp)
	if err != nil {
		dt = time.Time{}
	}

	return &models.Log{
		ID:        v.ID,
		Timestamp: dt,
		LogAttributes: models.LogAttributes{
			Level:  models.LogLevel(v.Level),
			Source: v.Source,
			Detail: v.Detail,
		},
	}
}
