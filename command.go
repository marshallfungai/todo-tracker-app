package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo, specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new tile. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Edit != "":
		edit := strings.SplitN(cf.Edit, ":", 2)
		if len(edit) != 2 {
			fmt.Println("Error, invalid format for edit. Use index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(edit[0])
		if err != nil {
			fmt.Println("Error, invalid format for edit. index should be an integer")
			os.Exit(1)
		}
		todos.edit(index, edit[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.List:
		todos.print()
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}