package tools

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	AccessTokenDuration  = 2 * time.Hour
	RefreshTokenDuration = 30 * 24 * time.Hour
	TokenIssuer          = "Library-Curator"
)

var Token VoteJwt

func GetSecret(s string) {
	b := []byte("secret-key")
	if s != "" {
		b = []byte(s)
	}
	Token = VoteJwt{Secret: b}
}

type VoteJwt struct {
	Secret []byte
}
type Claim struct {
	jwt.RegisteredClaims
	Name string `json:"username"`
	ID   int64  `json:"user_id"`
}

func (j *VoteJwt) getTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}
func (j *VoteJwt) keyFunc(token *jwt.Token) (interface{}, error) {
	return j.Secret, nil
}
func (j *VoteJwt) GetToken(id int64, name string) (aToken, rToken string, err error) {
	rc := jwt.RegisteredClaims{
		ExpiresAt: j.getTime(AccessTokenDuration),
		Issuer:    TokenIssuer,
	}
	claim := &Claim{
		Name:             name,
		ID:               id,
		RegisteredClaims: rc,
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(j.Secret)
	rc.ExpiresAt = j.getTime(RefreshTokenDuration)
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rc).SignedString(j.Secret)
	return
}
func (j *VoteJwt) VerifyToken(tokenID string) (*Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenID, claim, j.keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("access token 验证失效")
	}
	return claim, nil
}
func (j *VoteJwt) refreshToken(a, r string) (aToken, rToken string, err error) {
	if _, err = jwt.Parse(a, j.keyFunc); err != nil {
		return
	}
	claim := &Claim{}
	_, err = jwt.ParseWithClaims(a, claim, j.keyFunc)
	if errors.Is(err, jwt.ErrTokenExpired) {
		return j.GetToken(claim.ID, claim.Name)
	}
	return
}
