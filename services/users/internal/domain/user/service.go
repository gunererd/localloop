package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"

	"localloop/libs/pkg/errorbuilder"
	"localloop/services/users/internal/shared"
	apperr "localloop/services/users/internal/shared/error"
)

type Service struct {
	repo                  Repository // Use the Repository interface
	jwtSecret             []byte
	jwtExpirationDuration int
}

type ServiceConfig struct {
	JWTSecret            string
	JWTExpirationMinutes int
	SaltLength           int
}

type TokenClaims struct {
	Email string
	jwt.StandardClaims
}

func NewService(repo Repository, cfg ServiceConfig) *Service {
	return &Service{
		repo:                  repo,
		jwtSecret:             []byte(cfg.JWTSecret),
		jwtExpirationDuration: cfg.JWTExpirationMinutes,
	}
}

func (s *Service) Register(email, password, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return apperr.ErrUserExists(
			apperr.WithUser(email),
		)
	}

	salt, err := shared.GenerateSalt(25)
	if err != nil {
		return apperr.ErrGeneratingSalt(
			errorbuilder.WithOriginal(err),
		)
	}

	hashedPassword, err := shared.HashPassword(password, salt)
	if err != nil {
		return apperr.ErrHashingPassword(
			errorbuilder.WithOriginal(err),
		)
	}

	user := User{
		Email: email,
		Hash:  hashedPassword,
		Salt:  salt,
		Name:  name,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return apperr.ErrDatabaseOperation(
			errorbuilder.WithOriginal(err),
			apperr.WithUser(email),
		)
	}

	return nil
}

func (s *Service) Login(email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", apperr.ErrInvalidCredentials(
			apperr.WithUser(email),
		)
	}

	if !shared.CheckPasswordHash(password, user.Salt, user.Hash) {
		return "", apperr.ErrInvalidCredentials(
			apperr.WithUser(email),
		)
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Duration(s.jwtExpirationDuration) * time.Minute).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", apperr.ErrDatabaseOperation(
			errorbuilder.WithOriginal(err),
			apperr.WithUser(email),
		)
	}

	return tokenString, nil
}

func (s *Service) Get(email string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return User{}, apperr.ErrUserNotFound(
			apperr.WithUser(email),
		)
	}

	return user, nil
}

func (s *Service) Update(email, name, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updates := UpdateData{Name: name}
	if password != "" {
		salt, err := shared.GenerateSalt(25)
		if err != nil {
			return apperr.ErrGeneratingSalt(errorbuilder.WithOriginal(err))
		}
		hash, err := shared.HashPassword(password, salt)
		if err != nil {
			return apperr.ErrHashingPassword(errorbuilder.WithOriginal(err))
		}
		updates.Hash = hash
		updates.Salt = salt
	}

	return s.repo.Update(ctx, email, updates)
}

func (s *Service) ValidateToken(tokenString string) (*TokenClaims, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})

	if err != nil {
		switch v := err.(type) {
		case *jwt.ValidationError:
			switch v.Errors {
			case jwt.ValidationErrorExpired:
				return nil, apperr.ErrTokenExpired()
			default:
				return nil, apperr.ErrInvalidToken(
					errorbuilder.WithOriginal(err),
				)
			}
		default:
			return nil, apperr.ErrInvalidToken(
				errorbuilder.WithOriginal(err),
			)
		}
	}

	if !token.Valid {
		return nil, apperr.ErrInvalidToken()
	}

	return claims, nil
}
