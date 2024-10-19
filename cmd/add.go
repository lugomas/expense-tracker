package cmd

import (
	"roadmaps/projects/expense-tracker/internal"
	"roadmaps/projects/expense-tracker/pkg/fileutil"
	"time"
)

func Add(description string, amount int) (string, error) {

	newExpense := internal.Expense{
		ID:          time.Now().Format("200601021504105"),
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	var expenses []internal.Expense
	err := fileutil.CheckFileExists()
	if err != nil {
		expenses = append(expenses, newExpense)
		err := fileutil.CreateFile(expenses)
		if err != nil {
			return "", err
		}
		return newExpense.ID, nil
	}

	expenses, err = fileutil.ReadFile()
	if err != nil {
		return "", err
	}
	expenses = append(expenses, newExpense)
	err = fileutil.CreateFile(expenses)
	if err != nil {
		return "", err
	}
	return newExpense.ID, nil
}
