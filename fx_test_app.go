package golibtest

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
)

func SetupTestApp(options []fx.Option) (*fx.App, error) {
	// We need add RegisterHttpHandlerOnStartOpt here
	// to ensure this will invoke at last
	options = append(options, RegisterHttpHandlerOnStartOpt())

	// Validate our fx options is set up correctly
	err := fx.ValidateApp(options...)
	if err != nil {
		return nil, errors.WithMessage(err, "Fail to validate fx options")
	}
	log.Info("All fx options are valid")

	app := fx.New(options...)
	if err := app.Start(context.Background()); err != nil {
		return nil, errors.WithMessage(err, "Error when start fx app")
	}
	return app, nil
}
