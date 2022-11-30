// Package gops is the most simple terminal todo utility in the World.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

var file *os.File

const (
	addedMsg      = "Task has been added."
	completedMsg  = "Task \"%v\" is %scompleted."
	itemNotExists = "Item no exist."
)

func main() {
	creat := flag.String("n", "", "Set name of the new todo task.")
	compl := flag.Uint("c", 0, "Number of the task to complete.")
	fname := flag.String("f", "", "File of stored todo items."+
		" (default $HOME/.config/.gops)")
	all := flag.Bool("a", false, "Display also done items.")
	flag.Parse()

	if *fname != "" {
		storeFileName = *fname
	}

	file, err := storeFileResolver()
	if err != nil {
		exitErr(err)
	}
	items, err := ListItems(file)
	if err != nil {
		exitErr(err)
	}

	if *creat != "" {
		item := NewItem(time.Now(), itemStatusTodo, *creat)
		err := item.Save(file)
		if err != nil {
			exitErr(err)
		}
		exitSucces(addedMsg)
	}

	if *compl != 0 {
		if *compl > uint(len(items)) {
			exitErr(errors.New(itemNotExists))
		}
		completed, err := CompleteItem(*compl, items, file)
		if err != nil {
			exitErr(err)
		}
		prefix := ""
		if !completed.Status {
			prefix = "in" // incompleted
		}
		okMsg := fmt.Sprintf(completedMsg, completed.Task, prefix)
		exitSucces(okMsg)
	}

	if len(items) == 0 {
		exitSucces("There is no incomplete items, relax it is good for you.")
	}

	if !*all {
		items = FilterItemsByStatus(items, false)
	}
	for i, item := range items {
		fmt.Println(item.BeautifulString(i + 1))
	}
}

func exitErr(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func exitSucces(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
