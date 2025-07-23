package storage

import (
	"corporate/config"
	"io"

	"github.com/labstack/gommon/log"
	storage_go "github.com/supabase-community/storage-go"
)

type SupabaseInterface interface {
	UploadFile(path string, file io.Reader) (string, error)
}

type supabaseStruct struct {
	cfg *config.Config
}

func NewSupabase(cfg *config.Config) SupabaseInterface {
	return &supabaseStruct{
		cfg: cfg,
	}
}

func (s *supabaseStruct) UploadFile(path string, file io.Reader) (string, error) {
	client := storage_go.NewClient(s.cfg.Supabase.StorageUrl, s.cfg.Supabase.StorageKey, map[string]string{"Content-Type": "image/*"})
	_, err := client.UploadFile(s.cfg.Supabase.StorageBucket, path, file)

	if err != nil {
		log.Errorf("Error uploading file: %v", err)
		return "", err
	}

	result := client.GetPublicUrl(s.cfg.Supabase.StorageBucket, path)
	return result.SignedURL, nil
}
