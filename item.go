package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	itemStatusDone = true
	itemStatusTodo = false
	itemMarkDone   = 10003 // checkmark
	itemMarkTodo   = 32    // space
)

const dateFormat = time.RFC3339

type Item struct {
	Time   time.Time
	Status bool
	Task   string
}

func NewItem(tm time.Time, s bool, t string) *Item {
	return &Item{
		Time:   tm,
		Status: s,
		Task:   t,
	}
}

func ListItems(rd io.Reader) ([]Item, error) {
	items := make([]Item, 0, 10)
	bs, err := os.ReadFile(rd.(*os.File).Name())
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(bytes.NewReader(bs))
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		tim, err := time.Parse(dateFormat, rec[0])
		if err != nil {
			return nil, err
		}
		status, err := strconv.ParseBool(rec[1])
		if err != nil {
			return nil, err
		}
		items = append(items, Item{
			Time:   tim,
			Status: status,
			Task:   rec[2],
		})
	}
	return items, nil
}

func FilterItemsByStatus(items []Item, status bool) []Item {
	out := make([]Item, 0)
	for _, item := range items {
		if item.Status == status {
			out = append(out, item)
		}
	}
	return out
}

func CompleteItem(i uint, items []Item, wr io.Writer) (*Item, error) {
	fp, err := os.Create(wr.(*os.File).Name())
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	items[i-1].Complete()
	for _, item := range items {
		item.Save(fp)
	}
	return &items[i-1], nil
}

func (item *Item) Complete() {
	item.Status = !item.Status
}

func (item *Item) Save(wr io.Writer) error {
	w := csv.NewWriter(wr)
	slice := item.Slice()
	if err := w.Write(slice); err != nil {
		return err
	}
	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func (item *Item) Slice() []string {
	s := make([]string, 3)
	s[0] = item.Time.Format(dateFormat)
	s[1] = strconv.FormatBool(item.Status)
	s[2] = item.Task
	return s
}

func (item *Item) BeautifulString(i int) string {
	var mark []rune
	if item.Status {
		mark = []rune{itemMarkDone}
	} else {
		mark = []rune{itemMarkTodo}
	}

	return fmt.Sprintf("  %v#%v %d %v%s%v %v",
		Color.Blue, Color.Nul,
		i,
		Color.Green, string(mark), Color.Nul,
		item.Task)
}
