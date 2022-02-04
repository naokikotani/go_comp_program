// C099:折り紙の貼り合わせ

package main
import "fmt"
func main(){
	var N, D int
	fmt.Scanf("%d %d", &N, &D)
	inputs := make([]int, N - 1)
	for i := 0; i < N - 1; i++ {
		fmt.Scan(&inputs[i])
	}

	sum := 0
    for _, x := range inputs {
      sum += x
		}

	var answer int = 0
	answer = D * (D * N - sum)

  fmt.Println(answer)
}
