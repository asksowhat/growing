package main

import "fmt"

func quickSort(arr []int, left int, right int) {
    if left >= right {
        return
    }

    povit := arr[left]
    i, j := left, right

    for i < j {
        for i < j && arr[j] >= povit {
            j--
        }
        arr[i] = arr[j]

        for i < j && arr[i] <= povit {
            i++
        }
        arr[j] = arr[i]
    }
    arr[i] = povit

    quickSort(arr, left, i-1)
    quickSort(arr, i+1, right)
}

func main() {
    arr := []int{4,2,5,1,3}
    quickSort(arr, 0, len(arr)-1)
    fmt.Println(arr)

}
