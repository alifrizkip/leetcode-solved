type MinStack struct {
    min int
    len int
    q []int
}


func Constructor() MinStack {
    return MinStack{
        min: 2147483647,
        q: []int{},
    }
}


func (this *MinStack) Push(val int)  {
    this.q = append(this.q, val)
    this.len++

    if val < this.min {
        this.min = val
    }
}


func (this *MinStack) Pop()  {
    val := this.q[this.len-1]
    this.q = this.q[:this.len-1]
    this.len--
    
    // recalc min
    if val == this.min {
        this.min = 2147483647
        for _, v := range this.q {
            if v < this.min {
                this.min = v
            }
        }
    }
}


func (this *MinStack) Top() int {
    return this.q[this.len-1]
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */