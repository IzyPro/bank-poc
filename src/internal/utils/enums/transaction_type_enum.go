package utils

type TransactionType int

const (
	Credit TransactionType = iota + 1
	Debit
)

func (t TransactionType) String() string {
	return [...]string{"Credit", "Debit"}[t-1]
}

func (t TransactionType) EnumIndex() int {
	return int(t)
}
