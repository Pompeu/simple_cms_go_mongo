package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type Invalid struct {
	Input string
	Value string
}

func genereteCode(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func compare(hash, pass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err != nil
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func validImputs(name, email, password string) Invalid {
	if m, _ := regexp.MatchString("^[a-zA-Z ]+$", name); !m {
		return Invalid{"name", "required or invalid"}
	}
	if m := validateEmail(email); !m {
		return Invalid{"email", "required or invalid"}
	}
	if m, _ := regexp.MatchString("^[0-9a-zA-Z]{8}$", password); !m {
		return Invalid{"password", "required or invalid"}
	}
	return Invalid{"all", "done"}
}
