//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("Программа вычисления квадратного корня")
//	fmt.Println("Введите целое число:")
//
//	var value int
//	fmt.Scan(&value)
//
//	square := value * value
//
//	fmt.Println("Квадрат числа:", value, "равен", square)
//}


//-------------------------------------------------------
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("Введите стоимость товара:")
//	var product int
//	fmt.Scan(&product)
//
//	fmt.Println("Введите стоимость доставки:")
//	var delivery int
//	fmt.Scan(&delivery)
//
//	fmt.Println("Введите скидку:")
//	var discount int
//	fmt.Scan(&discount)
//
//	var result int = product + delivery - discount
//
//	fmt.Println("---------")
//	fmt.Println("Цена товара", product)
//	fmt.Println("Цена товара", delivery)
//	fmt.Println("Цена товара", discount)
//	fmt.Println("Итого:", result)
//}

//----------------------------------------------------

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("Здравствуйте, приветствуем Вас в нашей форме регистрации!")
//
//	fmt.Println("Введите Ваше имя:")
//	var login string
//	fmt.Scan(&login)
//
//	fmt.Println("Введите Ваш пароль:")
//	var password string
//	fmt.Scan(&password)
//
//	fmt.Println("Введите Ваш возраст:")
//	var age int
//	fmt.Scan(&age)
//
//	fmt.Println("Поздравляю", login, "теперь вы зарегистрированы!")
//	fmt.Println("Ваш пароль:", password)
//	fmt.Println("Ваш возраст:", age )
//}

//----------------------------------------------------------------

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("Введите Ваши данные!")
//
//	fmt.Println("Ваш логин:")
//	var name string
//	fmt.Scan( &name)
//
//	fmt.Println("Введите Ваш пароль")
//	var password string
//	fmt.Scan(&password)
//
//
//	fmt.Println(name, "Вы успешно зашли!")
//
//}

//-------------------------------------------

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Текстовый квест")
//
//	fmt.Println("Введите название планеты")
//	var planet string
//	fmt.Scan(&planet)
//
//	fmt.Println("Введите название звездной системы")
//	var star string
//	fmt.Scan(&star)
//
//	fmt.Println("Введите имя рейнджера")
//	var ranger string
//	fmt.Scan(&ranger)
//
//	fmt.Println("Введите количество вознаграждения")
//	var money int
//	fmt.Scan(&money)
//
//	fmt.Println("----------------------------------")
//
//	fmt.Println("Как вам", ranger, "известно, мы - раса мирная, поэтому на наших военных кораблях используются наемники с других планет. Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты", planet, "системы", star,".")
//
//	fmt.Print("\n")
//
//	fmt.Println("Но случилась неприятность - в связи с большими потерями, в последнее время, престиж профессии сильно упал, и теперь не так-то просто завербовать призывников. Главный комиссар планеты", planet, "впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!")
//
//	fmt.Print("\n")
//
//	fmt.Println(ranger, "вы должны прибыть на планету", planet, "системы", star, "и помочь выполнить план призыва. Мы готовы выплатить вам премию в", money, "кредитов за эту маленькую услугу.")
//}

//--------------------------------------------------------------------------------------------------


//
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//
//	var count1 int
//
//	fmt.Println("Прибываем на остановку Улица Программистов. В салоне пассажиров:")
//	var arrival_1 int
//	fmt.Scan(&arrival_1)
//
//	fmt.Println("Сколько пассажиров вышло на остановке?")
//	var getOff_1 int
//	fmt.Scan(&getOff_1)
//
//	fmt.Println("Сколько пассажиров зашло на остановке?")
//	var entrance_1 int
//	fmt.Scan(&entrance_1)
//
//
//	count1 = entrance_1 + arrival_1 - getOff_1
//	fmt.Println("Отправляемся с остановки Улица Программистов. В салоне пассажиров:", count1)
//
//
//	fmt.Print("-----------Едем--------- \n")
//
//	var count2 int
//
//
//	fmt.Println("Прибываем на остановку Проспект Алгоритмов. В салоне пассажиров:", count1)
//
//
//	var getOff_2 int
//	fmt.Println("Сколько пассажиров вышло на остановке?")
//	fmt.Scan(&getOff_2)
//
//
//	var entrance_2 int
//	fmt.Println("Сколько пассажиров зашло на остановке?")
//	fmt.Scan(&entrance_2)
//
//
//	count2 = count1 + entrance_2 - getOff_2
//	fmt.Println("Отправляемся с остановки Проспект Алгоритмов. В салоне пассажиров:", count2)
//
//	fmt.Print("-----------Едем--------- \n")
//
//
//	var count3 int
//
//	fmt.Println("Прибываем на остановку Площадь Golang. В салоне пассажиров:",count2)
//
//	fmt.Println("Сколько пассажиров вышло на остановке?")
//	var getOff_3 int
//	fmt.Scan(&getOff_3)
//
//	fmt.Println("Сколько пассажиров зашло на остановке?")
//	var entrance_3 int
//	fmt.Scan(&entrance_3)
//
//
//	count3 = count2 + entrance_3 - getOff_3
//	fmt.Println("Отправляемся с остановки Площадь Golang. В салоне пассажиров:", count3)
//
//	fmt.Print("-----------Едем--------- \n")
//
//
//
//	fmt.Println("Прибываем на конечную остановку Сквер Потоков. В салоне пассажиров:", count3)
//
//	fmt.Println("Сколько пассажиров вышло на остановке?")
//	var getOff_4 int
//	fmt.Scan(&getOff_4)
//
//	fmt.Println("Пассажиры покинувшие маршрутку:", count3)
//
//	sum := (entrance_1 + entrance_2 + entrance_3) * 20
//	salary := sum / 4
//	fuel := sum / 5
//	fix := sum / 5
//	tax := sum / 5
//	totalIncome := sum - (salary + fuel + fix + tax)
//
//	fmt.Println("Всего заработали:", sum, "руб." )
//	fmt.Println("Зарплата водителя:", salary, "руб.")
//	fmt.Println("Расходы на топливо:", fuel, "руб.")
//	fmt.Println("Налоги:", fix, "руб.")
//	fmt.Println("Расходы на ремонт машины:", fix, "руб.")
//	fmt.Println("Итого доход:", totalIncome, "руб.")
//}




//---------------------------------------------------------------------------------------------------
//package main
//
//import (
//"fmt"
//)
//
//func main() {
//
//a := 42
//
//b := 153
//
//a, b = b, a
//
//fmt.Println("a:", a)
//
//fmt.Println("b:", b)
//
//}

//--------------------------------------------

//package main

// import(
//   "fmt"
// )

// func main() {

//   fmt.Println("Злостные вредители")

//   // height := 100
//   // growth := 50
//   // eat := 20

//   // height = (growth - eat)*2 + growth/2 + 100


//   // fmt.Println("Высота бамбука к середине третьего дня:", height)


//   // height := 100
//   // growth := 50
//   // eat := 20
//   // day := 6
//   // night := 5


//   // target := height + day * growth - eat * night
//   // fmt.Println("Бамбук вырастит до", target, "см за", day, "полных дней")

//   fmt.Println("Ведите размер саженцов при высадке в см:")
//   var height int
//   fmt.Scan(&height)

//   fmt.Println("Введите скорость роста бамбука в см:")
//   var growth int
//   fmt.Scan(&growth)

//   fmt.Println("Введите скорость поедания бамбука гусеницами в см:")
//   var eat int
//   fmt.Scan(&eat)

//   fmt.Println("Введите колличество дней:")
//   var day int
//   fmt.Scan(&day)

//   fmt.Println("Введите колличество ночей:")
//   var night int
//   fmt.Scan(&night)

//   fmt.Println("Введите целевую высоту бамбука в см:")
//   var target int
//   fmt.Scan(&target)

//   target = height + day * growth - eat * night


//   fmt.Println(target)

// }