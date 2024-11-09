package responses

type TirResponse struct {
	Key        string  `json:"Key"`
	Value      float64 `json:"Value"`
	Price      float64 `json:"Price"`
	FinishDate string  `json:"FinishDate"`
	Company    string  `json:"Company"`
}
