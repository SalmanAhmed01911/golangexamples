package main

import "fmt"
import "strings"
import "math"
import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string{
  arr := ""
  for i := 0; i < len(sliceToConcat); i++ {
		arr = arr + string(sliceToConcat[i]) + "-"
	}
  return arr[0:len(arr)-1]
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int){
  str_sample := ""
  for i := 0; i < len(sliceToEncrypt); i++ {
		str_sample = str_sample + string(sliceToEncrypt[i])
	}

  arr := strings.ToLower(str_sample)
  arr1 := ""
  for i := 0; i < len(arr); i++ {
    arr1 = arr1 + string(int(math.Mod(float64(int(int(arr[i]) - 97) + ceaserCount), 26)) + 97)
	}

  for i := 0; i < len(sliceToEncrypt); i++ {
    sliceToEncrypt[i] = arr1[i]
	}

  fmt.Println("Encrypted String : ")
  fmt.Println(arr1)

}

func EZGreetings(name string) string{
  return greetings.PrintGreetings(name)
}


func main(){
  byteArray := []byte{'s', 'a', 'l', 'm', 'a', 'n'}
  fmt.Println("Concat Slice")
  fmt.Println(ConcatSlice(byteArray))
  fmt.Println("Encryption")
  Encrypt(byteArray, 3)
  fmt.Println("Encrypted Byte Array : ")
  fmt.Println(byteArray)
  fmt.Println("EZGreetings")
  fmt.Println(EZGreetings("Salman Ahmed"))
}
