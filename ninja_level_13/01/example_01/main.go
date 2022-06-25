package example_01

import (
	"GoLangCourseExercisesMore/ninja_level_13/01/example_01/dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func MainFunc() (canine, int) {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
	return fido, dog.YearsTwo(20)
}
