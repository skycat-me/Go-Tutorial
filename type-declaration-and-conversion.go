package main

import "fmt"

func main() {
    _type()
    readFunc()
}

type Dictionary struct{
    name string
    meaning string
}
type ReadFunc func(Dictionary) string

func readFunc() {
    var readFunc ReadFunc
    var dict Dictionary
    readFunc = readOut
    dict.name = "ラーメン"
    dict.meaning = "麺類の食べ物"
    fmt.Println(readFunc(dict))
}

func readOut(d Dictionary) string {
    return fmt.Sprintf("「%s」は「%s」という意味です", d.name, d.meaning)
}


// typeキーワードを使うことで新たな型を宣言できる  
type Score int
func _type() {
    var myScore Score = 100
    fmt.Printf("私の点数は%d点です。\n", myScore)

    // 型変換
    ShowInt(int(myScore))
}
// レシーバ型として仕様
func (s Score) Show() { fmt.Printf("私の点数は%d点です。\n", s) }

func ShowInt(i int) {
    fmt.Printf("value: %d\n", i);  
}