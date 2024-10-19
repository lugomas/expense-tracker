package cmd

import (
	"fmt"
	"roadmaps/projects/expense-tracker/internal"
	"roadmaps/projects/expense-tracker/pkg/fileutil"
	"time"
)

func List() ([]internal.Expense, error) {
	err := fileutil.CheckFileExists()
	if err != nil {
		return nil, err
	}
	return fileutil.ReadFile()
}

func PrintAll(expenses []internal.Expense) {
	for i, expense := range expenses {
		fmt.Printf("%d. ID: %s\n", i+1, expense.ID)
		fmt.Printf("   Description: %s\n", expense.Description)
		fmt.Printf("   Amount: %d\n", expense.Amount)
		fmt.Printf("   Created At: %s\n", expense.CreatedAt.Format(time.DateTime)) // Format time
		fmt.Println("   --------------------")
	}
}
