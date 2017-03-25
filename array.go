package main

func Intersection(a, b []int)[]int{
	mapA := map[int]int{}
	result := []int{}
	for _,i := range a{
		mapA[i]++
	}

	for _,j :=range b{

		if mapA[j] != 0{
			result = append(result,j)
		}

	}

	return result

}

// [4,5,7,89,34,64,2,91,9,13,24]
// hash table
func Difference(a,b []int){





}

func main(){

	Intersection([]int{1,2,3,3,2,2,4,34,222,78,9876},[]int{9876,132})

}