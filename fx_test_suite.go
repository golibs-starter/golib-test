package golibtest

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib/config"
	"go.uber.org/fx"
	"net"
	"os"
	"strconv"
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
	readyError  chan error
}

func NewFxTestSuite(bootstrap []fx.Option, options ...TsOption) *FxTestSuite {
	ts := FxTestSuite{readyError: make(chan error)}
	ts.options = bootstrap
	ts.profiles = []string{"testing"}
	for _, tsOption := range options {
		tsOption(&ts)
	}
	ReplaceFxOption(
		golib.PropertiesOpt(),
		fx.Provide(ts.NewPropertiesLoader),
	)(&ts)
	return &ts
}

func (s *FxTestSuite) SetupSuite() {
	options := s.collectOptions()
	err := fx.ValidateApp(options...)
	s.Require().NoErrorf(err, "Fail to validate fx options")

	go func() {
		app := fx.New(options...)
		if err := app.Err(); err != nil {
			s.readyError <- err
			return
		}
		startCtx, cancel := context.WithTimeout(context.Background(), app.StartTimeout())
		defer cancel()
		if err := app.Start(startCtx); err != nil {
			s.readyError <- err
			return
		}
		<-app.Done()
	}()

	if err := <-s.readyError; err != nil {
		s.FailNowf("Error when start application", "Error: %v", err)
	}
}

func (s *FxTestSuite) collectOptions() []fx.Option {
	options := append(
		s.options,
		s.networkPrepare(),
		s.invokePrepare(),
	)
	if s.invokeStart != nil {
		options = append(options, s.invokeStart)
	}
	return options
}

func (s *FxTestSuite) networkPrepare() fx.Option {
	return fx.Provide(func(app *golib.App) (net.Listener, error) {
		port := app.Port()
		if app.Port() <= 0 {
			port = 0
		}
		return net.Listen("tcp", fmt.Sprintf(":%d", port))
	})
}

func (s *FxTestSuite) invokePrepare() fx.Option {
	return fx.Invoke(func(lifecycle fx.Lifecycle, app *golib.App, ln net.Listener) error {
		_, portStr, _ := net.SplitHostPort(ln.Addr().String())
		port, err := strconv.Atoi(portStr)
		s.Require().NoErrorf(err, "Fail to select http port")
		s.T().Logf("Application [%s] will be served at %d", app.Name(), port)
		s.port = port
		s.baseUrl = fmt.Sprintf("http://localhost:%d", s.port)
		return nil
	})
}

func (s *FxTestSuite) NewPropertiesLoader(in golib.PropertiesLoaderIn) (config.Loader, error) {
	return golib.NewPropertiesLoader(in,
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

func (s *FxTestSuite) URL(path string) string {
	return fmt.Sprintf("%s/%s", s.baseUrl, strings.TrimLeft(path, "/"))
}
