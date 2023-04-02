package lib

import (
	"fmt"
)

// Language
type LanguageModel struct {
	// Name of the language model (e.g Alpaca.7B)
	name string `yaml:"name"`
	// URL to download this language model
	url string `yaml:"url"`
	// Magnet link for the torrent
	magnet string `yaml:"magnet"`
}

// Downloads the language model either from URL or torrent
func (model *LanguageModel) download() error {
	// Prioritize download from web
	err := model.downloadWeb()
	if err == nil {
		return err
	}
	// Try to download from torrent
	return model.downloadTorrent()
}

// Downloads the language model from the web
func (model *LanguageModel) downloadWeb() error {
	// TODO: Add download with web client
	return nil
}

// Downloads the language model from the torrent
func (model *LanguageModel) downloadTorrent() error {
	if model.magnet == "" {
		return nil
	} else {
		return fmt.Errorf("Language model: %s does not contain a torrent link", model.name)
	}
}
