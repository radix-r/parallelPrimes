package main

import "testing"

/*
Test yo shit
*/

func TestProbPrimeCheck(t *testing.T){
	res := ProbPrimeCheck(3929)

	if !res{
		t.Error("Expected true, got", res)
	}

}

func TestSimplePrimeCheck(t *testing.T){
	res:= SimplePrimeCheck(3929,0,3929)

	if !res{
		t.Error("Expected true, got", res)
	}

	res = SimplePrimeCheck(3929*127,0,3929*127)
	if res{
		t.Error("Expected false, got", res)
	}
}