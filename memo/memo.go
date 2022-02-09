// 標準入力
// str
	var str string
	fmt.Scan(&str)

// int
  var num int
	fmt.Scanf("%d", &num)

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

	//配列の要素を検索する
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

	// 日付の取り扱い
	// https://qiita.com/tez/items/1fdcbd9e2b6b575b6725

	// str → int, int → str
	// https://qiita.com/quicksort/items/c9522793a941edf074fd

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

	// for文
	// 無限ループ
	for {
		// 処理
	}


	// 四則演算
	// 剰余
	i % b // i割るbの余り

	// 割り算
	i / b // 整数 / 整数 の場合、小数点以下は切り捨てられる。




