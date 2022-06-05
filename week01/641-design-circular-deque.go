type MyCircularDeque struct {
    Head  int
	Tail  int
	Value []int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
		Head:  0,
		Tail:  0,
		Value: make([]int, k+1),
	}
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
		return false
	}
	this.Head = (this.Head - 1 + len(this.Value)) % len(this.Value)
	this.Value[this.Head] = value
	return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
		return false
	}
	this.Value[this.Tail] = value
	this.Tail = (this.Tail + 1) % len(this.Value)
	return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
		return false
	}
	this.Head = (this.Head + 1) % len(this.Value)
	return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
		return false
	}
	this.Tail = (this.Tail - 1 + len(this.Value)) % len(this.Value)
	return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
		return - 1
	}
	return this.Value[this.Head]
}


func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
		return - 1
	}
	return this.Value[(this.Tail-1+len(this.Value))%len(this.Value)]
}


func (this *MyCircularDeque) IsEmpty() bool {
    if this.Head == this.Tail {
		return true
	}
	return false
}


func (this *MyCircularDeque) IsFull() bool {
    if (this.Tail + 1) % len(this.Value) == this.Head {
		return true
	}
	return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */