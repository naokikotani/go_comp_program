// C052:ゲームの画面

package main
import "fmt"
import "math"
func main(){
    var H, W int
	fmt.Scanf("%d %d", &H, &W)
	var dy, dx int
	fmt.Scanf("%d %d", &dy, &dx)

	x := float64(dx)
	y := float64(dy)
	h := float64(H)
	w := float64(W)

	fmt.Println(h * w - (h - math.Abs(y)) * (w - math.Abs(x)))
}