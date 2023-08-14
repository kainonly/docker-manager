// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"github.com/weplanx/workflow/schedule/app"
	"github.com/weplanx/workflow/schedule/common"
)

// Injectors from wire.go:

func NewApp() (*app.App, error) {
	values, err := LoadStaticValues()
	if err != nil {
		return nil, err
	}
	logger, err := UseZap()
	if err != nil {
		return nil, err
	}
	conn, err := UseNats(values)
	if err != nil {
		return nil, err
	}
	jetStreamContext, err := UseJetStream(conn)
	if err != nil {
		return nil, err
	}
	keyValue, err := UseKeyValue(values, jetStreamContext)
	if err != nil {
		return nil, err
	}
	inject := &common.Inject{
		V:        values,
		Log:      logger,
		Nats:     conn,
		KeyValue: keyValue,
	}
	appApp := app.Initialize(inject)
	return appApp, nil
}
