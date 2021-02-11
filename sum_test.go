package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("Deu pala")
	}
}

func TestDivide(t *testing.T) {

	result := divide(20, 5)

	if result != 4 {
		t.Error("Deu Pala")
	}
}

func TestMultiply(t *testing.T) {

	result := multiply(2, 3)

	if result != 5 {
		t.Error("Deu pala")
	}
}