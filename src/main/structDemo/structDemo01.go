package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}
type Student struct {
	Human
	// embedded field, it means Student struct includes all fields that Human has.
	specialty string
}

func main() {
	// instantiate and initialize a student
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	// access fields
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His specialty is ", mark.specialty)
	// modify mark's specialty
	mark.specialty = "AI"
	fmt.Println("Mark changed his specialty")
	fmt.Println("His specialty is ", mark.specialty)
	fmt.Println("Mark become old. He is not an athlete anymore")
	mark.age = 46
	mark.weight += 60
	fmt.Println("His age is", mark.age)
	fmt.Println("His weight is", mark.weight)

	fmt.Println("----------------------------------")

	c1 := Circle{10}
	c2 := Circle{25}
	r1 := Rectangle{9, 4}
	r2 := Rectangle{12, 2}
	fmt.Println("Area of c1 is: ", c1.Area())
	fmt.Println("Area of c2 is: ", c2.Area())
	fmt.Println("Area of r1 is: ", r1.Area())
	fmt.Println("Area of r2 is: ", r2.Area())

	fmt.Println("----------------------------------")

	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm3")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())
	// Let's paint them all black
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())

	fmt.Println(WHITE)
	fmt.Println(YELLOW)
	fmt.Println(WHITE)
}
