package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    file, err := os.Create("nums.txt")
    if err != nil {
       log.Fatal(err)
    }

    file, err = os.Open("nums.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    var max, min, count int
    var nums []int
    for scanner.Scan() {
        number, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }

        nums = append(nums, number)
        if count == 0 {
            min = number
            max = number
        } else {
            if number > max {
                max = number
            } else if number < min {
                min = number
            }
        }
        count++
    }

    err = file.Close()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("max: %d\nmin: %d\nnums: %d", max, min, nums)
}

