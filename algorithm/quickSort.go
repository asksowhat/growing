package main

import "fmt"

func quickSort(arr []int, left int, right int) {
    if left >= right {
        return
    }
    povit := paritition(arr, left, right)
    quickSort(arr, left, povit-1)
    quickSort(arr, povit+1, right)
}

func paritition(arr []int, left int, right int) int {
        povit := arr[left]

        for left < right {
            for left < right && arr[right] >= povit {
                right--
            }
            arr[left] = arr[right]

            for left < right {
                left++
            }
            arr[right] = arr[left]
        }
        arr[left] = povit
        return left
}

func main() {
    arr := []int{5,2,1,4,3}
    quickSort(arr, 0, len(arr)-1)
    fmt.Println(arr)
}
