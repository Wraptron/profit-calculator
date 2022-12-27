package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	revenue, err := f.GetCellValue("Sheet1", "A2")
	if err != nil {
		fmt.Println(err)
		return
	}
	expense, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}

	i, err := strconv.ParseInt(revenue, 10, 64)
	if err != nil {
		panic(err)
	}
	j, err := strconv.ParseInt(expense, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Revenue :		", i)
	fmt.Println("Expense:		 ", j)
	var profit = i - j
	fmt.Println("Profit:			 ", profit)
	var tax = profit * 50 / 100
	fmt.Println("Tax (50%):		 ", tax)
	var profitaftertax = profit - tax
	fmt.Println("Profit after tax:	 ", profitaftertax)
}
