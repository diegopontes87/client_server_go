package entities

import (
	"encoding/json"
	"fmt"
)

type ServerCotation struct {
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

func (c ServerCotation) UnmarshalCotation(data []byte, cotation *ServerCotation) error {
	err := json.Unmarshal(data, &cotation)
	if err != nil {
		fmt.Printf("Error trying to Unmarshal json: %v\n", err)
		return err
	}
	return nil
}

func (cotation ServerCotation) ConvertToDBCotation() *DBCotation {
	return &DBCotation{
		Code:      cotation.USDBRL.Code,
		Codein:    cotation.USDBRL.Codein,
		Name:      cotation.USDBRL.Name,
		High:      cotation.USDBRL.High,
		Low:       cotation.USDBRL.Low,
		VarBid:    cotation.USDBRL.VarBid,
		PctChange: cotation.USDBRL.PctChange,
		Bid:       cotation.USDBRL.Bid,
		Ask:       cotation.USDBRL.Ask,
		Timestamp: cotation.USDBRL.Timestamp,
	}
}
