package list

type query int

const (
	queryAdd query = iota + 1
	queryGet
	queryGetByID
	queryUpdate
	queryDelete
)
