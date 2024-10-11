package services

import (
	"encoding/base64"
	"encoding/binary"
	"urlShorter/initializers"
	"urlShorter/models"

	"github.com/bwmarrin/snowflake"
	"github.com/skip2/go-qrcode"
)

type URLService struct {
	Node *snowflake.Node
}

func NewURLService(nodeID int64) (*URLService, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, err
	}

	return &URLService{Node: node}, nil
}

func (s *URLService) ShortenURL(longURL string) (string, error) {
	id := s.Node.Generate()

	shortURL := s.toBase64(id)

	url := models.URL{LongURL: longURL, ShortURL: shortURL}
	result := initializers.DB.Create(&url)
	if result.Error != nil {
		return "", result.Error
	}

	return shortURL, nil
}

func (s *URLService) toBase64(id snowflake.ID) string {
	byteID := make([]byte, 8)
	binary.BigEndian.PutUint64(byteID, uint64(id))
	encoded := base64.StdEncoding.EncodeToString(byteID)
	return encoded[:8]
}

func (s *URLService) GetOriginalURL(shortURL string) (string, error) {
	var url models.URL
	result := initializers.DB.First(&url, "short_url = ?", shortURL)
	if result.Error != nil {
		return "", result.Error
	}

	return url.LongURL, nil
}

func (s *URLService) GenerateQRCode(shortURL string) ([]byte, error) {
	var url models.URL
	result := initializers.DB.First(&url, "short_url = ?", shortURL)
	if result.Error != nil {
		return nil, result.Error
	}

	fullURL := "http://localhost:8081/" + shortURL
	qrCode, err := qrcode.Encode(fullURL, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return qrCode, nil
}
