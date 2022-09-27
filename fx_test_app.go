package golibtest

import (
	"github.com/pkg/errors"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
)

func SetupTestApp(options []fx.Option) (*fx.App, error) {
	// We need add HttpTestServerStartupOpt here
	// to ensure this will invoke at last
	options = append(options, HttpTestServerStartupOpt())

	// Validate our fx options is set up correctly
	err := fx.ValidateApp(options...)
	if err != nil {
		return nil, errors.WithMessage(err, "Fail to validate fx options")
	}
	log.Info("All fx options are valid")

	app := fx.New(options...)
	fxHandler := FxHttpHandler(app)
	RegisterHttpHandler(fxHandler)
	return app, nil
}
