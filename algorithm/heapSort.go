package main

import "fmt"

func heapSort(arr []int) {
    length := len(arr) - 1

    for i := length/2 - 1; i >= 0; i-- {
        heap(arr, i, length)
    }

    for i := length; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heap(arr, 0, i - 1)
    }
}

func heap(arr []int, i int, end int) {
    left := i*2 + 1
    if left > end {
        return
    }
    next := left
    right := 2*i + 2
    if right <= end && arr[next] < arr[right] {
        next = right
    }
    if arr[i] >= arr[next] {
        return
    }
    arr[i], arr[next] = arr[next], arr[i]
    heap(arr, next, end)
}

func main() {
     arr := []int{5,3,2,4,1}
     heapSort(arr)
     fmt.Println(arr)
}
