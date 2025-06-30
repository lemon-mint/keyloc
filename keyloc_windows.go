//go:build windows

package keyloc

import (
	"fmt"
	"syscall"
	"unsafe"
)

func getLanguages() ([]string, error) {
	user32 := syscall.NewLazyDLL("user32.dll")
	getKeyboardLayoutList := user32.NewProc("GetKeyboardLayoutList")

	var numLayouts int32
	ret, _, err := getKeyboardLayoutList.Call(0, uintptr(unsafe.Pointer(&numLayouts)))
	if ret == 0 {
		return nil, fmt.Errorf("failed to get number of keyboard layouts: %v", err)
	}

	// Handle the case where no keyboard layouts are present
	if numLayouts == 0 {
		return []string{}, nil
	}

	layouts := make([]uintptr, numLayouts)
	ret, _, err = getKeyboardLayoutList.Call(uintptr(numLayouts), uintptr(unsafe.Pointer(&layouts[0])))
	if ret == 0 {
		return nil, fmt.Errorf("failed to get keyboard layouts: %v", err)
	}

	langSet := make(map[string]bool)
	for _, layout := range layouts {
		langID := uint16(layout)
		lang := langID & 0x3ff
		code := langCode(langID) // Use the full langID for specific locales
		if code != "unknown" {
			langSet[code] = true
		} else {
			code = langCode(lang) // Fallback to primary language ID
			if code != "unknown" {
				langSet[code] = true
			}
		}
	}

	langs := make([]string, 0, len(langSet))
	for lang := range langSet {
		langs = append(langs, lang)
	}

	return langs, nil
}

// langCode maps a Windows LCID to a language tag.
// Based on [MS-LCID] v20240423.
func langCode(id uint16) string {
	switch id {
	case 0x0001:
		return "ar"
	case 0x0002:
		return "bg"
	case 0x0003:
		return "ca"
	case 0x0004:
		return "zh-CN" // zh-Hans
	case 0x0005:
		return "cs"
	case 0x0006:
		return "da"
	case 0x0007:
		return "de"
	case 0x0008:
		return "el"
	case 0x0009:
		return "en"
	case 0x000a:
		return "es"
	case 0x000b:
		return "fi"
	case 0x000c:
		return "fr"
	case 0x000d:
		return "he"
	case 0x000e:
		return "hu"
	case 0x000f:
		return "is"
	case 0x0010:
		return "it"
	case 0x0011:
		return "ja"
	case 0x0012:
		return "ko"
	case 0x0013:
		return "nl"
	case 0x0014:
		return "no"
	case 0x0015:
		return "pl"
	case 0x0016:
		return "pt"
	case 0x0017:
		return "rm"
	case 0x0018:
		return "ro"
	case 0x0019:
		return "ru"
	case 0x001a:
		return "hr"
	case 0x001b:
		return "sk"
	case 0x001c:
		return "sq"
	case 0x001d:
		return "sv"
	case 0x001e:
		return "th"
	case 0x001f:
		return "tr"
	case 0x0020:
		return "ur"
	case 0x0021:
		return "id"
	case 0x0022:
		return "uk"
	case 0x0023:
		return "be"
	case 0x0024:
		return "sl"
	case 0x0025:
		return "et"
	case 0x0026:
		return "lv"
	case 0x0027:
		return "lt"
	case 0x0028:
		return "tg"
	case 0x0029:
		return "fa"
	case 0x002a:
		return "vi"
	case 0x002b:
		return "hy"
	case 0x002c:
		return "az"
	case 0x002d:
		return "eu"
	case 0x002e:
		return "hsb"
	case 0x002f:
		return "mk"
	case 0x0036:
		return "af"
	case 0x0037:
		return "ka"
	case 0x0038:
		return "fo"
	case 0x0039:
		return "hi"
	case 0x003a:
		return "mt"
	case 0x003b:
		return "se"
	case 0x003c:
		return "ga"
	case 0x003e:
		return "ms"
	case 0x003f:
		return "kk"
	case 0x0040:
		return "ky"
	case 0x0041:
		return "sw"
	case 0x0042:
		return "tk"
	case 0x0043:
		return "uz"
	case 0x0044:
		return "tt"
	case 0x0045:
		return "bn"
	case 0x0046:
		return "pa"
	case 0x0047:
		return "gu"
	case 0x0048:
		return "or"
	case 0x0049:
		return "ta"
	case 0x004a:
		return "te"
	case 0x004b:
		return "kn"
	case 0x004c:
		return "ml"
	case 0x004d:
		return "as"
	case 0x004e:
		return "mr"
	case 0x004f:
		return "sa"
	case 0x0050:
		return "mn"
	case 0x0051:
		return "bo"
	case 0x0052:
		return "cy"
	case 0x0053:
		return "km"
	case 0x0054:
		return "lo"
	case 0x0056:
		return "gl"
	case 0x0057:
		return "kok"
	case 0x005a:
		return "syr"
	case 0x005b:
		return "si"
	case 0x005c:
		return "chr"
	case 0x005d:
		return "iu"
	case 0x005e:
		return "am"
	case 0x005f:
		return "tzm"
	case 0x0061:
		return "ne"
	case 0x0062:
		return "fy"
	case 0x0063:
		return "ps"
	case 0x0064:
		return "fil"
	case 0x0065:
		return "dv"
	case 0x0067:
		return "ff"
	case 0x0068:
		return "ha"
	case 0x006a:
		return "yo"
	case 0x006b:
		return "quz"
	case 0x006c:
		return "nso"
	case 0x006d:
		return "ba"
	case 0x006e:
		return "lb"
	case 0x006f:
		return "kl"
	case 0x0070:
		return "ig"
	case 0x0073:
		return "ti"
	case 0x0078:
		return "ii"
	case 0x007a:
		return "arn"
	case 0x007e:
		return "br"
	case 0x0080:
		return "ug"
	case 0x0081:
		return "mi"
	case 0x0082:
		return "oc"
	case 0x0083:
		return "co"
	case 0x0084:
		return "gsw"
	case 0x0085:
		return "sah"
	case 0x0087:
		return "rw"
	case 0x0088:
		return "wo"
	case 0x008c:
		return "prs"
	case 0x0091:
		return "gd"
	case 0x0092:
		return "ku"
	case 0x0401:
		return "ar" // ar-SA
	case 0x0402:
		return "bg" // bg-BG
	case 0x0403:
		return "ca" // ca-ES
	case 0x0404:
		return "zh-TW"
	case 0x0405:
		return "cs" // cs-CZ
	case 0x0406:
		return "da" // da-DK
	case 0x0407:
		return "de" // de-DE
	case 0x0408:
		return "el" // el-GR
	case 0x0409:
		return "en" // en-US
	case 0x040a:
		return "es" // es-ES_tradnl
	case 0x040b:
		return "fi" // fi-FI
	case 0x040c:
		return "fr" // fr-FR
	case 0x040d:
		return "he" // he-IL
	case 0x040e:
		return "hu" // hu-HU
	case 0x040f:
		return "is" // is-IS
	case 0x0410:
		return "it" // it-IT
	case 0x0411:
		return "ja" // ja-JP
	case 0x0412:
		return "ko" // ko-KR
	case 0x0413:
		return "nl" // nl-NL
	case 0x0414:
		return "nb" // nb-NO
	case 0x0415:
		return "pl" // pl-PL
	case 0x0416:
		return "pt" // pt-BR
	case 0x0417:
		return "rm" // rm-CH
	case 0x0418:
		return "ro" // ro-RO
	case 0x0419:
		return "ru" // ru-RU
	case 0x041a:
		return "hr" // hr-HR
	case 0x041b:
		return "sk" // sk-SK
	case 0x041c:
		return "sq" // sq-AL
	case 0x041d:
		return "sv" // sv-SE
	case 0x041e:
		return "th" // th-TH
	case 0x041f:
		return "tr" // tr-TR
	case 0x0420:
		return "ur" // ur-PK
	case 0x0421:
		return "id" // id-ID
	case 0x0422:
		return "uk" // uk-UA
	case 0x0423:
		return "be" // be-BY
	case 0x0424:
		return "sl" // sl-SI
	case 0x0425:
		return "et" // et-EE
	case 0x0426:
		return "lv" // lv-LV
	case 0x0427:
		return "lt" // lt-LT
	case 0x0428:
		return "tg" // tg-Cyrl-TJ
	case 0x0429:
		return "fa" // fa-IR
	case 0x042a:
		return "vi" // vi-VN
	case 0x042b:
		return "hy" // hy-AM
	case 0x042c:
		return "az" // az-Latn-AZ
	case 0x042d:
		return "eu" // eu-ES
	case 0x042e:
		return "hsb" // hsb-DE
	case 0x042f:
		return "mk" // mk-MK
	case 0x0436:
		return "af" // af-ZA
	case 0x0437:
		return "ka" // ka-GE
	case 0x0438:
		return "fo" // fo-FO
	case 0x0439:
		return "hi" // hi-IN
	case 0x043a:
		return "mt" // mt-MT
	case 0x043b:
		return "se" // se-NO
	case 0x043e:
		return "ms" // ms-MY
	case 0x043f:
		return "kk" // kk-KZ
	case 0x0440:
		return "ky" // ky-KG
	case 0x0441:
		return "sw" // sw-KE
	case 0x0442:
		return "tk" // tk-TM
	case 0x0443:
		return "uz" // uz-Latn-UZ
	case 0x0444:
		return "tt" // tt-RU
	case 0x0445:
		return "bn" // bn-IN
	case 0x0446:
		return "pa" // pa-IN
	case 0x0447:
		return "gu" // gu-IN
	case 0x0448:
		return "or" // or-IN
	case 0x0449:
		return "ta" // ta-IN
	case 0x044a:
		return "te" // te-IN
	case 0x044b:
		return "kn" // kn-IN
	case 0x044c:
		return "ml" // ml-IN
	case 0x044d:
		return "as" // as-IN
	case 0x044e:
		return "mr" // mr-IN
	case 0x044f:
		return "sa" // sa-IN
	case 0x0450:
		return "mn" // mn-MN
	case 0x0451:
		return "bo" // bo-CN
	case 0x0452:
		return "cy" // cy-GB
	case 0x0453:
		return "km" // km-KH
	case 0x0454:
		return "lo" // lo-LA
	case 0x0456:
		return "gl" // gl-ES
	case 0x0457:
		return "kok" // kok-IN
	case 0x045a:
		return "syr" // syr-SY
	case 0x045b:
		return "si" // si-LK
	case 0x045c:
		return "chr" // chr-Cher-US
	case 0x045d:
		return "iu" // iu-Cans-CA
	case 0x045e:
		return "am" // am-ET
	case 0x0461:
		return "ne" // ne-NP
	case 0x0462:
		return "fy" // fy-NL
	case 0x0463:
		return "ps" // ps-AF
	case 0x0464:
		return "fil" // fil-PH
	case 0x0465:
		return "dv" // dv-MV
	case 0x0467:
		return "ff" // ff-NG
	case 0x0468:
		return "ha" // ha-Latn-NG
	case 0x046a:
		return "yo" // yo-NG
	case 0x046b:
		return "quz" // quz-BO
	case 0x046c:
		return "nso" // nso-ZA
	case 0x046d:
		return "ba" // ba-RU
	case 0x046e:
		return "lb" // lb-LU
	case 0x046f:
		return "kl" // kl-GL
	case 0x0470:
		return "ig" // ig-NG
	case 0x0473:
		return "ti" // ti-ET
	case 0x0475:
		return "haw" // haw-US
	case 0x0478:
		return "ii" // ii-CN
	case 0x047a:
		return "arn" // arn-CL
	case 0x047c:
		return "moh" // moh-CA
	case 0x047e:
		return "br" // br-FR
	case 0x0480:
		return "ug" // ug-CN
	case 0x0481:
		return "mi" // mi-NZ
	case 0x0482:
		return "oc" // oc-FR
	case 0x0483:
		return "co" // co-FR
	case 0x0484:
		return "gsw" // gsw-FR
	case 0x0485:
		return "sah" // sah-RU
	case 0x0487:
		return "rw" // rw-RW
	case 0x0488:
		return "wo" // wo-SN
	case 0x048c:
		return "prs" // prs-AF
	case 0x0491:
		return "gd" // gd-GB
	case 0x0492:
		return "ku" // ku-Arab-IQ
	case 0x0801:
		return "ar" // ar-IQ
	case 0x0804:
		return "zh-CN"
	case 0x0807:
		return "de" // de-CH
	case 0x0809:
		return "en" // en-GB
	case 0x080a:
		return "es" // es-MX
	case 0x080c:
		return "fr" // fr-BE
	case 0x0810:
		return "it" // it-CH
	case 0x0813:
		return "nl" // nl-BE
	case 0x0814:
		return "nn" // nn-NO
	case 0x0816:
		return "pt" // pt-PT
	case 0x081a:
		return "sr" // sr-Latn-CS
	case 0x081d:
		return "sv" // sv-FI
	case 0x082c:
		return "az" // az-Cyrl-AZ
	case 0x082e:
		return "dsb" // dsb-DE
	case 0x083b:
		return "se" // se-SE
	case 0x083c:
		return "ga" // ga-IE
	case 0x083e:
		return "ms" // ms-BN
	case 0x0843:
		return "uz" // uz-Cyrl-UZ
	case 0x0845:
		return "bn" // bn-BD
	case 0x0846:
		return "pa" // pa-Arab-PK
	case 0x0849:
		return "ta" // ta-LK
	case 0x0850:
		return "mn" // mn-Mong-CN
	case 0x0859:
		return "sd" // sd-Arab-PK
	case 0x085d:
		return "iu" // iu-Latn-CA
	case 0x085f:
		return "tzm" // tzm-Latn-DZ
	case 0x0861:
		return "ne" // ne-IN
	case 0x0867:
		return "ff" // ff-Latn-SN
	case 0x086b:
		return "quz" // quz-EC
	case 0x0873:
		return "ti" // ti-ER
	case 0x0c01:
		return "ar" // ar-EG
	case 0x0c04:
		return "zh-HK"
	case 0x0c07:
		return "de" // de-AT
	case 0x0c09:
		return "en" // en-AU
	case 0x0c0a:
		return "es" // es-ES
	case 0x0c0c:
		return "fr" // fr-CA
	case 0x0c1a:
		return "sr" // sr-Cyrl-CS
	case 0x0c3b:
		return "se" // se-FI
	case 0x0c51:
		return "dz" // dz-BT
	case 0x0c6b:
		return "quz" // quz-PE
	case 0x1001:
		return "ar" // ar-LY
	case 0x1004:
		return "zh-SG"
	case 0x1007:
		return "de" // de-LU
	case 0x1009:
		return "en" // en-CA
	case 0x100a:
		return "es" // es-GT
	case 0x100c:
		return "fr" // fr-CH
	case 0x101a:
		return "hr" // hr-BA
	case 0x103b:
		return "smj" // smj-NO
	case 0x1401:
		return "ar" // ar-DZ
	case 0x1404:
		return "zh-MO"
	case 0x1407:
		return "de" // de-LI
	case 0x1409:
		return "en" // en-NZ
	case 0x140a:
		return "es" // es-CR
	case 0x140c:
		return "fr" // fr-LU
	case 0x141a:
		return "bs" // bs-Latn-BA
	case 0x143b:
		return "smj" // smj-SE
	case 0x1801:
		return "ar" // ar-MA
	case 0x1809:
		return "en" // en-IE
	case 0x180a:
		return "es" // es-PA
	case 0x180c:
		return "fr" // fr-MC
	case 0x181a:
		return "sr" // sr-Latn-BA
	case 0x183b:
		return "sma" // sma-NO
	case 0x1c01:
		return "ar" // ar-TN
	case 0x1c09:
		return "en" // en-ZA
	case 0x1c0a:
		return "es" // es-DO
	case 0x1c1a:
		return "sr" // sr-Cyrl-BA
	case 0x1c3b:
		return "sma" // sma-SE
	case 0x2001:
		return "ar" // ar-OM
	case 0x2009:
		return "en" // en-JM
	case 0x200a:
		return "es" // es-VE
	case 0x201a:
		return "bs" // bs-Cyrl-BA
	case 0x203b:
		return "sms" // sms-FI
	case 0x2401:
		return "ar" // ar-YE
	case 0x2409:
		return "en" // en-029
	case 0x240a:
		return "es" // es-CO
	case 0x240c:
		return "fr" // fr-CD
	case 0x241a:
		return "sr" // sr-Latn-RS
	case 0x243b:
		return "smn" // smn-FI
	case 0x2801:
		return "ar" // ar-SY
	case 0x2809:
		return "en" // en-BZ
	case 0x280a:
		return "es" // es-PE
	case 0x280c:
		return "fr" // fr-SN
	case 0x281a:
		return "sr" // sr-Cyrl-RS
	case 0x2c01:
		return "ar" // ar-JO
	case 0x2c09:
		return "en" // en-TT
	case 0x2c0a:
		return "es" // es-AR
	case 0x2c0c:
		return "fr" // fr-CM
	case 0x2c1a:
		return "sr" // sr-Latn-ME
	case 0x3001:
		return "ar" // ar-LB
	case 0x3009:
		return "en" // en-ZW
	case 0x300a:
		return "es" // es-EC
	case 0x300c:
		return "fr" // fr-CI
	case 0x301a:
		return "sr" // sr-Cyrl-ME
	case 0x3401:
		return "ar" // ar-KW
	case 0x3409:
		return "en" // en-PH
	case 0x340a:
		return "es" // es-CL
	case 0x340c:
		return "fr" // fr-ML
	case 0x3801:
		return "ar" // ar-AE
	case 0x380a:
		return "es" // es-UY
	case 0x380c:
		return "fr" // fr-MA
	case 0x3c01:
		return "ar" // ar-BH
	case 0x3c09:
		return "en" // en-HK
	case 0x3c0a:
		return "es" // es-PY
	case 0x3c0c:
		return "fr" // fr-HT
	case 0x4001:
		return "ar" // ar-QA
	case 0x4009:
		return "en" // en-IN
	case 0x400a:
		return "es" // es-BO
	case 0x4409:
		return "en" // en-MY
	case 0x440a:
		return "es" // es-SV
	case 0x4809:
		return "en" // en-SG
	case 0x480a:
		return "es" // es-HN
	case 0x4c09:
		return "en" // en-AE
	case 0x4c0a:
		return "es" // es-NI
	case 0x500a:
		return "es" // es-PR
	case 0x540a:
		return "es" // es-US
	case 0x580a:
		return "es" // es-419
	case 0x5c0a:
		return "es" // es-CU
	case 0x7c04:
		return "zh-TW" // zh-Hant
	case 0x7c14:
		return "nb"
	case 0x7c1a:
		return "sr"
	case 0x7c28:
		return "tg" // tg-Cyrl
	case 0x7c2e:
		return "dsb"
	case 0x7c3b:
		return "smj"
	case 0x7c43:
		return "uz" // uz-Latn
	case 0x7c46:
		return "pa" // pa-Arab
	case 0x7c50:
		return "mn" // mn-Mong
	case 0x7c59:
		return "sd" // sd-Arab
	case 0x7c5c:
		return "chr" // chr-Cher
	case 0x7c5d:
		return "iu" // iu-Latn
	case 0x7c5f:
		return "tzm" // tzm-Latn
	case 0x7c67:
		return "ff" // ff-Latn
	case 0x7c68:
		return "ha" // ha-Latn
	case 0x7c92:
		return "ku" // ku-Arab
	default:
		return "unknown"
	}
}
