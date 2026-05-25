package models

type BFFGenerateBillRequest struct {
	VisitRecordID uint64 `json:"visit_id" validate:"required"`
}

type BFFGenerateBillResponse struct {
	BillAmount float64 `json:"bill_amount"`
	Status     string  `json:"status" example:"Bill Generated Sucessfully"`
}
