const MAX_LEVEL = 32 

type SkiplistNode struct {
    Elem int 
    Level []*SkiplistNode 
}

type Skiplist struct {
    Head *SkiplistNode 
    CurMaxLevel int 
    Rand *rand.Rand
}

func Constructor() Skiplist {
    skl := Skiplist {
        Head: new(SkiplistNode),
        CurMaxLevel:0,
    }
    
    skl.Head.Level = make([]*SkiplistNode, MAX_LEVEL) 
    source := rand.NewSource(time.Now().UnixNano())
    skl.Rand = rand.New(source)

    return skl
}

func (this *Skiplist) Search(target int) bool {
    p := this.Head
    for i := this.CurMaxLevel-1; i >= 0; i-- {
        for p.Level[i] != nil {
            if p.Level[i].Elem == target {
                return true
            } 
            if p.Level[i].Elem > target { 
                break 
            }
            p = p.Level[i] 
        }
    }

    return false
}

func (this *Skiplist) Add(num int)  {
    update := make([]*SkiplistNode, MAX_LEVEL)
    p := this.Head
    for i := this.CurMaxLevel-1; i >= 0; i-- {
        for p.Level[i] != nil && p.Level[i].Elem < num { 
            p = p.Level[i]
        }
        update[i] = p
    }
    
    level := 1
    for this.Rand.Intn(2) == 1{
        level++
    }
    if level > MAX_LEVEL { 
        level = MAX_LEVEL
    }
    
    if level > this.CurMaxLevel {
        for i := this.CurMaxLevel; i < level; i++ {
            update[i] = this.Head
        }
        this.CurMaxLevel = level
    }
    
    newNode := new(SkiplistNode)
    newNode.Elem = num
    newNode.Level = make([]*SkiplistNode, level)
    
    for i := 0; i < level; i++ {
        newNode.Level[i] = update[i].Level[i]
        update[i].Level[i] = newNode
    }
}

func (this *Skiplist) Erase(num int) bool {
    if this.Head.Level[0] == nil { //Skiplist is null
        return false
    }
    
    update := make([]*SkiplistNode, MAX_LEVEL)
    p := this.Head
    for i := this.CurMaxLevel-1; i >= 0; i-- {
        for p.Level[i] != nil && p.Level[i].Elem < num {
            p = p.Level[i]
        }
        update[i] = p
    }
    
    if update[0].Level[0] == nil || update[0].Level[0].Elem != num { 
        return false
    }

    level := len(update[0].Level[0].Level) 
    for i := 0; i < level ; i++ {
        update[i].Level[i] = update[i].Level[i].Level[i]
    }

    for i := this.CurMaxLevel-1; this.Head.Level[i] == nil; i-- {
        this.CurMaxLevel--
    }
    
    return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */