package golibtest

import (
	"crypto/rsa"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
	"gitlab.com/golibs-starter/golib-message-bus/kafka/properties"
	"gitlab.com/golibs-starter/golib/web/log"
	"gorm.io/gorm"
	"testing"
	"time"
)

// TestSuite represent testing suite
type TestSuite struct {
	DB              *gorm.DB
	KafkaProperties *properties.Client
	jwtSignKey      *rsa.PrivateKey
}

// CreateJwtToken return a new jwt token
func (ts *TestSuite) CreateJwtToken(userId string) string {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	now := time.Now()
	token.Claims = &jwt.StandardClaims{
		Issuer:    "TESTER",
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Minute * 1).Unix(),
		Subject:   userId,
	}
	jwtToken, err := token.SignedString(ts.jwtSignKey)
	if err != nil {
		log.Fatalf("Could not create jwt token: %v", err)
	}
	return jwtToken
}

// LoadJwtPrivateKey load jwt config from properties
func (ts *TestSuite) LoadJwtPrivateKey(jwtProperties *JwtTestProperties) {
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(jwtProperties.PrivateKey))
	if err != nil {
		log.Fatalf("Could not load jwt private key: %v", err)
	}
	ts.jwtSignKey = signKey
}

// TruncateTables run truncate statement
func (ts *TestSuite) TruncateTables(tables []string) {
	for _, table := range tables {
		if err := ts.DB.Exec(fmt.Sprintf("TRUNCATE TABLE `%s`", table)).Error; err != nil {
			log.Fatalf("Could not truncate table %s: %v", table, err)
		}
	}
}

// Seed insert data to database
func (ts *TestSuite) Seed(model interface{}) {
	if err := ts.DB.Create(model).Error; err != nil {
		log.Fatalf("Could not create seed data, model: %v, err: %v")
	}
}

// AssertDatabaseCount assert database has number of row without query
func (ts *TestSuite) AssertDatabaseCount(t *testing.T, table string, expected int64) {
	var count int64
	ts.DB.Table(table).Count(&count)
	require.Equal(t, expected, count)
}

// AssertDatabaseHas assert database has more than a row with query params
func (ts *TestSuite) AssertDatabaseHas(t *testing.T, table string, conditions map[string]interface{}) {
	var count int64
	ts.DB.Table(table).Where(conditions).Count(&count)
	require.GreaterOrEqual(t, count, int64(1), "Record not found in database with query:", conditions)
}

// DeleteKafkaTopics delete topics kafka
func (ts *TestSuite) DeleteKafkaTopics() {
	config := sarama.NewConfig()
	config.Version = sarama.V1_1_0_0
	clusterAdmin, err := sarama.NewClusterAdmin(ts.KafkaProperties.BootstrapServers, config)
	if err != nil {
		log.Fatalf("Could not create sarama cluster admin: %v", err)
	}
	defer clusterAdmin.Close()
	topics, err := clusterAdmin.ListTopics()
	if err != nil {
		log.Fatalf("Could not list kafka topics", err)
	}
	for topic, _ := range topics {
		if err := clusterAdmin.DeleteTopic(topic); err != nil && err != sarama.ErrUnknownTopicOrPartition {
			log.Fatalf("Could not delete kafka topic", err)
		}
	}
	clusterAdmin.ListTopics()
}
