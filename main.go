package main

import(
	"fmt"
)

func main() {
    var input_word string
    char_count := make(map[string]int)

    fmt.Scanf("%s", &input_word)

    for _, value := range input_word {
        char_count[string(value)]++
    }

    fmt.Println(char_count)
}

