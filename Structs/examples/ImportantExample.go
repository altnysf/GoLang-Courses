package Examples

import (
	"fmt"
	"strconv"
)

func MainFunc() {
	fmt.Println("-------------------- ** --------------------")
	fmt.Println("First Way of Create a User!")

	user1 := &User{
		ID:        1,
		FirstName: "Yusuf",
		LastName:  "ALTUN",
		UserName:  "ysfaltn",
		Age:       28,
		Pay: &Payment{
			Salary: 4444,
			Bonus:  556,
		},
	}
	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetUserName())
	fmt.Println(user1.GetPayment())                                                 // OUTPUT "5000"
	fmt.Println("Salary : ", user1.GetPayment())                                    // OUTPUT "Salary :  5000"
	fmt.Println("Salary : " + strconv.FormatFloat(user1.GetPayment(), 'g', -1, 64)) // OUTPUT "Salary :  5000" (There is Conversion Operation)

	fmt.Println("-------------------- ** --------------------")

	fmt.Println("Second Way of Create a User!")

	user2 := NewUser()
	user2.FirstName = "Yusuf"
	user2.LastName = "ALTUN"
	user2.UserName = "GOPHER"
	user2.Age = 28
	user2.Pay = &Payment{Salary: 950, Bonus: 1050}

	fmt.Println(user2.GetFullName())
	fmt.Println(user2.GetUserName())
	fmt.Println(user2.GetPayment())                                                 // OUTPUT "2000"
	fmt.Println("Salary : ", user2.GetPayment())                                    // OUTPUT "Salary :  2000"
	fmt.Println("Salary : " + strconv.FormatFloat(user2.GetPayment(), 'g', -1, 64)) // OUTPUT "Salary : 2000" (There is Conversion Operation)

	fmt.Println("-------------------- ** --------------------")

}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

type Payment struct {
	Salary float64
	Bonus  float64
}

func NewPayment() *Payment {
	p := new(Payment)
	return p
}

func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}
