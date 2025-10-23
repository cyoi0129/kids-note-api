package services

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt"
)

const SECRET_KEY = "SECRET"

// 文字列の配列から数字配列へ変換
func convert2Int(strings []string) (ints []int) {
	var results []int
	for _, i := range strings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		results = append(results, j)
	}
	return results
}

// 暗号(Hash)化
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// メールの暫定トークン
func CreateMailToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": email,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// メール暫定トークンの有効化をチェック
func CheckMailToken(token string) bool {
	tokenResult, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	return err == nil && tokenResult.Valid
}

// ユーザートークン
func CreateToken(user_id int, family_id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":   user_id,
		"family": family_id,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// トークンの有効化をチェック
func CheckToken(token string) bool {
	tokenResult, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	return err == nil && tokenResult.Valid
}

type Auth struct {
	UserID   int
	FamilyID int
}

// トークンから情報取得
func ParseToken(signedString string) (*Auth, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("%s is expired: %w", signedString, err)
			} else {
				return nil, fmt.Errorf("%s is invalid: %w", signedString, err)
			}
		} else {
			return nil, fmt.Errorf("%s is invalid: %w", signedString, err)
		}
	}

	if token == nil {
		return nil, fmt.Errorf("not found token in %s", signedString)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("not found claims in %s", signedString)
	}
	user_id, ok := claims["user"].(float64)
	if !ok {
		return nil, fmt.Errorf("not found exp in %s", signedString)
	}
	family_id, ok := claims["family"].(float64)
	if !ok {
		return nil, fmt.Errorf("not found exp in %s", signedString)
	}

	return &Auth{
		UserID:   int(user_id),
		FamilyID: int(family_id),
	}, nil
}

// APIのトークンからの妥当な家族権限チェック
func CheckFamilyPermission(token string, id int) bool {
	auth, err := ParseToken(token)
	if err != nil {
		return false
	}
	return auth.FamilyID == id
}

// APIのトークンからの妥当な家族権限チェック
func CheckUserPermission(token string, id int) bool {
	auth, err := ParseToken(token)
	if err != nil {
		return false
	}
	return auth.UserID == id
}
