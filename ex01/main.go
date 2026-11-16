package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	firstName := "Abeh"
	lastName := "Abeh"
	dayOfBirth := 22
	monthOfBirth := 06
	yearOfBirth := 2005
	numberOfSiblings := 4
	heightInMeters := 1.65
	zodiacSign := 'T'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
