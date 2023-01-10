package utils

import (
	"testing"
)

func TestCheckParam(t *testing.T) {
	err := CheckParam("", "")
	if err.Err != nil {
		t.Errorf("got:%v", err)
	}
}
