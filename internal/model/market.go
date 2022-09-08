package model

import (
	"gopkg.in/validator.v2"
)

type StreetMarketDtos []*StreetMarketDto

type StreetMarkets []*StreetMarket
type StreetMarketDto struct {
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

type StreetMarket struct {
	Id         int64
	Long       string
	Lat        string
	Setcens    string
	Areap      string
	Coddist    int32
	Distrito   string
	Codsubpref int
	Subprefe   string
	Regiao5    string
	Regiao8    string
	NomeFeira  string
	Registro   string
	Logradouro string
	Numero     string
	Bairro     string
	Referencia string
}

func (market *StreetMarketDto) Validate() error {
	if err := validator.Validate(market); err != nil {
		return err
	}
	return nil
}

func (StreetMarket) TableName() string {
	return "markets"
}

func (s StreetMarket) ToDto() *StreetMarketDto {
	return &StreetMarketDto{
		Longitude:        s.Long,
		Latitude:         s.Lat,
		SetorCensitario:  s.Setcens,
		AreaPonderacao:   s.Areap,
		CodigoDistrito:   s.Coddist,
		Distrito:         s.Distrito,
		CodSubPrefeitura: s.Codsubpref,
		SubPrefeitura:    s.Subprefe,
		Regiao5:          s.Regiao5,
		Regiao8:          s.Regiao8,
		NomeFeira:        s.NomeFeira,
		Registro:         s.Registro,
		Logradouro:       s.Logradouro,
		Numero:           s.Numero,
		Bairro:           s.Bairro,
		Referencia:       s.Referencia,
	}
}

func (s StreetMarketDto) ToModel() *StreetMarket {
	return &StreetMarket{
		Long:       s.Longitude,
		Lat:        s.Latitude,
		Setcens:    s.SetorCensitario,
		Areap:      s.AreaPonderacao,
		Coddist:    s.CodigoDistrito,
		Distrito:   s.Distrito,
		Codsubpref: s.CodSubPrefeitura,
		Subprefe:   s.SubPrefeitura,
		Regiao5:    s.Regiao5,
		Regiao8:    s.Regiao8,
		NomeFeira:  s.NomeFeira,
		Logradouro: s.Logradouro,
		Numero:     s.Numero,
		Bairro:     s.Bairro,
		Referencia: s.Referencia,
	}
}

func (ss StreetMarkets) ToDto() StreetMarketDtos {
	dtos := make([]*StreetMarketDto, len(ss))
	for i, s := range ss {
		dtos[i] = s.ToDto()
	}

	return dtos
}
