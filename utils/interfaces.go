package utils

type TxParams interface {
	ToArgs() []any
}
