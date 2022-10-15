package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/smtp"
)

type UserClaims struct {
	Identity primitive.ObjectID `bson:"identity"`
	Email    string             `json:"email"`
	jwt.StandardClaims
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("im")

// GenerateToken
// 生成 token
func GenerateToken(identity, email string) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(identity)
	if err != nil {
		return "", err
	}
	UserClaim := &UserClaims{
		Identity:       objectID,
		Email:          email,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// SendCode
// 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <l18049444798@163.com>" // 发件人
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>") // 邮件内容
	return e.SendWithTLS("smtp.163.com:465",     // smtp.163.com:465,
		smtp.PlainAuth("", "l18049444798@163.com", "SHYAHIDBPPYKQQHL", "smtp.163.com"), // 这里的密码是授权码,不是邮箱密码,授权码在邮箱设置里面,163邮箱是POP3/SMTP服务
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})              // 这里的 ServerName 必须是 smtp.163.com
}
