package golibtest

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestSuite represent testing suite
type TestSuite struct {
	SecurityTestSuite *SecurityTestSuite
	DatabaseTestSuite *DatabaseTestSuite
	KafkaTestSuite    *KafkaTestSuite
}

// CreateJwtToken return a new jwt token
func (ts *TestSuite) CreateJwtToken(userId string) string {
	return ts.SecurityTestSuite.CreateJwtToken(userId)
}

// TruncateTables run truncate statement
func (ts *TestSuite) TruncateTables(tables []string) {
	for _, table := range tables {
		ts.DatabaseTestSuite.TruncateTable(table)
	}
}

// Seed insert data to database
func (ts *TestSuite) Seed(model interface{}) {
	ts.DatabaseTestSuite.Insert(model)
}

// AssertDatabaseCount assert database has number of row without query
func (ts *TestSuite) AssertDatabaseCount(t *testing.T, table string, expected int64) {
	count := ts.DatabaseTestSuite.CountWithoutQuery(table)
	require.Equal(t, expected, count)
}

// AssertDatabaseHas assert database has more than a row with query params
func (ts *TestSuite) AssertDatabaseHas(t *testing.T, table string, conditions map[string]interface{}) {
	count := ts.DatabaseTestSuite.CountWithQuery(table, conditions)
	require.GreaterOrEqual(t, count, int64(1), "Record not found in database with query:", conditions)
}

// RefreshKafkaTopics delete topics kafka
func (ts *TestSuite) RefreshKafkaTopics(topics []string) {
	for _, topic := range topics {
		ts.KafkaTestSuite.ClearMessages(topic)
	}
}

// AssertKafkaCount ...
func (ts *TestSuite) AssertKafkaCount(t *testing.T, topic string, expected int64)  {
	count := ts.KafkaTestSuite.Count(topic)
	require.Equal(t, expected, count)
}

// GetKafkaMessages ...
func (ts *TestSuite) GetKafkaMessages(topic string) []string {
	return ts.KafkaTestSuite.GetMessages(topic)
}
