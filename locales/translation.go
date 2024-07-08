package locales

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var localizer *i18n.Localizer

func init() {
	// 创建一个新的 Bundle
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// 加载翻译文件
	_, err := bundle.LoadMessageFile("locales/active.en.toml")
	if err != nil {
		log.Fatalf("failed to load message file: %v", err)
	}
	_, err = bundle.LoadMessageFile("locales/active.zh.toml")
	if err != nil {
		log.Fatalf("failed to load message file: %v", err)
	}
}
func NewLocalizer(lang string) {
	switch lang {
	default:
	case "zh-cn":
		localizer = i18n.NewLocalizer(bundle, "zh")
	case "en":
		localizer = i18n.NewLocalizer(bundle, "en")
	}
}

func GetTranslatedMessage(messageID string) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
}
