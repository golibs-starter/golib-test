package golibtest

import (
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib/config"
	"go.uber.org/fx"
	"strings"
)

var DefaultTestingProfile = "testing"

type TsOption func(ts *FxTestSuite)

func WithTestingDir(dir string) TsOption {
	return func(ts *FxTestSuite) {
		dir = strings.Trim(dir, "/")
		if dir == "" {
			return
		}
		ts.configPaths = []string{
			"../" + dir + "/" + config.DefaultConfigPath,
			dir + "/" + config.DefaultConfigPath,
		}
	}
}

func WithConfigPaths(paths []string) TsOption {
	return func(ts *FxTestSuite) {
		ts.configPaths = paths
	}
}

func WithActiveProfiles(profiles ...string) TsOption {
	return func(ts *FxTestSuite) {
		ts.options = append(ts.options, golib.ProvidePropsOption(golib.WithActiveProfiles(profiles)))
	}
}

func ReplaceFxOption(opt fx.Option, newOpt fx.Option) TsOption {
	return func(ts *FxTestSuite) {
		if !strings.HasPrefix(opt.String(), "fx.Provide") {
			ts.Require().FailNow("Replacement option only support for fx.Provide")
			return
		}
		for k, v := range ts.options {
			if v.String() == opt.String() {
				ts.options[k] = newOpt
				return
			}
		}
		ts.Require().FailNowf("Cannot replace option", "Replacement option %s not found, new option %s", opt, newOpt)
	}
}

func WithFxOption(opt fx.Option) TsOption {
	return func(ts *FxTestSuite) {
		for _, v := range ts.options {
			if v.String() == opt.String() {
				ts.Require().FailNowf("Cannot add new option", "Replacement option %s already exists", opt)
				return
			}
		}
		ts.options = append(ts.options, opt)
	}
}

func WithFxPopulate(targets ...interface{}) TsOption {
	return WithFxOption(fx.Populate(targets...))
}

func WithInvokeStart(invokeStart func(done func(err error)) fx.Option) TsOption {
	return func(ts *FxTestSuite) {
		ts.invokeStart = invokeStart(func(err error) {
			ts.readyError <- err
		})
	}
}
