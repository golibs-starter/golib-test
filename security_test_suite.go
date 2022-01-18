package golibtest

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type SecurityTestSuite struct {
	jwtSignKey    *rsa.PrivateKey
	jwtProperties *JwtTestProperties
}

func NewSecurityTestSuite(jwtProperties *JwtTestProperties) *SecurityTestSuite {
	ts := &SecurityTestSuite{jwtProperties: jwtProperties}
	ts.LoadJwtPrivateKey()
	return ts
}

// LoadJwtPrivateKey load jwt config from properties
func (s *SecurityTestSuite) LoadJwtPrivateKey() {
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(s.jwtProperties.PrivateKey))
	if err != nil {
		log.Fatalf("Could not load jwt private key: %v", err)
	}
	s.jwtSignKey = signKey
}

// CreateJwtToken return a new jwt token
func (s *SecurityTestSuite) CreateJwtToken(userId string) string {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	now := time.Now()
	token.Claims = &jwt.StandardClaims{
		Issuer:    "TESTER",
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Minute * 1).Unix(),
		Subject:   userId,
	}
	jwtToken, err := token.SignedString(s.jwtSignKey)
	if err != nil {
		log.Fatalf("Could not create jwt token: %v", err)
	}
	return jwtToken
}
