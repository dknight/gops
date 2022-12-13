package gops

import (
	"fmt"
	"io"
	"os"
)

// List represents lists.
type List struct {
	// Done counts completed items.
	Done int
	// Total counts total items.
	Total int
	// Name is the name of the list (file).
	Name string
}

// DisplayLists shows all lists.
func DisplayLists(rds []io.Reader) error {
	for _, rd := range rds {
		items, err := AllItems(rd)
		if err != nil {
			return err
		}

		done := len(FilterItemsByStatus(items, ItemStatusCompleted))
		total := len(items)
		var name string
		// More cases might be in future.
		switch rd.(type) {
		case *os.File:
			name = rd.(*os.File).Name()
		default:
			name = "buffer"
		}
		li := List{
			Done:  done,
			Total: total,
			Name:  name,
		}
		fmt.Println(li.BeautifulString())
	}
	return nil
}

// GetListsByPath gets all list by given path in file system.
func GetListsByPath(path string) ([]io.Reader, error) {
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
			path := MakeStoreFilePath(info.Name())
			file, err := os.Open(path)
			if err != nil {
				return nil, err
			}
			files = append(files, file)
		}
	}
	return files, nil
}

// BeautifulString prints list as the beautiful string.
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
