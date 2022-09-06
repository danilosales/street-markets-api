package model

import (
	"gopkg.in/validator.v2"
)

type StreetMarket struct {
	Id               int64  `json:"id"`
	Longitude        string `json:"long" validate:"nonzero, max=10"`
	Latitude         string `json:"lat" validate:"nonzero, max=10"`
	SetorCensitario  string `json:"setcens" validate:"nonzero, max=15"`
	AreaPonderacao   string `json:"areap" validate:"nonzero, max=13"`
	CodigoDistrito   int32  `json:"coddist" validate:"nonzero"`
	Distrito         string `json:"distrito" validate:"nonzero, max=18"`
	CodSubPrefeitura int    `json:"codsubpref" validate:"nonzero"`
	SubPrefeitura    string `json:"subprefe" validate:"nonzero, max=25"`
	Regiao5          string `json:"regiao5" validate:"nonzero, max=6"`
	Regiao8          string `json:"regiao8" validate:"nonzero, max=7"`
	NomeFeira        string `json:"nome_feira" validate:"nonzero, max=30"`
	Registro         string `json:"registro" validate:"nonzero, max=6"`
	Logradouro       string `json:"logradouro" validate:"nonzero, max=34"`
	Numero           string `json:"numero" validate:"max=5"`
	Bairro           string `json:"bairro" validate:"max=20"`
	Referencia       string `json:"referencia" validate:"max=60"`
}

func Validate(market *StreetMarket) error {
	if err := validator.Validate(market); err != nil {
		return err
	}
	return nil
}
