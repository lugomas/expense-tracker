package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"roadmaps/projects/expense-tracker/cmd"
	"time"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add an expense to the list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "description",
						Aliases:  []string{"d"},
						Usage:    "description of the expense",
						Required: true,
					},
					&cli.Float64Flag{
						Name:     "amount",
						Aliases:  []string{"a"},
						Usage:    "amount of the expense",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {

					description := cCtx.String("description")
					amount := cCtx.Int("amount")
					expenseID, err := cmd.Add(description, amount)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Printf("Expense added successfully (ID: %s)", expenseID)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all expenses",
				Action: func(cCtx *cli.Context) error {
					expenses, err := cmd.List()
					if err != nil {
						fmt.Println(err)
					}
					cmd.PrintAll(expenses)
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "delete an expense from the list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Aliases:  []string{"i"},
						Usage:    "deletion of the expense",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					expenseID := cCtx.String("id")
					err := cmd.Delete(expenseID)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("deleted expense: ", expenseID)
					return nil
				},
			},
			{
				Name:    "summary",
				Aliases: []string{"s"},
				Usage:   "options for task templates",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "month",
						Aliases:  []string{"m"},
						Usage:    "description of the expense for the give month",
						Required: false,
					},
				},
				Action: func(cCtx *cli.Context) error {
					if cCtx.IsSet("month") {
						month := cCtx.Int("month")
						totalMonth, err := cmd.SummaryMonth(month)
						if err != nil {
							fmt.Println(err)
						}
						fmt.Printf("Total expenses for %s: $%d\n", time.Month(month), totalMonth)
					} else {
						total, err := cmd.Summary()
						if err != nil {
							fmt.Println(err)
						}
						fmt.Printf("Total expenses: $%d\n", total)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
