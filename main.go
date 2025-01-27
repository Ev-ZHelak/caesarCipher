package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

type InputUser struct {
	Text string
	Key int
	EdCode bool
}

func (x *InputUser) inputText(s string) error {
	x.Text = s
	return nil
}

func (x *InputUser) inputKey(s string) error {
	if len(s) > 0 && (s[0] == '-' || unicode.IsDigit(rune(s[0]))) {
		for i := 1; i < len(s); i++ {
			if !unicode.IsDigit(rune(s[i])) {
				return fmt.Errorf("неверный формат ввода")
			}
		}
	} else {
		return fmt.Errorf("неверный формат ввода")
	}

	res, _ := strconv.Atoi(s)
	x.Key = res
	return nil
}

func (x *InputUser) inputEdCode(s string) error {
	if len(s) == 1 && (s == "1" || s == "0") {
		x.EdCode = map[string]bool{"1": true, "0": false}[s]
		return nil
	}
	return fmt.Errorf("неверный формат ввода")
}

var (
	allRune = [][]rune{
		[]rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдежзийклмнопрстуфхцчшщъыьэюя"),
		[]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"),
		[]rune("0123456789"),
		[]rune(`!"#%&'()*,-./:;?@[\]_{}`),
		[]rune(" "),
	}
)

//-------------------------------------------------
func main() {
	// encodeStr := caesarCipher("Привет мир", 3, true)
	// fmt.Println(encodeStr)
	// decoderStr := caesarCipher(encodeStr, 3, false)
	// fmt.Println(decoderStr)
	for {
		input := inputUser()
		result := caesarCipher(input.Text, input.Key, input.EdCode)
		fmt.Println(result)
	}
}
//-------------------------------------------------

func inputUser() InputUser {
	newScanner := bufio.NewScanner(os.Stdin)
	newInputUser := InputUser{}

	for {
		fmt.Print("Введите текст: ")
		newScanner.Scan() 
		input := newScanner.Text()
		err := newInputUser.inputText(input)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
	
	for {
		fmt.Print("Укажите ключ (сдвиг): ")
		newScanner.Scan() 
		input := newScanner.Text()
		err := newInputUser.inputKey(input)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
	
	for {
		fmt.Print("кодировать=1 декодировать=0: ")
		newScanner.Scan() 
		input := newScanner.Text()
		err := newInputUser.inputEdCode(input)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
	return newInputUser
}


func caesarCipher(text string, shift int, encode bool) string {
	shiftAllRune := make([][]rune, len(allRune))
	for i := 0; i < len(shiftAllRune); i++ {
		shiftAllRune[i] = rotate(allRune[i], shift)
	}

	result := ""
	textRune := []rune(text)
	if encode {
		for _, s := range textRune {
			for i, slr := range allRune {
				if slices.Contains(slr, s) {
					result += string(shiftAllRune[i][slices.Index(slr, s)])
				}
			}
		}
	} else {
		for _, s := range textRune {
			for i, slr := range shiftAllRune {
				if slices.Contains(slr, s) {
					result += string(allRune[i][slices.Index(slr, s)])
				}
			}
		}
	}
	return result
}

func rotate(slice []rune, steps int) []rune {
	n := len(slice)
	steps = steps % n

	if steps < 0 {
		steps = n + steps
	}

	return append(slice[steps:], slice[:steps]...)
}
