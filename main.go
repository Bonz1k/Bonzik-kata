package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A programmer friend of mine told me my crippled calculator was beautiful
var romToArabicDic = map[string]int64{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
}

var arabToRomaDic = map[int64]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	20:  "XX",
	30:  "XXX",
	40:  "XL",
	50:  "L",
	60:  "LX",
	70:  "LXX",
	80:  "LXXX",
	90:  "XC",
	100: "C",
	0:   "",
}

func main() {

	sumFunc := func(x, y int64) int64 {
		return x + y
	}
	subFunc := func(x, y int64) int64 {
		return x - y
	}
	multFunc := func(x, y int64) int64 {
		return x * y
	}
	divFunc := func(x, y int64) int64 {
		return x / y
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите выражение в виде (a [оператор: +, -, * или /] b): ")
		scanner.Scan()
		readExpression := scanner.Text()
		s := strings.Split(readExpression, " ")
		arabRef := "12345678910"
		romaRef := "IIIVIIIX"

		if len(s) < 3 {
			fmt.Println("Ошибка! Строка не является математической операцией.")
		}

		if len(s) > 3 {
			fmt.Println("Ошибка! Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}

		if len(s) == 3 {
			isFirstNumArab := strings.Contains(arabRef, s[0])
			isSecondNumArab := strings.Contains(arabRef, s[2])
			isFirstNumRoma := strings.Contains(romaRef, s[0])
			isSecondNumRoma := strings.Contains(romaRef, s[2])
			firstNum, ferr := strconv.ParseInt(s[0], 10, 64)
			secondNum, serr := strconv.ParseInt(s[2], 10, 64)

			if (isFirstNumArab == true && isSecondNumRoma == true) || (isFirstNumRoma == true && isSecondNumArab == true) {
				fmt.Println("Ошибка! Используются одновременно разные системы счисления.")
				break
			}

			if (ferr == nil && serr == nil) && ((firstNum < 1 || firstNum > 10) || (secondNum < 1 || secondNum > 10)) {
				fmt.Println("Я запрещаю тебе вводить в программу операнды х <= 0 и/или x > 10")
				break
			}

			if isFirstNumArab == true && isSecondNumArab == true {

				if s[1] == "+" {
					fmt.Println(sumFunc(firstNum, secondNum))
					break
				}

				if s[1] == "-" {
					fmt.Println(subFunc(firstNum, secondNum))
					break
				}

				if s[1] == "/" {
					fmt.Println(divFunc(firstNum, secondNum))
					break
				}

				if s[1] == "*" {
					fmt.Println(multFunc(firstNum, secondNum))
					break
				}
				break
			}

			if isFirstNumRoma == true && isSecondNumRoma == true {
				var arabRes int64

				if s[1] == "+" {
					arabRes = sumFunc(romToArabicDic[s[0]], romToArabicDic[s[2]])

					if arabRes == 100 {
						fmt.Println(arabToRomaDic[arabRes])
						break
					} else {
						fmt.Println(arabToRomaDic[(arabRes/10)*10] + arabToRomaDic[arabRes%10])
						break
					}
				}

				if s[1] == "-" {
					arabRes = subFunc(romToArabicDic[s[0]], romToArabicDic[s[2]])

					if arabRes > 0 {

						if arabRes == 100 {
							fmt.Println(arabToRomaDic[arabRes])
							break
						} else {
							fmt.Println(arabToRomaDic[(arabRes/10)*10] + arabToRomaDic[arabRes%10])
							break
						}
					}

					if arabRes == 0 {
						fmt.Println("Ошибка! В римской системе нет нуля.")
						break
					} else {
						fmt.Println("Ошибка! В римской системе нет отрицательных чисел.")
						break
					}
				}

				if s[1] == "/" {
					arabRes = divFunc(romToArabicDic[s[0]], romToArabicDic[s[2]])

					if arabRes == 0 {
						fmt.Println("Ошибка! В римской системе нет нуля.")
						break
					} else {

						if arabRes == 100 {
							fmt.Println(arabToRomaDic[arabRes])
							break
						} else {
							fmt.Println(arabToRomaDic[(arabRes/10)*10] + arabToRomaDic[arabRes%10])
							break
						}
					}
				}

				if s[1] == "*" {
					arabRes = multFunc(romToArabicDic[s[0]], romToArabicDic[s[2]])

					if arabRes == 100 {
						fmt.Println(arabToRomaDic[arabRes])
						break
					} else {
						fmt.Println(arabToRomaDic[(arabRes/10)*10] + arabToRomaDic[arabRes%10])
						break
					}
				}
				break
			}
		}
		break
	}
}
