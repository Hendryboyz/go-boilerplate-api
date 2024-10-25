package dto

import "time"

type TodoCreatedRequest struct {
	Description string    `json:"description" format:"string" example:"say hello to everyone"`
	StartDate   time.Time `json:"startDate" format:"string"`
	EndDate     time.Time `json:"endDate" format:"string"`
}
