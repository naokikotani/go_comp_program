// C062:回転寿司のメロン

package main
import "fmt"
func main(){
  var T int
	fmt.Scanf("%d", &T)
	var answer int = 0
	var d int = 0
	for i := 0; i < T; i++ {
		var n string
	  fmt.Scan(&n)
    if n != "melon" {
      d += 1
    } else if n == "melon" && d >= 10  {
      answer += 1
      d = 0
    }
	}
	fmt.Println(answer)
}