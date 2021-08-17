package repository

import "time"

type (
	SaveOperator struct {
		ID        int       `json:"id"`
		Code      string    `json:"code"`
		Label     string    `json:"label"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Operator struct {
		ID        int       `json:"id"`
		Code      string    `json:"code"`
		Label     string    `json:"label"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}
)
