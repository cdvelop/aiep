package user

// User represents a user account
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`    // Don't expose hash in JSON responses
	Role         string `json:"role"` // e.g., "admin", "user"
}
