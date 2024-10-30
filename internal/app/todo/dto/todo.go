package dto

import (
	"go-boilerplate-api/internal/model"
	"time"
)

type CreateTodoRequest struct {
	Description string          `json:"description" format:"string" binding:"required" example:"say hello to everyone"`
	StartDate   *model.DateTime `json:"startDate" format:"dateTime" binding:"required" time_format:"2024-01-01 00:00:00"`
	EndDate     *model.DateTime `json:"endDate" format:"dateTime" binding:"required" time_format:"2024-01-01 00:00:00"`
}

type UpdateTodoRequest struct {
	Description string    `json:"description" format:"string" example:"say hello to everyone" binding:"omitempty"`
	StartDate   time.Time `json:"startDate" format:"dateTime" binding:"omitempty"`
	EndDate     time.Time `json:"endDate" format:"dateTime" binding:"omitempty"`
}
