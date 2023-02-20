package conv

import (
	// "reflect"
	"math"
	"testing"
)
/*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/

func withinTolerance(a, b, error float64) bool {
	// Først sjekker om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
		return difference < error
	}

	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference / math.Abs(b)) < error
}

// Konverterer Farhenheit til Celsius
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Tester inputen og sammenligner den med ønsket
	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her
// ...

// Konverterer Celcius til Farhenheit
func TestCelsiusToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Tester inputen og sammenligner den med ønsket
	tests := []test{
		{input: 56.67, want: 134},
	}
	for _, tc := range tests {
		got := CelsiusToFarhenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Konverterer Kelvin til Celcius
func TestKelvinToCelcius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
// Tester inputen og sammenligner den med ønsket
	tests := []test{
		{input: 329.82, want: 56.67},
	}
	for _, tc := range tests {
		got := KelvinToCelcius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Konverterer Celcius til Kelvin
func TestCelciusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Tester inputen og sammenligner den med ønsket
	tests := []test{
		{input: 56.67, want: 329.82},
	}
	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Konverterer Kelvin til Farhenheit
func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Tester inputen og sammenligner den med ønsket
	tests := []test{
		{input: 329.82, want: 134.00},
	}
	for _, tc := range tests {
		got := KelvinToFarhenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Konverterer Farhenheit til Kelvin
func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Tester inputen og sammenligner den med ønsket
	tests := []test{
		{input: 134.00, want: 329.82},
	}
	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
