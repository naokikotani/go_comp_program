// C104:虫食い算

package main
import "fmt"

func main(){
    var a, b int
    var c = false
	fmt.Scanf("%d %d", &a, &b)

	for y := 0; y < 10; y++ {
	    for x := 1; x < 10;x++ {
	    result := y * y + 10 * y * x - 10 * x
	    if result == 100 * a + b {
	        c = true
	        fmt.Println(x, y)
	        }
	    }
	}
	if c == false {
	    fmt.Println("No")
	}
}
