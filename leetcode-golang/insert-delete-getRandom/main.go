package main

import "math/rand"

type RandomizedSet struct {
	dic map[int]int
}

func Constructor() RandomizedSet {
	dic := make(map[int]int)
	rs := RandomizedSet {
		dic: dic,
	}
	return rs
}

func (rs *RandomizedSet) Insert(val int) bool {
	_, ok := rs.dic[val] 
		if !ok {
			rs.dic[val]++ 
			return true //return true if it does not exist yet
		}
		return false
}

func (rs *RandomizedSet) Remove(val int) bool {
	_, ok:= rs.dic[val]
		if ok {
			delete(rs.dic, val)
			return true //return true when it exists
		}
		return true
}

//either get the return key from the for-loop(will always return a value)
//or get it from the a as it will always be given a value
func (rs *RandomizedSet) GetRandom() int {
	num := rand.Intn(len(rs.dic))
	i := 0
	var a int
	for key, _ := range rs.dic {
		a = key
		if i == num{
			return key
		}
	}
	return a
}

func main(){
	
}