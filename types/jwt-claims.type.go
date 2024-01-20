package types

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	jwt.MapClaims
	Email	string
	Exp		time.Time
	Ref		string
	Sub		string
	Version	string
}