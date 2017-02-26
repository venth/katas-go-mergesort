package main


func main() {
}

func mergeSort(array *[]int) *[]int {
    performInPlaceMergeSort(array, 0, len(*array) - 1)
    return array
}

func performInPlaceMergeSort(array *[]int, start, end int) {
    if (start < end) {
        medium := start + (end - start) / 2

        performInPlaceMergeSort(array, start, medium)
        performInPlaceMergeSort(array, medium + 1, end)

        mergeTwoRangesWithSorting(array, start, medium, medium + 1, end)
    }
}

func mergeTwoRangesWithSorting(array *[]int, start1, end1, start2, end2 int) {
    leftIdx := start1
    rightIdx := start2
    var left, right int

    for leftIdx <= end1 && rightIdx <= end2 {
        left = (*array)[leftIdx]
        right = (*array)[rightIdx]

        if left > right {
            moveRightInsideRanges(array, leftIdx, end1, start2, rightIdx)
            rightIdx++
        } else  {
            leftIdx++
        }
    }
}

func moveRightInsideRanges(array *[]int, start1, end1, start2, end2 int) {
    endValue := (*array)[end2]
    moveRight(array, start2, end2)
    (*array)[start2] = (*array)[end1]
    moveRight(array, start1, end1)
    (*array)[start1] = endValue
}

func moveRight(array *[]int, start, end int) {
    for start < end {
        (*array)[end] = (*array)[end - 1]
        end--
    }
}
