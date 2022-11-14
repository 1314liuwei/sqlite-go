package consts

const (
	TableMaxRows         = 4096
	TableMaxStringLength = 255
)

type StatementType int

const (
	StUnknown StatementType = iota
	StInsert
	StSelect
)
