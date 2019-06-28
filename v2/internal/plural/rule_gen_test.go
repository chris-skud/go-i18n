// This file is generated by i18n/plural/codegen/generate.sh; DO NOT EDIT

package plural

import "testing"

func TestBmBoDzIdIgIiInJaJboJvJwKdeKeaKmKoLktLoMsMyNqoRootSahSesSgThToViWoYoYueZh(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, Other, []string{"0~15", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"bm", "bo", "dz", "id", "ig", "ii", "in", "ja", "jbo", "jv", "jw", "kde", "kea", "km", "ko", "lkt", "lo", "ms", "my", "nqo", "root", "sah", "ses", "sg", "th", "to", "vi", "wo", "yo", "yue", "zh"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestAmAsBnFaGuHiKnZu(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"0", "1"})
	tests = appendDecimalTests(tests, One, []string{"0.0~1.0", "0.00~0.04"})

	tests = appendIntegerTests(tests, Other, []string{"2~17", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"1.1~2.6", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"am", "as", "bn", "fa", "gu", "hi", "kn", "zu"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestFfFrHyKab(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"0", "1"})
	tests = appendDecimalTests(tests, One, []string{"0.0~1.5"})

	tests = appendIntegerTests(tests, Other, []string{"2~17", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"2.0~3.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"ff", "fr", "hy", "kab"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestPt(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"0", "1"})
	tests = appendDecimalTests(tests, One, []string{"0.0~1.5"})

	tests = appendIntegerTests(tests, Other, []string{"2~17", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"2.0~3.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"pt"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestAstCaDeEnEtFiFyGlIaIoItJiNlPt_PTScScnSvSwUrYi(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})

	tests = appendIntegerTests(tests, Other, []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"ast", "ca", "de", "en", "et", "fi", "fy", "gl", "ia", "io", "it", "ji", "nl", "pt_PT", "sc", "scn", "sv", "sw", "ur", "yi"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestSi(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"0", "1"})
	tests = appendDecimalTests(tests, One, []string{"0.0", "0.1", "1.0", "0.00", "0.01", "1.00", "0.000", "0.001", "1.000", "0.0000", "0.0001", "1.0000"})

	tests = appendIntegerTests(tests, Other, []string{"2~17", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.2~0.9", "1.1~1.8", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"si"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestAkBhGuwLnMgNsoPaTiWa(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"0", "1"})
	tests = appendDecimalTests(tests, One, []string{"0.0", "1.0", "0.00", "1.00", "0.000", "1.000", "0.0000", "1.0000"})

	tests = appendIntegerTests(tests, Other, []string{"2~17", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.1~0.9", "1.1~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"ak", "bh", "guw", "ln", "mg", "nso", "pa", "ti", "wa"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestTzm(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"0", "1", "11~24"})
	tests = appendDecimalTests(tests, One, []string{"0.0", "1.0", "11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "17.0", "18.0", "19.0", "20.0", "21.0", "22.0", "23.0", "24.0"})

	tests = appendIntegerTests(tests, Other, []string{"2~10", "100~106", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.1~0.9", "1.1~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"tzm"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestAfAsaAzBemBezBgBrxCeCggChrCkbDvEeElEoEsEuFoFurGswHaHawHuJgoJmcKaKajKcgKkKkjKlKsKsbKuKyLbLgMasMgoMlMnMrNahNbNdNeNnNnhNoNrNyNynOmOrOsPapPsRmRofRwkSaqSdSdhSehSnSoSqSsSsyStSyrTaTeTeoTigTkTnTrTsUgUzVeVoVunWaeXhXog(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "1.00", "1.000", "1.0000"})

	tests = appendIntegerTests(tests, Other, []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0~0.9", "1.1~1.6", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"af", "asa", "az", "bem", "bez", "bg", "brx", "ce", "cgg", "chr", "ckb", "dv", "ee", "el", "eo", "es", "eu", "fo", "fur", "gsw", "ha", "haw", "hu", "jgo", "jmc", "ka", "kaj", "kcg", "kk", "kkj", "kl", "ks", "ksb", "ku", "ky", "lb", "lg", "mas", "mgo", "ml", "mn", "mr", "nah", "nb", "nd", "ne", "nn", "nnh", "no", "nr", "ny", "nyn", "om", "or", "os", "pap", "ps", "rm", "rof", "rwk", "saq", "sd", "sdh", "seh", "sn", "so", "sq", "ss", "ssy", "st", "syr", "ta", "te", "teo", "tig", "tk", "tn", "tr", "ts", "ug", "uz", "ve", "vo", "vun", "wae", "xh", "xog"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestDa(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"0.1~1.6"})

	tests = appendIntegerTests(tests, Other, []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0", "2.0~3.4", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"da"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestIs(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"})
	tests = appendDecimalTests(tests, One, []string{"0.1~1.6", "10.1", "100.1", "1000.1"})

	tests = appendIntegerTests(tests, Other, []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0", "2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"is"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestMk(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"})
	tests = appendDecimalTests(tests, One, []string{"0.1", "1.1", "2.1", "3.1", "4.1", "5.1", "6.1", "7.1", "10.1", "100.1", "1000.1"})

	tests = appendIntegerTests(tests, Other, []string{"0", "2~16", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0", "0.2~1.0", "1.2~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"mk"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestCebFilTl(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"0~3", "5", "7", "8", "10~13", "15", "17", "18", "20", "21", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, One, []string{"0.0~0.3", "0.5", "0.7", "0.8", "1.0~1.3", "1.5", "1.7", "1.8", "2.0", "2.1", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	tests = appendIntegerTests(tests, Other, []string{"4", "6", "9", "14", "16", "19", "24", "26", "104", "1004"})
	tests = appendDecimalTests(tests, Other, []string{"0.4", "0.6", "0.9", "1.4", "1.6", "1.9", "2.4", "2.6", "10.4", "100.4", "1000.4"})

	locales := []string{"ceb", "fil", "tl"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestLvPrg(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, Zero, []string{"0", "10~20", "30", "40", "50", "60", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Zero, []string{"0.0", "10.0", "11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	tests = appendIntegerTests(tests, One, []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"})
	tests = appendDecimalTests(tests, One, []string{"0.1", "1.0", "1.1", "2.1", "3.1", "4.1", "5.1", "6.1", "7.1", "10.1", "100.1", "1000.1"})

	tests = appendIntegerTests(tests, Other, []string{"2~9", "22~29", "102", "1002"})
	tests = appendDecimalTests(tests, Other, []string{"0.2~0.9", "1.2~1.9", "10.2", "100.2", "1000.2"})

	locales := []string{"lv", "prg"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestLag(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, Zero, []string{"0"})
	tests = appendDecimalTests(tests, Zero, []string{"0.0", "0.00", "0.000", "0.0000"})

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"0.1~1.6"})

	tests = appendIntegerTests(tests, Other, []string{"2~17", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"2.0~3.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"lag"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestKsh(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, Zero, []string{"0"})
	tests = appendDecimalTests(tests, Zero, []string{"0.0", "0.00", "0.000", "0.0000"})

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "1.00", "1.000", "1.0000"})

	tests = appendIntegerTests(tests, Other, []string{"2~17", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.1~0.9", "1.1~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"ksh"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestIuNaqSeSmaSmiSmjSmnSms(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "1.00", "1.000", "1.0000"})

	tests = appendIntegerTests(tests, Two, []string{"2"})
	tests = appendDecimalTests(tests, Two, []string{"2.0", "2.00", "2.000", "2.0000"})

	tests = appendIntegerTests(tests, Other, []string{"0", "3~17", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0~0.9", "1.1~1.6", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"iu", "naq", "se", "sma", "smi", "smj", "smn", "sms"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestShi(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"0", "1"})
	tests = appendDecimalTests(tests, One, []string{"0.0~1.0", "0.00~0.04"})

	tests = appendIntegerTests(tests, Few, []string{"2~10"})
	tests = appendDecimalTests(tests, Few, []string{"2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "9.0", "10.0", "2.00", "3.00", "4.00", "5.00", "6.00", "7.00", "8.00"})

	tests = appendIntegerTests(tests, Other, []string{"11~26", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"1.1~1.9", "2.1~2.7", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"shi"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestMoRo(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})

	tests = appendIntegerTests(tests, Few, []string{"0", "2~16", "102", "1002"})
	tests = appendDecimalTests(tests, Few, []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	tests = appendIntegerTests(tests, Other, []string{"20~35", "100", "1000", "10000", "100000", "1000000"})

	locales := []string{"mo", "ro"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestBsHrShSr(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"})
	tests = appendDecimalTests(tests, One, []string{"0.1", "1.1", "2.1", "3.1", "4.1", "5.1", "6.1", "7.1", "10.1", "100.1", "1000.1"})

	tests = appendIntegerTests(tests, Few, []string{"2~4", "22~24", "32~34", "42~44", "52~54", "62", "102", "1002"})
	tests = appendDecimalTests(tests, Few, []string{"0.2~0.4", "1.2~1.4", "2.2~2.4", "3.2~3.4", "4.2~4.4", "5.2", "10.2", "100.2", "1000.2"})

	tests = appendIntegerTests(tests, Other, []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0", "0.5~1.0", "1.5~2.0", "2.5~2.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"bs", "hr", "sh", "sr"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestGd(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "11"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "11.0", "1.00", "11.00", "1.000", "11.000", "1.0000"})

	tests = appendIntegerTests(tests, Two, []string{"2", "12"})
	tests = appendDecimalTests(tests, Two, []string{"2.0", "12.0", "2.00", "12.00", "2.000", "12.000", "2.0000"})

	tests = appendIntegerTests(tests, Few, []string{"3~10", "13~19"})
	tests = appendDecimalTests(tests, Few, []string{"3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "9.0", "10.0", "13.0", "14.0", "15.0", "16.0", "17.0", "18.0", "19.0", "3.00"})

	tests = appendIntegerTests(tests, Other, []string{"0", "20~34", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0~0.9", "1.1~1.6", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"gd"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestSl(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "101", "201", "301", "401", "501", "601", "701", "1001"})

	tests = appendIntegerTests(tests, Two, []string{"2", "102", "202", "302", "402", "502", "602", "702", "1002"})

	tests = appendIntegerTests(tests, Few, []string{"3", "4", "103", "104", "203", "204", "303", "304", "403", "404", "503", "504", "603", "604", "703", "704", "1003"})
	tests = appendDecimalTests(tests, Few, []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	tests = appendIntegerTests(tests, Other, []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"})

	locales := []string{"sl"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestDsbHsb(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "101", "201", "301", "401", "501", "601", "701", "1001"})
	tests = appendDecimalTests(tests, One, []string{"0.1", "1.1", "2.1", "3.1", "4.1", "5.1", "6.1", "7.1", "10.1", "100.1", "1000.1"})

	tests = appendIntegerTests(tests, Two, []string{"2", "102", "202", "302", "402", "502", "602", "702", "1002"})
	tests = appendDecimalTests(tests, Two, []string{"0.2", "1.2", "2.2", "3.2", "4.2", "5.2", "6.2", "7.2", "10.2", "100.2", "1000.2"})

	tests = appendIntegerTests(tests, Few, []string{"3", "4", "103", "104", "203", "204", "303", "304", "403", "404", "503", "504", "603", "604", "703", "704", "1003"})
	tests = appendDecimalTests(tests, Few, []string{"0.3", "0.4", "1.3", "1.4", "2.3", "2.4", "3.3", "3.4", "4.3", "4.4", "5.3", "5.4", "6.3", "6.4", "7.3", "7.4", "10.3", "100.3", "1000.3"})

	tests = appendIntegerTests(tests, Other, []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0", "0.5~1.0", "1.5~2.0", "2.5~2.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"dsb", "hsb"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestHeIw(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})

	tests = appendIntegerTests(tests, Two, []string{"2"})

	tests = appendIntegerTests(tests, Many, []string{"20", "30", "40", "50", "60", "70", "80", "90", "100", "1000", "10000", "100000", "1000000"})

	tests = appendIntegerTests(tests, Other, []string{"0", "3~17", "101", "1001"})
	tests = appendDecimalTests(tests, Other, []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"he", "iw"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestCsSk(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})

	tests = appendIntegerTests(tests, Few, []string{"2~4"})

	tests = appendDecimalTests(tests, Many, []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	tests = appendIntegerTests(tests, Other, []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"})

	locales := []string{"cs", "sk"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestPl(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})

	tests = appendIntegerTests(tests, Few, []string{"2~4", "22~24", "32~34", "42~44", "52~54", "62", "102", "1002"})

	tests = appendIntegerTests(tests, Many, []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"})

	tests = appendDecimalTests(tests, Other, []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"pl"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestBe(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "21.0", "31.0", "41.0", "51.0", "61.0", "71.0", "81.0", "101.0", "1001.0"})

	tests = appendIntegerTests(tests, Few, []string{"2~4", "22~24", "32~34", "42~44", "52~54", "62", "102", "1002"})
	tests = appendDecimalTests(tests, Few, []string{"2.0", "3.0", "4.0", "22.0", "23.0", "24.0", "32.0", "33.0", "102.0", "1002.0"})

	tests = appendIntegerTests(tests, Many, []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Many, []string{"0.0", "5.0", "6.0", "7.0", "8.0", "9.0", "10.0", "11.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	tests = appendDecimalTests(tests, Other, []string{"0.1~0.9", "1.1~1.7", "10.1", "100.1", "1000.1"})

	locales := []string{"be"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestLt(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "21.0", "31.0", "41.0", "51.0", "61.0", "71.0", "81.0", "101.0", "1001.0"})

	tests = appendIntegerTests(tests, Few, []string{"2~9", "22~29", "102", "1002"})
	tests = appendDecimalTests(tests, Few, []string{"2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "9.0", "22.0", "102.0", "1002.0"})

	tests = appendDecimalTests(tests, Many, []string{"0.1~0.9", "1.1~1.7", "10.1", "100.1", "1000.1"})

	tests = appendIntegerTests(tests, Other, []string{"0", "10~20", "30", "40", "50", "60", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0", "10.0", "11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"lt"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestMt(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "1.00", "1.000", "1.0000"})

	tests = appendIntegerTests(tests, Few, []string{"0", "2~10", "102~107", "1002"})
	tests = appendDecimalTests(tests, Few, []string{"0.0", "2.0", "3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "10.0", "102.0", "1002.0"})

	tests = appendIntegerTests(tests, Many, []string{"11~19", "111~117", "1011"})
	tests = appendDecimalTests(tests, Many, []string{"11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "17.0", "18.0", "111.0", "1011.0"})

	tests = appendIntegerTests(tests, Other, []string{"20~35", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.1~0.9", "1.1~1.7", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"mt"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestRuUk(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "21", "31", "41", "51", "61", "71", "81", "101", "1001"})

	tests = appendIntegerTests(tests, Few, []string{"2~4", "22~24", "32~34", "42~44", "52~54", "62", "102", "1002"})

	tests = appendIntegerTests(tests, Many, []string{"0", "5~19", "100", "1000", "10000", "100000", "1000000"})

	tests = appendDecimalTests(tests, Other, []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"ru", "uk"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestBr(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "21", "31", "41", "51", "61", "81", "101", "1001"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "21.0", "31.0", "41.0", "51.0", "61.0", "81.0", "101.0", "1001.0"})

	tests = appendIntegerTests(tests, Two, []string{"2", "22", "32", "42", "52", "62", "82", "102", "1002"})
	tests = appendDecimalTests(tests, Two, []string{"2.0", "22.0", "32.0", "42.0", "52.0", "62.0", "82.0", "102.0", "1002.0"})

	tests = appendIntegerTests(tests, Few, []string{"3", "4", "9", "23", "24", "29", "33", "34", "39", "43", "44", "49", "103", "1003"})
	tests = appendDecimalTests(tests, Few, []string{"3.0", "4.0", "9.0", "23.0", "24.0", "29.0", "33.0", "34.0", "103.0", "1003.0"})

	tests = appendIntegerTests(tests, Many, []string{"1000000"})
	tests = appendDecimalTests(tests, Many, []string{"1000000.0", "1000000.00", "1000000.000"})

	tests = appendIntegerTests(tests, Other, []string{"0", "5~8", "10~20", "100", "1000", "10000", "100000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0~0.9", "1.1~1.6", "10.0", "100.0", "1000.0", "10000.0", "100000.0"})

	locales := []string{"br"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestGa(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "1.00", "1.000", "1.0000"})

	tests = appendIntegerTests(tests, Two, []string{"2"})
	tests = appendDecimalTests(tests, Two, []string{"2.0", "2.00", "2.000", "2.0000"})

	tests = appendIntegerTests(tests, Few, []string{"3~6"})
	tests = appendDecimalTests(tests, Few, []string{"3.0", "4.0", "5.0", "6.0", "3.00", "4.00", "5.00", "6.00", "3.000", "4.000", "5.000", "6.000", "3.0000", "4.0000", "5.0000", "6.0000"})

	tests = appendIntegerTests(tests, Many, []string{"7~10"})
	tests = appendDecimalTests(tests, Many, []string{"7.0", "8.0", "9.0", "10.0", "7.00", "8.00", "9.00", "10.00", "7.000", "8.000", "9.000", "10.000", "7.0000", "8.0000", "9.0000", "10.0000"})

	tests = appendIntegerTests(tests, Other, []string{"0", "11~25", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.0~0.9", "1.1~1.6", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"ga"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestGv(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, One, []string{"1", "11", "21", "31", "41", "51", "61", "71", "101", "1001"})

	tests = appendIntegerTests(tests, Two, []string{"2", "12", "22", "32", "42", "52", "62", "72", "102", "1002"})

	tests = appendIntegerTests(tests, Few, []string{"0", "20", "40", "60", "80", "100", "120", "140", "1000", "10000", "100000", "1000000"})

	tests = appendDecimalTests(tests, Many, []string{"0.0~1.5", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	tests = appendIntegerTests(tests, Other, []string{"3~10", "13~19", "23", "103", "1003"})

	locales := []string{"gv"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestArArs(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, Zero, []string{"0"})
	tests = appendDecimalTests(tests, Zero, []string{"0.0", "0.00", "0.000", "0.0000"})

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "1.00", "1.000", "1.0000"})

	tests = appendIntegerTests(tests, Two, []string{"2"})
	tests = appendDecimalTests(tests, Two, []string{"2.0", "2.00", "2.000", "2.0000"})

	tests = appendIntegerTests(tests, Few, []string{"3~10", "103~110", "1003"})
	tests = appendDecimalTests(tests, Few, []string{"3.0", "4.0", "5.0", "6.0", "7.0", "8.0", "9.0", "10.0", "103.0", "1003.0"})

	tests = appendIntegerTests(tests, Many, []string{"11~26", "111", "1011"})
	tests = appendDecimalTests(tests, Many, []string{"11.0", "12.0", "13.0", "14.0", "15.0", "16.0", "17.0", "18.0", "111.0", "1011.0"})

	tests = appendIntegerTests(tests, Other, []string{"100~102", "200~202", "300~302", "400~402", "500~502", "600", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.1~0.9", "1.1~1.7", "10.1", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"ar", "ars"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestCy(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, Zero, []string{"0"})
	tests = appendDecimalTests(tests, Zero, []string{"0.0", "0.00", "0.000", "0.0000"})

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "1.00", "1.000", "1.0000"})

	tests = appendIntegerTests(tests, Two, []string{"2"})
	tests = appendDecimalTests(tests, Two, []string{"2.0", "2.00", "2.000", "2.0000"})

	tests = appendIntegerTests(tests, Few, []string{"3"})
	tests = appendDecimalTests(tests, Few, []string{"3.0", "3.00", "3.000", "3.0000"})

	tests = appendIntegerTests(tests, Many, []string{"6"})
	tests = appendDecimalTests(tests, Many, []string{"6.0", "6.00", "6.000", "6.0000"})

	tests = appendIntegerTests(tests, Other, []string{"4", "5", "7~20", "100", "1000", "10000", "100000", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.1~0.9", "1.1~1.7", "10.0", "100.0", "1000.0", "10000.0", "100000.0", "1000000.0"})

	locales := []string{"cy"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}

func TestKw(t *testing.T) {
	var tests []pluralFormTest

	tests = appendIntegerTests(tests, Zero, []string{"0"})
	tests = appendDecimalTests(tests, Zero, []string{"0.0", "0.00", "0.000", "0.0000"})

	tests = appendIntegerTests(tests, One, []string{"1"})
	tests = appendDecimalTests(tests, One, []string{"1.0", "1.00", "1.000", "1.0000"})

	tests = appendIntegerTests(tests, Two, []string{"2", "22", "42", "62", "82", "102", "122", "142", "1002"})
	tests = appendDecimalTests(tests, Two, []string{"2.0", "22.0", "42.0", "62.0", "82.0", "102.0", "122.0", "142.0", "1002.0"})

	tests = appendIntegerTests(tests, Few, []string{"3", "23", "43", "63", "83", "103", "123", "143", "1003"})
	tests = appendDecimalTests(tests, Few, []string{"3.0", "23.0", "43.0", "63.0", "83.0", "103.0", "123.0", "143.0", "1003.0"})

	tests = appendIntegerTests(tests, Many, []string{"21", "41", "61", "81", "101", "121", "141", "161", "1001"})
	tests = appendDecimalTests(tests, Many, []string{"21.0", "41.0", "61.0", "81.0", "101.0", "121.0", "141.0", "161.0", "1001.0"})

	tests = appendIntegerTests(tests, Other, []string{"4~19", "100", "1000000"})
	tests = appendDecimalTests(tests, Other, []string{"0.1~0.9", "1.1~1.7", "10.0", "100.0", "1000000.0"})

	locales := []string{"kw"}
	for _, locale := range locales {
		runTests(t, locale, tests)
	}
}
