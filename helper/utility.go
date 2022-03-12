package helper

import (
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func RandomSolution() string {
	rand.Seed(time.Now().UnixNano())
	solutionIndex := rand.Intn(len(Words))
	solutionWord := strings.ToUpper(Words[solutionIndex])
	return solutionWord
}
func OfficialSolution() string {
	// Official intital date for wordle is 2021/06/19
	// so calculate no of days from initial date to today date
	// modulus of total days to number of array is index of official word.
	firstWordleDay := time.Date(2021, time.Month(6), 19, 0, 0, 0, 0, time.UTC) // 2021/06/19
	day := int(time.Now().Sub(firstWordleDay).Hours() / 24)                    // 2021/06/19 minus today_date
	officialWord := Words[day%len(Words)]
	return strings.ToUpper(officialWord)
}
func ValidateWord(guess string) (string, bool) {
	message := ""
	if !include(guess, Words) && !include(guess, ValidWord) && len(guess) == 5 {
		message = "Not in Word List"
	}
	if len(guess) != 5 {
		message = "words should be exactly 5"
	}
	return message, (include(guess, Words) || include(guess, ValidWord)) && len(guess) == 5
}
func CheckAccuracy(solution string, guessed string) (string, string) {

	solutionWordArr := strings.Split(solution, "")
	guessedWordArr := strings.Split(guessed, "")
	finalWord := ""
	for i, value := range solutionWordArr {
		alphaIndex := getAlphabetIndex(guessedWordArr[i])
		if !include(guessedWordArr[i], solutionWordArr) {
			alphabets[alphaIndex] = PaintGrey(guessedWordArr[i])
			finalWord += coloredString("grey").Render(guessedWordArr[i])
		}
		if guessedWordArr[i] == value {
			alphabets[alphaIndex] = PaintGreen(guessedWordArr[i])
			solutionWordArr[i] = "*"
			finalWord += coloredString("green").
				Render(guessedWordArr[i])
		}
		if guessedWordArr[i] != value && include(guessedWordArr[i], solutionWordArr) {
			alphabets[alphaIndex] = PaintYellow(guessedWordArr[i])
			solutionWordArr[index(guessedWordArr[i], solutionWordArr)] = "*"
			finalWord += coloredString("yellow").
				Render(guessedWordArr[i])
		}
	}
	return lipgloss.NewStyle().
		Margin(1).
		MarginLeft(5).
		Render(finalWord), GetAlphabetKeyboard(alphabets)
}

// use pointer concept
func GetRightAnswer(solution string) string {
	return coloredString("cyan").Margin(2).Render(solution)
}

// todo: remove this logic
func InProgress(round int, solutionWord string, guessWord string) bool {
	return round <= 6 && solutionWord != guessWord
}

func WinCondition(round int, solutionWord string, guessWord string) bool {
	return round <= 6 && solutionWord == guessWord
}

func LoseCondition(round int, solutionWord string, guessWord string) bool {
	// fmt.Print(round == 6 && solutionWord != guessWord) //remove
	return round == 6 && solutionWord != guessWord
}
func include(a string, list []string) bool {
	for _, b := range list {
		if strings.ToUpper(b) == strings.ToUpper(a) {
			return true
		}
	}
	return false
}

func index(a string, arr []string) int {
	for i, value := range arr {
		if value == a {
			return i
		}
	}
	return -1
}

func getAlphabetIndex(alphabet string) int {
	alphabets := []string{"A", "B", "C", "D", "E", "F", "G", "H",
		"I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return index(strings.ToUpper(alphabet), alphabets)
}
