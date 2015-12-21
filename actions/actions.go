package actions

type Action interface {
	Execute(parameters string) string
}
