package operator

type CreateRequestOperator struct {
	Code  string `json:"code" validate:"required"`
	Label string `json:"label" validate:"required"`
}

type UpdateRequestOperator struct {
	Code  string `json:"code" validate:"required"`
	Label string `json:"label" validate:"required"`
}
