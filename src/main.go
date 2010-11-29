package main

import (
  
  "fmt"
  "os"
  
)

func DocReader()( i int64, err os.Error){
  
  document_file , err := os.Open("../data/cran.all.1400",os.O_RDONLY, 0666)
  defer document_file.Close()

  file_info, err := document_file.Stat()
  
  fmt.Printf("The File %s is of length %d\n",file_info.Name, file_info.Size)
  
  i = file_info.Size;
  
  
  return i, err
}


func main(){
  
  i ,_ := DocReader()
  println (i)
}