package order

type Order struct {
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name" validate:"required,min=2"`
	Status       string `json:"status" validate:"lowercase,oneof=created shipped delivered"`
}

type UpdateReq struct {
	Status string `json:"status" validate:"required,lowercase,oneof=created shipped delivered"`
}
