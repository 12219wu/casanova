package loopQueue

import (
	"fmt"
	"strconv"
)

type LoopQueue struct {
	Items []string //利用数组实现队列
	Head  uint16   //头索引，恒为0
	Tail  uint16   //尾索引
	Len   uint16   //限定队列长度
}

//队列初始化，将[底层数组]初始化
func (this *LoopQueue) Init() {
	this.Len = this.Len + 1 //因为最后一个位置不能使用，所以len要比声明的大一
	for i := 0; i < int(this.Len); i++ {
		this.Items = append(this.Items, "")
	}
	fmt.Println("Init:", this.Len, len(this.Items))
}

//入列
func (this *LoopQueue) Enqueue(v interface{}) {
	if ((this.Tail + 1) % this.Len) == this.Head {
		fmt.Println("the queue is full!", this.Head, this.Tail)
		return
	}
	this.Items[this.Tail] = v.(string)
	fmt.Printf("%s enqueued, now length: tail:%d \n", v, this.Tail)
	this.Tail = (this.Tail + 1) % this.Len
}

//出列
func (this *LoopQueue) Dequeue() interface{} {
	if this.Head == this.Tail {
		fmt.Println("the queue is empty!", this.Head, this.Tail)
		return nil
	}
	ret := this.Items[this.Head]
	fmt.Printf("Dequeue: %s! the surplus len: head:%d \n", ret, this.Head)
	this.Head = (this.Head + 1) % this.Len

	return ret
}

func Example() {
	Q := &LoopQueue{Len: 10} //设定队列长度

	Q.Init() //初始化队列

	for i := 0; i < 12; i++ { //12>10 会提示full
		s := strconv.Itoa(i)
		Q.Enqueue(s + s + s) //入列（但只能入10个）
	}

	for i := 0; i < 3; i++ {
		Q.Dequeue() //出列3个
	}

	for i := 0; i < 12; i++ { //12>10 会提示full
		s := strconv.Itoa(i)
		Q.Enqueue(s + s + s) //入列12个（但只能入3个）
	}

	for i := 0; i < 12; i++ {
		Q.Dequeue() //出列12次（但只能出10个）
	}
}
