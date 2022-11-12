package consts

const (
	TableMaxRows         = 4096
	TableMaxStringLength = 255
)

type MetaCommandState int

const (
	McsSuccess MetaCommandState = iota
	McsUnrecognizedCommand
)

type StatementType int

const (
	StUnknown StatementType = iota
	StInsert
	StSelect
)
