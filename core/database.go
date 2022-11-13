package core

type DB interface {
	ExecuteStatement(state UserTableStatement) error
}

func Open(name string) (DB, error) {
	return nil, nil
}
