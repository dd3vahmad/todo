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

func NewCommandFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by specifying it index and the new title")
	flag.IntVar(&cf.Del, "del", -1, "Specify todo to delete by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify the index of the todo to toggle it status")
	flag.BoolVar(&cf.List, "list", false, "List our todos in a tabular form")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) execute(todos *Todos) {
	switch {
	case cf.List:
		todos.list()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error invalid argument format, use id:title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error invalid index")
			os.Exit(1)
		}

		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	default:
		fmt.Println("Invalid command")
	}
}
