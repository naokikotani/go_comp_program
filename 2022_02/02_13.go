// ABC086A - Product

package main
import "fmt"
import "strings"
import "strconv"

func Split(input string){
	result := 0
	slice := strings.Split(input, "")
	len := len(slice)
	for i := 0; i < len; i++ {
		if slice[i] == "1" {
			result += 1
		}
	}
	fmt.Println(result)
}

func main(){
	var a int
	fmt.Scanf("%d", &a)
	b := strconv.Itoa(a)
	Split(b)
}


// 回答例１
package main
import "fmt"

func main() {
	var (
		a string
		count int = 0
	)
	fmt.Scan(&a)
	for i := 0; i < len(a); i++ {
		if a[i] == '1' {
			count++
		}
	}
	fmt.Println(count)
}


// 回答例2
package main

import ("fmt"
       )

func main() {
  var str string
  var result int

  fmt.Scan(&str)
  for _, r := range str {
    result += int(r - '0')
  }
  fmt.Println(result)
}

// 回答例3
package main
import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(strings.Count(s, "1"))
}


// ABC081B - Shift only
package main
import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	result := 0
	is_odd := false
	inputs := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&inputs[i])
	}
	for {
		for i := 0; i < N; i++ {
			if inputs[i] % 2 == 1 {
				is_odd = true
			} else {
				inputs[i] = inputs[i] / 2
			}
		}
		if is_odd == true {
			break
		} else {
			result++
		}
	}
	fmt.Println(result)
}




package main
import "fmt"

func main() {
	var (
		a int
		b int
		c int
		x int
		result int = 0
	)

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	for num_a := 0; num_a < a + 1; num_a++ {
		for num_b := 0; num_b < b + 1; num_b++ {
			for num_c := 0; num_c < c + 1; num_c++ {
				if num_a * 500 + num_b * 100 + num_c * 50 == x {
					result++
				}
			}
		}
	}
	fmt.Println(result)
}
