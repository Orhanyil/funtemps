package conv

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	
	return (value - 32) * 5 / 9
}

// De andre konverteringsfunksjonene implementeres her
// ...

// Konverterer Celcius til Farhenheit
func CelsiusToFarhenheit(value float64) float64 {

	return value*9/5 + 32
}

// Konverterer Kelvin til Celcius
func KelvinToCelcius(value float64) float64 {

	return value - 273.15
}

// Konverterer Celcius til Kelvin
func CelsiusToKelvin(value float64) float64 {

	return value + 273.15
}

// Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(value float64) float64 {

	return (value-273.15)*9/5 + 32
}

// Konverterer Farhenheit til Kelvin
func FarhenheitToKelvin(value float64) float64 {

	return (value-32)*5/9 + 273.15
}
