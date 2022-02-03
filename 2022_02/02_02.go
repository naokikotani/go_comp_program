// C097 プレゼント応募企画の実施

package main
import "fmt"
func main(){
  var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	for i := 1; i <= a; i++ {
		if i % b == 0 && i % c == 0 {
				fmt.Println("AB")
		} else if i % b == 0 {
				fmt.Println("A")
		} else if i % c == 0 {
				fmt.Println("B")
		} else {
				fmt.Println("N")
		}
	}
}