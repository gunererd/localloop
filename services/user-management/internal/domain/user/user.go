package user

// User represents a user in the system
type User struct {
	Email string `bson:"email"`    // Tags for MongoDB serialization
	Hash  string `bson:"password"` // In production, this would be a hashed password
	Salt  string
	Name  string `bson:"name"`
}
