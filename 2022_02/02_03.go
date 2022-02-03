package main
import "fmt"

//配列の中に特定の文字列が含まれるかを返す
func arrayContains(arr []bool, boolean bool) bool{
	for _, v := range arr{
		if v == boolean{
			return true
		}
	}
	return false
}

func main(){
  var num_s, num_g int
	fmt.Scanf("%d %d", &num_s, &num_g)
	var answer int = 0
  seats := make([]bool, num_s)

  for i := 0; i <= num_g; i++ {
		var num, init int
		fmt.Scanf("%d %d", &num, &init)

		if arrayContains(seats[init - 1: init + num - 1], true) {
		} else {
			for i := init - 1; i <= num; i++ {
				seats[i] = true
				}
		}
  }
  for i := 0; i <= len(seats); i++ {
    if seats[i] == true {
			answer += 1
		}
  }
  fmt.Println(answer)
}
