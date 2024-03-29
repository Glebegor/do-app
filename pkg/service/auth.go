package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	todo "github.com/Glebegor/do-app"
	"github.com/Glebegor/do-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	Salt       = "lkj41340ewrlkj4324ldsf"
	SingingKey = "dsfu0u423jefs8dufp32ioj#wqe12"
	TokenTTl   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json: "user_id"`
	Username string `json: "username"`
}
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		username,
	})
	return token.SignedString([]byte(SingingKey))
}

func (s *AuthService) ParseToken(accesToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}
		return []byte(SingingKey), nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, claims.Username, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(Salt)))
}
