package internal

import (
	"errors"
	"reflect"
)

var AttributeIndexer map[string]AttributeSchema

type AttributeSchema struct{
	Name string
	Type string
	Indexer bool
	SecondaryIndex map[string][]string
}

func (obs AttributeSchema) Validate(value interface{}) error {
	if obs.Type != reflect.TypeOf(value).Name() {
		return errors.New("error type mismatch")
	}

	return nil
}

func GetOrCreateAttributeSchema(AttributeKey string, AttributeValue interface{}) AttributeSchema {

	object, found := AttributeIndexer[AttributeKey]

	if !found {
		object = AttributeSchema{
			Name: AttributeKey,
			Type: reflect.TypeOf(AttributeValue).Name(),
			Indexer: true,
			SecondaryIndex: map[string][]string{},
		}

		AttributeIndexer[AttributeKey] =  object
	}

	return object
}

func GetAttributeSchema(AttributeKey string) AttributeSchema {
	object, found := AttributeIndexer[AttributeKey]
	if !found {
		err := errors.New("attribute not found in indexer")
		panic(err)
	}

	return object
}