package backend

import (
	"1314liuwei/sqlite.go/consts"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"unsafe"
)

type UserTable struct {
	counter   uint
	rowSize   int
	RowLength int
	Rows      []UserTableRow
}

type UserTableRow struct {
	ID       uint
	Username string
	Email    string
}

type UserTableStatement struct {
	Type consts.StatementType
	Row  UserTableRow
}

type Database struct {
	conn   *os.File
	table  *UserTable
	cursor *UserTableRow
}

func Open(name string) (*Database, error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	db := &Database{
		conn: file,
		table: &UserTable{
			counter: 1,
			rowSize: GetTableSizeof(UserTableRow{}),
		},
	}

	err = db.GetDBDataFromFile()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db *Database) ExecuteInsert(row UserTableRow) error {
	table := db.table
	if table.RowLength+1 > consts.TableMaxRows {
		return errors.New("the number of current database storage rows has reached the maximum")
	}

	if row.ID != 0 {
		for i := 0; i < table.RowLength; i++ {
			if row.ID == table.Rows[i].ID {
				return errors.New("ID value cannot be repeated")
			}
		}
	} else {
		row.ID = table.counter
	}

	table.counter++
	table.Rows = append(table.Rows, row)
	table.RowLength++

	err := db.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) ExecuteSelect() error {
	table := db.table

	fmt.Println("ID\t|Username\t|Email")
	for i := 0; i < table.RowLength; i++ {
		row := table.Rows[i]
		db.cursor = &row
		fmt.Printf("%d\t|%s\t|%s\n", row.ID, row.Username, row.Email)
	}
	return nil
}

func (db *Database) Flush() error {
	var (
		offset = 0
	)

	row := db.table.Rows[db.table.RowLength-1]
	data := make([]byte, GetTableSizeof(db.table.Rows[0]))
	*(*uint)(unsafe.Pointer(&data[0])) = row.ID
	offset += int(unsafe.Sizeof(row.ID))

	stringBuff := make([]byte, consts.TableMaxStringLength)
	copy(stringBuff, []byte(row.Username)[:])
	copy(data[offset:offset+consts.TableMaxStringLength], stringBuff)
	offset += consts.TableMaxStringLength

	stringBuff = make([]byte, consts.TableMaxStringLength)
	copy(stringBuff, []byte(row.Email)[:])
	copy(data[offset:offset+consts.TableMaxStringLength], stringBuff)
	offset += consts.TableMaxStringLength

	_, err := db.conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetDBDataFromFile() error {
	buff, err := io.ReadAll(db.conn)
	if err != nil {
		return err
	}

	if len(buff) >= db.table.rowSize {
		for i := 0; i < len(buff); i += db.table.rowSize {
			data := buff[i:db.table.rowSize]

			id := *(*uint)(unsafe.Pointer(&data[0]))
			offset := unsafe.Sizeof(id)

			s1 := data[offset : offset+consts.TableMaxStringLength]
			name := string(bytes.Trim(s1, "\x00"))
			offset += consts.TableMaxStringLength

			s1 = data[offset : offset+consts.TableMaxStringLength]
			email := string(bytes.Trim(s1, "\x00"))
			offset += consts.TableMaxStringLength

			db.table.Rows = append(db.table.Rows, UserTableRow{
				ID:       id,
				Username: name,
				Email:    email,
			})
			db.table.counter++
			db.table.RowLength++
		}
	}

	return nil
}

func (db *Database) Close() error {
	return db.conn.Close()
}

func (db *Database) ExecuteStatement(state UserTableStatement) error {
	switch state.Type {
	case consts.StInsert:
		fmt.Println("exec insert!")
		err := db.ExecuteInsert(state.Row)
		if err != nil {
			return err
		}
	case consts.StSelect:
		fmt.Println("exec select!")
		err := db.ExecuteSelect()
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTableSizeof(row UserTableRow) int {
	size := 0
	size += int(unsafe.Sizeof(row.ID))
	size += consts.TableMaxStringLength
	size += consts.TableMaxStringLength
	return size
}
