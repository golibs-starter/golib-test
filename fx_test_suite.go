package golibtest

import (
	"github.com/golibs-starter/golib"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
)

type FxTestSuite struct {
	suite.Suite
	app       *fx.App
	tsOptions []TsConfig
	fxOptions []fx.Option
}

func (s *FxTestSuite) Config(opts ...TsConfig) {
	s.tsOptions = append(s.tsOptions, opts...)
}

func (s *FxTestSuite) Options(opts []fx.Option) {
	s.Option(opts...)
}

func (s *FxTestSuite) Option(opts ...fx.Option) {
	for _, opt := range opts {
		s.Config(WithFxOption(opt))
	}
}

func (s *FxTestSuite) ProfilePath(paths ...string) {
	s.Option(golib.ProvidePropsOption(golib.WithPaths(paths)))
}

func (s *FxTestSuite) Profile(profiles ...string) {
	s.Option(golib.ProvidePropsOption(golib.WithActiveProfiles(profiles)))
}

func (s *FxTestSuite) Populate(targets ...interface{}) {
	s.Option(fx.Populate(targets...))
}

func (s *FxTestSuite) Decorate(decorators ...interface{}) {
	s.Option(fx.Decorate(decorators...))
}

func (s *FxTestSuite) Provide(constructors ...interface{}) {
	s.Option(fx.Provide(constructors...))
}

func (s *FxTestSuite) Invoke(funcs ...interface{}) {
	s.Option(fx.Invoke(funcs...))
}

func (s *FxTestSuite) StartApp() {
	// Apply all TsConfig
	for _, tsOption := range s.tsOptions {
		tsOption(s)
	}

	var err error
	s.app, err = StartFxAppT(s.T(), s.fxOptions...)
	s.Require().NoError(err, "Error when setup test app")
}
