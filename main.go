// Package gops is the most simple terminal todo utility in the World.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

const version = "1.0.4-rc1"

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
		" (default "+getDefaultStoreFile()+")")
	all := flag.Bool("a", false, "Display also done items.")
	today := flag.Bool("t", false, "Set list to today's date.")
	list := flag.Bool("l", false, "Display todo-lists.")
	ver := flag.Bool("v", false, "Displays the version")
	flag.Parse()

	if *ver {
		fmt.Println(version)
		exitSucces("")
	}

	if *fname != "" {
		storeFileName = *fname
	}

	if *today {
		t := time.Now()
		storeFileName = t.Format("2006-01-02")
	}

	file, err := storeFileResolver()
	if err != nil {
		exitErr(err)
	}
	items, err := AllItems(file)
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

	if *list {
		DisplayLists(getConfigPath())
		exitSucces("")
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
	if msg != "" {
		fmt.Println(msg)
	}
	os.Exit(0)
}
