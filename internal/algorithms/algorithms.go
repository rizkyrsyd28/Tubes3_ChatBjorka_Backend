package algorithms

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func computeBorder(pattern string) []int {
	length := len(pattern)
	lps := make([]int, length)
	lps[0] = 0

	i, j := 1, 0

	for i < length {
		if pattern[j] == pattern[i] {
			lps[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = lps[j-1]
		} else {
			lps[i] = 0
			i++
		}
	}

	return lps
}

func kmpSearch(text string, pattern string) int {
	lenText := len(text)
	lenPattern := len(pattern)

	lps := computeBorder(pattern)

	i, j := 0, 0

	for i < lenText {
		if pattern[j] == text[i] {
			if j == lenPattern-1 {
				return i - lenPattern + 1
				// fmt.Println("Found at", i-lenPattern+1)
				// j = lps[j-1]
			}
			i++
			j++
		} else if j > 0 {
			j = lps[j-1]
		} else {
			i++
		}
	}

	// Not found
	return -1
}

func buildLast(pattern string) [256]int {
	var lastOcc [256]int

	for i := range lastOcc {
		lastOcc[i] = -1
	}

	for i := 0; i < len(pattern); i++ {
		lastOcc[pattern[i]] = i
	}

	return lastOcc
}

func bmSearch(text string, pattern string) int {
	lastOcc := buildLast(pattern)
	lenText := len(text)
	lenPattern := len(pattern)

	i := lenPattern - 1
	j := lenPattern - 1

	for true {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
				// fmt.Println("Found at", i)
				// lo := lastOcc[text[i]]
				// i = i + lenPattern - min(j, i+lo)
				// j = lenPattern - 1
			} else {
				i--
				j--
			}
		} else {
			lo := lastOcc[text[i]]
			i = i + lenPattern - min(j, i+lo)
			j = lenPattern - 1
		}

		if i > lenText-1 {
			break
		}
	}

	// Not found
	return -1
}

func levenshteinDistance(str1 string, str2 string) int {
	rows, cols := len(str2)+1, len(str1)+1
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	for i := 0; i < cols; i++ {
		matrix[0][i] = i
	}
	for i := 0; i < rows; i++ {
		matrix[i][0] = i
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if str1[j-1] != str2[i-1] {
				matrix[i][j] = min(min(matrix[i][j-1]+1, matrix[i-1][j]+1), matrix[i-1][j-1]+1)
			} else {
				matrix[i][j] = min(min(matrix[i][j-1]+1, matrix[i-1][j]+1), matrix[i-1][j-1])
			}
		}
	}

	return matrix[rows-1][cols-1]
}

func distToPercentage(levDist int, str1 string, str2 string) float64 {
	maxlen := max(len(str1), len(str2))
	return (float64(maxlen-levDist) / float64(maxlen)) * 100
}

func dateToDay(str string) (string, error) {
	reDate := regexp.MustCompile(`[0-9]{4}/[0-9]{2}/[0-9]{2}`)
	strippedDate := reDate.FindStringSubmatch(str)
	if len(strippedDate) > 0 {
		strippedDate[0] = strings.ReplaceAll(strippedDate[0], "-", "/")
		fmt.Println(strippedDate[0])
		t, _ := time.Parse("2006/01/02", strippedDate[0])
		return t.Weekday().String(), nil
	} else {
		return time.Monday.String(), fmt.Errorf("Does not contain date pattern")
	}
}

func solveExpression(str string) string {
	reArithmeticStrip := regexp.MustCompile(`[^0-9\/\*+\-\(\)]`)
	strExpression := reArithmeticStrip.ReplaceAllString(str, "")

	expression, err := govaluate.NewEvaluableExpression(strExpression)
	if err != nil {
		return "Sintaks persamaan tidak valid"
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return "Sintaks persamaan tidak valid"
	}

	if fmt.Sprintf("%v", result) == "+Inf" || fmt.Sprintf("%v", result) == "-Inf" {
		return "Result undefined"
	} else {
		return fmt.Sprintf("Hasilnya adalah %v", result)
	}
}

func preprocessQuery(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, "\n", " ")

	reSpaces := regexp.MustCompile(`\s+`)
	str = reSpaces.ReplaceAllString(str, " ")

	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")
	return str
}

func handleQueries(str string) string {
	// str1 := "Apakah 1+2 sama dengan 3? Siapa wakil presiden indonesia ke-3? Apakah dia benar atau salah. Halo semuanya\n Halohalo"
	// db := "Test kemiripan dengan string ini sekarang juga"
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Input here: ")
	// str1, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// reSpaces := regexp.MustCompile(`\s+`)
	// str1 = reSpaces.ReplaceAllString(str1, " ")
	// fmt.Println(str1)
	// str1 = string.ReplaceAll(str1)

	str = strings.ReplaceAll(str, "\n", "")
	separators := func(sep rune) bool {
		return sep == '?' || sep == '.' || sep == ';'
	}
	queries := strings.FieldsFunc(str, separators)
	// fmt.Println("Queries:", queries)

	reArithmetic := regexp.MustCompile(`^(berapa|hasil dari|hitunglah|hitung|berapakah)?[0-9+\-*/()\s]+$`)
	reDate := regexp.MustCompile(`^\s*(hari|hari apa)?\s*[0-9]{4}/[0-9]{2}/[0-9]{2}\s*\?*\s*$`)
	reAddQuestion := regexp.MustCompile(`^\s*tambah pertanyaan (.+) dengan jawaban (.+)$`)
	reDeleteQuestion := regexp.MustCompile(`^\s*hapus pertanyaan (.+)$`)

	for _, query := range queries {
		query = preprocessQuery(query)
		fmt.Println("preprocesed query:", query)
		if reDate.MatchString(query) {
			dateStr, _ := dateToDay(query)
			return dateStr
		} else if reArithmetic.MatchString(query) {
			return solveExpression(query)
		} else if reAddQuestion.MatchString(query) {
			matches := reAddQuestion.FindStringSubmatch(str)
			// TODO: Query to db
			question := matches[1]
			answer := matches[2]
			return fmt.Sprintf("Pertanyaan %s telah ditambah dengan jawaban %s", question, answer)

		} else if reDeleteQuestion.MatchString(query) {
			match := reDeleteQuestion.FindStringSubmatch(str)

			// TODO: Query to db
			question := match[1]
			return fmt.Sprintf("Pertanyaan %s telah dihapus", question)

		} else {
			// TODO: Match from database

			// fmt.Println("Matching...")
			// fmt.Println("KMP")
			// if len(db) != len(query) {
			// 	levDist := levenshteinDistance(db, query)
			// 	perct := distToPercentage(levDist, db, query)
			// 	fmt.Println("Percentage:", perct)
			// 	if perct < 90 {
			// 		fmt.Println("tidak mirip")
			// 	} else {
			// 		fmt.Println("mirip")
			// 	}
			// } else {
			// 	if kmpSearch(db, query) == -1 {
			// 		levDist := levenshteinDistance(db, query)
			// 		perct := distToPercentage(levDist, db, query)
			// 		fmt.Println("Percentage:", perct)
			// 		if perct < 90 {
			// 			fmt.Println("tidak mirip")
			// 		} else {
			// 			fmt.Println("mirip")
			// 		}
			// 	} else {
			// 		fmt.Println("Exact match !!!")
			// 	}
			// }
			return ""
		}
	}

	return "Empty query"
}
