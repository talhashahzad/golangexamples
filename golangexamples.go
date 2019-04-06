package golangexamples

import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string {
  var conc = ""
  for i := range sliceToConcat {
    conc+=string(sliceToConcat[i])
    if i != len(sliceToConcat)-1{
      conc+="-"
    }
  }
  return conc
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int){
  for i := range sliceToEncrypt {
    sliceToEncrypt[i]=byte((int(sliceToEncrypt[i]) + ceaserCount) % 26)
  }
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
