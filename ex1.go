package main

import (
	"fmt"
	"strings"
)

func checkAccess(licensePlate string) string {
	allowedLicensePlates := []string{"ABC123", "XYZ789"} // hardcoded list

	for _, plate := range allowedLicensePlates {
		if strings.ToUpper(licensePlate) == plate {
			return "Welcome to Fonteyn Vakantieparken"
		}
	}

	return "You unfortunately do not have access to the parking lot"
}

func main() {
	var licensePlate string
	fmt.Print("Enter your license plate: ")
	fmt.Scan(&licensePlate)

	result := checkAccess(licensePlate)
	fmt.Println(result)
}
