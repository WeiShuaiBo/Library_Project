package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	AccessTokenDuration  = 2 * time.Hour
	RefreshTokenDuration = 7 * 24 * time.Hour
	TokenIssuer          = "library by wsb"
)

type LibraryJwt struct {
	Secret []byte
}

var Token LibraryJwt

type MyClaims struct {
	UserID   uint64 `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

func NewToken(s string) {
	b := []byte("wsb")
	if s != "" {
		b = []byte(s)
	}
	Token = LibraryJwt{Secret: b}
}

// 获取时间
func (l *LibraryJwt) getTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

func (l *LibraryJwt) keyFunc(token *jwt.Token) (interface{}, error) {
	return l.Secret, nil
}

// 生成token
func (l *LibraryJwt) GetToken(id uint64, name string) (aToken, rToken string, err error) {
	rc := jwt.RegisteredClaims{
		ExpiresAt: l.getTime(AccessTokenDuration),
		Issuer:    TokenIssuer,
	}
	claim := MyClaims{
		UserID:           id,
		UserName:         name,
		RegisteredClaims: rc,
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(l.Secret)
	//refresh token 不需要保存任何用户信息
	rc.ExpiresAt = l.getTime(RefreshTokenDuration)
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(l.Secret)
	return
}

//刷新token
func (l *LibraryJwt) RefreshToken(a, r string) (aToken, rToken string, err error) {
	if _, err = jwt.Parse(r, l.keyFunc); err != nil {
		return
	}
	//从旧access token 中解析出 claims 数据
	claim := &MyClaims{}
	_, err = jwt.ParseWithClaims(a, claim, l.keyFunc)
	//判断错误是不是因为access token 正常过期导致的
	if errors.Is(err, jwt.ErrTokenExpired) {
		return l.GetToken(claim.UserID, claim.UserName)
	}
	return
}

//验证token
func (l *LibraryJwt) VerifyToken(tokenID string) (*MyClaims, error) {
	claim := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenID, claim, l.keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("access token 验证失败")
	}
	return claim, nil
}