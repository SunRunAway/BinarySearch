package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// need to call rand.Read(maxKey) during init
var maxKey = make([]byte, 256)

type Message struct {
	val byte
	great bool
}

type Message2 struct {
	val int
	great bool
}

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




func ParallelSearch() []byte {
	var now = make([]byte,256)
	var res = make([]byte,256)
	var l,r byte = 0,255
	for i:=0;i<=255;i++ {
		l,r = 0,255
		c := make(chan Message)
		for l=0;l<=r;l++ {

			//now[i]=l
			newnow := make([]byte,256)
			copy(newnow,now)
			newnow[i]=l
			go func(l byte,now []byte) {
				//now[i]=l
				res = Search(now)
				//fmt.Println(l)
				if (res == nil) { // mid is greater than goal
					c <- Message{
						val:l,
						great:true,
					}

				} else {
					//fmt.Println("i am ge")
					c <- Message{
						val:l,
						great:false,
					}
				}
			}(l,newnow)
			if (l==r) {
				break
			}
			//fmt.Println(l,r)


		}

		l,r = 0,255
		for j:=0;j<=255;j++ {
			//fmt.Println("jih",j)
			m := <- c
			//fmt.Println("m valsu", m.great,m.val)
			if (m.great) {
				if(m.val-1<r) {
					r=m.val-1
				}
			} else {
				if(m.val > l) {
					l=m.val
				}
			}
			//fmt.Printf("%d %d\n",l,r)
		}
		now[i]=l
		//return now
		//fmt.Println(i)
	}
	return now
}

func main() {
	t := time.Now()
	res := ParallelSearch()
	fmt.Println(res)
	fmt.Println(maxKey)
	fmt.Println(time.Now().Sub(t).Seconds())


}

/*


func ParallelSearch2() []byte {
	var now = make([]byte,256)
	var res = make([]byte,256)
	var l,r byte = 0,255
	for i:=0;i<=255;i++ {
		for bit:=1;bit<=2;bit++ {
			l,r = 0,15
			c := make(chan Message)
			for l=0;l<=r;l++ {

				//now[i]=l
				newnow := make([]byte,256)
				copy(newnow,now)
				if bit==1 {
					newnow[i]=l*16
				} else {
					newnow[i]=now[i]+l
				}
				//newnow[i]=l
				go func(l byte,now []byte) {
					//now[i]=l
					res = Search(now)
					//fmt.Println(l)
					if (res == nil) { // mid is greater than goal
						c <- Message{
							val:l,
							great:true,
						}

					} else {
						//fmt.Println("i am ge")
						c <- Message{
							val:l,
							great:false,
						}
					}
				}(l,newnow)
				if (l==r) {
					break
				}
				//fmt.Println(l,r)


			}

			l,r = 0,15
			for j:=0;j<=15;j++ {
				//fmt.Println("jih",j)
				m := <- c
				//fmt.Println("m valsu", m.great,m.val)
				if (m.great) {
					if(m.val-1<r) {
						r=m.val-1
					}
				} else {
					if(m.val > l) {
						l=m.val
					}
				}
				//fmt.Printf("%d %d\n",l,r)
			}
			if bit==1 {
				now[i]=l*16
			} else {
				now[i]=now[i]+l
			}
			//now[i]=l
			//return now
			//fmt.Println(i)
		}
	}
	return now
}

func ParallelSearch3() []byte {
	var now = make([]byte,256)
	var res = make([]byte,256)
	var l,r int = 0,255
	for i:=0;i<=255;i+=2 {
		l,r = 0,256*256-1
		c := make(chan Message2)
		for l=0;l<=r;l++ {

			//now[i]=l
			newnow := make([]byte,256)
			copy(newnow,now)
			newnow[i]=byte(l/256)
			newnow[i+1]=byte(l%256)
			go func(l int,now []byte) {
				//now[i]=l
				res = Search(now)
				//fmt.Println(l)
				if (res == nil) { // mid is greater than goal
					c <- Message2{
						val:l,
						great:true,
					}

				} else {
					//fmt.Println("i am ge")
					c <- Message2{
						val:l,
						great:false,
					}
				}
			}(l,newnow)
			if (l==r) {
				break
			}
			//fmt.Println(l,r)


		}

		l,r = 0,256*256-1
		for j:=0;j<=256*256-1;j++ {
			//fmt.Println("jih",j)
			m := <- c
			//fmt.Println("m valsu", m.great,m.val)
			if (m.great) {
				if(m.val-1<r) {
					r=m.val-1
				}
			} else {
				if(m.val > l) {
					l=m.val
				}
			}
			//fmt.Printf("%d %d\n",l,r)
		}
		now[i]=byte(l/256)
		now[i+1]=byte(l%256)
		//return now
		//fmt.Println(i)
	}
	return now
}
*/

