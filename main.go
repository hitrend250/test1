package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	commits := map[int]string{
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
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	}
	key := map[int]string{
		1: "+",
		2: "-",
		3: "/",
		4: "*",
	}
	var is1 bool = false
	var is2 bool = false

	// Получение данных
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("Input: ")

		text, _ := reader.ReadString('\n') // Ждет ввода данных в формате строки
		//text = strings.ReplaceAll(text, " ", "")
		//text = strings.ReplaceAll(text, "	", "")
		text = strings.TrimSpace(text)
		sum := 0
		for q := 1; q < 5; q++ {
			mass := strings.Split(text, key[q])

			if len(mass) == 2 {
				sum += q
				for i := 1; i < 11; i++ {
					if commits[i] == strings.TrimSpace(mass[0]) {
						mass[0] = strconv.Itoa(i)
						is1 = true
					}
				}
				for a := 1; a < 11; a++ {
					if commits[a] == strings.TrimSpace(mass[1]) {
						mass[1] = strconv.Itoa(a)
						is2 = true
					}
				}
				if is1 == true && is2 != true || is1 != true && is2 == true {
					fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
					return
				}
				number1, _ := strconv.Atoi(strings.TrimSpace(mass[0]))
				operator := key[q]
				number2, _ := strconv.Atoi(strings.TrimSpace(mass[1]))

				if number1 < 1 || number1 > 10 {

					fmt.Println("Значение первого аргумента неверное")
					return
				}
				if number2 < 1 || number2 > 10 {
					fmt.Println("Значение второго аргумента неверное")
					return
				}
				switch operator {

				case "+":
					fmt.Println("Output: ")
					if is1 == true && is2 == true {
						rim, _ := commits[number1+number2]
						fmt.Println(rim)
						is1 = false
						is2 = false
					} else {
						fmt.Println(number1 + number2)
					}
				case "-":
					fmt.Println("Output: ")
					if is1 == true && is2 == true {
						if number1-number2 < 1 {
							fmt.Println("Вывод ошибки, так как в римской системе вышел результат работы меньше единицы.")
							return
						}
						rim, _ := commits[number1-number2]
						fmt.Println(rim)
						is1 = false
						is2 = false
					} else {
						fmt.Println(number1 - number2)
					}
				case "*":
					fmt.Println("Output: ")
					if is1 == true && is2 == true {
						rim, _ := commits[number1*number2]
						fmt.Println(rim)
						is1 = false
						is2 = false
					} else {
						fmt.Println(number1 * number2)
					}
				case "/":
					fmt.Println("Output: ")
					if is1 == true && is2 == true {
						rim, _ := commits[number1/number2]
						fmt.Println(rim)
						is1 = false
						is2 = false
					} else {
						fmt.Println(number1 / number2)
					}
				default:
					fmt.Println("Неверный ввод данных")
				}
			}
		}
		if sum == 0 {
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}
	}
}
