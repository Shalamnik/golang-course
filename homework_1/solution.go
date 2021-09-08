package main

import(
	"fmt"
)

func main() {
// 1
//     var distance, gasoline_consumption, price float64
//     fmt.Println("Distance to point:")
// 	fmt.Scanf("%d", &distance)
//
//     fmt.Println("Gasoline consumption per 100 km:")
// 	fmt.Scanf("%d", &gasoline_consumption)
//
//     fmt.Println("Enter price:")
// 	fmt.Scanf("%d", &price)
//
// 	fmt.Println((distance / 100) * gasoline_consumption * price)

    // 2
//     var number int
//     fmt.Scanf("%d", &number)
//     first := number / 100
//     sec := number / 10 % 10
//     third := number % 10
//     fmt.Println(third, sec, first)

// 3
//     var (
//         num_1 int
//         num_2 int
//         num_3 int
//         max_num int
//         middle_num int
//         low_num int
//     )
//     fmt.Scanf("%d", &num_1)
//     fmt.Scanf("%d", &num_2)
//     fmt.Scanf("%d", &num_3)
//     if (num_3 < num_1) && (num_1 > num_2){
//         max_num = num_1
//         if (num_3 > num_2) {
//             middle_num = num_3
//             low_num = num_2
//         } else {
//             middle_num = num_2
//             low_num = num_3
//         }
//     } else if (num_1 < num_3) && (num_3 > num_2) {
//         max_num = num_3
//         if (num_1 > num_2) {
//             middle_num = num_1
//             low_num = num_2
//         } else {
//             middle_num = num_2
//             low_num = num_1
//         }
//     } else {
//         max_num = num_2
//         if (num_1 > num_3) {
//             middle_num = num_1
//             low_num = num_3
//         } else {
//             middle_num = num_3
//             low_num = num_1
//         }
//     }
//     fmt.Println(low_num, middle_num, max_num)

//4
    var palindrome int
    fmt.Scanf("%d", &palindrome)
    out_compare := (palindrome / 1000 == palindrome % 10)
    in_compare := (palindrome / 100 % 10 == palindrome / 10 % 10)
    fmt.Println(out_compare && in_compare)
}

