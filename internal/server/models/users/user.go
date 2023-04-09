package users

const (
	TableName  = "users"
	PrimaryKey = "email"
)

type User struct {
	Email        string `gorm:"primaryKey"`
	PasswordHash string
}
