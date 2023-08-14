// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"github.com/weplanx/workflow/excel/api"
	"github.com/weplanx/workflow/excel/common"
)

// Injectors from wire.go:

func NewAPI() (*api.API, error) {
	values, err := LoadStaticValues()
	if err != nil {
		return nil, err
	}
	client, err := UseCos(values)
	if err != nil {
		return nil, err
	}
	inject := &common.Inject{
		V:      values,
		Client: client,
	}
	apiAPI := &api.API{
		Inject: inject,
	}
	return apiAPI, nil
}
