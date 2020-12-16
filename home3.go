//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Программа которая ищет минимум из двух чисел")
//
//	fmt.Println("Введите первое число:")
//	var numberOne int
//	fmt.Scan(&numberOne)
//
//	fmt.Println("Введите второе число:")
//	var numberTwo int
//	fmt.Scan(&numberTwo)
//
//	if numberOne <= numberTwo {
//		fmt.Println("Число", numberOne, "минимальное")
//	} else if numberTwo <= numberOne {
//		fmt.Println("Число", numberTwo, "минимальное")
//	}
//
//}

//----------------------------------------------------

//
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Складываем в уме")
//
//	numberOne := 86
//	numberTwo := 17
//	result := numberOne + numberTwo
//
//	fmt.Print(numberOne, " + ", numberTwo, "\n")
//	fmt.Println("Введите результат сложения этих чисел:")
//
//	var user int
//	fmt.Scan(&user)
//
//	if user == result {
//		fmt.Println("Вы правильно посчитали! Поздравляем!")
//	} else {
//		fmt.Println("Вы посчитали не правильно! Правильный ответ:", result)
//	}
//
//}

//----------------------------------------------------------

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Проверяем четное ли число")
//
//	fmt.Println("Введите число:")
//	var numberInput int
//	fmt.Scan(&numberInput)
//
//	if numberInput % 2 == 0 {
//		fmt.Println("Число четное!")
//	} else {
//		fmt.Println("Число нечетное")
//	}
//
//}

//-------------------------------------------------------------

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Программа, которая определяет уровень персонажа в компьютерной игре.")
//	fmt.Println("Новый уровень дается при достижении 1000, 2500 и 5000 очков опыта")
//
//	fmt.Println("Введите число:")
//	var experiencePoints int
//	fmt.Scan(&experiencePoints)
//
//	if experiencePoints < 1000 {
//		fmt.Println("Персонаж 1-го уровня!")
//
//	} else if experiencePoints == 1000 {
//		fmt.Println("Персонаж 2-го уровня!")
//
//	} else if experiencePoints == 2500 {
//		fmt.Println("Персонаж 3-го уровня!")
//
//	} else if experiencePoints >= 5000 {
//		fmt.Println("Персонаж 4-го уровня!")
//	}
//
//}

//----------------------------------------------------------------

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Программа, которая проверяет, делится ли одно число на другое без остатка.")
//
//	fmt.Println("Введите первое число:")
//	var firstNumber int
//	fmt.Scan(&firstNumber)
//
//	fmt.Println("Введите второе число")
//	var secondNumber int
//	fmt.Scan(&secondNumber)
//
//	result := firstNumber % secondNumber
//
//	fmt.Println(result)
//
//}

//----------------------------------------------------------------