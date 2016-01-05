package controllers

import (
	"github.com/pompeu/Godeps/_workspace/src/golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
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

func compare(hash, pass []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err
}

func validName(name string) bool {
	v, _ := regexp.MatchString("^[a-zA-Z0-9 ]+$", name)
	return v
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func validPost(title, body, tags string) Invalid {
	if !validName(title) {
		return Invalid{"title", "required or invalid"}
	}
	if !validName(body) {
		return Invalid{"body", "required or invalid"}
	}
	if sTags := strings.Split(tags, " "); len(sTags) != 3 {
		return Invalid{"tags", "required and need 3 tagas separate by single space"}
	}
	return Invalid{"all", "done"}
}
func validImputs(name, email, password string) Invalid {
	if !validName(name) {
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
