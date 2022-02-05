// C102:行きたいライブのスケジュール

package main
import "fmt"

func contains(target interface{}, list interface{}) (bool) {
	switch list.(type) {
		default:
			return false
		case []int:
			revert := list.([]int)
			for _, r := range revert {
				if target == r {
					return true
				}
			}
			return false

		case []uint64:
			revert := list.([]uint64)
			for _, r := range revert {
				if target == r {
					return true
				}
			}
			return false

		case []string:
			revert := list.([]string)
			for _, r := range revert {
				if target == r {
					return true
				}
			}
			return false
	}
	return false
}

func main(){
  var A int
	fmt.Scanf("%d", &A)
	A_date := make([]int, A)
	for i := 0; i < A; i++ {
		fmt.Scan(&A_date[i])
	}
	var B int
	fmt.Scanf("%d", &B)
	B_date := make([]int, B)
	for i := 0; i < B; i++ {
		fmt.Scan(&B_date[i])
	}
	var is_A bool = true

	for i := 1; i <= 31; i++ {
		if contains(i, A_date) == true && contains(i, B_date) == true {
			if is_A == true {
				fmt.Println("A")
				is_A = false
			} else {
				fmt.Println("B")
				is_A = true
			}
		} else if contains(i, A_date) == true {
			fmt.Println("A")
		} else if contains(i, B_date) == true {
			fmt.Println("B")
		} else {
			fmt.Println("x")
		}
	}
}
