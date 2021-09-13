package main

import(
    "fmt"
    "regexp"
)

func readInput(prompt string) string {
    var input string
    fmt.Println(prompt)
    fmt.Scanf("%s", &input)
    return input
}

func splitCodeIntoLetters(code string) []string {
    re := regexp.MustCompile(`\w\w`)
    splittedCode := re.FindAllString(code, -1)
    return splittedCode
}

func getCipherCharAssociation() map[string]string {
    cipherCharAssociation := make(map[string]string)
    lettersToDecipherCode := "abcdefghijkmnopqrstuvwxyz "

    for i, value := range lettersToDecipherCode {
        var index string
        value := string(value)
        if i < 10 {
            index = fmt.Sprintf("0%d", i)
        } else {
            index = fmt.Sprintf("%d", i)
        }
        cipherCharAssociation[index] = value
    }
    return cipherCharAssociation
}

func translateCipherLetters(cipherLetters []string) string {
    var decipherCode string
    cipherCharAssociation := getCipherCharAssociation()

    for _, value := range cipherLetters {
        decipherCode = fmt.Sprintf("%s%s", decipherCode, cipherCharAssociation[value])
    }
    return decipherCode
}

func codeToString() string {
    code := readInput("Enter a code to decipher:")
    splittedCode := splitCodeIntoLetters(code)
    decipherCode := translateCipherLetters(splittedCode)
    return decipherCode
}

func main() {
    fmt.Println(codeToString())
}


