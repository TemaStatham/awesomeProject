package models

type RemoveResponse struct {
	ID         int    `json:"id"`
	CompaignID string `json:"compaignId"`
	Removed    bool   `json:"removed"`
}
