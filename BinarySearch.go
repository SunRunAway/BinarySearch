package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// need to call rand.Read(maxKey) during init
var maxKey = make([]byte, 256)

func init() {
	rand.Read(maxKey)
}

func Search(key []byte) []byte {
	time.Sleep(time.Millisecond*10)
	if bytes.Compare(key, maxKey) > 0 {
		return nil
	}
	return key
}

func mid(index int, l byte, r byte, cur []byte) []byte {
	m := (l+r)/2
	cur[index] = m
	return cur
}
func BinarySearch() []byte {
	var now = make([]byte,256)
	var res = make([]byte,256)
	var l,r int = 0,255
	for i:=0;i<=255;i++ {
		l,r = 0,255
		for l!=r {
			mid := (byte)((l+r+1)/2)
			now[i]=mid
			res = Search(now)
			if (res == nil) { // mid is greater than goal
				r = (int)(mid-1)

			} else {
				l = int(mid)
			}
			//fmt.Printf("%d %d\n",l,r)
		}
		now[i]= byte(l)
	}

	return now


}

func main() {
	t := time.Now()
	res := BinarySearch()
	fmt.Println(res)
	fmt.Println(maxKey)
	fmt.Println(time.Now().Sub(t).Seconds())


}