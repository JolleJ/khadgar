package balance

type GetBalanceResponse struct {
	Balance Balance `json:"Balance"`
}

type Balance struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
