package terminal_test

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var scn = bufio.NewScanner(os.Stdin)

func getInputString() string {
	scn.Scan()
	return strings.TrimSpace(scn.Text())
}

func getInputInt() (int, error) {
	return strconv.Atoi(getInputString())
}

func getInputFloat() (float64, error) {
	return strconv.ParseFloat(getInputString(), 64)
}

func invalidInput() {
	fmt.Println("Invalid input!")
}

func y_or_n() bool {
	fmt.Print("Are you sure? (y or n): ")
	for {
		switch getInputString() {
		case "y", "Y":
			return true
		case "n", "N":
			return false
		default:
			fmt.Print("Please enter y or n: ")
		}
	}
}

func ShowMenu(title string, options []MenuOption) error {
	for {
		for _, opt := range options {
			fmt.Printf("|%d| %s\n", opt.ID, opt.Label)
		}
		fmt.Printf("%s (-1 to go back): ", title)

		val, err := getInputInt()
		if err != nil {
			fmt.Println(err)
		}

		if val == -1 {
			fmt.Println()
			return nil
		}

		found := false
		for _, opt := range options {
			if opt.ID == val {
				found = true
				if err := opt.Action(); err != nil {
					fmt.Println(err)
				}
				break
			}
		}

		if !found {
			invalidInput()
		}
	}
}

func AskQ(max int) int {
	fmt.Print("(-1 to leave without saving)\n|0| save and leave\n|1| see status\nPlease select an option and click \"Enter\": ")
	for {
		val, err := getInputInt()
		if err != nil {
			invalidInput()
			continue
		}

		if val <= max || val >= -1 {
			return val
		} else {
			fmt.Printf("\nPlease choose a number between [-1, %d]: ", max)
			continue
		}
	}
}

func leaveSafe() {

}

func getDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./pos.db?_parse_time=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func seeRestaurants() error {
	db, err := getDB()
	if err != nil {
		return err
	}

	query := `
		SELECT
			r.id, r.name, r.capacity, r.state, COUNT(t.id)
		FROM restaurants r
		LEFT JOIN tables t ON r.id = t.restaurant_id 
		GROUP BY r.id, r.name
		ORDER BY r.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("===========================================================")
	fmt.Printf("%-4s: %-20s | %-8s | %-6s | %-10s\n", "ID", "NAME", "CAPACITY", "TABLES", "STATE")
	fmt.Println("-----------------------------------------------------------")
	for rows.Next() {
		var id, capacity, tCount int
		var name string
		var state bool

		if err := rows.Scan(&id, &name, &capacity, &state, &tCount); err != nil {
			return err
		}

		stateText := "FREE"
		if state {
			stateText = "OCCUPIED"
		}

		fmt.Printf(": %-20s | %-8d | %-6d | %-10s", name, capacity, tCount, stateText)
	}
	fmt.Println("===========================================================")
	return nil
}
func seeOrderOf(id int) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	query := `
		SELECT 
			o.id, t.number, m.name, o.quantity, o.order_time
		FROM orders o
		JOIN tables t ON t.id = o.table_id
		JOIN menus m ON m.id = o.menu_item_id
		WHERE t.restaurant_id = ? AND m.restaurant_id = ?
	`

	rows, err := db.Query(query, id, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("==============================================================")
	fmt.Printf("%-4s: %-6s | %-20s | %-5s | %-10s\n", "ID", "TABLE", "ITEM", "PIECE", "ORDER")
	fmt.Printf("%-4s: %-6s | %-20s | %-5s | %-10s\n", " ", "NUMBER", "NAME", "     ", "TIME ")
	fmt.Println("--------------------------------------------------------------")

	for rows.Next() {
		var id, tNumber, piece int
		var iName string
		var oTime time.Time

		if err := rows.Scan(&id, &tNumber, &iName, &piece, &oTime); err != nil {
			return err
		}

		fmt.Printf("%-4d: %-6d | %-20s | %-5d | %-10s\n", id, tNumber, iName, piece, oTime.Format("02/01/2006 15:04"))

	}
	fmt.Println("==============================================================")
	return nil
}
func seeOrders() error {
	options := []MenuOption{}

	db, err := getDB()
	if err != nil {
		return err
	}

	query := `
		SELECT
			r.id, r.name, r.capacity, r.state, COUNT(t.id)
		FROM restaurants r
		LEFT JOIN tables t ON r.id = t.restaurant_id 
		GROUP BY r.id, r.name
		ORDER BY r.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("===========================================================")
	fmt.Printf("%-4s: %-20s | %-8s | %-6s | %-10s\n", "ID", "NAME", "CAPACITY", "TABLES", "STATE")
	fmt.Println("-----------------------------------------------------------")
	for rows.Next() {
		var id, capacity, tCount int
		var name string
		var state bool

		if err := rows.Scan(&id, &name, &capacity, &state, &tCount); err != nil {
			return err
		}

		stateText := "FREE"
		if state {
			stateText = "OCCUPIED"
		}

		text := fmt.Sprintf(": %-20s | %-8d | %-6d | %-10s", name, capacity, tCount, stateText)
		newOption := MenuOption{
			ID:    id,
			Label: text,
			Action: func() error {
				return seeOrderOf(id)
			},
		}
		options = append(options, newOption)
	}
	ShowMenu("Which restaurant's orders should be displayed?", options)
	return nil
}
func seeMenus() error {
	options := []MenuOption{}

	db, err := getDB()
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT id, name, capacity, state FROM restaurants ORDER BY id")
	if err != nil {
		return err
	}

	fmt.Println("==================================================")
	fmt.Printf("%-4s: %-20s | %-8s | %-10s\n", "ID", "NAME", "CAPACITY", "STATE")
	fmt.Println("--------------------------------------------------")
	for rows.Next() {
		var id, capacity int
		var name string
		var state bool

		if err := rows.Scan(&id, &name, &capacity, &state); err != nil {
			return err
		}

		stateText := "FREE"
		if state {
			stateText = "OCCUPIED"
		}

		text := fmt.Sprintf(": %-20s | %-8d | %-10s", name, capacity, stateText)
		newOption := MenuOption{
			ID:    id,
			Label: text,
			Action: func() error {
				return seeOrderOf(id)
			},
		}
		options = append(options, newOption)
	}
	ShowMenu("Which restaurant's orders should be displayed?", options)

	rows.Close()
	return nil
}
func seeTempMemory() error {
	options := []MenuOption{
		{1, "restaurants", seeRestaurants},
		{2, "orders", seeOrders},
		{3, "menus", seeMenus},
	}

	ShowMenu("Which one would you like to view information about?", options)

	return nil
}
func seeDB() error {
	db, err := getDB()
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT id, name, capacity, state FROM restaurants ORDER BY id")
	if err != nil {
		return err
	}
	fmt.Println(rows)

	return nil
}

func seeDetails() error {
	options := []MenuOption{
		{1, "temporary memory", seeTempMemory},
		{2, "database", seeDB},
	}

	ShowMenu("Which records would you like to view?", options)

	return nil
}
