package main 

import (
	"fmt"
	"testing"
)

func TestSayHi(t *testing.T){
	if SayHi() != "Hi"{
		t.Fatal("Error in function SayHi")
	}
}