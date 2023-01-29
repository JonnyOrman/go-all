package command

type ModNameProvider interface {
	Get(modFile string) string
}
