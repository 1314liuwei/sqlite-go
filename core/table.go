package core

import "1314liuwei/sqlite.go/consts"

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
