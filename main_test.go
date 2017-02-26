package main

import (
    "reflect"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestMergeSort(t *testing.T) {
    Convey("empty array is sorted", t, func() {
        Convey("given an empty array", func() {
            givenEmptyArray := []int{}
            Convey("when sorts given array", func() {
                sorted := mergeSort(&givenEmptyArray)
                Convey("then returns empty array", func() {
                    emptyArray := []int{}
                    So(reflect.DeepEqual(*sorted, emptyArray), ShouldBeTrue)
                })
            })
        })
    })

    Convey("one element array is sorted", t, func() {
        Convey("given one element array", func() {
            givenOneElementArray := []int{ 1 }
            Convey("when sorts given array", func() {
                sorted := mergeSort(&givenOneElementArray)
                Convey("then returns empty array", func() {
                    sortedOneElementArray := []int{ 1 }
                    So(reflect.DeepEqual(*sorted, sortedOneElementArray), ShouldBeTrue)
                })
            })
        })
    })

    Convey("two element array is sorted", t, func() {
        Convey("given two element array", func() {
            givenTwoElementUnsortedArray := []int{ 2, 1 }
            Convey("when sorts given array", func() {
                sorted := mergeSort(&givenTwoElementUnsortedArray)
                Convey("then returns sorted array", func() {
                    sortedTwoElementArray := []int{ 1, 2 }
                    So(reflect.DeepEqual(*sorted, sortedTwoElementArray), ShouldBeTrue)
                })
            })
        })
    })

    Convey("three element array is sorted", t, func() {
        Convey("given three element array", func() {
            givenThreeElementUnsortedArray := []int{ 3, 2, 1 }
            Convey("when sorts given array", func() {
                sorted := mergeSort(&givenThreeElementUnsortedArray)
                Convey("then returns sorted array", func() {
                    sortedThreeElementArray := []int{ 1, 2, 3 }
                    So(reflect.DeepEqual(*sorted, sortedThreeElementArray), ShouldBeTrue)
                })
            })
        })
    })

    Convey("five element array is sorted", t, func() {
        Convey("given five element array", func() {
            givenFiveElementUnsortedArray := []int{ 5, 4, 3, 2, 1 }
            Convey("when sorts given array", func() {
                sorted := mergeSort(&givenFiveElementUnsortedArray)
                Convey("then returns sorted array", func() {
                    sortedFiveElementArray := []int{ 1, 2, 3, 4, 5 }
                    So(reflect.DeepEqual(*sorted, sortedFiveElementArray), ShouldBeTrue)
                })
            })
        })
    })

    Convey("four element array is sorted", t, func() {
        Convey("given four element array", func() {
            givenFourElementUnsortedArray := []int{ 4, 3, 2, 1 }
            Convey("when sorts given array", func() {
                sorted := mergeSort(&givenFourElementUnsortedArray)
                Convey("then returns sorted array", func() {
                    sortedFourElementArray := []int{ 1, 2, 3, 4 }
                    So(reflect.DeepEqual(*sorted, sortedFourElementArray), ShouldBeTrue)
                })
            })
        })
    })

    Convey("array containing same elements is sorted", t, func() {
        Convey("given four element array with same elements", func() {
            givenSameElementArray := []int{ 4, 4, 4, 4, 4 }
            Convey("when sorts given array", func() {
                sorted := mergeSort(&givenSameElementArray)
                Convey("then returns sorted array", func() {
                    sortedSameElementArray := []int{ 4, 4, 4, 4, 4 }
                    So(reflect.DeepEqual(*sorted, sortedSameElementArray), ShouldBeTrue)
                })
            })
        })
    })
}
