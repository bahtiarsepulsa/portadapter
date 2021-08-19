package controller

type CreateRequestOperator struct {
	Code  string `json:"code" validate:"required"`
	Label string `json:"label" validate:"required"`
}

type UpdateRequestOperator struct {
	Code  string `json:"code" validate:"required"`
	Label string `json:"label" validate:"required"`
}

type ReadDataResponseOperator struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Label string `json:"label"`
}

type ListDataResponseOperator struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Label string `json:"label"`
}
