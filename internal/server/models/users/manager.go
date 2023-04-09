package users

type Storage interface {
	Find(v any, keyName, keyValues string) error
	Save(v any, keyValue string) error
}

type UserManager struct {
	DB Storage
}

func (um *UserManager) Find(email string) (*User, error) {
	user := &User{}
	if err := um.DB.Find(user, PrimaryKey, email); err != nil {
		return nil, err
	}
	return user, nil
}

func (um *UserManager) Save(u User) error {
	return um.DB.Save(u, u.Email)
}

func (um *UserManager) PasswordHashByEmail(email string) (string, error) {
	user, err := um.Find(email)
	if user == nil {
		return "", err
	}
	return user.PasswordHash, err
}

func (um *UserManager) SetPasswordHashByEmail(email, phash string) error {
	user := User{
		Email:        email,
		PasswordHash: phash,
	}
	return um.Save(user)
}

var Manager *UserManager

func Init(db Storage) *UserManager {
	if Manager != nil {
		return Manager
	}
	Manager = &UserManager{db}
	return Manager
}
