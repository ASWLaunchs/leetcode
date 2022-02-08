package main

import "fmt"

func main() {
	var res int
	var resTail int
	var resHead int
	fmt.Println("迎接冬天算式可能的结果如下所示：")
	for i := 10; i <= 99; i++ {
		for j := 10; j <= 99; j++ {
			res = i * j
			resTail = res % 10
			resHead = res % 10000 / 1000
			if res > 999 && resHead == int(i%100/10) && resTail == j%10 && i%10 == int(j%100/10) && i*int(j%10) < 100 && i*int(j%100/10) < 100 && (i*int(j%100/10))%10 != (i*int(j%10))%10 {
				fmt.Println(i, "*", j, "=", res)
			}
		}
	}
}
