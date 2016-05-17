package main 

import (
	"fmt"
	"testing"
)

func TestSayHi(t *testing.T){
	if SayHi() != "Hi1"{
		t.Fatal("Error in function SayHi")
	}else{
		fmt.Println("Ok")
	}
}