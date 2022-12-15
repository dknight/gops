// Check usage by command:
//
//	gops -h
//
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/dknight/gops"
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
		" (default "+gops.GetDefaultStoreFilePath()+")")
	today := flag.Bool("t", false, "Set list to today's date.")
	list := flag.Bool("l", false, "Display todo-lists.")
	ver := flag.Bool("v", false, "Displays the version")
	undone := flag.Bool("u", false, "Display only incomplete items.")
	sort := flag.Bool("s", false, "Sort items by date and status.")
	flag.Parse()

	if *ver {
		fmt.Println(gops.Version())
		exitSucces()
	}

	if *fname != "" {
		gops.SetStoreFileName(*fname)
	}

	if *today {
		t := time.Now()
		gops.SetStoreFileName(t.Format("2006-01-02"))
	}

	fp, err := gops.GetStoreFile()
	if *creat != "" {
		if err != nil {
			fp, err = gops.CreateStoreFile()
			if err != nil {
				exitErr(err)
			}
		}
		item := gops.NewItem(time.Now(), gops.ItemStatusIncompleted, *creat)
		err := item.Save(fp)
		if err != nil {
			exitErr(err)
		}
		exitSucces(addedMsg)
	}

	if err != nil {
		exitErr(err)
	}
	items, err := gops.AllItems(fp)
	if err != nil {
		exitErr(err)
	}

	if *compl != 0 {
		if *compl > uint(len(items)) {
			exitErr(errors.New(itemNotExists))
		}
		completed, err := gops.CompleteItem(*compl, items, fp)
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

	if *sort {
		var ans rune
		fmt.Printf("Items will be sorted in %s."+
			" You cannot revert thisoperation. [y/N]", fp.Name())
		_, err := fmt.Fscanf(os.Stdin, "%c\n", &ans)
		if err != nil {
			exitErr(err)
		}
		if ans == 'y' || ans == 'Y' {
			gops.SortItems(items, fp)
		} else {
			exitSucces()
		}
	}

	if *list {
		lists, err := gops.GetListsByPath(gops.GetConfigPath())
		if err != nil {
			exitErr(err)
		}
		err = gops.DisplayLists(lists)
		if err != nil {
			exitErr(err)
		}
		exitSucces()
	}

	if *undone {
		items = gops.FilterItemsByStatus(items, gops.ItemStatusIncompleted)
	}
	for i, item := range items {
		fmt.Println(item.BeautifulString(i + 1))
	}
}

func exitErr(errs ...error) {
	for _, e := range errs {
		fmt.Println(e.Error())
	}
	os.Exit(1)
}

func exitSucces(msgs ...string) {
	for _, m := range msgs {
		if m != "" {
			fmt.Println(m)
		}
	}
	os.Exit(0)
}
