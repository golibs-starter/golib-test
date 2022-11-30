package golibtest

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func RequireFxApp(options ...fx.Option) *fx.App {
	app, err := StartFxApp(options...)
	if err != nil {
		panic(fmt.Errorf("error when setup test app: [%v]", err))
	}
	return app
}

func StartFxApp(options ...fx.Option) (*fx.App, error) {
	// We need add RegisterHttpHandlerOnStartOpt here
	// to ensure this will invoke at last
	//options = append(options, RegisterHttpHandlerOnStartOpt())

	app := fx.New(options...)
	if err := app.Err(); err != nil {
		return nil, errors.WithMessage(err, "fx.New failed")
	}
	if err := app.Start(context.Background()); err != nil {
		return nil, errors.WithMessage(err, "Error when start fx app")
	}
	return app, nil
}

func RequireFxAppT(tb testing.TB, options ...fx.Option) *fx.App {
	app, err := StartFxAppT(tb, options...)
	if err != nil {
		panic(fmt.Errorf("error when setup test app: [%v]", err))
	}
	return app
}

func StartFxAppT(tb testing.TB, options ...fx.Option) (*fx.App, error) {
	options = append(
		options,
		// Wrap current logger with testing logger
		WrapTestingLoggerOpt(tb),

		// We need add RegisterHttpHandlerOnStartOpt here
		// to ensure this will invoke at last
		//RegisterHttpHandlerOnStartOpt(),
	)

	app := fxtest.New(tb, options...)
	if err := app.Start(context.Background()); err != nil {
		return nil, errors.WithMessage(err, "Error when start fxtest app")
	}
	return app.App, nil
}

func WrapTestingLoggerOpt(tb testing.TB) fx.Option {
	return fx.Decorate(fx.Annotate(
		func(oldLogger log.Logger) (coreLogger log.Logger, webLogger log.Logger, err error) {
			var testingLogger *log.TestingLogger
			if defaultLogger, ok := oldLogger.(*log.DefaultLogger); ok {
				testingLogger = log.NewTestingLoggerFromDefault(tb, defaultLogger)
			} else {
				if testingLogger, err = log.NewTestingLogger(tb, &log.Options{CallerSkip: 2}); err != nil {
					return nil, nil, errors.WithMessage(err, "init testing logger failed")
				}
			}
			coreLogger = testingLogger
			webLogger = testingLogger.Clone(log.AddCallerSkip(1))
			return
		},
		fx.ResultTags(``, `name:"web_logger"`),
	))
}
