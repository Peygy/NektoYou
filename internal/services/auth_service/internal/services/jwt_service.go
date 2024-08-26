package services

import "time"

type TokenManager interface {
	NewAccessToken(userId string, ttl time.Duration) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	secretKey string `yaml:"secretKey"`
}

func (m *Manager) NewAccessToken(userId string, ttl time.Duration) (string, error) {

}

func (m *Manager) NewRefreshToken() (string, error) {

}