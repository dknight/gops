package main

import (
	"fmt"
	"os"
)

// List repsents lists.
type List struct {
	Done  int
	Total int
	Name  string
}

// DisplayLists shows all lists.
func DisplayLists(path string) {
	lists, err := AllLists(path)
	if err != nil {
		exitErr(err)
	}

	for _, list := range lists {
		fpath := fmt.Sprintf("%v%c%v", path, os.PathSeparator, list.Name())
		fp, err := os.Open(fpath)
		if err != nil {
			exitErr(err)
		}
		defer fp.Close()

		items, err := AllItems(fp)
		if err != nil {
			exitErr(err)
		}

		done := len(FilterItemsByStatus(items, itemStatusDone))
		total := len(items)
		li := List{
			Done:  done,
			Total: total,
			Name:  list.Name(),
		}
		fmt.Printf("%s\n", li)
	}
}

// AllLists gets all lists.
func AllLists(path string) ([]os.DirEntry, error) {
	lists, err := os.ReadDir(path)
	if err != nil {
		exitErr(err)
	}
	return lists, nil
}

func (li List) String() string {
	color := ""
	switch {
	case li.Done == 0:
		color = Color.Red
	case li.Done == li.Total:
		color = Color.Green
	default:
		color = Color.Yellow
	}
	return fmt.Sprintf("[%v%v/%v%v]\t%v",
		color, li.Done, li.Total, Color.Nul, li.Name)
}
