package server

import "errors"

func validate(email, password string) error {
	if email == "" || password == "" {
		return errors.New("field cannot be empty")
	}
	return nil
}
