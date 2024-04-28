package cipher

import (
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

type UserPassword struct {
	cfg Config
}

func NewUserPassword(cfg Config) *UserPassword {
	return &UserPassword{
		cfg: cfg,
	}
}

func (u UserPassword) HashPassword(password string, salt string) string {
	hash := argon2.IDKey([]byte(password), []byte(u.cfg.Paper+salt), u.cfg.Time, u.cfg.Memory, u.cfg.Thread, u.cfg.KeyLen)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	return encodedHash
}

func (u UserPassword) ComparePasswords(encryptedPassword string, userPassword string, salt string) bool {
	decodedEncryptedPassword, err := base64.RawStdEncoding.Strict().DecodeString(encryptedPassword)
	if err != nil {
		//TODO: might need to send a log
		return false
	}
	decodedUserPassword, err := base64.RawStdEncoding.Strict().DecodeString(u.HashPassword(userPassword, salt))
	if err != nil {
		//TODO: might need to send a log
		return false
	}
	return subtle.ConstantTimeCompare(decodedEncryptedPassword, decodedUserPassword) == 1
}
