package service

import (
	"math/rand"
	"short-url-app/pkg/repository"
)

type Service struct {
	repo *repository.Repository
	host string
}

func NewService(repo *repository.Repository, host string) *Service {
	return &Service{repo: repo, host: host}
}

func (s *Service) SaveLongUrl(longUrl string) (string, error) {
	shortUrl := generateShortUrl()
	if err := s.repo.SaveUrl(longUrl, shortUrl); err != nil {
		return "", err
	}
	return s.host + shortUrl, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789_")

func generateShortUrl() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *Service) GetLongUrl(longUrl string) (string, error) {
	return s.repo.GetLongUrl(longUrl)
}
