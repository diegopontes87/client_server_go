package entities

import (
	"encoding/json"
	"fmt"
)

type Cotation struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func (c Cotation) UnmarshalCotation(body []byte, cotation *Cotation) error {
	err := json.Unmarshal(body, &cotation)
	if err != nil {
		fmt.Printf("Error trying to Unmarshal json: %v\n", err)
		return err
	}
	return nil
}
