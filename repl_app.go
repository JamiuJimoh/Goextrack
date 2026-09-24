package repl

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"uuid"

	expense "github.com/JamiuJimoh/Goextrack/internal"
)

const title = `
╔═══════════════════════════════════════════════════════════════════════════╗
║                                                                           ║
║                                                                           ║
║ ██████╗  ██████╗ ███████╗██╗  ██╗████████╗██████╗  █████╗  ██████╗██╗  ██╗║
║██═════╝ ██╔═══██╗██╔════╝╚██╗██╔╝╚══██╔══╝██╔══██╗██╔══██╗██╔════╝██╗ ██╔╝║
║██║  ███╗██║   ██║█████╗   ╚███╔╝    ██║   ██████╔╝███████║██║     █████╔╝ ║
║██║   ██║██║   ██║██╔══╝   ██╔██╗    ██║   ██╔══██╗██╔══██║██║     ██╔═██╗ ║
║╚██████╔╝╚██████╔╝███████╗██╔╝ ██╗   ██║   ██║  ██║██║  ██║╚██████╗██║  ██╗║
║ ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝║
║                                                                           ║
║                           Your Expense Tracker                            ║
║                                                                           ║
╚═══════════════════════════════════════════════════════════════════════════╝`

type REPL struct {
	t *expense.Tracker
	s *bufio.Scanner
}

func New(t *expense.Tracker) *REPL {
	return &REPL{t: t,
		s: bufio.NewScanner(os.Stdin),
	}
}

func (r *REPL) Run() error {
	fmt.Println(title)
	showMenu()

scanner:
	for r.s.Scan() {
		option, err := strconv.Atoi(r.s.Text())
		if err != nil || option > 6 || option < 1 {
			fmt.Println("Invalid option. Please select a valid option")
			showMenu()
			continue
		}
		switch option {
		case 1:
			r.add()
			fmt.Print("\nSelect an option: ")
		case 2:
			r.list()
			fmt.Print("\nSelect an option: ")
		case 3:
			r.view()
			fmt.Print("\nSelect an option: ")
		case 4:
			r.delete()
			fmt.Print("\nSelect an option: ")
		case 5:
			showMenu()
		case 6:
			break scanner
		default:
			showMenu()
			continue
		}

	}
	return r.s.Err()
}

func (r *REPL) add() {
	buffer := make([]expense.Expense, 0, 5)

	buffer = append(buffer, newExpense(r.s))
	fmt.Printf("You have added %d new entries. Do you want to;\n", len(buffer))
	fmt.Printf("1. Add more\n2. Save\n")

	for r.s.Scan() {
		option, err := strconv.Atoi(r.s.Text())
		if err != nil || option < 1 || option > 2 {
			fmt.Println("Invalid option")
			continue
		}
		switch option {
		case 1:
			buffer = append(buffer, newExpense(r.s))
			fmt.Printf("You have added %d new entries. Do you want to;\n", len(buffer))
			fmt.Printf("1. Add more\n2. Save\n")
		case 2:
			r.t.AddAll(buffer)
			err := r.t.Save()
			if err != nil {
				fmt.Println("Internal error occured while saving data")
			}
			fmt.Printf("You saved %d item(s)\n", len(buffer))
			return
		default:
			continue
		}
	}
}

func (r *REPL) list() {
	data := r.t.List()
	if len(data) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}

	fmt.Println("\n╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         YOUR EXPENSES                            ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "ID\tCATEGORY\tAMOUNT\tDESCRIPTION\tDATE")

	fmt.Fprintln(w, "──\t────────\t──────\t───────────\t────")

	for _, expense := range data {
		fmt.Fprintf(
			w,
			"%s\t%s\t₦%.2f\t%s\t%v\n",
			expense.ID,
			expense.Category,
			expense.Amount,
			expense.Description,
			expense.CreatedAt.Format("2006-01-02"),
		)
	}

	w.Flush()

	summary := r.t.Summarize()
	fmt.Printf("\nExpenses Count: %d\n", summary.Count)
	fmt.Printf("Total Expenses: ₦%.2f\n", float64(summary.Total))
	fmt.Printf("Expenses Average: %.2f\n", summary.Average)
}

func (r *REPL) view() {
	fmt.Print("Enter the expense id: ")
	for r.s.Scan() {
		id, err := uuid.Parse(r.s.Text())
		if err != nil {
			fmt.Println("Invalid id")
			fmt.Print("Enter the expense id: ")
			continue
		}
		e, err := r.t.Get(id)
		if err != nil {
			if errors.Is(err, expense.ErrExpenseNotFound) {
				fmt.Println(expense.ErrExpenseNotFound)
				fmt.Print("Enter the expense id: ")
				continue
			}
			fmt.Println("Internal error")
			return
		}

		fmt.Println()
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		fmt.Fprintln(w, "ID\tCATEGORY\tAMOUNT\tDESCRIPTION\tCREATEDAT\tUPDATEDAT")

		fmt.Fprintln(w, "──\t────────\t──────\t───────────\t─────────\t─────────")

		timeFormat := "Mon January 02, 2006 by 15:04:05"
		fmt.Fprintf(
			w,
			"%s\t%s\t₦%.2f\t%s\t%v\t%v\n",
			e.ID,
			e.Category,
			e.Amount,
			e.Description,
			e.CreatedAt.Format(timeFormat),
			e.UpdatedAt.Format(timeFormat),
		)

		w.Flush()
		return
	}
}

func (r *REPL) delete() {
	fmt.Print("Enter the expense id: ")
	for r.s.Scan() {
		id, err := uuid.Parse(r.s.Text())
		if err != nil {
			fmt.Println("Invalid id")
			fmt.Print("Enter the expense id: ")
			continue
		}
		err = r.t.Delete(id)
		if err != nil {
			if errors.Is(err, expense.ErrExpenseNotFound) {
				fmt.Println(expense.ErrExpenseNotFound)
				fmt.Print("Enter the expense id: ")
				continue
			}
			fmt.Println("Internal error")
			return
		}

		err = r.t.Save()
		if err != nil {
			fmt.Println("Internal error")
			continue
		}
		fmt.Println("Deleted one entry")
		return
	}
}

func showMenu() {
	menu := `
╔══════════════════════════════════════════════════╗
║                  MAIN MENU                       ║
╠══════════════════════════════════════════════════╣
║                                                  ║
║  1. Add Expense                                  ║
║  2. List Expenses                                ║
║  3. View Expense                                 ║
║  4. Delete Expense                               ║
║  5. Help                                         ║
║  6. Exit                                         ║
║                                                  ║
╚══════════════════════════════════════════════════╝

Select an option: `

	fmt.Print(menu)
}

func amountCollector(scanner *bufio.Scanner) float64 {
	fmt.Print("Enter amount: ")
	for scanner.Scan() {
		amount, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil || amount <= 0 {
			fmt.Println("Invalid amount. Please enter a valid amount")
			fmt.Print("\nEnter amount: ")
			continue
		}
		return amount
	}
	return 0
}

func showCategoriesMenu(categories []expense.Category, scanner *bufio.Scanner) {
	fmt.Println()
	fmt.Println("-------------------------------------")
	fmt.Println("Categories")
	fmt.Println("-------------------------------------")
	for i, category := range categories {
		fmt.Printf("%d. %s\n", i+1, category)
	}
	fmt.Print("\nEnter category: ")
}
func categoryCollector(scanner *bufio.Scanner) expense.Category {
	categories := expense.AllCategories()
	showCategoriesMenu(categories, scanner)

	for scanner.Scan() {
		option, err := strconv.Atoi(scanner.Text())
		position := option - 1
		if err != nil || position < 0 || position > len(categories)-1 {
			fmt.Println("Invalid category. Please select a valid category")
			showCategoriesMenu(categories, scanner)
			continue
		}
		return categories[position]
	}
	return expense.Category("")
}

func descriptionCollector(scanner *bufio.Scanner) string {
	fmt.Print("\nEnter description (optional): ")
	scanner.Scan()
	return scanner.Text()
}

func newExpense(scanner *bufio.Scanner) expense.Expense {
	amount := amountCollector(scanner)
	category := categoryCollector(scanner)
	description := descriptionCollector(scanner)
	return expense.NewExpense(amount, category, description)
}
