package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)


//check a string has all unique characters
func IsUnique(str string)bool{
	// O(N)
	// use a dictionary to map character and its occurences in string
	counter := map[rune]int{}
	for _,j := range str{
		counter[j]++
	}
	fmt.Printf("len(str):%v\tlen(count):%v\n",len([]rune(str)),len(counter))
        // if the str has all unique unicode characters, its counter will
	// has the same amount of elements(no characters have more than two counts)
        return	len(counter) == len([]rune(str))


}

// check if stringA  is a permutation of stringB
func isPermutationOf(strA,strB string) bool{
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


func FindPermutation(strA,strB string){
	//b := make(map[rune]int)
	//a := make(map[int]map[rune]int)
	//for i,j :=range strB{
	//	b[strB[i]] ++
	//	b[strB[i+1]] ++
	//	b[strB[i+2]] ++
	//	b[strB[i+3]] ++
	//	a[i] = strB[i:i+4]
	//
	//
	//}





}

//Write a method to replace all spaces in a string with '%20'. You may assume that the string
//has sufficient space at the end to hold the additional characters, and that you are given the "true"
//length of the string. (Note: If implementing in Java, please use a character array so that you can
//perform this operation in place.)
//EXAMPLE
//Input: "Mr John Smith ", 13
//Output: "Mr%20John%20Smith"

func URLify(){




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
	counter := make(map[rune]int)
	// oddCounter stores characters that appear odd times in str
	// for a str to be palindrome, the number of these characters
	// should be less than one, either all occur even time ,or one of them occur
	// odd times to make sure the string is symmetric from both sides
	oddCounter := 0
	for _,j := range str{
		counter[j]++

	}
	fmt.Println(counter)
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




}


//Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
//column are set to 0.

func ZeroMatrix(){



}


//String Rotation:Assumeyou have a method isSubstringwhich checks if one word is a substring
//of another. Given two strings, sl and s2, write code to check if s2 is a rotation of sl using only one
//call to isSubstring (e.g., "waterbottle" is a rotation of"erbottlewat").
func RotateString(){




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
}

func main(){





}