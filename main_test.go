package main

import (
	"codes/gogin/protos"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestPingRoute(t *testing.T) {
	r := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestProtobuf(t *testing.T) {
	r := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/someproto", nil)
	r.ServeHTTP(w, req)

	var data protos.Greeting
	proto.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, data.Message, "hello")
	assert.Equal(t, len(data.Label), 3)
	assert.Equal(t, data.Label[0], int64(1))
}
