type MyStack struct {
    length int
    datas []int
}


func Constructor() MyStack {
    return MyStack{
        datas: []int{},
    }
}


func (this *MyStack) Push(x int)  {
    this.datas = append(this.datas, x)
    this.length++
}


func (this *MyStack) Pop() int {
    res := this.datas[len(this.datas)-1]
    this.datas = this.datas[:len(this.datas)-1]
    if this.length > 0 {
        this.length--
    }

    return res
}


func (this *MyStack) Top() int {
    return this.datas[len(this.datas)-1]
}


func (this *MyStack) Empty() bool {
    return this.length == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */