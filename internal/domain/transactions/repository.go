package transactions

type Repository interface {
	Save(transactions *Transactions) error
}
