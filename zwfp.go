package zwfp

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ZWSP = '\u200B' /** Zero-width space **/
	ZWNJ = '\u200C' /** Zero-width non-joiner **/
	ZWJ  = '\u200D' /** Zero-width joiner **/
	ZWNB = '\uFEFF' /**	Zero-width no-break space **/
)

func toZeroWidth(keyInfo string) []rune {
	var result []rune
	for i, c := range keyInfo {
		if c == ' ' {
			result = append(result, ZWNB)
			continue
		}
		bits := fmt.Sprintf("%b", c)
		for _, b := range bits {
			if b == '0' {
				result = append(result, ZWNJ)
			} else {
				result = append(result, ZWSP)
			}
		}
		if i != len(keyInfo)-1 {
			result = append(result, ZWJ)
		}
	}
	return result
}

func toString(keyInfo []rune) (string, error) {
	var sb strings.Builder
	var cl strings.Builder
	for _, c := range keyInfo {
		if c == ZWJ || c == ZWNB {
			d, err := strconv.ParseInt(cl.String(), 2, 32)
			if err != nil {
				return "", fmt.Errorf("failed to convert [%v] to letter: %w", cl.String(), err)
			}
			sb.WriteRune(rune(d))
			cl.Reset()
			if c == ZWNB {
				sb.WriteString(" ")
			}
		} else if c == ZWNJ {
			cl.WriteString("0")
		} else if c == ZWSP {
			cl.WriteString("1")
		}
	}
	if cl.Len() > 0 {
		d, err := strconv.ParseInt(cl.String(), 2, 32)
		if err != nil {
			return "", fmt.Errorf("failed to convert [%v] to letter: %w", cl.String(), err)
		}
		sb.WriteRune(rune(d))
	}
	return sb.String(), nil
}

func checkEmbedInput(data, keyInfo string) error {
	if len(data) < 2 {
		return fmt.Errorf("the length of data should be greater than 2 bytes")
	}

	if len(keyInfo) < 1 {
		return fmt.Errorf("the length of keyInfo should be greater than 1 bytes")
	}
	return nil
}

func Embed(data, keyInfo string) (string, error) {
	if err := checkEmbedInput(data, keyInfo); err != nil {
		return "", err
	}

	t := 1
	k := toZeroWidth(keyInfo)
	var embed []rune

	embed = append(embed, rune(data[0]), k[0])
	for i, c := range data[1:] {
		if i == len(data)-2 && t < len(k) {
			embed = append(embed, k[t:]...)
		} else if t < len(k) {
			embed = append(embed, k[t])
			t++
		}
		embed = append(embed, c)
	}
	return string(embed), nil
}

func Extract(embed string) (string, string, error) {
	var data, keyInfo []rune
	for _, v := range embed {
		switch v {
		case ZWNB, ZWNJ, ZWSP, ZWJ:
			keyInfo = append(keyInfo, v)
		default:
			data = append(data, v)
		}
	}
	if key, err := toString(keyInfo); err != nil {
		return string(data), "", err
	} else {
		return string(data), key, nil
	}
}
