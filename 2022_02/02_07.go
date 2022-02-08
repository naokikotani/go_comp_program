// C021:暴風域にいますか？

package main
import "fmt"

func main() {
  var xc, yc, r_1, r_2 int
	fmt.Scanf("%d %d %d %d", &xc, &yc, &r_1, &r_2)
	var n int
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var x_i, y_i int
		fmt.Scanf("%d %d", &x_i, &y_i)
		sum := (x_i - xc) * (x_i - xc) + (y_i - yc) * (y_i - yc)
		if r_1 * r_1 <= sum && sum <= r_2 * r_2 {
		  fmt.Println("yes")
		} else {
		  fmt.Println("no")
		}
	}
}