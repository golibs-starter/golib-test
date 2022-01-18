package golibtest

import (
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/golibs-starter/golib-migrate"
	"go.uber.org/fx"
)

func NeedSecurityTestSuite() fx.Option {
	return fx.Options(
		golib.ProvideProps(NewJwtTestProperties),
		fx.Provide(NewSecurityTestSuite),
	)
}

func NeedDatabaseTestSuite() fx.Option {
	return fx.Options(
		fx.Provide(NewDatabaseTestSuite),
	)
}

func NeedKafkaTestSuite() fx.Option {
	return fx.Options(
		golibmsg.KafkaConsumerOpt(),
		fx.Provide(NewKafkaTestSuite),
	)
}

func NeedMigration() fx.Option {
	return fx.Options(
		golibmigrate.MigrationOpt(),
	)
}
