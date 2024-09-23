package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Luhn Algorithm...")
	server := gin.Default()

	server.GET("/:CreditCardNumber", ValidatorHandler)
	server.Run()
}

func ValidatorHandler(c *gin.Context) {
	CreditCardNumber := c.Param("CreditCardNumber")
	log.Print("Card Number is:", CreditCardNumber)
	CardNumber, err := strconv.Atoi(CreditCardNumber[0 : len(CreditCardNumber)-1])
	if err != nil {
		log.Panic("Error occured: ", err)
	}
	checkDigit, err := strconv.Atoi(CreditCardNumber[len(CreditCardNumber)-1:])
	if err != nil {
		log.Panic("Error occured: ", err)
	}

	check := LuhnAlgorithm(CardNumber, checkDigit)
	// fmt.Print(check)
	var message string
	if check {
		message = "Valid Credit Card Number"
	} else {
		message = "Invalid Credit Card Number"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})

}

func LuhnAlgorithm(CardNumber int, checkDigit int) bool {
	var number []int
	for i := 0; i < 16; i++ {
		rightMostNumber := CardNumber % 10
		CardNumber = int(CardNumber / 10)
		number = append(number, rightMostNumber)

	}
	fmt.Print(number)
	counter := 2
	for index, val := range number {
		if counter%2 == 0 {
			res := val * 2
			if res > 9 {
				res = res - 9
			}
			number[index] = res
		}
		counter += 1
	}
	fmt.Print(number)

	Summation := 0

	for _, val := range number {
		Summation += val
	}

	Summation += checkDigit
	fmt.Print("Summation: ", Summation)
	return Summation%10 == 0

}
