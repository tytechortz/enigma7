package main

import "fmt"

var wheel2Position rune
var str string
var char rune

func main() {
	fmt.Print("Enter message : ")
	fmt.Scan(&str)
	length := len(str)
	fmt.Println(length)
	//	length1 := rune(length)

	s := []rune(str)
	//	fmt.Println(s)

	for _, char := range s {

		wheel1 := map[rune]rune{
			'a': 'e',
			'b': 'k',
			'c': 'm',
			'd': 'f',
			'e': 'l',
			'f': 'g',
			'g': 'd',
			'h': 'q',
			'i': 'v',
			'j': 'z',
			'k': 'n',
			'l': 't',
			'm': 'o',
			'n': 'w',
			'o': 'y',
			'p': 'h',
			'q': 'x',
			'r': 'u',
			's': 's',
			't': 'p',
			'u': 'a',
			'v': 'i',
			'w': 'b',
			'x': 'r',
			'y': 'c',
			'z': 'j',
		}
		//	fmt.Println(char1)
		wheel1output := wheel1[char]
		//	fmt.Println(wheel1output)

		wheel2 := map[rune]rune{
			'a': 'a',
			'b': 'j',
			'c': 'd',
			'd': 'k',
			'e': 's',
			'f': 'i',
			'g': 'r',
			'h': 'u',
			'i': 'x',
			'j': 'b',
			'k': 'l',
			'l': 'h',
			'm': 'w',
			'n': 't',
			'o': 'm',
			'p': 'c',
			'q': 'q',
			'r': 'g',
			's': 'z',
			't': 'n',
			'u': 'p',
			'v': 'y',
			'w': 'f',
			'x': 'v',
			'y': 'o',
			'z': 'e',
		}
		wheel2output := wheel2[wheel1output]
		//	fmt.Println(wheel2output)

		reflector := map[rune]rune{
			'a': 'b',
			'b': 'd',
			'c': 'f',
			'd': 'h',
			'e': 'j',
			'f': 'l',
			'g': 'c',
			'h': 'p',
			'i': 'r',
			'j': 't',
			'k': 'x',
			'l': 'v',
			'm': 'z',
			'n': 'n',
			'o': 'y',
			'p': 'e',
			'q': 'i',
			'r': 'w',
			's': 'g',
			't': 'a',
			'u': 'k',
			'v': 'm',
			'w': 'u',
			'x': 's',
			'y': 'q',
			'z': 'o',
		}
		reflectoroutput := reflector[wheel2output]
		//	fmt.Println(reflectoroutput)

		reflector2 := map[rune]rune{
			'a': 'e',
			'b': 'j',
			'c': 'm',
			'd': 'z',
			'e': 'a',
			'f': 'l',
			'g': 'y',
			'h': 'x',
			'i': 'v',
			'j': 'b',
			'k': 'w',
			'l': 'f',
			'm': 'c',
			'n': 'r',
			'o': 'q',
			'p': 'u',
			'q': 'o',
			'r': 'n',
			's': 't',
			't': 's',
			'u': 'p',
			'v': 'i',
			'w': 'k',
			'x': 'h',
			'y': 'g',
			'z': 'd',
		}
		reflector2output := reflector2[reflectoroutput]
		//	fmt.Println(reflector2output)

		wheel2reverse := map[rune]rune{
			'a': 't',
			'b': 'a',
			'c': 'g',
			'd': 'b',
			'e': 'p',
			'f': 'c',
			'g': 's',
			'h': 'd',
			'i': 'q',
			'j': 'e',
			'k': 'u',
			'l': 'f',
			'm': 'v',
			'n': 'n',
			'o': 'z',
			'p': 'h',
			'q': 'y',
			'r': 'i',
			's': 'x',
			't': 'j',
			'u': 'w',
			'v': 'l',
			'w': 'r',
			'x': 'k',
			'y': 'o',
			'z': 'm',
		}
		wheel2reverseoutput := wheel2reverse[reflector2output]
		//fmt.Println(wheel2reverseoutput)

		wheel1reverse := map[rune]rune{
			'a': 'a',
			'b': 'j',
			'c': 'p',
			'd': 'c',
			'e': 'z',
			'f': 'w',
			'g': 'r',
			'h': 'l',
			'i': 'f',
			'j': 'b',
			'k': 'd',
			'l': 'k',
			'm': 'o',
			'n': 't',
			'o': 'y',
			'p': 'u',
			'q': 'q',
			'r': 'g',
			's': 'e',
			't': 'n',
			'u': 'h',
			'v': 'x',
			'w': 'm',
			'x': 'i',
			'y': 'v',
			'z': 's',
		}

		wheel1reverseoutput := wheel1reverse[wheel2reverseoutput]
		//	fmt.Println(wheel1reverseoutput)

		finaloutput := map[rune]rune{
			'a': 'u',
			'b': 'w',
			'c': 'y',
			'd': 'g',
			'e': 'a',
			'f': 'd',
			'g': 'f',
			'h': 'p',
			'i': 'v',
			'j': 'z',
			'k': 'b',
			'l': 'e',
			'm': 'c',
			'n': 'k',
			'o': 'm',
			'p': 't',
			'q': 'h',
			'r': 'x',
			's': 's',
			't': 'l',
			'u': 'r',
			'v': 'i',
			'w': 'n',
			'x': 'q',
			'y': 'o',
			'z': 'j',
		}
		output := finaloutput[wheel1reverseoutput]
		//	fmt.Println(output)

		output1 := int(output)
		fmt.Println(string(output1))
		//	fmt.Println(char)
	}
}
