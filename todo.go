package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// lowercase struct casue we donot export it
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// list of to-do items -- upercase casue we want it to export
type List []item

// add create a new todo itme and append it to list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// complete method to mark an item as complete
// by setting done as true and completedat to current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

//delete metod to delete an item from list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exit", i)

	}
	// adjust the item index for 0 based index
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// save method encodes List as json and saves it
func (l *List) Save(fileName string) error {
	json, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, json, 0644)
}

// get method opens provided filename and
//deocdes the json data and parset it into a list
func (l *List) Get(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		// if the given file does not exist
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
