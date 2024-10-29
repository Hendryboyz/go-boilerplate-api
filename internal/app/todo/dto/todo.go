package dto

import (
	"go-boilerplate-api/internal/pkg/utils"
	"time"
)

type CreateTodoRequest struct {
	Description string          `json:"description" format:"string" example:"say hello to everyone"`
	StartDate   *utils.DateTime `json:"startDate" format:"dateTime" time_format:"2024-01-01 00:00:00"`
	EndDate     *utils.DateTime `json:"endDate" format:"dateTime" time_format:"2024-01-01 00:00:00"`
}

type UpdateTodoRequest struct {
	Description string    `json:"description" format:"string" example:"say hello to everyone" binding:"-"`
	StartDate   time.Time `json:"startDate" format:"dateTime" binding:"-"`
	EndDate     time.Time `json:"endDate" format:"dateTime" binding:"-"`
}
