type MyCircularQueue struct {
    cap int
    length int
    datas []int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        cap: k,
        datas: []int{},
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.length == this.cap {
        return false
    }

    this.datas = append(this.datas, value)
    this.length++

    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.length == 0 {
        return false
    }

    this.datas = this.datas[1:]
    this.length--

    return true
}


func (this *MyCircularQueue) Front() int {
    if this.length == 0 {
        return -1
    }

    return this.datas[0]
}


func (this *MyCircularQueue) Rear() int {
    if this.length == 0 {
        return -1
    }

    return this.datas[this.length-1]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.length == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.length == this.cap
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */