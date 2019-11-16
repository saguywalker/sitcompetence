package model

// TxResponse contains txID and its log
type TxResponse struct {
	TxID string                 `json:"transaction_id"`
	Log  map[string]interface{} `json:"evidence"`
}
