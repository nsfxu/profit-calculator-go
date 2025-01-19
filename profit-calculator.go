package main

import "fmt"

func main() {
	// Ask for revenue, expenses & tax rate
	// Calculate earnings before tax(EBT) and earnings after tax (profit)
	// Calculate ratio (EBT / profit)
	// Output EBT, profit and the ratio
	var revenue, expenses, taxRate, ebt, profit, ratio float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	// EBT = Revenue - Expenses
	ebt = revenue - expenses

	// EAT = EBT - (EBT * Tax Rate)
	profit = ebt - (ebt * taxRate)

	// Ratio = (EBT / profit)
	ratio = ebt / profit

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
