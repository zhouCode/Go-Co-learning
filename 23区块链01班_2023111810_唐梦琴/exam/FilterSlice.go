package main

import "fmt"

func FilterSliceNotGreater(nums []int, threshold int) []int {

    var res []int

    for _, num := range nums {

        if num <= threshold {

            res = append(res, num)

        }

    }

    return res

}