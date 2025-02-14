package requests

type CreateProduct struct {
		CategoryID          string    `json:"category_id"`
		BranchID            string    `json:"branch_id"`
		Name                string         `json:"name"`
		ProductType         string         `json:"product_type"`
		ServicePricingModel string    `json:"service_pricing_model"`
		DefaultUnit         string    `json:"default_unit"`
		IsBillable          bool    `json:"is_billable"`
		Sku                 string         `json:"sku"`
		Description         string    `json:"description"`
		BasePrice           float64 `json:"base_price"`
}