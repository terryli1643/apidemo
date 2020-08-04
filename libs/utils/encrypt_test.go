package utils

import (
	"fmt"
	"testing"
)

var key = []byte("example key 1234")

func TestEncrypt(t *testing.T) {
	originalText := "this will be encrypt"
	fmt.Println(Encrypt(key, originalText))
}

func TestDecrypt(t *testing.T) {
	cryotoText := "q8-xX9-h4UtdD5KOxjEKSj-hN04uqUjuA6c4KbrY4FyJFUbh"
	cryotoText2 := "4hP_Oo0sZqLSdL1l_bYj5LeuxJfYAX-BNPeLD-6wTiNGfQL7"
	fmt.Println(Decrypt(key, cryotoText))
	fmt.Println(Decrypt(key, cryotoText2))
}
