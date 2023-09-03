type RandomizedSet struct {
    data map[int]int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        data: make(map[int]int),
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.data[val]; !ok {
        this.data[val] = val
        return true
    }

    return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.data[val]; ok {
        delete(this.data, val)
        return true
    }

    return false
}


func (this *RandomizedSet) GetRandom() int {
    l := len(this.data)
    idx := rand.Intn(l)
    c := 0

    for k, _ := range this.data {
        if c == idx {
            return k
        }
        c++
    }

    return -1
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */