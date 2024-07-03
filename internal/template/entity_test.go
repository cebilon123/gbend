package template

import (
	"bytes"
	"context"
	"io"
	"testing"
)

func TestWriteEntity(t *testing.T) {
	var buf []byte

	writer := bytes.NewBuffer(buf)

	if err := WriteEntity(context.Background(), "test", map[string]string{
		"id":        "string",
		"age":       "int",
		"createdAt": "datetime",
	}, writer); err != nil {
		t.Errorf("unexpected err: %s", err.Error())
	}

	_, err := io.ReadAll(writer)
	if err != nil {
		t.Errorf("unexpected err: %s", err.Error())
	}
}
