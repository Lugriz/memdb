package domain

type Engine interface {
	Execute(Command) (OperationResult, error)
}
