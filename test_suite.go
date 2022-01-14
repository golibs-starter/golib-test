package golibtest

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/golibs-starter/golib/web/log"
	"gorm.io/gorm"
	"time"
)

type TestSuite struct {
	db         *gorm.DB
	jwtSignKey *rsa.PrivateKey
}

func NewTestSuite(db *gorm.DB, jwtProperties *JwtTestProperties) *TestSuite {
	testSuite := &TestSuite{
		db: db,
	}
	testSuite.loadJwtPrivateKey(jwtProperties)
	return testSuite
}

func (ts *TestSuite) DB() *gorm.DB {
	return ts.db
}

func (ts *TestSuite) CreateJwtToken(userId string) string {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	now := time.Now()
	token.Claims = &jwt.StandardClaims{
		Issuer:    "TESTER",
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Minute * 1).Unix(),
		Subject:   userId,
	}
	jwtToken, err := token.SignedString(ts.jwtSignKey)
	if err != nil {
		log.Fatalf("Could not create jwt token: %v", err)
	}
	return jwtToken
}

func (ts *TestSuite) loadJwtPrivateKey(jwtProperties *JwtTestProperties) {
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(jwtProperties.PrivateKey))
	if err != nil {
		log.Fatalf("Could not load jwt private key: %v", err)
	}
	ts.jwtSignKey = signKey
}

func (ts *TestSuite) TruncateTables(tables []string) {
	for _, table := range tables {
		if err := ts.db.Exec(fmt.Sprintf("TRUNCATE TABLE `%s`", table)).Error; err != nil {
			log.Fatalf("Could not truncate table %s: %v", table, err)
		}
	}
}

func (ts *TestSuite) Seed(model interface{}) {
	if err := ts.db.Create(model).Error; err != nil {
		log.Fatalf("Could not create seed data, model: %v, err: %v")
	}
}
