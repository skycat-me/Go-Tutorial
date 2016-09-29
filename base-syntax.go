package main

import "fmt"

func main() {
	// Goでは変数名の後に型がくる
	var num int
	var pow	int
	var result = 1

	// 文末のセミコロンは省略できるが、
	// 以下のように1行に複数の分を書く場合は必要。
	num = 2; pow = 4

	// 宣言と同時に初期化を行う場合は以下のように省略できる。
	// var result = 1
	// i := 0

	for i := 0; i < pow; i++ {
		result *= num
	}

	fmt.Printf("%dの%d乗の%dです\n", num, pow, result);

	// if, for, switchに丸括弧は不要(省略可！)
	if result > 10 { fmt.Println("10よりおおきい") }

}

