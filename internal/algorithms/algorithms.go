package algorithms

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/rizkyrsyd28/internal/repository"
	"golang.org/x/net/context"

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

func dateToDay(str string) string {
	reDate := regexp.MustCompile(`[0-9]{4}/[0-9]{2}/[0-9]{2}`)
	strippedDate := reDate.FindStringSubmatch(str)
	if len(strippedDate) > 0 {
		strippedDate[0] = strings.ReplaceAll(strippedDate[0], "-", "/")
		fmt.Println(strippedDate[0])
		year, _ := strconv.Atoi(strippedDate[0][0:4])
		month, _ := strconv.Atoi(strippedDate[0][5:7])
		day, _ := strconv.Atoi(strippedDate[0][8:10])

		if month < 1 || month > 12 {
			return "Date not valid"
		}
		if day < 1 {
			return "Date not valid"
		}
		if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
			if day > 31 {
				return "Date not valid"
			}
		}
		if month == 4 || month == 6 || month == 9 || month == 11 {
			if day > 30 {
				return "Date not valid"
			}
		}
		if month == 2 {
			if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
				if day > 29 {
					return "Date not valid"
				}
			} else {
				if day > 28 {
					return "Date not valid"
				}
			}
		}
		t, _ := time.Parse("2006/01/02", strippedDate[0])
		return t.Weekday().String()
	} else {
		return "Date not valid"
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

	re := regexp.MustCompile(`[!?;.]`)
	str = re.ReplaceAllString(str, "")

	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")
	return str
}

type QnaDistance struct {
	IDQna    int
	Question string
	Distance float64
	Answer   string
}

// Functions for sorting QnaDistance
type ByDistance []QnaDistance

func (a ByDistance) Len() int {
	return len(a)
}

func (a ByDistance) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDistance) Less(i, j int) bool {
	return a[i].Distance > a[j].Distance
}

func HandleQueries(r repository.Repo, c context.Context, str string, algo string) string {
	fmt.Println("TESTING1")
	str = strings.ReplaceAll(str, "\n", "")
	separators := func(sep rune) bool {
		return sep == '?' || sep == '.' || sep == ';' || sep == '!'
	}
	queries := strings.FieldsFunc(str, separators)

	reArithmetic := regexp.MustCompile(`^(berapa|hasil dari|hitunglah|hitung|berapakah)?[0-9+\-*/()\s]+$`)
	reDate := regexp.MustCompile(`^\s*(hari|hari apa)?\s*[0-9]{4}/[0-9]{2}/[0-9]{2}\s*\?*\s*$`)
	reAddQuestion := regexp.MustCompile(`^\s*tambah pertanyaan (.+) dengan jawaban (.+)$`)
	reDeleteQuestion := regexp.MustCompile(`^\s*hapus pertanyaan (.+)$`)

	result := ""
	for _, query := range queries {
		fmt.Println("TESTING1")
		allData, err := r.GetAllData(c)
		if err != nil {
			return "Fetch data error"
		}

		query = preprocessQuery(query)
		if reDate.MatchString(query) {
			dateStr := dateToDay(query)
			result += fmt.Sprintf("%s\n", dateStr)
		} else if reArithmetic.MatchString(query) {
			result += fmt.Sprintf("%s\n", solveExpression(query))
		} else if reAddQuestion.MatchString(query) {
			matches := reAddQuestion.FindStringSubmatch(query)
			question := matches[1]
			answer := matches[2]

			var distances []QnaDistance

			foundExact := false
			for _, qna := range allData {
				qnaProcessed := preprocessQuery(qna.Question)
				if len(qnaProcessed) != len(question) {
					levDist := levenshteinDistance(qnaProcessed, question)
					perct := distToPercentage(levDist, qnaProcessed, question)
					distances = append(distances, QnaDistance{qna.IDQna, qnaProcessed, perct, qna.Answer})
				} else {
					if algo == "KMP" {
						if kmpSearch(qnaProcessed, question) == -1 {
							levDist := levenshteinDistance(qnaProcessed, question)
							perct := distToPercentage(levDist, qnaProcessed, question)
							distances = append(distances, QnaDistance{qna.IDQna, qnaProcessed, perct, qna.Answer})
						} else {
							r.DeleteDataById(c, qna.IDQna)
							r.AddData(c, preprocessQuery(question), preprocessQuery(answer))
							result += fmt.Sprintf("Pertanyaan %s sudah ada! Jawaban di update ke %s.\n", question, answer)
							foundExact = true
							break
						}
					} else {
						if bmSearch(qnaProcessed, question) == -1 {
							levDist := levenshteinDistance(qnaProcessed, question)
							perct := distToPercentage(levDist, qnaProcessed, question)
							distances = append(distances, QnaDistance{qna.IDQna, qnaProcessed, perct, qna.Answer})
						} else {
							r.DeleteDataById(c, qna.IDQna)
							r.AddData(c, preprocessQuery(question), preprocessQuery(answer))
							result += fmt.Sprintf("Pertanyaan %s sudah ada! Jawaban di update ke %s.\n", question, answer)
							foundExact = true
							break
						}
					}
				}
			}

			if !foundExact {
				sort.Sort(ByDistance(distances))
				if distances[0].Distance >= 90 {
					r.DeleteDataById(c, distances[0].IDQna)
					r.AddData(c, preprocessQuery(distances[0].Question), preprocessQuery(answer))
					result += fmt.Sprintf("Pertanyaan %s sudah ada! Jawaban di update ke %s.\n", distances[0].Question, answer)
				} else {
					r.AddData(c, preprocessQuery(question), preprocessQuery(answer))
					result += fmt.Sprintf("Pertanyaan %s telah ditambah dengan jawaban %s.\n", question, answer)
				}
			}

		} else if reDeleteQuestion.MatchString(query) {
			match := reDeleteQuestion.FindStringSubmatch(query)
			question := match[1]
			fmt.Println("Question yg mau didelet", question)

			var distances []QnaDistance

			foundExact := false
			for _, qna := range allData {
				qnaProcessed := preprocessQuery(qna.Question)
				if len(qnaProcessed) != len(question) {
					levDist := levenshteinDistance(qnaProcessed, question)
					perct := distToPercentage(levDist, qnaProcessed, question)
					distances = append(distances, QnaDistance{qna.IDQna, qnaProcessed, perct, qna.Answer})
				} else {
					if algo == "KMP" {
						if kmpSearch(qnaProcessed, question) == -1 {
							levDist := levenshteinDistance(qnaProcessed, question)
							perct := distToPercentage(levDist, qnaProcessed, question)
							distances = append(distances, QnaDistance{qna.IDQna, qnaProcessed, perct, qna.Answer})
						} else {
							r.DeleteDataById(c, qna.IDQna)
							result += fmt.Sprintf("Pertanyaan %s telah dihapus.\n", question)
							foundExact = true
							break
						}
					} else {
						if bmSearch(qnaProcessed, question) == -1 {
							levDist := levenshteinDistance(qnaProcessed, question)
							perct := distToPercentage(levDist, qnaProcessed, question)
							distances = append(distances, QnaDistance{qna.IDQna, qnaProcessed, perct, qna.Answer})
						} else {
							r.DeleteDataById(c, qna.IDQna)
							result += fmt.Sprintf("Pertanyaan %s telah dihapus.\n", question)
							foundExact = true
							break
						}
					}
				}
			}

			if !foundExact {
				sort.Sort(ByDistance(distances))
				fmt.Println(distances)
				if distances[0].Distance >= 90 {
					r.DeleteDataById(c, distances[0].IDQna)
					result += fmt.Sprintf("Pertanyaan %s telah dihapus.\n", distances[0].Question)
				} else {
					result += fmt.Sprintf("Tidak ada pertanyaan %s pada database!\nApakah maksud anda:\n", question)
					for i := 0; i < len(distances); i++ {
						result += fmt.Sprintf("%d. %s\n", i+1, distances[i].Question)
						if i == 2 {
							break
						}
					}
				}
			}

		} else {
			fmt.Println("TESTING MATCHING")
			var distances []QnaDistance

			foundExact := false
			for _, qna := range allData {
				qnaProcessed := preprocessQuery(qna.Question)
				if len(qnaProcessed) != len(query) {
					levDist := levenshteinDistance(qnaProcessed, query)
					perct := distToPercentage(levDist, qnaProcessed, query)
					distances = append(distances, QnaDistance{qna.IDQna, qnaProcessed, perct, qna.Answer})
				} else {
					if algo == "KMP" {
						if kmpSearch(qnaProcessed, query) == -1 {
							levDist := levenshteinDistance(qnaProcessed, query)
							perct := distToPercentage(levDist, qnaProcessed, query)
							distances = append(distances, QnaDistance{qna.IDQna, qnaProcessed, perct, qna.Answer})
						} else {
							result += fmt.Sprintf("%s\n", qna.Answer)
							foundExact = true
							break
						}
					} else {
						if bmSearch(qnaProcessed, query) == -1 {
							levDist := levenshteinDistance(qnaProcessed, query)
							perct := distToPercentage(levDist, qnaProcessed, query)
							distances = append(distances, QnaDistance{qna.IDQna, qnaProcessed, perct, qna.Answer})
						} else {
							result += fmt.Sprintf("%s\n", qna.Answer)
							foundExact = true
							break
						}
					}
				}
			}

			if !foundExact {
				sort.Sort(ByDistance(distances))
				fmt.Println(distances)
				if distances[0].Distance >= 90 {
					result += fmt.Sprintf("%s\n", distances[0].Answer)
				} else if distances[0].Distance <= 30 {
					result += "Pertanyaan tidak dapat dihandle oleh fitur yang ada (pertanyaan terlalu acak jika dibandingkan dengan database).\n"
				} else {
					result += "Pertanyaan tidak ditemukan di database.\nApakah maksud anda:\n"
					for i := 0; i < len(distances); i++ {
						result += fmt.Sprintf("%d. %s\n", i+1, distances[i].Question)
						if i == 2 {
							break
						}
					}
				}
			}

		}
	}
	return result
}
