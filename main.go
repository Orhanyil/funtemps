// for å klare denne oppgaven brukte jeg på github, chatGBT, forelesningene til Janis,
//og mye google.

package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Definisjon
var (
	fahr    float64
	kelvin  float64
	celsius float64
	out     string
	funf    string
	t       string
)

func init() {

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader Fahrenheit")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader Kelvin")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader Celsius")
	flag.StringVar(&out, "out", "C", "beregne temperatur i F - farhenheit")
	flag.StringVar(&funf, "funfacts", "sun", " \"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "f", "temperature unit to convert to (c, f, k)")

}

func main() {
	flag.Parse()

	//Fahrenheit to Celsius

	if out == "C" && isFlagPassed("F") {
		celsius = conv.FarhenheitToCelsius(fahr)
		fmt.Printf("%.2f°F er %.2f°C", fahr, celsius)
	}

	//Fahrenheit to Kelvin
	if out == "K" && isFlagPassed("F") {
		kelvin = conv.FarhenheitToKelvin(fahr)
		fmt.Printf("%.2f °F er %.2f °K", fahr, kelvin)
	}

	//Celsius to Fahrenheit
	if out == "F" && isFlagPassed("C") {
		fahr = conv.CelsiusToFarhenheit(celsius)
		fmt.Printf("%.2f°C er %.2f°F", celsius, fahr)
	}

	//Celsius to Kelvin
	if out == "K" && isFlagPassed("C") {
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Printf("%.2f°C er %.2f°K", celsius, kelvin)
	}

	//Kelvin to Fahrenheit
	if out == "F" && isFlagPassed("K") {
		fahr = conv.KelvinToFarhenheit(kelvin)
		fmt.Printf("%.2f°K er %.2f°F", kelvin, fahr)
	}

	//Kelvin to Celsius
	if out == "C" && isFlagPassed("K") {
		celsius = conv.KelvinToCelcius(kelvin)
		fmt.Printf("%.2f°K er %.2f°C", kelvin, celsius)
	}

	if funf == "terra" && isFlagPassed("funfacts") {
		terrafact := funfacts.GetFunFacts(funf)
		if t == "C" {
			fmt.Printf("%s %s °C. %s %s °C. \n%s %s °C.", terrafact[0], formatNumber(56.7), terrafact[1], formatNumber(-89.4), terrafact[2], formatNumber(9118))
		} else if t == "K" {
			fmt.Printf("%s %s K. %s %s \n%s %s K.", terrafact[0], formatNumber(conv.CelsiusToKelvin(56.7)), terrafact[1], formatNumber(conv.CelsiusToKelvin(-89.4)), terrafact[2], formatNumber(conv.CelsiusToKelvin(9118)))
		} else if t == "F" {
			fmt.Printf("%s %s °F. %s %s \n%s %s °F.", terrafact[0], formatNumber(conv.CelsiusToFarhenheit(56.7)), terrafact[1], formatNumber(conv.CelsiusToFarhenheit(-89.4)), terrafact[2], formatNumber(conv.CelsiusToFarhenheit(9118)))
		}
	}

	if funf == "sun" && isFlagPassed("funfacts") {
		sunfact := funfacts.GetFunFacts(funf)
		if t == "C" {
			fmt.Printf("%s %s °C.\n%s %s °C.", sunfact[0], formatNumber(15000000), sunfact[1], formatNumber(conv.KelvinToCelcius(5778)))
		} else if t == "K" {
			fmt.Printf("%s %s K.\n%s %s K.", sunfact[0], formatNumber(conv.CelsiusToKelvin(15000000)), sunfact[1], formatNumber(5778))
		} else if t == "F" {
			fmt.Printf("%s %s °F.\n%s %s °F.", sunfact[0], formatNumber(conv.CelsiusToFarhenheit(15000000)), sunfact[1], formatNumber(conv.CelsiusToFarhenheit(5778)))
		}
	}

	if funf == "luna" && isFlagPassed("funfacts") {
		lunafact := funfacts.GetFunFacts(funf)
		if t == "C" {
			fmt.Printf("%s %s °C.\n%s %s °C.", lunafact[0], formatNumber(-183), lunafact[1], formatNumber(106))
		} else if t == "K" {
			fmt.Printf("%s %s K.\n%s %s K.", lunafact[0], formatNumber(conv.CelsiusToKelvin(-183)), lunafact[1], formatNumber(conv.CelsiusToKelvin(106)))
		} else if t == "F" {
			fmt.Printf("%s %s °F.\n%s %s °F.", lunafact[0], formatNumber(conv.CelsiusToFarhenheit(-183)), lunafact[1], formatNumber(conv.CelsiusToFarhenheit(106)))
		}
	}
}

func formatNumber(num float64) string {
	str := strconv.FormatFloat(num, 'f', 2, 64)

	for str[len(str)-1] == '0' {
		str = str[:len(str)-1]
	}

	if math.Abs(num) >= 1000 {
		i := strings.Index(str, ".")
		for j := i - 3; j > 0; j -= 3 {
			str = str[:j] + " " + str[j:]
		}
	}

	return str
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
