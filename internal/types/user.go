package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"first_name" json:"firstName"`
	LastName  string `bson:"last_name" json:"lastName"`
}
