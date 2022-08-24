package Examples

import (
	"fmt"
	"strconv"
	"time"
)

func CarExecute(c Carface) {
	fmt.Println("\n")
	fmt.Println("Information of Car Status : \n" + c.Information())
	fmt.Println("\n")

	msg := ""
	msg2 := ""

	isRun := c.Run()

	if isRun {
		msg = " is starting."
		time.Sleep(3 * time.Second)
		msg2 = " started."
		time.Sleep(3 * time.Second)
	} else {
		msg = " is not starting."
		time.Sleep(3 * time.Second)
		msg2 = " has a FAULT..!!"
		time.Sleep(3 * time.Second)
	}

	fmt.Println("The Car " + msg + ".")
	time.Sleep(3 * time.Second)
	fmt.Println("The Car " + msg2 + ".")
	time.Sleep(3 * time.Second)

	isStop := c.Stop()

	if isStop {
		msg = " stopped."
		time.Sleep(3 * time.Second)
	} else {
		msg = " could not be stopped. There is a FAULT.!!!"
		time.Sleep(3 * time.Second)
	}

	fmt.Println("The Car " + msg + ".")
	time.Sleep(3 * time.Second)
}

func Example1() {
	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "F50"
	ferr.Color = "Red"
	ferr.Speed = 340
	ferr.Price = 1.75
	ferr.Special = true

	fmt.Println(ferr.Information())

	fmt.Println("-*------*------*------*------*------*------*------*-")

	merc := new(Merc)
	merc.Brand = "Mercedes"
	merc.Model = "CLX"
	merc.Color = "Black"
	merc.Speed = 290
	merc.Price = 0.65
	merc.Special = true

	fmt.Println(merc.Information())

	/* Car Executing Operations */

	CarExecute(ferr)
	CarExecute(merc)
}

/* Car Struct */
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

/* Special Production Info */

type SpecialProduction struct {
	Special bool
}

/* Struct for FERRARI */
type Ferrari struct {
	Car               /* Composition */
	SpecialProduction /* Composition */
}

/* Struct for MERCEDES */
type Merc struct {
	Car               /* Composition */
	SpecialProduction /* Composition */
}

/* Carface Interface Definition */

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return false
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : " + x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million $"
	Add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special : " + Add
	}
	return ret
}

func (_ Merc) Run() bool {
	return true
}

func (_ Merc) Stop() bool {
	return true
}

func (x *Merc) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : " + x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million $"
	Add := "No"
	if x.Special {
		ret += "\n" + "\t" + "Special : " + Add
	}
	return ret
}
