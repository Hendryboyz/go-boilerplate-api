package dto

import "time"

type CreateTodoRequest struct {
	Description string    `json:"description" format:"string" example:"say hello to everyone"`
	StartDate   time.Time `json:"startDate" format:"string"`
	EndDate     time.Time `json:"endDate" format:"string"`
}

type UpdateTodoRequest struct {
	Description string    `json:"description" format:"string" example:"say hello to everyone" binding:"-"`
	StartDate   time.Time `json:"startDate" format:"string" binding:"-"`
	EndDate     time.Time `json:"endDate" format:"string" binding:"-"`
}
