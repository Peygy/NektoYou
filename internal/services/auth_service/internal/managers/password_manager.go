package managers

import (
	"errors"
	"strconv"

	"github.com/peygy/nektoyou/internal/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type iPasswordManager interface {
	hashPassword(password string) (string, error)
	validPassword(password string) error
	checkPasswordHash(password, hash string) bool
}

type passwordManager struct {
	minLen int
	log    logger.ILogger
}

func newPasswordManager(minLen int, log logger.ILogger) iPasswordManager {
	log.Info("PasswordManager created")
	return &passwordManager{minLen: minLen, log: log}
}

func (p passwordManager) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		p.log.Errorf("Can't create hashed password with error: %v", err)
		return "", errors.New("managers-password: can't create hashed password")
	}

	p.log.Info("Password is hashed successfully")
	return string(bytes), nil
}

func (p passwordManager) validPassword(password string) error {
	if len(password) < p.minLen {
		p.log.Error("Password length less than minimum length")
		return errors.New("managers-password: user password is not valid: password length less than " + strconv.Itoa(p.minLen))
	}

	p.log.Info("Password is valided")
	return nil
}

func (p passwordManager) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
