package main

import "fmt"

func main() {

	passingScore := 275

	fmt.Print("Введите колличество баллов за 1ый предмет: \n")
	var exam1 int
	fmt.Scan(&exam1)

	fmt.Print("Введите колличество баллов за 2ой предмет: \n")
	var exam2 int
	fmt.Scan(&exam2)

	fmt.Print("Введите колличество баллов за 3ий предмет: \n")
	var exam3 int
	fmt.Scan(&exam3)

	totalScore := exam1 + exam2 + exam3

	if totalScore >= passingScore {
		fmt.Println("Вы успешно сдали все экзамены и поступили. Ваш общий бал:", totalScore)
	} else {
		fmt.Println("У Вас не хватает баллов для поступления! Ваш общий бал всего:", totalScore, "из", passingScore)
	}

}

//--------------------------------------------------------------

package main

import "fmt"

func main() {

	fmt.Println("Программа, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число, большее, чем 5")

	number := 5

	fmt.Println("Введите первое число:")
	var numberOne int
	fmt.Scan(&numberOne)

	fmt.Println("Введите второе число:")
	var numberTwo int
	fmt.Scan(&numberTwo)

	fmt.Println("Введите третье число:")
	var numberThree int
	fmt.Scan(&numberThree)

	if numberOne > number {
		fmt.Println("Число", numberOne, "больше чем 5")
	} else if numberTwo > number {
		fmt.Println("Число", numberTwo, "больше чем 5")
	} else if numberThree > number {
		fmt.Println("Число", numberThree, "больше чем 5")
	} else {
		fmt.Println("Нет числа, больше чем", number)
	}


}

//-------------------------------------------------------------------------------

package main

import "fmt"

func main() {

	fmt.Println("Банкомат")


	billOne := 100
	billTwo := 100000

	fmt.Print("Какую сумму наличных Вы хотите снять с карты?\n")
	var user int
	fmt.Scan(&user)

	if user < billOne {
		fmt.Println("Извините, но минимальная сумма для снятия равна", billOne,"руб.")

	} else if user % 100 != 0 {
		fmt.Println("Извините, но купюра минимального номинала", billOne, "руб.")

	} else if user >= billTwo {
		fmt.Println("Извините, но максимальная сумма для снятия не может превышать", billTwo,"руб.")

	} else {
		fmt.Println("Уважаемый пользователь, только что Вы сняли", user,"руб.")
	}

}

//-----------------------------------------------------------------------------------

package main

import "fmt"

func main() {

	number := 5
	var count int

	fmt.Println("Программа, которая запрашивает у пользователя три числа и выводит количество чисел, которые больше, либо равны 5")

	fmt.Println("Введите первое число:")
	var numberOne int
	fmt.Scan(&numberOne)

	fmt.Println("Введите второе число:")
	var numberTwo int
	fmt.Scan(&numberTwo)

	fmt.Println("Введите третье число:")
	var numberThree int
	fmt.Scan(&numberThree)



	if numberOne > number {

		count++
	}
	if numberTwo > number {

		count++
	}
	if numberThree > number {

		count++
	}

	fmt.Println("Чисел больше чем:", number, count)

}

//---------------------------------------------------------

package main

import "fmt"

func main() {

	fmt.Println("---Ресторан---")

	discount1 := 10
	discount2 := 5
	bill := 10000
	costumers := 5
	service := 10
	monday := "понедельник"
	friday := "пятница"
	var day string


	fmt.Println("Введите день недели:")
	fmt.Scan(&day)

	fmt.Println("Введите число гостей:")
	fmt.Scan(&costumers)

	fmt.Println("Введите сумму чека:")
	fmt.Scan(&bill)


	if day == monday {
		fmt.Println("У Вас скидка по счету",bill, " руб. в размере", discount1, "% на всё меню, потому что понедельник — день тяжёлый!")
		fmt.Println("Итоговая сумма оплаты:",bill - (bill * discount1 / 100))
	}

	if day == friday {
		fmt.Println("по пятницам, если сумма чека превышает", bill, "руб., включается дополнительная скидка в размере", discount2,"%")

		if bill >= bill {
			fmt.Println("Итоговая сумма оплаты:",bill - (bill * discount2 / 100))
		}
	}

	if costumers > costumers {

		fmt.Print("Если число гостей в одной компании превышает", costumers,"человек, автоматически включается надбавка на обслуживание", service, "%.")
		bill = (bill/100) * service + bill

	} else {
		fmt.Println("Вы заказали столик на", day, ", число Ваших гостей", costumers, ", Предпологаемая сумма заказа", bill, "руб.")
	}


}

//---------------------------------------------------

package main

import "fmt"

func main() {

	fmt.Println("---Студенты---")

	var nameOfStudent string
	var nameOfGroup string
	var serialNumber int
	groupFirst := "Группа 1"
	groupSecond := "Группа 2"

	fmt.Print("Введите имя cтудента:\n")
	fmt.Scan(&nameOfStudent)

	fmt.Print("Введите название группы:\n")
	fmt.Scan(&nameOfGroup)

	fmt.Print("Введите порядковый номер студента:\n")
	fmt.Scan(&serialNumber)

	if serialNumber / 2 == 0 {
		fmt.Println("Студент", nameOfStudent, "попадает в группу", groupSecond)


	} else {
		fmt.Println("Студент", nameOfStudent, "попадает в группу", groupFirst)
	}

}

//-----------------------------------------------------------

package main

import "fmt"

func main() {

	fmt.Println("---Прогрессивный налог---")

	var tax float64

	fmt.Println("Сумма Вашего дохода?:")

	var income float64
	fmt.Scan(&income)

	if income > 10000 {

		tax = 10000 * 0.13

		if income > 50000 {

			tax += (50000 - 10000) * 0.2 + (income - 50000) * 0.3

		} else {

			tax += (income - 10000) * 0.2
		}

	} else {

		tax = income * 0.13

	}
	fmt.Println("Сумма Вашего налога составит:",tax)
}