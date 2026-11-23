package main

// First implementation. 125ms runtime and 20.11MB memory use.
type ProductOfNumbers struct {
    prefixes []int
}

func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        prefixes: []int{1},
    }
}



func (this *ProductOfNumbers) Add(num int)  {
    if num == 0{
        this.prefixes = []int{1}
    } else {
        mul := this.prefixes[len(this.prefixes)-1] * num
        this.prefixes = append(this.prefixes, mul)
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    n := len(this.prefixes)
    if k >= n {
        return 0
    }
    last := n -1
    kSteps := n-1-k
    return this.prefixes[last]/this.prefixes[kSteps]
}




/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
