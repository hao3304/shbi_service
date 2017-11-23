package util

import (
	"shbi_service/models"
	jwt "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"github.com/astaxie/beego/logs"
	"fmt"
	"errors"
)
var (
	key []byte = []byte("flyone@soft")
)


func GenToken(user *models.User) (string,int64) {
	expires := int64(time.Now().Add(time.Hour * time.Duration(24*5)).Unix())
	claims := jwt.StandardClaims{
		Audience:user.UserName,
		Id:strconv.Itoa(user.Id),
		NotBefore:int64(time.Now().Unix()),
		ExpiresAt:expires,
		Issuer:"jack",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logs.Error(err)
		return "", 0
	}
	return ss,expires
}


func ParseJwt(token string) (*jwt.StandardClaims, error) {
	var jclaim = &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, jclaim, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err, token)
		return nil, errors.New("parse with claims failed")
	}
	return jclaim, nil
}
