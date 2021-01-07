package quiz_test

import (
	"fmt"

	"github.com/babygoat/quiz"
)

func ExampleConstArray() {
	fmt.Println(quiz.ConstArray())
	// Output: 2
}

func ExampleEmbeddedJson() {
	fmt.Printf("%s", quiz.EmbeddedJson())
	// Output: "2021-01-07T17:19:09Z"
}
