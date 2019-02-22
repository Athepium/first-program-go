package main

import (
  "fmt"
  "strconv"
  "strings"
  "os"
)
func PrintSquare(size int){
    for i := 0;i<size;i++{
      if i== 0 || i==size-1{
        fmt.Printf("%c%s%c\n",'+',strings.Repeat("-",size-2),'+')
      }else{
          fmt.Printf("%c%s%c\n",'|',strings.Repeat(" ",size-2),'|')
      }
        }

}

func main(){
  if len(os.Args) ==2{
    i, _:= strconv.Atoi(os.Args[1])
    PrintSquare(i)
  }
}
