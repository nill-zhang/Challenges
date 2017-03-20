package main

import (
	"fmt"
)


//5.1 Insertion: You are given two 32-bit numbers, N and M, and two bit positions, i and
//j. Write a method to insert M into N such that M starts at bit j and ends at bit i. You
//can assume that the bits j through i have enough space to fit all of M. That is, if
//M = 10011, you can assume that there are at least 5 bits between j and i. You would not, for
//example, have j = 3 and i = 2, because M could not fully fit between bit 3 and bit 2.
//EXAMPLE
//Input: N 10000000000, M
//Output: N = 10001001100
//Hints: #137, #769, #215

func Insertion(M,N,j int) int{

	return N | M << uint(j)



}

// You have an integer and you can flip exactly one bit from a 0 to a 1. Write code to
// find the length of the longest sequence of ls you could create.
// EXAMPLE
// Input: 1775
// Output: 8
// (or: 11011101111)


// FlipToWin flips one and only one 0 bit of a integer x, and return
// the longest 1s it get
func FlipToWin(x int) int{

	_,_,msb := Count(x)
	// sections stores counters
	sections :=[]int{}
	// counter stores how many consecutive 1s are encountered
	counter := 0
	// if x is zero, early return
	// msb:Most Significant bit
	if msb ==  -1 {
		return 1
	}
	for i:=0;i<=msb;i++{
                // two things here:
		// 1. counter will do self-addition until current bit is not set
		// so, don' forget continue
		// 2. when the msb is 1, you will leave the loop without update the sections
		// update it outside the loop by checking counter recently increased or not
		if isBitSet(x,uint(i)){
			counter++
			continue

		}
		// reset counter
		sections =  append(sections,counter)
		counter = 0
	}
	// if counter self-adds all the way up until msb(inclusive) which means x is all 1s
	// the longest 1s we can get by swapping 0 to 1 is x itself
	if counter == msb+1{
		return counter
	}

	if counter != 0{
		sections =  append(sections,counter)
	}

	max := sections[0]
	for j:=1;j<len(sections);j++{
		// find the longest sum of two sections
		if sections[j]+sections[j-1] > max {
			max = sections[j]+sections[j-1]
		}
	}
	// besides the 1s, don't forget the 0 in the middle
	return max + 1





}

// 5.6 Conversion: Write a function to determine the number of bits you would need to flip to convert
// integer A to integer B.
// EXAMPLE
// Input: 29 (or: 11101), 15 (or: 01111)
// Output: 2
// Hints: #336, #369
// pg286

// BitsForAToB returns the number of bits need to be flipped to convert A to B
func BitsForAToB(x, y int) int{

	// bits that need to be converted is where they are
	// opposite of each other, so first we can xor x and y
	// the 1s in the result is how many bits need to be flipped
	// then call Count to get the number of 1s in the result integer

	_, num, _ := Count(x ^ y)
	return num

}



// SetBit sets the bit at the position of x's binary form
func SetBit(x int, position uint) int{
	return x | 1<<position

}

// GetBit gets the bit at the position of x's binary string form
func GetBit(x int, position uint) int{

	return x & (1<<position)

}

// FlipBit switches the bit at the position of x
func FlipBit(x int, position uint) int{
	return x ^ (1<<position)


}

// ClearBit clears the bit at the position of x's binary string form
func ClearBit(x int, position uint) int{
	return x &^(1<<position)


}
// right shift x
// or leftshift 1
// isBitSet checks if x's position is set to 1
func isBitSet(x int, position uint) bool{

	return ((x >> position) & 1) == 1


}


// clearBitsMSBthroughI clears bits from MSB through i(i from 0)
func ClearBitsMSBThroughI(x int, i uint) int{

	return (x & (1<<i)-1)



}

// clearBitsithrough0 clears bits  0 through i
func ClearBitsIThrough0(x int, i uint) int{

	// -1 is all 1s in two-complement representation
	return x & (-1 << (i+1))


}

func Count(x int) (num0,num1,MSB int){
	PostMSB := uint(0)

	// get the Most Significant bit
	for {
		y := 1 << PostMSB
		if y > x{
			break
		}
		PostMSB++
		//
	}
        //  Watch out, i must convert PreMSB-1 to int type
	// because if I don't,j will be type of unsigned integer, which is always positive
	// so the for loop will not break
	// j starts with the MSB
	for j:=int(PostMSB-1);j>=0;j--{
		if x & (1 << uint(j)) == (1 << uint(j)){
			num1++
			continue
		}
		num0++
	}
	MSB = int(PostMSB-1)
	return


}

func main(){

	cnt0,cnt1,_ := Count(243)
	fmt.Printf("Number of 0S:%v,Number of 1s:%v\n",cnt0,cnt1)
	cnt0,cnt1,_ = Count(8)
	fmt.Printf("Number of 0S:%v,Number of 1s:%v\n",cnt0,cnt1)

	fmt.Printf("Bits need to be flipped:%v\n",BitsForAToB(29,15))
	fmt.Printf("Before Insertion: %b\n After Insertion: %b\n",2048,Insertion(19,2048,2,0))
	fmt.Printf("Number of longest 1s if swap one 0 --> 1: %v(%#b)\n",FlipToWin(1775),1775)
	fmt.Printf("Number of longest 1s if swap one 0 --> 1: %v(%#b)\n",FlipToWin(0),0)
	fmt.Printf("Number of longest 1s if swap one 0 --> 1: %v(%#b)\n",FlipToWin(1),1)
	fmt.Printf("Number of longest 1s if swap one 0 --> 1: %v(%#b)\n",FlipToWin(255),255)
	fmt.Printf("Number of longest 1s if swap one 0 --> 1: %v(%#b)\n",FlipToWin(56),56)
	fmt.Printf("Number of longest 1s if swap one 0 --> 1: %v(%#b)\n",FlipToWin(556),556)
	fmt.Printf("Number of longest 1s if swap one 0 --> 1: %v(%#b)\n",FlipToWin(85962),85962)

}