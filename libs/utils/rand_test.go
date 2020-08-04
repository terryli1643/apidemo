package utils

import (
	"fmt"
	"testing"
)

func TestRandomSecret(t *testing.T) {
	fmt.Println(RandString(32))
}
