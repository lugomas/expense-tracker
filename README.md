# Expense-Tracker CLI
`expense-tracker` is a command-line interface (CLI) tool designed to help users manage their personal expenses. This tool allows users to add, update, delete, and view expenses, while also providing summary features. Additionally, users can set budgets, categorize expenses, and export data to CSV format.

## Features
- Add new expenses with a description and amount.
- Delete expenses.
- View all expenses.
- View a summary of total expenses.
- View a summary of expenses for a specific month of the current year

## Installation
To install expense-tracker, clone the repository and build the Go binary:
```
git clone https://github.com/lugomas/expense-tracker.git
cd expense-tracker
go build -o expense-tracker
```

You can now use the `expense-tracker` binary to manage your tasks.
Note: The Expense ID is dynamically generated based on the current time, with a precision of seconds.

## Usage
```
$ ./expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ ./expense-tracker add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

$ ./expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

$ ./expense-tracker summary
# Total expenses: $30

$ ./expense-tracker delete --id 1
# Expense deleted successfully

$ ./expense-tracker summary
# Total expenses: $20

$ ./expense-tracker summary --month 8
# Total expenses for August: $20
```

## License
This project is licensed under the MIT License.

## Project Inspiration
This project was developed based on the guidelines provided by [roadmap.sh's Expense Tracker project](https://roadmap.sh/projects/expense-tracker)
