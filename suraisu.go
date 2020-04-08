//3つの変数しか使わないように修正してください
//プログラムの動作はそのままにすること

package main

func main() {
	ns := []int{19, 86, 1, 12} //スライスの作成//スライスは[]に要素数入ってない
	var sum int
	for i := 0; i < len(ns); i++ {
		sum += ns[i]
	}
	println(sum)
}

//削除
//a = append(a[:i], a[i+1:]...)

/*
package main
func main() {
	n1 := 19 n2 := 86 n3 := 1 n4 := 12
 sum := n1 + n2 + n3 + n4
    println(sum)
}
*/
