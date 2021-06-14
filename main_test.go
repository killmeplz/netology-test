package main

import (
	"log"
	"testing"
)

func TestDivision(t *testing.T) {
	got, _ := Division(2,1)
	if got != 2 {
		t.Errorf("Division(2,1) = %f; want 2", got)
	}
	got, err := Division(2,0)
	if err == nil {
		t.Errorf("No divizion by zero exception Division(2,0) want error")
	}
	got, err = Division(1,2)
	if got != 0.5 {
		log.Println(err)
		t.Errorf("Division(1,2) = %f; want 0.5", got)
	}
}