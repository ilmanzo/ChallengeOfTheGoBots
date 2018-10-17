package main

import "testing"

var result string

func TestLeftpad1(t *testing.T) {
	s := leftpad1("test", 10, '-')
	if s != "------test" {
		t.Error("test failed")
	}
}

func TestLeftpad1_small(t *testing.T) {
	s := leftpad1("test", 3, '-')
	if s != "test" {
		t.Error("test failed")
	}
}

func TestLeftpadComparison(t *testing.T) {
	s1 := leftpad1("test", 10, '-')
	s2 := leftpad2("test", 10, '-')
	if s1 != s2 {
		t.Error("test failed")
	}
}

func BenchmarkLeftpad1(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = leftpad1("a", 50, '*')
	}
	result = r //per disattivare l'ottimizzazione
}

/*
func BenchmarkLeftpad2(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = leftpad2("a", 50, '*')
	}
	result = r
}
*/
