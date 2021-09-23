package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Toolnado/SwalloW/internal/repository"
	"github.com/Toolnado/SwalloW/model"
	"github.com/Toolnado/SwalloW/pkg/utils"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt             = "IGTBFIGINILHyghtjfdkfsoeifje#$%^&*())ijg__hash__password"
	defaultAccessKey = "qwertyuiopasdfghjklzxcvbnmWERTYUIOPASDFGHJKLZXCVBNM"
	tokenTTL         = 2 * time.Hour
)

type AuthService struct {
	repo repository.Authentication
}

type swallowClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func NewAuthService(r repository.Authentication) *AuthService {
	return &AuthService{
		repo: r,
	}
}

func (s *AuthService) CreateUser(u *model.User) (int, error) {
	u.Password = generateHashPassword(u.Password)
	u.CreateAt = time.Now()
	id, err := s.repo.CreateUser(u)

	if err != nil {
		log.Printf("[not found id, error create user: %s]\n", err)
		return id, err
	}

	return id, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	accessKey := utils.GetEnvStr("ACCESS_KEY", defaultAccessKey)
	hashPassword := generateHashPassword(password)
	user, err := s.repo.GetUser(username, hashPassword)

	if err != nil {
		log.Printf("[user not found: %s]\n", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &swallowClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(accessKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	accessKey := utils.GetEnvStr("ACCESS_KEY", defaultAccessKey)

	token, err := jwt.ParseWithClaims(accessToken, &swallowClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(accessKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*swallowClaims)

	if !ok {
		return 0, errors.New("token claims are not type *swallowClaims")
	}

	return claims.UserID, nil
}

func generateHashPassword(password string) string {
	c := sha1.New()
	c.Write([]byte(password))
	return fmt.Sprintf("%x", c.Sum([]byte(salt)))
}
