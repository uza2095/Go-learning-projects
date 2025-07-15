package main
import ("fmt"
		"strings"
		"unicode"
)
func counter(word string) map[rune]int{
	count:=make(map[rune]int)
	for _,r:=range word{
		if unicode.IsLetter(r){
			count[r]++
		}
		
	}
	return count
}

func main(){
	var word string
	fmt.Println("please enter the word you want to count")
	fmt.Scanln(&word)
	edited:=strings.ToLower(word)
	freq:=counter(edited)
	for i,r:= range freq{
		fmt.Printf("%c:%d\n",i,r)
	}


}