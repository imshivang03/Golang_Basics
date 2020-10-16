package main 
  
import( 
    "fmt"
    "strings"
) 
  
func joinstr(element...string)string{ 
    return strings.Join(element, "-") 
} 
  
func main() { 
     
   element:= []string{"Welcome", "To", "golang"} 
   fmt.Println(joinstr(element...)) 
} 
