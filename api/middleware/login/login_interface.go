package login

import (
	"net/http"
)

type Login interface {
	Authenticate(next http.HandlerFunc) http.HandlerFunc
	GetToken(user LoginData, interval int) (map[string]string, error)
}

type LoginStorage interface {
	SetLastAccess(userId int) error
}
