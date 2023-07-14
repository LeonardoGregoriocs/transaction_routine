package client

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	documentNumber = "123456789"
)

func TestNewClient(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(reflect.TypeOf(documentNumber).Kind(), reflect.String)
}

func TestDateCreateUpdateMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	newClient, _ := NewClient(documentNumber)

	assert.Greater(newClient.DateCreateUpdate, now)
}

func TestValidateDocumentNumber(t *testing.T) {
	assert := assert.New(t)

	_, err := NewClient("")

	assert.Equal("Document Number is required", err.Error())
}
