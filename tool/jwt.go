package tool

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	AccessToken struct {
		Token string
		Secret string
		Expire int64
		jwt.StandardClaims
	}
)

//加密key
var jwtSecret = []byte("0^2SljwaYzlRU7*u")

//生成token
func (a *AccessToken)GenerateToken()  error {

	claims := AccessToken{
		Token: a.Secret,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: a.Expire,
			// 指定token发行人
			Issuer: "ganganlee",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return  err
	}

	a.Token = token
	return  nil
}

//验证token
func (a *AccessToken)ValidateToken() (string,error) {

	accessToken := &AccessToken{}
	token, err := jwt.ParseWithClaims(a.Token, accessToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return "",err
	}

	if claims,ok:=token.Claims.(*AccessToken);ok&&token.Valid{
		return claims.Token,nil
	}else {
		return "",nil
	}
}