package mojicon

import (
	"fmt"
	"unicode/utf8"
)

const (
	HiraganaKatakanaDiff = 96
	HiraganaStart        = 12353
	HiraganaEnd          = HiraganaStart + HiraganaKatakanaDiff - 1
	KataganaStart        = HiraganaStart + HiraganaKatakanaDiff
	KataganaEnd          = KataganaStart + HiraganaKatakanaDiff - 1

	HankakuAsciiStart       = '!'
	HankakuAsciiEnd         = '~'
	ZenkakuAsciiStart       = 65281 //！
	ZenkakuAsciiEnd         = ZenkakuAsciiStart + ('~' - '!') - 1
	HankakuZenkakuAsciiDiff = ZenkakuAsciiStart - HankakuAsciiStart

	//	半角カタカナと全角カタカナの順番は濁点などの影響によりずれているため、単純な置換方法は困難
//	HankakuKatakanaStart       = 65383//	ｧ
//	HankakuKatakanaEnd         = 65439//	ﾟ
//	ZenkakuKatakanaStart       = 65281//！
//	ZenkakuKatakanaEnd         = ZenkakuKatakanaStart + ('~' - '!') - 1
//	HankakuZenkakuKatakanaDiff = ZenkakuKatakanaStart - HankakuKatakanaStart
)

func ConvHiragana(src string) (dst string) {
	buf := make([]rune, utf8.RuneCountInString(src))
	index := 0
	for _, r := range src {
		if KataganaStart <= r && r <= KataganaEnd {
			r -= HiraganaKatakanaDiff
		}
		buf[index] = r
		index++
	}
	dst = string(buf)
	return
}

func ConvKatakana(src string) (dst string) {
	buf := make([]rune, utf8.RuneCountInString(src))
	index := 0
	for _, r := range src {
		if HiraganaStart <= r && r <= HiraganaEnd {
			r += HiraganaKatakanaDiff
		}
		buf[index] = rune(r)
		index++
	}
	dst = string(buf)
	return
}

func ConvAsciiZenkaku(src string) (dst string) {
	buf := make([]rune, utf8.RuneCountInString(src))
	index := 0
	for _, r := range src {
		if HankakuAsciiStart <= r && r <= HankakuAsciiEnd {
			r += HankakuZenkakuAsciiDiff
		}
		buf[index] = rune(r)
		index++
	}
	dst = string(buf)
	return
}

func ConvAsciiHankaku(src string) (dst string) {
	buf := make([]rune, utf8.RuneCountInString(src))
	index := 0
	for _, r := range src {
		if ZenkakuAsciiStart <= r && r <= ZenkakuAsciiEnd {
			r -= HankakuZenkakuAsciiDiff
		}
		buf[index] = rune(r)
		index++
	}
	dst = string(buf)
	return
}

func example() {
	fmt.Println("ひらがな	：", ConvHiragana("ぁあｱァアをヲaAａＡzZｚＺー。"))
	fmt.Println("アスキー半角	：", ConvAsciiHankaku("ぁあｱァアをヲaAａＡzZｚＺー。"))
	fmt.Println("アスキー全角	：", ConvAsciiZenkaku("ぁあｱァアをヲaAａＡzZｚＺー。"))
	fmt.Println("ひらがな	：", ConvHiragana("カタカナをひらがなへへんかんします。"))
	fmt.Println("カタカナ	：", ConvKatakana("ひらがなをカタカナへへんかんします。"))
	//	output
	/*
		ひらがな	： ぁあｱぁあををaAａＡzZｚＺ゜。
		アスキー半角	： ぁあｱァアをヲaAaAzZzZー。
		アスキー全角	： ぁあｱァアをヲａＡａＡｚＺｚＺー。
		ひらがな	： かたかなをひらがなへへんかんします。
		カタカナ	： ヒラガナヲカタカナヘヘンカンシマス。
	*/
}
