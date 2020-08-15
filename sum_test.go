package main

import "testing"

func TestSum(t *testing.T) {

    numbers := [5]int{1, 2, 3, 4, 5}

    got := Sum(numbers)
    want := 15

    if want != got {
        t.Errorf("got %d want %d given, %v", got, want, numbers)
    }
}

func Sum(num [5]int) int{
    var sum int = 0
    for _,n:= range num{
        sum+=n
    }
    return sum
}
