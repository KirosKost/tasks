package main
//
//import "fmt"
//
//func main() {
//
//	fmt.Println("Определение координатной плоскости точки")
//
//	var x int
//	var y int
//
//	fmt.Println("Введите число Х:")
//	fmt.Scan(&x)
//	fmt.Println("Введите число Y:")
//	fmt.Scan(&y)
//
//
//
//	if x > 0 && y > 0 {
//		fmt.Println("Точка лежит в первой четверти")
//	}else if x < 0 && y > 0 {
//		fmt.Println("Точка лежит во второй четверти")
//	}else if x < 0 && y < 0 {
//		fmt.Println("Точка лежит в третьей четверти")
//	}else if x > 0 && y < 0 {
//		fmt.Println("Точка лежит в четвертой четверти")
//	}
//}
//-----------------------------------------------------------------------

//package main
//
//import "fmt"
//
//func main() {
//
//	fmt.Println("Проверить, что одно из чисел положительное")
//
//	var firstNumber int
//	fmt.Println("Введите первое число:")
//	fmt.Scan(&firstNumber)
//
//	var secondNumber int
//	fmt.Println("Введите второе число:")
//	fmt.Scan(&secondNumber)
//
//	var thirdNumber int
//	fmt.Println("Введите третье число:")
//	fmt.Scan(&thirdNumber)
//
//	positiveNumber := false
//
//	if firstNumber > 0 || secondNumber > 0 || thirdNumber > 0 {
//		positiveNumber = true
//	}
//
//	if positiveNumber {
//		fmt.Println("Среди введенных Вами чисел, есть положительное.")
//
//	}else{
//		fmt.Println("Среди введенных Вами чисел, нет положительного.")
//	}
//}

//--------------------------------------------------------------------
//
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Проверить, что есть совпадающие числа")
//
//	fmt.Println("Введите первое число:")
//	var firstNumber int
//	fmt.Scan(&firstNumber)
//
//	fmt.Println("Введите второе число:")
//	var secondNumber int
//	fmt.Scan(&secondNumber)
//
//	fmt.Println("Введите третье число:")
//	var thirdNumber int
//	fmt.Scan(&thirdNumber)
//
//	if firstNumber == secondNumber || firstNumber == thirdNumber || secondNumber == thirdNumber {
//		fmt.Println("Вводимые Вами числа не должны повтаряться.")
//	}
//
//
//}

//--------------------------------------------------------------
//
//package main

//func main() {
//	fmt.Println("Сумма без сдачи")
//
//	fmt.Println("Введите стоимость товара:")
//	var sum int
//	fmt.Scan(&sum)
//
//	fmt.Println("Введите номинал первой монеты:")
//	var firstCoin int
//	fmt.Scan(&firstCoin)
//
//	fmt.Println("Введите номинал второй монеты:")
//	var secondCoin int
//	fmt.Scan(&secondCoin)
//
//	fmt.Println("Введите номинал третьей монеты:")
//	var thirdCoin int
//	fmt.Scan(&thirdCoin)
//
//	if firstCoin == sum || secondCoin == sum || thirdCoin == sum {
//		fmt.Println("Вы можете оплатить без сдачи")
//	} else if firstCoin + secondCoin == sum || secondCoin + thirdCoin == sum|| thirdCoin + firstCoin == sum {
//		fmt.Println("Вы можете оплатить без сдачи")
//	} else if firstCoin + secondCoin + thirdCoin == sum {
//		fmt.Println("Вы можете оплатить без сдачи")
//	} else {
//		fmt.Println("Вы не можете оплатить без сдачи")
//	}
//
//}

//----------------------------------------------
//
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Определение максимальных процентов")
//
//	fmt.Println("Первая ставка:")
//	var rateOne int
//	fmt.Scan(&rateOne)
//
//	fmt.Println("Вторая ставка:")
//	var rateTwo int
//	fmt.Scan(&rateTwo)
//
//	fmt.Println("Третья ставка:")
//	var rateThree int
//	fmt.Scan(&rateThree)
//
//	if rateOne < rateThree && rateTwo < rateThree {
//		fmt.Println("Процентные ставки", rateOne, "%", "и", rateTwo, "% выгоднее")
//	} else if rateOne < rateTwo && rateThree < rateTwo {
//		fmt.Println("Процентные ставки", rateOne, "%", "и", rateThree, "% выгоднее")
//	}else if rateTwo < rateOne && rateThree < rateOne {
//		fmt.Println("Процентные ставки", rateTwo, "%", "и", rateThree, "% выгоднее")
//	}
//
//}

//-------------------------------------------------------------

//package main
//
//import (
//"fmt"
//"math"
//)
//
//func main() {
//	fmt.Println("Решение квадратного уравнения")
//
//	fmt.Println("Введите первый коэффициент:")
//	var coefficientFirst float64
//	fmt.Scan(&coefficientFirst)
//
//	fmt.Println("Введите второй коэффициент:")
//	var coefficientSecond float64
//	fmt.Scan(&coefficientSecond)
//
//	fmt.Println("Введите третий коэффициент:")
//	var coefficientThird float64
//	fmt.Scan(&coefficientThird)
//
//
//	var discriminant = math.Pow(coefficientSecond, 2) - 4 * coefficientFirst * coefficientThird
//	var root1 = (-(coefficientSecond) + math.Sqrt(discriminant)) / (2 * coefficientFirst)
//	var root2 = (-(coefficientSecond) - math.Sqrt(discriminant)) / (2 * coefficientFirst)
//
//	if (math.IsNaN(root1) && math.IsNaN(root2)) {
//		fmt.Println(discriminant, 0)
//	} else if (root1 == root2) {
//		fmt.Println(discriminant, root1)
//	} else {
//		fmt.Println(discriminant, root1, root2);
//	}
//}

//-------------------------------------

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Счастливый билет")
//
//	fmt.Println("Введите четырехзначное число:")
//	var number string
//	fmt.Scan(&number)
//
//	if number[0] == number[3] && number[1] == number[2] {
//		fmt.Println("Зеркальный билет")
//	}else if number[0] + number[1] == number[2] + number[3] {
//		fmt.Println("Счастливый билет")
//	}else if number[0] + 1 == number[1] && number[1] + 1 ==number[2] && number[2] + 1 == number[3] {
//		fmt.Println("Обычный билет")
//	}
//}
