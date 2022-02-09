// C050:オークションの結果

package main
import "fmt"

func main() {
  var S, a, b int
	fmt.Scanf("%d %d %d", &S, &a, &b)
  var is_a bool = true

	for {
		if is_a == true && S + 10 <= a  {
			S += 10
			is_a = false
		} else if is_a == false && S + 1000 <= b {
			S += 1000
			is_a = true
		} else if is_a == true && S + 10 > a {
			fmt.Println("B", S)
			break
		} else if is_a == false && S + 1000 > b {
			fmt.Println("A", S)
			break
		}
	}
}