package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
	"strings"
)


//check a string has all unique characters
func IsUnique(str string)bool{
	// O(N)
	// use a dictionary to map characters and its occurrences
	counter := map[rune]int{}
	for _,j := range str{
		counter[j]++
	}
	fmt.Printf("len(str):%v\tlen(count):%v\n",len([]rune(str)),len(counter))
        // if the str has all unique unicode characters, its element counter will
	// has the same amount of elements(no characters have more than two counts)
        return	len(counter) == len([]rune(str))


}

// check if stringA  is a permutation of stringB
func isPermutationOf(strA,strB string) bool{
	// O(A+B)
	// early return
	if len([]rune(strA)) != len([]rune(strB)){
		return false
	}

	// contruct character-count mapping for each string
	counterA := map[rune]int{}
	for _,i := range strA{
		counterA[i]++
	}
	counterB := map[rune]int{}
	for _,j := range strA{
		counterB[j]++
	}

	// check equality
	for key := range counterA{
		if counterA[key] != counterB[key]{
			return false
		}

	}

	return true
}

// FindPermutation finds all strA's permutation in strB,return its indices
func FindPermutation(strA,strB string)[]int{
        outputIndices := []int{}
	counterA := map[rune]int{}
	lengthA := len([]rune(strA))
	for _,v := range strA{

		counterA[v]++
	}
	span := []int{}
        for i,v := range strB{
		if counterA[v] == 0{
			span = append(span,i)
		}
	}
	span = append(span,lengthA)

	for j:=0;j<len(span)-1;j++{
		if span[j+1] - span[j]-1 < lengthA{
			continue
		}
		for i := span[j]+1;i+lengthA-1<span[j+1];i++{

			counterTemp := map[rune]int{}

			for _,v := range strB[i:i+lengthA]{
			     counterTemp[v]++

			}
			permu := true
			for key,_ := range counterTemp{
				if counterTemp[key] != counterA[key]{
					permu = false
					break
				}

			}
			if permu {
				outputIndices = append(outputIndices, i)
			}
		}






	}
	return outputIndices

}


//Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome.
//A palindrome is a word or phrase that is the same forwards and backwards. A permutation
//is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
//1.5
//1.6
//EXAMPLE
//Input: Tact Coa
//Output: True (permutations: "taco cat", "atco eta", etc.)
func isPalindromePermutation(str string) bool{
	// O(N)
	counter := make(map[rune]int)
	// oddCounter stores characters that appear odd times in str
	// for a str to be palindrome, the number of these characters
	// should be less than one, either all occur even time ,or one of them occur
	// odd times to make sure the string is symmetric from both sides
	oddCounter := 0
	for _,j := range str{
		counter[j]++

	}
	for _,value := range counter{

		if value % 2 == 1{
			oddCounter ++
		}

	}
	// the code below this line of code is redundant
	return oddCounter <= 1

	//if oddCounter > 1{
	//	return false
	//}
	//
	//return true

}

//One Away: There are three types of edits that can be performed on strings: insert a character,
//remove a character, or replace a character. Given two strings, write a function to check if they are
//one edit (or zero edits) away.
//EXAMPLE
//pale, ple -> true
//pales, pale -> true
//pale, bale -> true
//pale, bake -> false
func OneAway(strA, strB string) bool{
	// O(A+B)
	// early return, length difference should be within range of 1
	if math.Abs(float64(len([]rune(strA))-len([]rune(strB)))) > 1 {
		return false
	}
	counterA := make(map[rune]int)
	counterB := make(map[rune]int)
	missingCounter := 0
	for _,i := range strA{
		counterA[i]++
	}

	for _,j := range strB{
		counterB[j]++
	}

	for key,_:= range counterA{
		if counterA[key] != counterB[key]{
			missingCounter++
		}

	}

	return missingCounter <= 1
}

//String Compression: Implement a method to perform basic string compression using the counts
//of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
//"compressed" string would not become smaller than the original string, your method should return
//the original string. You can assume the string has only uppercase and lowercase letters (a - z).
func Compress(str string)string{
	// O(N)
	// added unicode support
	outputstr := ""
	if len(str) == 0{
		return outputstr
	}

	if len(str) == 1 || len(str) == 2{
		return str

	}
	current,size := utf8.DecodeRuneInString(str)
	span := 1
	indices := []interface{}{}
	for _,rune := range str[size:]{
		if rune != current {

			indices = append(indices, current,span)
			//fmt.Println(string(current),span)
			current = rune

			span = 1
			continue
		}

		span++
	}
	indices = append(indices, current,span)
	// if compressed string is no shorter than the original, then no-op
	if len(indices) >= len(str){
		return str
	}

	for j:=1;j<len(indices);j+=2{

                // type assertion
		v1,_ := indices[j-1].(rune)
		v2,_ := indices[j].(int)
                outputstr += string(v1)+strconv.Itoa(v2)
	}
        return  outputstr
}

//Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
//bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
func RotateMatrix(){
	// O(N2)
	const N = 5
	img := [N][N]rune{[N]rune{0,0,0,0,0},
		         [N]rune {1,1,1,1,1},
		         [N]rune{2,2,2,2,2},
		         [N]rune{3,3,3,3,3},
	                 [N]rune{4,4,4,4,4}}



     	for i:=0;i<N;i++{
		fmt.Println(img[i])
	}
	for i:=0;i<N;i++{

		// J must start with i, if j starts from 0, you do nothing
		// you reverse what you have swapped
		for j:=i;j<N;j++{
			img[i][j],img[j][i] = img[j][i],img[i][j]
		}

	}

	for i:=0;i<N;i++{
		// swap symmetric columns
		for j:=0;j<N/2;j++{
			img[i][j],img[i][N-1-j] = img[i][N-1-j],img[i][j]
		}

	}
	fmt.Println(strings.Repeat("*",180))
	for i:=0;i<N;i++{
		fmt.Println(img[i])
	}

}


//Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
//column are set to 0.

func ZeroMatrix(){
	array := [][]int{[]int{1,1,1,1,1,3,0,1},
		         []int{2,2,2,2,2,2,2,2},
		         []int{3,3,3,3,3,3,3,3},
	                 []int{4,4,4,4,4,5,6,7}}
        i,j := 0,0
	// keep i,j'scope out of for loop for later use
	for ;i<len(array);i++{
		for ;j<len(array[0]);j++{
			if array[i][j] == 0{
				// get out of inner and outer for loop
				goto next
			}
		}

	}

        next:
		for l:=0;l<len(array[0]);l++{
			array[i][l] = 0
		}

		for m:=0;m<len(array);m++{
			array[m][j] = 0
		}
		for ;i<len(array);i++{
			fmt.Println(array[i])
		}

}


//String Rotation:Assume you have a method isSubstring which checks if one word is a substring
//of another. Given two strings, sl and s2, write code to check if s2 is a rotation of sl using only one
//call to isSubstring (e.g., "waterbottle" is a rotation of"erbottlewat").
func isRotatedString(s2,s1 string)bool{
	// O(N)
	// early return
	if s1==s2{
		return true
	}
	i:= 0
	for ;i< len(s1);i++{
		// find longest prefix substring of s1 that is a
		// substring of s2, stops at i-1
		if !strings.Contains(s2,s1[:i]){
			break
		}


	}
        // test the rest
	return s1[i-1:] == s2[:len(s2)-i+1]



}

// string_test tests all the previous functions
func string_test(){



	p := fmt.Printf
	strA := "fff我是一个好人吗人好个一是我fff"
	strB := "abcdefghijklm我是"
	strC := "abcd好ef"
	strD := "febdac好"
	strE := "fff我是一个好人吗人好个一是我fff"
	strF:= "fff一个好人吗人我是好个一是我fff"
	strH := "fff个好人吗人好个一是我ff我是一f"
	strI := "我是一f个好人人好个一是我f"
	p("is %v,Unique? %v\n",strA,IsUnique(strA))
	p("is %v,Unique? %v\n",strB,IsUnique(strB))
	p("%v is Permutation of %v? %t\n",strA,strB,isPermutationOf(strA,strB))
	p("%v is Permutation of %v? %t\n",strC,strD,isPermutationOf(strC,strD))
	p("%v is Palindrome Permutation ? %t\n",strE,isPalindromePermutation(strE))
	p("%v is Palindrome Permutation ? %t\n",strF,isPalindromePermutation(strF))
	p("%v is Palindrome Permutation ? %t\n",strD,isPalindromePermutation(strD))
	p("%v is Palindrome Permutation ? %t\n",strH,isPalindromePermutation(strH))
	p("%v is Palindrome Permutation ? %t\n",strI,isPalindromePermutation(strI))

	oneawayA :=[]string{"fff","fff","fffi","abc","pale","pales","pale","pale","pale","我是、"}
	oneawayB := []string{"ffb","fffj","fff","def","ple","pale","palewe","pale","bake","我是+"}
	for i:=0;i<len(oneawayA);i++{
		p("\033[35m%v\033[0m is one edit way from \033[33m%v\033[0m? %t\n",oneawayA[i],
			oneawayB[i], OneAway(oneawayA[i], oneawayB[i]))
	}

	compressStrs :=[]string{"fff","fffsfefeeesfe22222yyy","fffiiiii","我我我我","++++yyyyyyyyyabccew","bb","c","bcdefcc","abc"}
        for _, str := range compressStrs{
		p("Before: \033[35m%v\033[0m After compression: \033[33m%v\033[0m\n",str,Compress(str))
	}

	ZeroMatrix()
	p("%t\n",isRotatedString("alexisavery","isaveryalex"))
	p("%t\n",isRotatedString("alexisavery","averyalexis"))
	p("%t\n",isRotatedString("alexisavery","veryalexisa"))
	p("%t\n",isRotatedString("alexisavery","yalexisaver"))
	RotateMatrix()
	p("%v\n",FindPermutation("abbc","cbabadcbbabbcbabaabccbabc"))
}

func main(){


	string_test()


}