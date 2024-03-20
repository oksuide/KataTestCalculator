package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
		println(definingOperands())
	}
	 
func removeSpaces(text string) string {
    result := make([]rune, 0, len(text))
    for _, char := range text {
        if char != ' ' {
            result = append(result, char)
        }
    }
    return string(result)
	}

func receivedData()string{	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите значения")
	text,_:=reader.ReadString('\n') 							//ждет ввода данных в формате строки
	text = strings.TrimSpace(text) 								//Очищает все пустоты
	return removeSpaces(text) 
}

var a = receivedData()

func operatorDefinition()string{ 									//определение оператора
	chip:=a
	if (strings.Contains(chip,"+")){
		return "+"
		} else if (strings.Contains(chip,"-")){
			return "-"
		} else if (strings.Contains(chip,"*")){
				return "*"
		} else if (strings.Contains(chip,"/")){
					return "/"
		} else {return ""}
				}


				
func operandSeparation()[]string{
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("+-*/", r) 			//Функция разделения на отдельные операнды
		}
	noNumbers := strings.FieldsFunc(a, splitFunc)
		if len(noNumbers)>2{
			panic("Введено более 2 значений")					//Проверка на количество операндов
		}
		return noNumbers
	}

func definingOperands()string{										//Определение и сравнение операндов
	numArray:= operandSeparation()
	noNum1,noNum2 := numArray[0],numArray[1] 					
	if isRoman(noNum1) && isRoman(noNum2){
		num1 := convertToArabian(noNum1)
		num2 := convertToArabian(noNum2)
		if num1 < 1||num1>10{					
		panic("Нарушен диапазон чисел")
	}
	if num2 < 1||num2>10{
		panic("Нарушен диапазон чисел")										
	}
	arabRes:=count(num1,num2)
	if arabRes < 1{
		panic("Результатом выражения римских чисел должно быть положительным")
	}
	rimRes:= convertToRim(arabRes)
	return rimRes

} else if !isRoman(noNum1) && !isRoman(noNum2){
		num1,_ := strconv.Atoi(noNum1)  								
	if num1 < 1||num1>10{
		panic("Нарушен диапазон чисел")
	}
	num2,_ := strconv.Atoi(noNum2)
	if num2 < 1||num2>10{
		panic("Нарушен диапазон чисел")
	}
	resInt:=count(num1,num2)
	resString:=strconv.Itoa(resInt)
	return resString	

} else {panic("Введены некорректные значения. Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
}

}

func count(num1 int,num2 int)int{										//проведение подсчетов
		operator:=operatorDefinition()
		switch{
		case operator == "+":
			return (num1 + num2)
		case operator == "-":
			return (num1 - num2)
		case operator == "*":
			return (num1 * num2)
		case operator == "/":
			return (num1 / num2)
			default:
				return 0
		}
	}

	var romanArray = []string{"0","I","II","III","IV","V","VI","VII","VIII","IX","X","XI","XII","XIII","XIV","XV","XVI","XVII","XVIII","XIX","XX","XXI","XXII","XXIII","XXIV","XXV","XXVI","XXVII","XXVIII","XXIX","XXX","XXXI","XXXII","XXXIII","XXXIV","XXXV","XXXVI","XXXVII","XXXVIII","XXXIX","XL","XLI","XLII","XLIII","XLIV","XLV","XLVI","XLVII","XLVIII","XLIX","L","LI","LII","LIII","LIV","LV","LVI","LVII","LVIII","LIX","LX","LXI","LXII","LXIII","LXIV","LXV","LXVI","LXVII","LXVIII","LXIX","LXX","LXXI","LXXII","LXXIII","LXXIV","LXXV","LXXVI","LXXVII","LXXVIII","LXXIX","LXXX","LXXXI","LXXXII","LXXXIII","LXXXIV","LXXXV","LXXXVI","LXXXVII","LXXXVIII","LXXXIX","XC","XCI","XCII","XCIII","XCIV","XCV","XCVI","XCVII","XCVIII","XCIX","C"}

	func isRoman(val string)bool{											//Проверка принадлежности чисел к арабским\римским
		for	i:=0; i < len(romanArray); i++{
			if (slices.Contains(romanArray, val)){
			return true
			}}
			return false
	}

	func convertToArabian(rimVal string)int{					//Конвертация римских чисел в арабские
		for i:=0; i < len(romanArray); i++{
			if (romanArray[i] == rimVal){
			return i
		}}
		return -1
	}

	func convertToRim(arabVal int)string{							//Конвертация арабских чисел в римские
		rimVal:=romanArray[arabVal]
		return rimVal
	}
