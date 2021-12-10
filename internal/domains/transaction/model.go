package transaction

type Transaction struct {
	ID   int    		`json:"id,omitempty"`
	Name string 		`json:"name,omitempty"`
}
