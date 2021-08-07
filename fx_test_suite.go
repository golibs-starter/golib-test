package golibtest

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib/config"
	"go.uber.org/fx"
	"os"
	"strings"
	"syscall"
)

type FxTestSuite struct {
	suite.Suite
	testingDir  string
	profiles    []string
	options     []fx.Option
	invokeStart fx.Option
	port        int
	baseUrl     string
	ready       chan bool
}

func NewFxTestSuite(bootstrap []fx.Option, options ...TsOption) *FxTestSuite {
	ts := FxTestSuite{ready: make(chan bool)}
	ts.options = bootstrap
	ts.profiles = []string{"testing"}
	for _, tsOption := range options {
		tsOption(&ts)
	}
	ReplaceFxOption(
		fx.Provide(golib.NewPropertiesAutoLoad),
		fx.Provide(ts.NewPropertiesAutoLoad),
	)(&ts)
	return &ts
}

func (s *FxTestSuite) SetupSuite() {
	options := append(s.options, s.invokePrepare())
	if s.invokeStart != nil {
		options = append(options, s.invokeStart)
	}
	err := fx.ValidateApp(options...)
	s.Require().NoErrorf(err, "Cannot start application")
	go fx.New(options...).Run()
	<-s.ready
}

func (s *FxTestSuite) invokePrepare() fx.Option {
	return fx.Invoke(func(lifecycle fx.Lifecycle, app *golib.App) {
		s.port = app.Port()
		s.baseUrl = fmt.Sprintf("http://localhost:%d", s.port)
	})
}

func (s *FxTestSuite) NewPropertiesAutoLoad() (config.Loader, *config.AppProperties, error) {
	return golib.NewPropertiesAutoLoad(
		golib.WithActiveProfiles(s.profiles),
		golib.WithPaths([]string{
			"../" + s.testingDir + "/config",
			s.testingDir + "/config",
		}),
	)
}

func (s *FxTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	_ = p.Signal(syscall.SIGINT)
}

func (s FxTestSuite) URL(path string) string {
	return fmt.Sprintf("%s/%s", s.baseUrl, strings.TrimLeft(path, "/"))
}
