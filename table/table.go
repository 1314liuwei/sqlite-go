package table

import (
	"1314liuwei/sqlite.go/consts"
	"errors"
	"fmt"
	"sync"
)

type UserTable struct {
	counter   uint
	RowLength int
	Rows      []UserTableRow
}

type UserTableRow struct {
	ID       uint
	Username string
	Email    string
}

var (
	userTable *UserTable
	once      sync.Once
)

func (t *UserTable) ExecuteInsert(row UserTableRow) error {
	if t.RowLength+1 > consts.TableMaxRows {
		return errors.New("the number of current table storage rows has reached the maximum")
	}

	if row.ID != 0 {
		for i := 0; i < t.RowLength; i++ {
			if row.ID == t.Rows[i].ID {
				return errors.New("ID value cannot be repeated")
			}
		}
	} else {
		row.ID = t.counter
	}

	t.counter++
	t.Rows = append(t.Rows, row)
	t.RowLength++
	return nil
}

func (t *UserTable) ExecuteSelect() error {
	fmt.Println("ID\t|Username\t|Email")
	for i := 0; i < t.RowLength; i++ {
		row := t.Rows[i]
		fmt.Printf("%d\t|%s\t|%s\n", row.ID, row.Username, row.Email)
	}
	return nil
}

func Table() *UserTable {
	once.Do(func() {
		userTable = &UserTable{
			counter: 1,
		}
	})
	return userTable
}
