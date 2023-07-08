package tools

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	//两个小时
	AccessTokenDuration = 2 * time.Hour
	//一个月
	RefreshTokenDuration = 30 * 24 * time.Hour
	TokenIssuer          = "xinxuecheng-vote" //token的发布者
)

var Token VoteJwt //用于处理JWT相关操作

func NewToken(s string) {
	b := []byte("香香编程喵喵喵")
	if s != "" {
		b = []byte(s)
	}

	Token = VoteJwt{Secret: b}
}

type VoteJwt struct {
	Secret []byte //JWT密钥
}

// Claim 自定义的数据结构，这里使用了结构体的组合
type Claim struct {
	jwt.RegisteredClaims
	ID   int64  `json:"user_id"`
	Name string `json:"username"`
}

// 用于获取指定时间间隔后的NumericDate对象,返回一个未来的时间
func (j *VoteJwt) getTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

// 用于返回用于验证签名的密钥
func (j *VoteJwt) keyFunc(token *jwt.Token) (interface{}, error) {
	return j.Secret, nil
}

// GetToken 颁发token access token 和 refresh token
func (j *VoteJwt) GetToken(id int64, name string) (aToken, rToken string, err error) {
	rc := jwt.RegisteredClaims{
		ExpiresAt: j.getTime(AccessTokenDuration), //jwt的过期时间
		Issuer:    TokenIssuer,                    //签发者名称
	}
	claim := Claim{
		ID:               id,
		Name:             name,
		RegisteredClaims: rc,
	}
	//使用HMAC-SHA256算法对声明对象进行签名，返回一个签名后的JWT对象
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(j.Secret)

	// refresh token 不需要保存任何用户信息
	rc.ExpiresAt = j.getTime(RefreshTokenDuration)
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rc).SignedString(j.Secret)
	return
}

// VerifyToken 验证Token的有效性，并返回解析后的Claim（载荷）对象
func (j *VoteJwt) VerifyToken(tokenID string) (*Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenID, claim, j.keyFunc)
	if err != nil {
		return nil, err
	}
	//如果值为true,则表示JWT通过了签名验证，并且再有效期被，否则表示无效，原因可能是签名验证失败，已经过期等
	if !token.Valid {
		return nil, errors.New("access token 验证失败")
	}

	return claim, nil
}

// RefreshToken 通过 refresh token（刷新令牌） 刷新 access token（访问令牌）
func (j *VoteJwt) RefreshToken(a, r string) (aToken, rToken string, err error) {
	// r 无效直接返回
	if _, err = jwt.Parse(r, j.keyFunc); err != nil {
		return
	}
	// 从旧access token 中解析出claims数据
	claim := &Claim{}
	_, err = jwt.ParseWithClaims(a, claim, j.keyFunc)
	// 判断错误是不是因为access token 正常过期导致的
	if errors.Is(err, jwt.ErrTokenExpired) {
		return j.GetToken(claim.ID, claim.Name)
	}
	return
}
