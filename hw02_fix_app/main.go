package main

import (
	"fmt"

	"github.com/IlyaPol99/go_basic/hw02_fix_app/printer"
	"github.com/IlyaPol99/go_basic/hw02_fix_app/reader"
	"github.com/IlyaPol99/go_basic/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
