package entity

import (
	"log"
	"os"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"hacktiv8-msib-final-project-3/pkg/errs"
)

var jwtSecret = os.Getenv("JWT_SECRET")

type User struct {
	gorm.Model
	FullName string `gorm:"not null" binding:"required"`
	Email    string `gorm:"unique;not null" binding:"email,required"`
	Password string `gorm:"not null" binding:"required,min=6"`
	Role     string `gorm:"not null" binding:"required,oneof=admin member"`
	Tasks    []Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) HashPassword() errs.MessageErr {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return errs.NewInternalServerError("Failed to hash password")
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) errs.MessageErr {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errs.NewBadRequest("Password is not valid!")
	}

	return nil
}

func (u *User) CreateToken() (string, errs.MessageErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": u.ID,
			"exp":    time.Now().Add(1 * time.Hour).Unix(),
		})

	signedString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("Error:", err.Error())
		return "", errs.NewInternalServerError("Failed to sign jwt token")
	}

	return signedString, nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	if isBearer := strings.HasPrefix(bearerToken, "Bearer"); !isBearer {
		return errs.NewUnauthenticated("Token type should be Bearer")
	}

	splitToken := strings.Fields(bearerToken)
	if len(splitToken) != 2 {
		return errs.NewUnauthenticated("Token is not valid")
	}

	tokenString := splitToken[1]
	token, err := u.ParseToken(tokenString)
	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errs.NewUnauthenticated("Token is not valid")
	}
	mapClaims = claims

	return u.bindTokenToUserEntity(mapClaims)
}

func (u *User) ParseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthenticated("Token method is not valid")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, errs.NewUnauthenticated("Token is not valid")
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	id, ok := claim["userId"].(float64)
	if !ok {
		return errs.NewUnauthenticated("Token doesn't contains userId")
	}
	u.ID = uint(id)

	return nil
}
