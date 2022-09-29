package golibtest

import "go.uber.org/fx"

type TsConfig func(ts *FxTestSuite)

func WithFxOptions(opts ...fx.Option) TsConfig {
	return func(ts *FxTestSuite) {
		for _, opt := range opts {
			WithFxOption(opt)(ts)
		}
	}
}

func WithFxOption(opt fx.Option) TsConfig {
	return func(ts *FxTestSuite) {
		ts.fxOptions = append(ts.fxOptions, opt)
	}
}
