type ParkingSystem struct {
    slots [3]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    slots := [3]int{big, medium, small}
    return ParkingSystem{
        slots: slots,
    }
}


func (this *ParkingSystem) AddCar(carType int) bool {
    if this.slots[carType - 1] == 0 {
        return false
    } else {
        this.slots[carType - 1] -= 1
        return true
    }
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */