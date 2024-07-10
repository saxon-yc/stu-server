package locales

import (
	_ "embed"
	"log"
	"student-server/internal/model"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle    *i18n.Bundle
	localizer *i18n.Localizer

	//go:embed active.en.toml
	localesENFiles []byte

	//go:embed active.zh.toml
	localesZHFiles []byte
)

func init() {
	// 创建一个新的 Bundle
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
}

// 加载翻译文件
func loadLanguageFile(lang string) {
	switch lang {
	default:
	case model.Language_zh:
		if _, err := bundle.ParseMessageFileBytes(localesZHFiles, "locales/active.zh.toml"); err != nil {
			log.Fatalf("Failed to parse zh_locales message file: %v", err)
		}
	case model.Language_en:
		if _, err := bundle.ParseMessageFileBytes(localesENFiles, "locales/active.en.toml"); err != nil {
			log.Fatalf("Failed to parse en_locales message file: %v", err)
		}
	}

}
func setLanguageEnv(lang string) {
	switch lang {
	default:
	case model.Language_zh:
		localizer = i18n.NewLocalizer(bundle, "zh")
	case model.Language_en:
		localizer = i18n.NewLocalizer(bundle, "en")
	}
}
func NewLocalizer(lang string) {
	loadLanguageFile(lang)
	setLanguageEnv(lang)
}

func GetTranslatedMessage(messageID string) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
}
