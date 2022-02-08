// C032:お得な買い物

package main
import "fmt"

func main() {
  var N int
	fmt.Scanf("%d", &N)
	var sum_0 int = 0
	var sum_1 int = 0
	var sum_2 int = 0
	var sum_3 int = 0

	for i := 0; i < N; i++ {
		var v, p int
		fmt.Scanf("%d %d", &v, &p)
		if v == 0 {
			sum_0 += p
		} else if v == 1 {
			sum_1 += p
		} else if v == 2 {
			sum_2 += p
		} else if v == 3 {
			sum_3 += p
		}
	}
	fmt.Println(sum_0 / 100 * 5 + sum_1 / 100 * 3 + sum_2 / 100 * 2 + sum_3 / 100)
}