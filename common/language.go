package common

import "strings"

type Language string // 語言
const (
	// 中文
	LANGUAGE_CHINESE                Language = "zh"
	LANGUAGE_SIMPLIFIED_CHINESE     Language = "zh_hans"
	LANGUAGE_TRADITIONAL_CHINESE_TW Language = "zh_hant"
	LANGUAGE_TRADITIONAL_CHINESE_HK Language = "zh_hk"
	LANGUAGE_SIMPLIFIED_CHINESE_SG  Language = "zh_sg"
	LANGUAGE_ENGLISH                Language = "en"
	LANGUAGE_ENGLISH_US             Language = "en_us"
	LANGUAGE_ENGLISH_GB             Language = "en_gb"
	LANGUAGE_ENGLISH_AU             Language = "en_au"
	LANGUAGE_ENGLISH_CA             Language = "en_ca"
	LANGUAGE_GERMAN                 Language = "de"
	LANGUAGE_GERMAN_DE              Language = "de_de"
	LANGUAGE_GERMAN_AT              Language = "de_at"
	LANGUAGE_GERMAN_CH              Language = "de_ch"
	LANGUAGE_FRENCH                 Language = "fr"
	LANGUAGE_FRENCH_FR              Language = "fr_fr"
	LANGUAGE_FRENCH_CA              Language = "fr_ca"
	LANGUAGE_SPANISH                Language = "es"
	LANGUAGE_SPANISH_ES             Language = "es_es"
	LANGUAGE_SPANISH_MX             Language = "es_mx"
	LANGUAGE_ITALIAN                Language = "it"
	LANGUAGE_PORTUGUESE             Language = "pt"
	LANGUAGE_PORTUGUESE_BR          Language = "pt_br"
	LANGUAGE_DUTCH                  Language = "nl"
	LANGUAGE_RUSSIAN                Language = "ru"
	LANGUAGE_POLISH                 Language = "pl"
	LANGUAGE_TURKISH                Language = "tr"
	LANGUAGE_JAPANESE               Language = "ja"
	LANGUAGE_KOREAN                 Language = "ko"
	LANGUAGE_THAI                   Language = "th"
	LANGUAGE_VIETNAMESE             Language = "vi"
	LANGUAGE_HINDI                  Language = "hi"
	LANGUAGE_ARABIC                 Language = "ar"
	LANGUAGE_AFRIKAANS              Language = "af"
	LANGUAGE_BULGARIAN              Language = "bg"
	LANGUAGE_CATALAN                Language = "ca"
	LANGUAGE_CZECH                  Language = "cs"
	LANGUAGE_DANISH                 Language = "da"
	LANGUAGE_GREEK                  Language = "el"
	LANGUAGE_FINNISH                Language = "fi"
	LANGUAGE_HEBREW                 Language = "he"
	LANGUAGE_HUNGARIAN              Language = "hu"
	LANGUAGE_INDONESIAN             Language = "id"
	LANGUAGE_NORWEGIAN              Language = "no"
	LANGUAGE_ROMANIAN               Language = "ro"
	LANGUAGE_SLOVAK                 Language = "sk"
	LANGUAGE_SWEDISH                Language = "sv"
	LANGUAGE_UKRAINIAN              Language = "uk"
)

func (Language) FromString(s string) Language {
	sLower := strings.ToLower(s)
	switch sLower {
	case "zh":
		return LANGUAGE_CHINESE
	case "zh_hans":
		return LANGUAGE_SIMPLIFIED_CHINESE
	case "zh_hant":
		return LANGUAGE_TRADITIONAL_CHINESE_TW
	case "zh_hk":
		return LANGUAGE_TRADITIONAL_CHINESE_HK
	case "zh_sg":
		return LANGUAGE_SIMPLIFIED_CHINESE_SG
	case "en":
		return LANGUAGE_ENGLISH
	case "en_us":
		return LANGUAGE_ENGLISH_US
	case "en_gb":
		return LANGUAGE_ENGLISH_GB
	case "en_au":
		return LANGUAGE_ENGLISH_AU
	case "en_ca":
		return LANGUAGE_ENGLISH_CA
	case "de":
		return LANGUAGE_GERMAN
	case "de_de":
		return LANGUAGE_GERMAN_DE
	case "de_at":
		return LANGUAGE_GERMAN_AT
	case "de_ch":
		return LANGUAGE_GERMAN_CH
	case "fr":
		return LANGUAGE_FRENCH
	case "fr_fr":
		return LANGUAGE_FRENCH_FR
	case "fr_ca":
		return LANGUAGE_FRENCH_CA
	case "es":
		return LANGUAGE_SPANISH
	case "es_es":
		return LANGUAGE_SPANISH_ES
	case "es_mx":
		return LANGUAGE_SPANISH_MX
	case "it":
		return LANGUAGE_ITALIAN
	case "pt":
		return LANGUAGE_PORTUGUESE
	case "pt_br":
		return LANGUAGE_PORTUGUESE_BR
	case "nl":
		return LANGUAGE_DUTCH
	case "ru":
		return LANGUAGE_RUSSIAN
	case "pl":
		return LANGUAGE_POLISH
	case "tr":
		return LANGUAGE_TURKISH
	case "ja":
		return LANGUAGE_JAPANESE
	case "ko":
		return LANGUAGE_KOREAN
	case "th":
		return LANGUAGE_THAI
	case "vi":
		return LANGUAGE_VIETNAMESE
	case "hi":
		return LANGUAGE_HINDI
	case "ar":
		return LANGUAGE_ARABIC
	case "af":
		return LANGUAGE_AFRIKAANS
	case "bg":
		return LANGUAGE_BULGARIAN
	case "ca":
		return LANGUAGE_CATALAN
	case "cs":
		return LANGUAGE_CZECH
	case "da":
		return LANGUAGE_DANISH
	case "el":
		return LANGUAGE_GREEK
	case "fi":
		return LANGUAGE_FINNISH
	case "he":
		return LANGUAGE_HEBREW
	case "hu":
		return LANGUAGE_HUNGARIAN
	case "id":
		return LANGUAGE_INDONESIAN
	case "no":
		return LANGUAGE_NORWEGIAN
	case "ro":
		return LANGUAGE_ROMANIAN
	case "sk":
		return LANGUAGE_SLOVAK
	case "sv":
		return LANGUAGE_SWEDISH
	case "uk":
		return LANGUAGE_UKRAINIAN
	}
	return ""
}

func (lang Language) String() string {
	switch lang {
	case LANGUAGE_CHINESE:
		return "zh"
	case LANGUAGE_SIMPLIFIED_CHINESE:
		return "zh_hans"
	case LANGUAGE_TRADITIONAL_CHINESE_TW:
		return "zh_hant"
	case LANGUAGE_TRADITIONAL_CHINESE_HK:
		return "zh_hk"
	case LANGUAGE_SIMPLIFIED_CHINESE_SG:
		return "zh_sg"
	case LANGUAGE_ENGLISH:
		return "en"
	case LANGUAGE_ENGLISH_US:
		return "en_us"
	case LANGUAGE_ENGLISH_GB:
		return "en_gb"
	case LANGUAGE_ENGLISH_AU:
		return "en_au"
	case LANGUAGE_ENGLISH_CA:
		return "en_ca"
	case LANGUAGE_GERMAN:
		return "de"
	case LANGUAGE_GERMAN_DE:
		return "de_de"
	case LANGUAGE_GERMAN_AT:
		return "de_at"
	case LANGUAGE_GERMAN_CH:
		return "de_ch"
	case LANGUAGE_FRENCH:
		return "fr"
	case LANGUAGE_FRENCH_FR:
		return "fr_fr"
	case LANGUAGE_FRENCH_CA:
		return "fr_ca"
	case LANGUAGE_SPANISH:
		return "es"
	case LANGUAGE_SPANISH_ES:
		return "es_es"
	case LANGUAGE_SPANISH_MX:
		return "es_mx"
	case LANGUAGE_ITALIAN:
		return "it"
	case LANGUAGE_PORTUGUESE:
		return "pt"
	case LANGUAGE_PORTUGUESE_BR:
		return "pt_br"
	case LANGUAGE_DUTCH:
		return "nl"
	case LANGUAGE_RUSSIAN:
		return "ru"
	case LANGUAGE_POLISH:
		return "pl"
	case LANGUAGE_TURKISH:
		return "tr"
	case LANGUAGE_JAPANESE:
		return "ja"
	case LANGUAGE_KOREAN:
		return "ko"
	case LANGUAGE_THAI:
		return "th"
	case LANGUAGE_VIETNAMESE:
		return "vi"
	case LANGUAGE_HINDI:
		return "hi"
	case LANGUAGE_ARABIC:
		return "ar"
	case LANGUAGE_AFRIKAANS:
		return "af"
	case LANGUAGE_BULGARIAN:
		return "bg"
	case LANGUAGE_CATALAN:
		return "ca"
	case LANGUAGE_CZECH:
		return "cs"
	case LANGUAGE_DANISH:
		return "da"
	case LANGUAGE_GREEK:
		return "el"
	case LANGUAGE_FINNISH:
		return "fi"
	case LANGUAGE_HEBREW:
		return "he"
	case LANGUAGE_HUNGARIAN:
		return "hu"
	case LANGUAGE_INDONESIAN:
		return "id"
	case LANGUAGE_NORWEGIAN:
		return "no"
	case LANGUAGE_ROMANIAN:
		return "ro"
	case LANGUAGE_SLOVAK:
		return "sk"
	case LANGUAGE_SWEDISH:
		return "sv"
	case LANGUAGE_UKRAINIAN:
		return "uk"
	}
	return ""
}

func (l Language) IsValid() bool {
	switch l {
	case
		LANGUAGE_ENGLISH,
		LANGUAGE_SIMPLIFIED_CHINESE,
		LANGUAGE_TRADITIONAL_CHINESE_TW,
		LANGUAGE_KOREAN,
		LANGUAGE_JAPANESE,
		LANGUAGE_FRENCH,
		LANGUAGE_GERMAN,
		LANGUAGE_SPANISH,
		LANGUAGE_PORTUGUESE,
		LANGUAGE_RUSSIAN,
		LANGUAGE_ARABIC,
		LANGUAGE_VIETNAMESE:
		return true
	}
	return false
}
