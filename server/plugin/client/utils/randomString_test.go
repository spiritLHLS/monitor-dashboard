package utils

import (
	"testing"
	"fmt"
)

func TestRandomString(t *testing.T) {
	str := RandomString(6)
	if len(str) == 6 {
		fmt.Println("Generated random string:", str)
		t.Log("Test passed.")
	} else {
		t.Error("Generated string length is not 4.")
	}
}