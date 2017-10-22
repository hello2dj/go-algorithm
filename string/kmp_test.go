package string

import (
	"fmt"
	"strings"
	"testing"
)

func Test_generateIndexArr(t *testing.T) {
	// fmt.Println(strings.Filed("fjl sajki  jidf "))
	fmt.Println(strings.Fields("jasio sjdif jaiof sjfi"))
	if KMPSearch("12312", "312") != 2 {
		t.Error("wrong match")
	}
}
