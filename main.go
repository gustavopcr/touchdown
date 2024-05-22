package main

import(
	"fmt"
)

var count = make([]int, 17, 17)
var points = [...]int{ 3, 6, 7, 8}

func max( nums ...int) int{
	if len(nums)==0{
		return 0
	}
	max := nums[0]
	
	for i:=0; i<len(nums); i++{
		if nums[i] > max{
			max = nums[i]
		}
	}
	return max
}

func min( nums ...int) int{
	if len(nums)==0{
		return 0
	}
	min := nums[0]
	
	for i:=0; i<len(nums); i++{
		if nums[i] < min{
			min = nums[i]
		}
	}
	return min
}


func countPos(x int) int{
	for i:=1; i<len(count); i++{
		for _, p := range points{
			if i-p >= 0{
				count[i] += count[i-p]
			}
		}
	}
	return count[x]
	//return max(countPos(x-3), countPos(x-6), countPos(x-7), countPos(x-8))
}
/*  0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
	1 0 0 0 0 0 0 0 0 0  0  0  0  0  0  0
    1 0 0 1 0 0 2 1 1 2  1  1  
*/

/*
  12
  3 3 3 3
  3 3 6
  6 6
*/

/*
 15
 3 3 3 3 3
 3 3 3 6
 6 6 3
 7 8 
*/

func countPos2(x int) int{
	for i:=1; i<len(count); i++{
		for _, p := range points{
			if i-p >= 0{
				count[i] += count[i-p]
			}
		}
	}
	return count[x]
	//return max(countPos(x-3), countPos(x-6), countPos(x-7), countPos(x-8))
}

func main(){
	count[0] = 1

	x := 5
	fmt.Println(countPos(x))
	for i, e := range count{
		fmt.Printf("count[%d]: %d\n", i, e)
	}
}

//7

// countPos(4)+1, countPos(1)+1, countPos(0)+1, countPos(-1)+1

// countPos(4) -> countPos(1)+1, countPos(-2)+1, countPos(-3)+1, countPos(-4)+1



// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
// 0 0 0 1 0 0 2 1 1 2 1   1  2 2   2  


// 9
/*
	3 3 3
	3 6
*/
//