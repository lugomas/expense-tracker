package cmd

import (
	"fmt"
	"roadmaps/projects/expense-tracker/pkg/fileutil"
)

func Delete(expenseId string) error {
	expenses, err := fileutil.ReadFile()
	if err != nil {
		return err
	}
	expenseFound := false
	for i := range expenses {
		if expenses[i].ID == expenseId {
			expenseFound = true
			expenses = append(expenses[:i], expenses[i+1:]...)
			break
		}
	}
	if !expenseFound {
		return fmt.Errorf("expense with id %s not found", expenseId)
	}
	err = fileutil.CreateFile(expenses)
	if err != nil {
		return err
	}
	return nil
}
