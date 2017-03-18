package Challenges_and_DataStructures
The following operations are very important to know, but do not simply memorize them. Memorizing leads
to mistakes that are impossible to recover from. Rather, understand how to implement these methods, so
that you can implement these, and other, bit problems.
Get Bit
This method shifts 1 over by i bits, creating a value that looks like 00010000. By performing an AND with
num, we clear all bits other than the bit at bit i. Finally, we compare that to 0. If that new value is not zero,
then bit i must have a 1. Otherwise, biti is a 0.
1 boolean getBit(int num, inti) {
2 return (( num & (1 « i)) != 0);
3 }
Set Bit
Set Bit shifts 1 over byi bits, creating a value like 00010000. By performing an OR with num, only the
value at bit i will change. All other bits of the mask are zero and will not affect num.
1 int setBit(int num, inti) {
2 return num I ( 1 « i);
3 }
Clear Bit
This method operates in almost the reverse of setBi t. First, we create a number like 11101111 by creating
the reverse of it (00010000) and negating it. Then, we perform an AND with num. This will clear the ith bit
and leave the remainder unchanged.
1 int clearBit(int num, int i) {
2 int mask = =(1 << i);
3 return num & mask;
4 }
To clear all bits from the most significant bit through i (inclusive), we create a mask with a 1 at the ith bit (1
< < i). Then, we subtract 1 from it, giving us a sequence of 0s followed by i ls. We then AND our number
with this mask to leave just the last i bits.
1 int clearBitsMSBthroughI(int num, inti) {
2 int mask = (1 << i) - 1;
3 return num & mask;
4 }
To clear all bits from i through 0 (inclusive), we take a sequence of all ls (which is -1) and shift it left by i
+ 1 bits. This gives us a sequence of 1 s (in the most significant bits) followed by i 0 bits.
1 int clearBitsithrough0(int num, int i) {
2 int mask = (-1 << (i + 1));
3 return num & mask;
4 }
Update Bit
To set the ith bit to a valuev, we first clear the bit at position i by using a mask that looks like 11101111.
Then, we shift the intended value, v, left by i bits. This will create a number with bit i equal to v and all
other bits equal to 0. Finally, we OR these two numbers, updating the ith bit if v is 1 and leaving it as 0
otherwise.
1 int updateBit(int num, int i, boolean bitisl) {
2 int value = bitisl? 1: 0;
3 int mask = =(1 << i);
4 return (num & mask) I (value << i);
5 }


5.1 Insertion: You are given two 32-bit numbers, N and M, and two bit positions, i and
j. Write a method to insert M into N such that M starts at bit j and ends at bit i. You
can assume that the bits j through i have enough space to fit all of M. That is, if
M = 10011, you can assume that there are at least 5 bits between j and i. You would not, for
example, have j = 3 and i = 2, because M could not fully fit between bit 3 and bit 2.
EXAMPLE
Input: N 10000000000, M
Output: N = 10001001100
Hints: #137, #769, #215


5.2 Binary to String: Given a real number between O and 1 (e.g., 0.72) that is passed in as a double, print
the binary representation. If the number cannot be represented accurately in binary with at most 32
characters, print "ERROR:'
Hints: #743, #767, #7 73, #269, #297
5.3 Flip Bit to Win: You have an integer and you can flip exactly one bit from a 0 to a 1. Write code to
find the length of the longest sequence of ls you could create.
EXAMPLE
Input: 1775
Output: 8
(or: 11011101111)
Hints: #759, #226, #374, #352
5.4 Next Number: Given a positive integer, print the next smallest and the next largest number that
have the same number of 1 bits in their binary representation.
Hints: #747, #7 75, #242, #372, #339, #358, #375, #390
5.5 Debugger: Explain what the following code does: ( ( n & ( n-1)) == 0).
Hints:#757,#202,#267,#302,#346,#372,#383,#398
...- -----· .................................... --- · _______ pg 285
5.6 Conversion: Write a function to determine the number of bits you would need to flip to convert
integer A to integer B.
EXAMPLE
Input: 29 (or: 11101), 15 (or: 01111)
Output: 2
Hints: #336, #369
pg286
5.7 Pairwise Swap: Write a program to swap odd and even bits in an integer with as few instructions as
possible (e.g., bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, and so on).
Hints: #745, #248, #328, #355
. ... .... . . . . pg 286
5.8 Draw Line: A monochrome screen is stored as a single array of bytes, allowing eight consecutive
pixels to be stored in one byte. The screen has width w, where w is divisible by 8 (that is, no byte will
be split across rows). The height of the screen, of course, can be derived from the length of the array
and the width. Implement a function that draws a horizontal line from ( xl, y) to ( x2, y).
The method signature should look something like:
drawline(byte[] screen, int width, int xl, int x2, int y)
Hints: #366, #387, #384, #397

func SetBit(x, position int) int{
	return x | 1<<position

}

func GetBit(x, position int) int{

	return x & (1<<position)

}


func FlipBit(x, position int){
	return x ^ 1<<position


}

// right shift
// or 1 leftshift
func isBitSet(){



}