package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Address struct {
	postcode int
	city string
	street string
	flatNum int
}

type Waybill struct {
	name string
	amount int
	customerName string
	contactPhone string
	Address
}

func (w Waybill) enterDataToWaybill() Waybill {
	waybill := Waybill{}
	waybill.name = readInput("Enter a product name:")
	productAmount, err := strconv.Atoi(readInput("Enter product amount:"))
	if err != nil {
		log.Fatalf("An error has occured while transform string to digit: %s", err)
	}
	waybill.amount = productAmount
	waybill.customerName = readInput("Enter a customer name:")
	waybill.contactPhone = readInput("Enter a contact phone:")
	return waybill
}

func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("The error has occured: %s", err)
	}
	return input
}

func main() {
	waybill := Waybill{}
	filledWaybill := waybill.enterDataToWaybill()
	fmt.Println(filledWaybill)
}
