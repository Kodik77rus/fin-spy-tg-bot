package models

type Market struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	Mic        string `json:"mic"`
	YoohooCode string `json:"yoohooCode"`
	Location   string `json:"location"`
	Country    string `json:"country"`
	Delay      uint   `json:"delay"`
	Hour       string `json:"hour"`
}
