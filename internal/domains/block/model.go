package block

import (
	"clean_arhitecture_golang/internal/domains/transaction"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Block struct {
	ID primitive.ObjectID    				`bson:"_id" json:"id,omitempty"`
	Name string 							`bson:"name" json:"name"`
	CreatedAt time.Time 					`bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time 					`bson:"updated_at" json:"updated_at,omitempty"`
	Transactions []transaction.Transaction  `bson:"transactions" json:"transactions,omitempty"`
}
