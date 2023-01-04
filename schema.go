package schemas

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"reflect"
	"time"

	commodityv3 "github.com/neilmcgibbon/eddn-schemas/schemas/commodity/v3"
	journalv1 "github.com/neilmcgibbon/eddn-schemas/schemas/journal/v1"
)

type Envelope struct {
	SchemaRef  string `json:"$schemaRef"`
	Header     Header `json:"header"`
	Data       Data
	RawMessage json.RawMessage `json:"message"`
}

type Header struct {
	Gamebuild        string    `json:"gamebuild"`
	Gameversion      string    `json:"gameversion"`
	GatewayTimestamp time.Time `json:"gatewayTimestamp"`
	SoftwareName     string    `json:"softwareName"`
	SoftwareVersion  string    `json:"softwareVersion"`
	UploaderID       string    `json:"uploaderID"`
}

type Data interface {
	GetSchema() string
	PostDecode()
}

var schemaRegister map[string]reflect.Type = map[string]reflect.Type{
	"https://eddn.edcd.io/schemas/commodity/3": reflect.TypeOf(commodityv3.Schema{}),
	"https://eddn.edcd.io/schemas/journal/1":   reflect.TypeOf(journalv1.Schema{}),
}

func Decode(raw []byte) (*Envelope, error) {
	var envelope Envelope

	// The zeromq message is compressed, so we decompress it.
	r, err := zlib.NewReader(bytes.NewReader(raw))
	if err != nil {
		return &envelope, err
	}
	defer r.Close()

	// Decode the message into our Envelope struct
	err = json.NewDecoder(r).Decode(&envelope)
	if err != nil {
		return &envelope, err
	}

	// Make sure the schema type (defined in the envelope) is of a type we handle
	t, ok := schemaRegister[envelope.SchemaRef]
	if !ok {
		// Note we don'r return an error - we just skip the message if we don't have the
		// schema files to handle it
		return &envelope, nil
	}

	// Cast the envelope data property to be the type of the schema/structure
	envelope.Data = reflect.New(t).Interface().(Data)

	// If the data property of envelop is nil, then we (for some reason) don't have the
	// strcut types to handle this. This should never happen as test in the schema register
	// (above) before we get here
	if envelope.Data == nil {
		// @todo we might want to return an acrual error here
		return &envelope, nil
	}

	// Unmarshal the envelope contents into our type-casted data property
	if err = json.Unmarshal(envelope.RawMessage, envelope.Data); err != nil {
		// Error with JSON decoding
		return &envelope, err
	}

	// Do any schema specific post-decode actions
	envelope.Data.PostDecode()

	return &envelope, nil

}
