package fileutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"roadmaps/projects/expense-tracker/internal"
)

const expensesFile = "json.txt"

func CheckFileExists() error {
	_, err := os.Stat(expensesFile)
	if os.IsNotExist(err) {
		return errors.New("No expenses found\nUse 'expense-tracker add' to add a new expense\n")
	}
	return nil
}

func ReadFile() ([]internal.Expense, error) {
	fileData, err := os.ReadFile(expensesFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	var expenses []internal.Expense
	err = json.Unmarshal(fileData, &expenses)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling data: %w", err)
	}
	return expenses, nil
}

func CreateFile(expenses []internal.Expense) error {
	newData, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}
	err = os.WriteFile(expensesFile, newData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to json file: %w", err)
	}
	return nil
}
