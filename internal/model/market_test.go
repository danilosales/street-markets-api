package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/validator.v2"
)

func TestValidStreetMarket(t *testing.T) {

	market := StreetMarket{
		Id: 1, Longitude: "-46550164", Latitude: "-23558733", SetorCensitario: "355030885000091",
		AreaPonderacao: "3550308005040", CodigoDistrito: 87, Distrito: "VILA FORMOSA",
		CodSubPrefeitura: 26, SubPrefeitura: "ARICANDUVA-FORMOSA-CARRAO", Regiao5: "Leste",
		Regiao8: "Leste 1", NomeFeira: "VILA FORMOSA", Registro: "4041-0",
		Logradouro: "RUA MARAGOJIPE", Numero: "S/N", Bairro: "VL FORMOSA", Referencia: "TV RUA PRETORIA",
	}

	err := Validate(&market)

	assert.Nil(t, err)

}

func TestFailWithBlankValues(t *testing.T) {
	market := StreetMarket{}

	err := Validate(&market)
	errs := err.(validator.ErrorMap)
	assert.Equal(t, validator.ErrZeroValue, errs["CodigoDistrito"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["Logradouro"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["Longitude"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["AreaPonderacao"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["Regiao8"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["SetorCensitario"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["Distrito"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["CodSubPrefeitura"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["Regiao5"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["Latitude"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["SubPrefeitura"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["NomeFeira"][0])
	assert.Equal(t, validator.ErrZeroValue, errs["Registro"][0])
}

func TestFieldsWithSizeOverAllowed(t *testing.T) {

	market := StreetMarket{
		Id: 1, Longitude: "Lorem Ipsum is simply dummy text of the printing ",
		Latitude:        "Lorem Ipsum is simply dummy text of the printing ",
		SetorCensitario: "Lorem Ipsum is simply dummy text of the printing ",
		AreaPonderacao:  "Lorem Ipsum is simply dummy text of the printing ", CodigoDistrito: 87,
		Distrito:         "Lorem Ipsum is simply dummy text of the printing ",
		CodSubPrefeitura: 26, SubPrefeitura: "Lorem Ipsum is simply dummy text of the printing ",
		Regiao5:    "Lorem Ipsum is simply dummy text of the printing ",
		Regiao8:    "Lorem Ipsum is simply dummy text of the printing ",
		NomeFeira:  "Lorem Ipsum is simply dummy text of the printing ",
		Registro:   "Lorem Ipsum is simply dummy text of the printing ",
		Logradouro: "Lorem Ipsum is simply dummy text of the printing ",
		Numero:     "Lorem Ipsum is simply dummy text of the printing ",
		Bairro:     "Lorem Ipsum is simply dummy text of the printing ",
		Referencia: "Lorem Ipsum is simply dummy text of the printing Lorem Ipsum is simply dummy text of the printing Lorem Ipsum is simply dummy text of the printing Lorem Ipsum is simply dummy text of the printing Lorem Ipsum is simply dummy text of the printing Lorem Ipsum is simply dummy text of the printing Lorem Ipsum is simply dummy text of the printing Lorem Ipsum is simply dummy text of the printing Lorem Ipsum is simply dummy text of the printing ",
	}

	err := Validate(&market)
	fmt.Println(err)
	errs := err.(validator.ErrorMap)
	fmt.Println(len(errs))

	assert.Equal(t, validator.ErrMax, errs["Logradouro"][0])
	assert.Equal(t, validator.ErrMax, errs["Longitude"][0])
	assert.Equal(t, validator.ErrMax, errs["AreaPonderacao"][0])
	assert.Equal(t, validator.ErrMax, errs["Regiao8"][0])
	assert.Equal(t, validator.ErrMax, errs["SetorCensitario"][0])
	assert.Equal(t, validator.ErrMax, errs["Distrito"][0])
	assert.Equal(t, validator.ErrMax, errs["Regiao5"][0])
	assert.Equal(t, validator.ErrMax, errs["Latitude"][0])
	assert.Equal(t, validator.ErrMax, errs["SubPrefeitura"][0])
	assert.Equal(t, validator.ErrMax, errs["NomeFeira"][0])
	assert.Equal(t, validator.ErrMax, errs["Registro"][0])
	assert.Equal(t, validator.ErrMax, errs["Numero"][0])
	assert.Equal(t, validator.ErrMax, errs["Bairro"][0])
	assert.Equal(t, validator.ErrMax, errs["Referencia"][0])
}
