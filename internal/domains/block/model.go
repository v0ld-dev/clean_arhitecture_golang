package block

import "clean_arhitecture_golang/internal/domains/transaction"

type Block struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Transactions []transaction.Transaction
}
