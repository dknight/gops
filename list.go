package main

import (
	"fmt"
	"io"
	"os"
)

// List repsents lists.
type List struct {
	Done  int
	Total int
	Name  string
}

// DisplayLists shows all lists.
func DisplayLists(rds []io.Reader) {
	for _, rd := range rds {
		items, err := AllItems(rd)
		if err != nil {
			exitErr(err)
		}

		done := len(FilterItemsByStatus(items, itemStatusDone))
		total := len(items)
		name := "buffer" // TODO rename?
		switch rd.(type) {
		case *os.File:
			name = rd.(*os.File).Name()
		}
		li := List{
			Done:  done,
			Total: total,
			Name:  name,
		}
		fmt.Println(li.BeautifulString())
	}
}

func getListsByPath(path string) ([]io.Reader, error) {
	files := make([]io.Reader, 0)
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		typ := entry.Type()
		if typ.IsRegular() {
			info, err := entry.Info()
			if err != nil {
				return nil, err
			}
			file, err := os.Open(info.Name())
			if err != nil {
				return nil, err
			}
			files = append(files, file)
		}
	}
	return files, nil
}

func (li List) BeautifulString() string {
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
