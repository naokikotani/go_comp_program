package main
import "fmt"
import "time"
import "strconv"

func main(){
  var N int
	fmt.Scanf("%d", &N)

	for i := 0; i < N; i++ {
	  var place string
    fmt.Scan(&place)
    var time int
	  fmt.Scanf("%d", &time)
	}

	var st_place string
	fmt.Scan(&st_place)
	var input_time string
	fmt.Scan(&input_time)
	var hour int
	hour, _ = strconv.Atoi(input_time[0:2])
	var minute int
	minute, _ = strconv.Atoi(input_time[3:5])
	var t Time
	t = time.Date(2000, 1, 1, hour, minute, 0, 0, time.UTC)
	fmt.Println(input_time[3:5])
}