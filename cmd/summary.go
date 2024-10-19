package cmd

import "roadmaps/projects/expense-tracker/pkg/fileutil"

func Summary() (int, error) {
	err := fileutil.CheckFileExists()
	if err != nil {
		return 0, err
	}

	expenses, err := fileutil.ReadFile()
	if err != nil {
		return 0, err
	}

	var total int
	for _, expense := range expenses {
		total += expense.Amount
	}

	return total, nil
}

func SummaryMonth(month int) (int, error) {
	err := fileutil.CheckFileExists()
	if err != nil {
		return 0, err
	}

	expenses, err := fileutil.ReadFile()
	if err != nil {
		return 0, err
	}

	var total int
	for _, expense := range expenses {
		if month != int(expense.CreatedAt.Month()) {
			continue
		}
		total += expense.Amount
	}
	return total, nil
}
