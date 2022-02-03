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

	// 配列
	arr := [...] string{"Golang", "Java"}
	fmt.Println(arr[0], arr[1]) //=> Golange Java
	fmt.Println(arr) //=> [Golange Java]
	// 配列について
	// https://ashitani.jp/golangtips/tips_slice.html

	//配列の中に特定の文字列が含まれるかを返す
	func arrayContains(arr []string, str string) bool{
		for _, v := range arr{
			if v == str{
				return true
			}
		}
		return false
	}

	// if文
	if i % b == 0 && i % c == 0 {
			fmt.Println("AB")
	} else if i % b == 0 {
			fmt.Println("A")
	} else if i % c == 0 {
			fmt.Println("B")
	} else {
			fmt.Println("N")
	}

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

	// 四則演算
	// 剰余
	i % b // i割るbの余り




