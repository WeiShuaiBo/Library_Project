package tools

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	AccessTokenDuration  = 30 * time.Hour
	RefreshTokenDuration = 30 * 40 * time.Hour
	TokenIssuer          = "XinXueCheng-library"
	TokenAdmin           = "AdminXinXueCheng-library"
)

type LibraryJwt struct {
	Secret []byte
}

var Token LibraryJwt
var AdToken LibraryJwt

type Claim struct {
	jwt.RegisteredClaims
	ID   int64  `json:"user_id"`
	Name string `json:"username"`
	Role string `json:"role"`
}

func NewToken(sa, sb string) {
	a := []byte("新学橙管理员")
	if sa != "" {
		a = []byte(sa)
	}
	AdToken = LibraryJwt{Secret: a}

	b := []byte("新学橙")
	if sb != "" {
		b = []byte(sb)
	}
	Token = LibraryJwt{Secret: b}
}

// GetToken 获取用户token
func (j *LibraryJwt) GetToken(id int64, name string) (aToken, rToken string, err error) {
	rc := jwt.RegisteredClaims{
		ExpiresAt: j.getTime(AccessTokenDuration),
		Issuer:    TokenIssuer,
	}
	claim := Claim{
		ID:               id,
		Name:             name,
		RegisteredClaims: rc,
		Role:             "user",
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(j.Secret)

	//refresh token 不需要保存任何用户信息
	rc.ExpiresAt = j.getTime(RefreshTokenDuration)
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rc).SignedString(j.Secret)
	return
}

func (j *LibraryJwt) getTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}
func (j *LibraryJwt) keyFunc(token *jwt.Token) (interface{}, error) {
	return j.Secret, nil
}

// VerifyToken 验证Token
func (j *LibraryJwt) VerifyToken(tokenID string) (*Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenID, claim, j.keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid || claim.Role != "user" {
		return nil, errors.New("access token验证失败")
	}
	return claim, nil
}

// GetIDFromToken 根据Token获取ID
func GetIDFromToken(tokenString string) (int64, error) {
	// 解析token

	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return Token.Secret, nil
	})
	if err != nil {
		return 0, err
	}

	// 验证token
	claim, ok := token.Claims.(*Claim)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// 返回id
	return claim.ID, nil
}

// AdminGetToken 获取管理员登录token
func (j *LibraryJwt) AdminGetToken(id int64, name string) (aToken, rToken string, err error) {
	// 在这里进行管理员登录验证逻辑
	rc := jwt.RegisteredClaims{
		Issuer:    TokenAdmin,
		ExpiresAt: j.getTime(AccessTokenDuration),
	}
	claim := Claim{
		RegisteredClaims: rc,
		ID:               id,
		Name:             name,
		Role:             "admin",
	}

	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(j.Secret)
	//refresh token 不需要保存任何用户信息
	rc.ExpiresAt = j.getTime(RefreshTokenDuration)
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rc).SignedString(j.Secret)
	return
}

func (j *LibraryJwt) AdminVerifyToken(tokenID string) (*Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenID, claim, j.keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid || claim.Role != "admin" {
		return nil, errors.New("管理员access token验证失败")
	}
	return claim, nil
}
