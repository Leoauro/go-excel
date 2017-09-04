package test

import (
	"testing"

	"github.com/szyhf/go-excel/internal"
)

func TestOpen(t *testing.T) {
	// file
	conn := internal.NewConnect()
	err := conn.Open(filePath)
	if err != nil {
		t.Error(err)
	}
	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestNewReader(t *testing.T) {
	// file
	conn := internal.NewConnect()
	err := conn.Open(filePath)
	if err != nil {
		t.Error(err)
	}

}