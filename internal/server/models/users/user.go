package users

type User struct {
	Email        string `gorm:"primaryKey"`
	PasswordHash string
}
