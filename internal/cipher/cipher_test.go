package cipher

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserPassword(t *testing.T) {
	cfg := Config{
		MoreSalt: "test123",
		Time:     1,
		Memory:   64 * 1024,
		Thread:   1,
		KeyLen:   64,
	}
	password := "1234qwer"
	salt := "salt123"

	p := NewUserPassword(cfg)
	encryptedPassword := p.HashPassword(password, salt)
	ok := p.ComparePasswords(encryptedPassword, password, salt)
	assert.True(t, ok)
	log.Println("encryptedPassword: ", string(encryptedPassword), " password: ", password, " isMatch:", ok)

}
