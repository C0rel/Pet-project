package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func saveResult(outcome string, attempts int) error {
   
    file, err := os.Open("results.json")
    var results []GameResult

    if err == nil {
        defer file.Close()
        json.NewDecoder(file).Decode(&results)
    } else if !os.IsNotExist(err) {
        return err 
    }

    
    newResult := GameResult{
        Date:     time.Now().Format("2006-01-02 15:04:05"),
        Outcome:  outcome,
        Attempts: attempts,
    }
    results = append(results, newResult)


    file, err = os.Create("results.json")
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(results)
}
type GameResult struct {
    Date     string `json:"date"`
    Outcome  string `json:"outcome"`
    Attempts int    `json:"attempts"`
}

func pastAttempts(slInputValue []int, inputValue int) []int {
    slInputValue = append(slInputValue, inputValue)
    fmt.Printf("Прошлые попытки: %v\n", slInputValue)
    return slInputValue
}
func hints(randNum int, userValue int) {
    diff := int64(math.Abs(float64(randNum - userValue)))
   	if diff <= 5 {
    	color.Red("Горячо")
	} else if diff <= 15 {
    	color.Yellow("Тепло")
	} else {
    	color.Blue("Холодно")
	}
}
	func choosingTheDifficulty()(int,int){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Выбор сложности:")
		fmt.Println("1)Easy: 1–50, 15 попыток")
		fmt.Println("2)Medium: 1–100, 10 попыток")
		fmt.Println("3)Hard: 1–200, 5 попыток")
		var selected int
		var err error
		for{
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			selected, err = strconv.Atoi(input) 
		
			if err==nil&&selected>=1&&selected<=3{
				break
			}
			fmt.Println("Введите корректное значение")
			
		}
		switch selected{
		case 1:
			fmt.Print("\033[H\033[2J")
			return rand.Intn(50) + 1, 15
		case 2:
			fmt.Print("\033[H\033[2J")
			return rand.Intn(100) + 1, 10
		case 3:
			fmt.Print("\033[H\033[2J")
			return rand.Intn(200) + 1, 5
		}
		fmt.Print("\033[H\033[2J")
		return 0,0
	}
	func getUserInput() int {
    reader := bufio.NewReader(os.Stdin)
    for {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        value, err := strconv.Atoi(input)
        if err == nil {
            return value
        }
        fmt.Println("Введите корректное число")
    }
}
	func compareAndHint(secret, guess int) bool {
    if guess > secret {
        hints(secret, guess)
        color.Yellow("Секретное число меньше")
        return false
    } else if guess < secret {
        hints(secret, guess)
        color.Yellow("Секретное число больше")
        return false
    } else {
        return true 
    }
}
	func guessGame(secretNumber int, maxAttempts int) (string, int) {
    var allAttempts []int
    var attempts int

    for i := 0; i < maxAttempts; i++ {
        color.Yellow("\nПопытка %d из %d", i+1, maxAttempts)
        fmt.Println("Введите число:")

        userGuess := getUserInput()           
        allAttempts = pastAttempts(allAttempts, userGuess) 
        attempts++

        if compareAndHint(secretNumber, userGuess) { 
            outcome := "Победа"
            color.Green(outcome)
            fmt.Printf("Попыток: %d\n", attempts)
            saveResult(outcome, attempts)
            return outcome, attempts
        }
    }

   
    outcome := "Проигрыш"
    color.Red(outcome)
    fmt.Printf("Попыток: %d\n", attempts)
    saveResult(outcome, attempts)
    return outcome, attempts
}

func main(){
	replay:=true
	for replay{
	randNumber,numOfAttepts:=choosingTheDifficulty()
	guessGame(randNumber,numOfAttepts)
	fmt.Printf("Загаданное число:%d\n",randNumber)
	fmt.Println("Повторить игру?")
	fmt.Println("1-Да")
	fmt.Println("2-Нет")
	reader := bufio.NewReader(os.Stdin)
	var inpUser int
	var err error
	for{
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		inpUser, err = strconv.Atoi(input) 
		if err==nil&&(inpUser==1||inpUser==2) {break}
		fmt.Println("Введите корректное значение")	
		}

    if inpUser == 2 {
    replay = false
    break 
	}else {
        fmt.Println("Введите 1 или 2")
	}
	fmt.Print("\033[H\033[2J")
}

}