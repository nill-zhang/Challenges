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

func Insertion(M,N,j,i int) int{

	return N | M << uint(j)



}











//5.2 Binary to String: Given a real number between O and 1 (e.g., 0.72) that is passed in as a double, print
//the binary representation. If the number cannot be represented accurately in binary with at most 32
//characters, print "ERROR:'
//Hints: #743, #767, #7 73, #269, #297





//5.3 Flip Bit to Win: You have an integer and you can flip exactly one bit from a 0 to a 1. Write code to
//find the length of the longest sequence of ls you could create.
//EXAMPLE
//Input: 1775
//Output: 8
//(or: 11011101111)
//Hints: #759, #226, #374, #352




//5.4 Next Number: Given a positive integer, print the next smallest and the next largest number that
//have the same number of 1 bits in their binary representation.
//Hints: #747, #7 75, #242, #372, #339, #358, #375, #390





//5.6 Conversion: Write a function to determine the number of bits you would need to flip to convert
//integer A to integer B.
//EXAMPLE
//Input: 29 (or: 11101), 15 (or: 01111)
//Output: 2
//Hints: #336, #369
//pg286

// BitsForAToB returns the number of bits need to be flipped to convert A to B
func BitsForAToB(x, y int) int{
	// when the according bits need to be converted is they are
	// opposite of each other, so first we can xor the x and y
	// the 1s in the result is how many bits need to be flipped
	// I then call Count to get the number

	_, num := Count(x ^ y)
	return num

}



//5.7 Pairwise Swap: Write a program to swap odd and even bits in an integer with as few instructions as
//possible (e.g., bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, and so on).
//Hints: #745, #248, #328, #355
//. ... .... . . . . pg 286




//5.8 Draw Line: A monochrome screen is stored as a single array of bytes, allowing eight consecutive
//pixels to be stored in one byte. The screen has width w, where w is divisible by 8 (that is, no byte will
//be split across rows). The height of the screen, of course, can be derived from the length of the array
//and the width. Implement a function that draws a horizontal line from ( xl, y) to ( x2, y).
//The method signature should look something like:
//drawline(byte[] screen, int width, int xl, int x2, int y)
//Hints: #366, #387, #384, #397

// SetBit sets the bit at the position of x's binary form
func SetBit(x int, position uint) int{
	return x | 1<<position

}

// GetBit gets the bit at the position of x's binary string form
func GetBit(x int, position uint) int{

	return x & (1<<position)

}

// FlipBit switch the bit at the position of x's binary string form
func FlipBit(x int, position uint) int{
	return x ^ (1<<position)


}

// ClearBit clears the bit at the position of x's binary string form
func ClearBit(x int, position uint) int{
	return x &^(1<<position)


}
// right shift
// or 1 leftshift
// isBitSet checks if x's position is set
func isBitSet(x int, position uint) bool{

	return ((x >>(position-1)) & 1) == 1


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

func Count(x int) (num0,num1 int){
	PreMSB := uint(0)

	// get the Most Significant bit
	for {
		y := 1 << PreMSB
		if y > x{
			break
		}
		PreMSB++
		//
	}
        //  Watch out, i must convert PreMSB-1 to int type
	// because if I don't,j will be type of unsigned integer, which is always positive
	// so the for loop will not break
	// j starts with the MSB
	for j:=int(PreMSB-1);j>=0;j--{
		if x & (1 << uint(j)) == (1 << uint(j)){
			num1++
			continue
		}
		num0++
	}
	return


}

func main(){

	cnt0,cnt1 := Count(243)
	fmt.Printf("Number of 0S:%v,Number of 1s:%v\n",cnt0,cnt1)
	cnt0,cnt1 = Count(8)
	fmt.Printf("Number of 0S:%v,Number of 1s:%v\n",cnt0,cnt1)

	fmt.Printf("Bits need to be flipped:%v\n",BitsForAToB(29,15))
	fmt.Printf("Before Insertion: %b\n After Insertion: %b\n",2048,Insertion(19,2048,2,0))

}