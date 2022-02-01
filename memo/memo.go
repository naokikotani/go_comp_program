// 標準入力

// 複数行の標準入力
// 入力例
// 1
// 2
// 3
// 4
var count int
	fmt.Scan(&count)
	inputs := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&inputs[i]) // => [1, 2, 3, 4]
	}

	// if文

	// case文
	a := "2"
	switch a {
	case "1":
		fmt.Println("1つめ")
	case "2":
		fmt.Println("2つめ") // 2つめが出力される
	default:
		fmt.Println("3つめ")
}


