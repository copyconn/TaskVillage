package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Names = []string{"Kassandra", "John", "Alexa", "Mike", "Vanessa", "Leonard", "Emily", "Brandon", "Beth", "Andrew"}
var Seconds = []string{"Hutt", "Brown", "Scott", "Stone", "Ride", "Rowling", "Craud", "Faze", "Loutt", "Longhead"}
var Classes = []string{"Peasant", "Witch", "Vampire", "Werewolf"}

type man struct {
	name   int
	second int
	class  int
}

func CreatePerson() []man { //тип возвращаемого результата: [] - массив, man - со структурами man
	men := []man{} //создаем пустой массив людей
	for i := 0; i < 50; i++ {
		Person := man{name: rand.Intn(10), second: rand.Intn(10), class: rand.Intn(4)} //генерируем рандомные индексы в имя, фамилию и класс, по которым потом подтянем имя, фамилию и класс из соответствующих массивов
		//fmt.Println("man", i, Names[Person.name], Seconds[Person.second])
		men = append(men, Person) //добавляем в массив людей нового жителя
	}
	return men //возвращаем готовый массив людей
}

func DrinkTea(men []man) []man { //тип принимаемого значения: men - название массива, [] - показывает, что это тип массив, man - со структурами man
	PersonIndex1 := rand.Intn(len(men))
	Person1 := men[PersonIndex1] //случайно выбираем 1 человека из массива людей
	PersonIndex2 := rand.Intn(len(men))
	Person2 := men[PersonIndex2] //случайно выбираем 2 человека из массива людей

	p := rand.Intn(100) //генерируем рандомный процент из 100

	if Person2.class == 2 && p <= 5 { //если второй житель вампир и процент попадает в 5%
		for Person1.class == 2 { //перевыбираем пока не выпадет не вампир
			PersonIndex1 = rand.Intn(len(men))
			Person1 = men[PersonIndex1]
		}

		if p <= 1 { //если выпал 1%
			fmt.Println("DAWNEXCEPTION", Names[Person2.name], Seconds[Person2.second], "(", Classes[Person2.class], ") is killed by", Names[Person1.name], Seconds[Person1.second], "(", Classes[Person1.class], ")")
			men = append(men[:PersonIndex2], men[PersonIndex2+1:]...) //склеиваем новый массив, исключая из него выбранный ранее элемент по индексу
			return men
		} else if Person1.class == 0 || Person1.class == 1 { //если первый житель Peasant или Witch
			Person1.class = 2 //становится вампиром
		} else { //если попали в 5% и не прошли первые два условия, то вампир убивает жителя
			fmt.Println("MURDEREXCEPTION", Names[Person1.name], Seconds[Person1.second], "(", Classes[Person1.class], ") is killed by", Names[Person2.name], Seconds[Person2.second], "(", Classes[Person2.class], ")")
			men = append(men[:PersonIndex1], men[PersonIndex1+1:]...) //склеиваем новый массив, исключая из него выбранный ранее элемент по индексу
			return men
		}
	} else { //вот всех остальных случаях, если сгенерированный процент не попал в 5%
		fmt.Println(Names[Person1.name], Seconds[Person1.second], "(", Classes[Person1.class], ") drinks tea with ", Names[Person2.name], Seconds[Person2.second], "(", Classes[Person2.class], ")")
	}
	return men
}

func AddPerson(men []man) []man {
	if len(men) < 50 { //проверка на то, полная ли деревня или можно еще кого добавить
		Person := man{name: rand.Intn(10), second: rand.Intn(10), class: rand.Intn(4)} //генерируем рандомные индексы в имя, фамилию и класс, по которым потом подтянем имя, фамилию и класс из соответствующих массивов
		fmt.Println(Names[Person.name], Seconds[Person.second], "(", Classes[Person.class], ") appeared in the Village")
		men = append(men, Person) //добавляем в массив людей нового жителя
		return men
	} else {
		fmt.Println("Everybody sleeps. Nothing is happening. Village is full")
	}
	return men
}

func DelPerson(men []man) []man {
	PersonIndex := rand.Intn(len(men)) //выбираем случайный элемент в массиве по индексу
	Person := men[PersonIndex]         //добавляем в переменную все атрибуты выбранного ранее элемента
	fmt.Println(Names[Person.name], Seconds[Person.second], "(", Classes[Person.class], ") left the Village")
	men = append(men[:PersonIndex], men[PersonIndex+1:]...) //склеиваем новый массив, исключая из него выбранный ранее элемент по индексу
	return men
}

func MeetWitch(men []man) []man {
	PersonIndex1 := rand.Intn(len(men))
	Person1 := men[PersonIndex1] //случайно выбираем 1 человека из массива людей
	PersonIndex2 := rand.Intn(len(men))
	Person2 := men[PersonIndex2]

	for Person1.class == 1 {
		PersonIndex1 = rand.Intn(len(men))
		Person1 = men[PersonIndex1]
	}

	for Person2.class != 1 { //перевыбираем человека пока не выпадет ведьма
		PersonIndex2 = rand.Intn(len(men))
		Person2 = men[PersonIndex2]
	}

	fmt.Println("MURDEREXCEPTION", Names[Person1.name], Seconds[Person1.second], "(", Classes[Person1.class], ") is killed by", Names[Person2.name], Seconds[Person2.second], "(", Classes[Person2.class], ")")
	men = append(men[:PersonIndex1], men[PersonIndex1+1:]...) //склеиваем новый массив, исключая из него выбранный ранее элемент по индексу
	return men
}

func MeetWolf(men []man) []man {
	PersonIndex1 := rand.Intn(len(men))
	Person1 := men[PersonIndex1] //случайно выбираем 1 человека из массива людей
	PersonIndex2 := rand.Intn(len(men))
	Person2 := men[PersonIndex2]

	for Person1.class == 3 {
		PersonIndex1 = rand.Intn(len(men))
		Person1 = men[PersonIndex1]
	}

	for Person2.class != 3 { //перевыбираем человека пока не выпадет ведьма
		PersonIndex2 = rand.Intn(len(men))
		Person2 = men[PersonIndex2]
	}

	fmt.Println("MURDEREXCEPTION", Names[Person1.name], Seconds[Person1.second], "(", Classes[Person1.class], ") is killed by", Names[Person2.name], Seconds[Person2.second], "(", Classes[Person2.class], ")")
	men = append(men[:PersonIndex1], men[PersonIndex1+1:]...) //склеиваем новый массив, исключая из него выбранный ранее элемент по индексу
	return men
}

func main() {
	startDate := time.Date(1650, 1, 1, 0, 0, 0, 0, time.UTC).Unix() //Unix превращает всю дату в одно число секунд
	finishDate := time.Date(1655, 12, 31, 23, 59, 59, 0, time.UTC).Unix()
	//step := startDate + 86400 //прибавляем 1 день к стартовой дате
	//fmt.Println(time.Unix(startDate, 0))
	//fmt.Println(time.Unix(finishDate, 0))
	//fmt.Println(time.Unix(step, 0))

	men := CreatePerson() //присваиваем возвращаемый функцией массив men в массив men, который будет виден в основной функции main
	p := 0                //счетчик порядкового номера дня в целом
	n := 28               //счетчик полнолуний

	for i := startDate; i < finishDate; i += 86400 {
		date := time.Unix(i, 0)   //превращаем обратно число секунд в дату
		num := date.Day()         //номер дня в месяце
		month := date.Month()     //название месяца
		year := date.Year()       //номер года
		weekday := date.Weekday() //название дня недели
		fmt.Println(num, month, year, "(", weekday, ")")

		t := rand.Intn(4) //возвращает значения от 0 до n, не включая ее
		p = p + 1         //считаем порядковый номер дня в целом

		if len(men) == 0 {
			men = CreatePerson()
		}

		if p == n {
			fmt.Println("MEETING WEREWOLF")
			men = MeetWolf(men)
			n = n + 28
		} else if (num == 31 && int(month) == 10) || (num == 13 && int(weekday) == 5) { //если пятница 13 или 31 октября
			fmt.Println("MEETING WITCH")
			men = MeetWitch(men)
		} else if t == 0 {
			fmt.Println("Everybody sleeps. Nothing is happening")
		} else if t == 1 {
			men = DrinkTea(men)
		} else if t == 2 {
			men = AddPerson(men)
		} else if t == 3 {
			men = DelPerson(men)
		}
	}
	fmt.Println("Total number of people:", len(men))
	fmt.Println("Total number of days:", p)
}
