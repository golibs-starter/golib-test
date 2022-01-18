package golibtest

import (
	"gitlab.com/golibs-starter/golib-message-bus/kafka/core"
	"gitlab.com/golibs-starter/golib-message-bus/kafka/properties"
)

// KafkaTestSuite ...
type KafkaTestSuite struct {
	kafkaProperties *properties.Client
	messages        map[string][]string
}

// NewKafkaTestSuite ...
func NewKafkaTestSuite(kafkaProperties *properties.Client) *KafkaTestSuite {
	return &KafkaTestSuite{
		kafkaProperties: kafkaProperties,
		messages:        map[string][]string{},
	}
}

// PushMessage ...
func (k *KafkaTestSuite) PushMessage(message *core.ConsumerMessage) {
	if _, ok := k.messages[message.Topic]; ok {
		k.messages[message.Topic] = append(k.messages[message.Topic], string(message.Value))
	} else {
		k.messages[message.Topic] = []string{string(message.Value)}
	}
}

// ClearMessages ...
func (k *KafkaTestSuite) ClearMessages(topic string)  {
	if _, ok := k.messages[topic]; ok {
		k.messages[topic] = []string{}
	}
}

// Count ...
func (k *KafkaTestSuite) Count(topic string) int64 {
	if val, ok := k.messages[topic]; ok {
		return int64(len(val))
	}
	return 0
}

// GetMessages ...
func (k *KafkaTestSuite) GetMessages(topic string) []string {
	if messages, ok := k.messages[topic]; ok {
		return messages
	}
	return []string{}
}
