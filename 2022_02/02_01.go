// paiza B103:カブトムシの誘導

package main
import "fmt"
func main(){
    var count int
	fmt.Scan(&count)
	var x_R, x_G, x_B int
	fmt.Scanf("%d %d %d", &x_R, &x_G, &x_B)
	inputs := make([]string, count * 2)
	for i := 0; i < count * 2; i++ {
		fmt.Scan(&inputs[i])
	}

	for i := 0; i < count * 2 + 1 ; count++ {
    	switch i {
    	case "R":
    		fmt.Println("1つめ")
    	case "2":
    		fmt.Println("2つめ") // 2つめが出力される
    	default:
    		fmt.Println("3つめ")
        }

	}
	if c == false {
	    fmt.Println("No")
	}
}