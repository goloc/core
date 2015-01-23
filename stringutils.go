package goloc

import (
	"bytes"
	"regexp"
	"strings"
)

var splitSpacePunctRegex = regexp.MustCompile("[[:space:][:punct:]]")

var accentsMap = map[rune]string{
	'À': "A",
	'Á': "A",
	'Â': "A",
	'Ã': "A",
	'Ä': "A",
	'Å': "AA",
	'Æ': "AE",
	'Ç': "C",
	'È': "E",
	'É': "E",
	'Ê': "E",
	'Ë': "E",
	'Ì': "I",
	'Í': "I",
	'Î': "I",
	'Ï': "I",
	'Ð': "D",
	'Ł': "L",
	'Ñ': "N",
	'Ò': "O",
	'Ó': "O",
	'Ô': "O",
	'Õ': "O",
	'Ö': "O",
	'Ø': "OE",
	'Ù': "U",
	'Ú': "U",
	'Ü': "U",
	'Û': "U",
	'Ý': "Y",
	'Þ': "Th",
	'ß': "ss",
	'à': "a",
	'á': "a",
	'â': "a",
	'ã': "a",
	'ä': "a",
	'å': "aa",
	'æ': "ae",
	'ç': "c",
	'è': "e",
	'é': "e",
	'ê': "e",
	'ë': "e",
	'ì': "i",
	'í': "i",
	'î': "i",
	'ï': "i",
	'ð': "d",
	'ł': "l",
	'ñ': "n",
	'ń': "n",
	'ò': "o",
	'ó': "o",
	'ô': "o",
	'õ': "o",
	'ō': "o",
	'ö': "o",
	'ø': "oe",
	'ś': "s",
	'ù': "u",
	'ú': "u",
	'û': "u",
	'ū': "u",
	'ü': "u",
	'ý': "y",
	'þ': "th",
	'ÿ': "y",
	'ż': "z",
	'Œ': "OE",
	'œ': "oe",
}

func splitSpacePunct(s string) []string {
	n := 0
	splited := splitSpacePunctRegex.Split(s, -1)
	for _, s := range splited {
		if len(s) > 0 {
			n++
		}
	}
	r := make([]string, n)
	i := 0
	for _, s := range splited {
		if len(s) > 0 {
			r[i] = s
			i++
		}
	}
	return r
}

func stripAccents(s string) string {
	b := bytes.NewBufferString("")
	for _, c := range s {
		if val, ok := accentsMap[c]; ok {
			b.WriteString(val)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func min(a int, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}
	return
}

func levenshteinDistance(source string, target string) int {
	if source == target {
		return 0
	}
	r1 := []rune(source)
	r2 := []rune(target)
	if len(r1) == 0 {
		return len(r2)
	}
	if len(r2) == 0 {
		return len(r1)
	}
	rows := len(r1) + 1
	cols := len(r2) + 1
	var d1 int
	var d2 int
	var d3 int
	var i int
	var j int
	dist := make([]int, rows*cols)
	for i = 0; i < rows; i++ {
		dist[i*cols] = i
	}
	for j = 0; j < cols; j++ {
		dist[j] = j
	}
	for j = 1; j < cols; j++ {
		for i = 1; i < rows; i++ {
			if r1[i-1] == r2[j-1] {
				dist[(i*cols)+j] = dist[((i-1)*cols)+(j-1)]
			} else {
				d1 = dist[((i-1)*cols)+j] + 1
				d2 = dist[(i*cols)+(j-1)] + 1
				d3 = dist[((i-1)*cols)+(j-1)] + 1
				dist[(i*cols)+j] = min(d1, min(d2, d3))
			}
		}
	}
	return dist[(cols*rows)-1]
}

func partialphone(source string) string {
	r := []rune(strings.ToUpper(stripAccents(source)))

	if len(r) == 0 {
		return ""
	}

	b := bytes.NewBufferString("")
	lastRune := ' '

	for _, currentRune := range r {
		switch currentRune {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			switch lastRune {
			case ' ', 'A', 'E', 'I', 'O', 'U', 'Y':
				b.WriteRune(currentRune)
				lastRune = currentRune
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				lastRune = currentRune
			default:
				if lastRune != 0 {
					b.WriteRune(lastRune)
				}
				lastRune = 0
			}

		case 'A', 'E', 'I', 'O', 'U', 'Y':
			switch lastRune {
			case ' ', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				b.WriteRune('A')
				lastRune = 'A'
			case 'A', 'E', 'I', 'O', 'U', 'Y':
				lastRune = 'A'
			default:
				if lastRune != 0 {
					b.WriteRune(lastRune)
				}
				lastRune = 0
			}

		case 'B':
			lastRune = 'B'

		case 'C':
			lastRune = 'S'

		case 'D':
			lastRune = 'D'

		case 'F':
			lastRune = 'F'

		case 'G':
			lastRune = 'J'

		case 'H':
			if lastRune == 'P' {
				lastRune = 'F'
			}

		case 'J':
			lastRune = 'J'

		case 'K':
			lastRune = 'K'

		case 'L':
			lastRune = 'L'

		case 'M':
			lastRune = 'M'

		case 'N':
			lastRune = 'N'

		case 'P':
			lastRune = 'P'

		case 'Q':
			lastRune = 'K'

		case 'R':
			lastRune = 'R'

		case 'S':
			lastRune = 'S'

		case 'T':
			lastRune = 'S'

		case 'V':
			lastRune = 'F'

		case 'W':
			lastRune = 'A'

		case 'X':
			lastRune = 'S'

		case 'Z':
			lastRune = 'S'

		default:
			switch lastRune {
			case ' ':

			default:
				b.WriteRune(' ')
				lastRune = ' '
			}
		}
	}

	return b.String()
}
