package requests

type CreateCustomerReq struct {
	Name string `json:"name"`
	Email string `json:"email"`
	PhobeNumber string `json:"phone_number"`
	BillingAddress map[string]interface{}
	BranchID string `json:"branch_id"`
}

type DeleteCustomerReq struct {
	ID string `json:"customer_id"`
}

