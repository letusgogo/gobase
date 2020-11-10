package kafka

import "encoding/json"

type Marshaler interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
	String() string
}

type jsonCodec struct{}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (jsonCodec) String() string {
	return "json"
}
