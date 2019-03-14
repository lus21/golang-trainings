package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type user struct {
	name     string
	birthday string
	gender   string
}

func (user *user) setBirthday(date string) {
	user.birthday = date
}

func (user *user) setGender(gender string) {
	user.gender = gender
}

func abs(integer int) int {
	if integer < 0 {
		return -integer
	}

	return integer
}

func calcAgeDifference(user1, user2 user) int {
	birthYear1, _, _ := parseBirthday(user1.birthday)
	birthYear2, _, _ := parseBirthday(user2.birthday)
	return abs(birthYear2 - birthYear1)
}

func parseBirthday(birthday string) (int, int, int) {
	birthDate := strings.Split(birthday, "/")
	birthYearString := birthDate[0]
	birthMonthString := birthDate[1]
	birthDayString := birthDate[2]
	birthYear, _ := strconv.Atoi(birthYearString)
	birthMonth, _ := strconv.Atoi(birthMonthString)
	birthDay, _ := strconv.Atoi(birthDayString)
	return birthYear, birthMonth, birthDay
}

func (user *user) getAge() int {
	year, month, day := time.Now().Date()
	monthInt := int(month)
	birthYear, birthMonth, birthDay := parseBirthday(user.birthday)

	if birthMonth > monthInt && birthDay > day {
		return year - birthYear - 1
	}

	return year - birthYear
}

func main() {
	user1 := user{
		name:     "Bob",
		birthday: "1989/11/19",
		gender:   "male",
	}

	user2 := user{
		name:     "Alice",
		birthday: "1994/08/12",
	}

	fmt.Println("user1", user1)
	fmt.Println("user2", user2)
	user1.setBirthday("1990/11/19")
	fmt.Println("user1 after changing birthday", user1)
	user2.setGender("female")
	fmt.Println("user2 after setting gender", user2)

	user1Age := user1.getAge()
	user2Age := user2.getAge()
	fmt.Println(user1.name, "age is", user1Age, user2.name, "age is", user2Age)

	ageDifference := calcAgeDifference(user1, user2)
	fmt.Println("age difference is", ageDifference)

}
