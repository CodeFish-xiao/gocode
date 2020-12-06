package main

// 空结构体
var Exists = struct{}{}

// Set is the main interface
type Set struct {
	// struct为结构体类型的变量
	m map[interface{}]struct{}
}

//初始化
//
//Set类型数据结构的初始化操作，在声明的同时可以选择传入或者不传入进去。声明Map切片的时候，Key可以为任意类型的数据，用空接口来实现即可。Value的话按照上面的分析，用空结构体即可：
func New(items ...interface{}) *Set {
	// 获取Set的地址
	s := &Set{}
	// 声明map类型的数据结构
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

//添加
//
//简化操作可以添加不定个数的元素进入到Set中，用变长参数的特性来实现这个需求即可，因为Map不允许Key值相同，所以不必有排重操作。同时将Value数值指定为空结构体类型。
func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}

//包含
//Contains操作其实就是查询操作，看看有没有对应的Item存在，可以利用Map的特性来实现，但是由于不需要Value的数值，所以可以用 _,ok来达到目的：
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

///获取Set长度很简单，只需要获取底层实现的Map的长度即可：
func (s *Set) Size() int {
	return len(s.m)
}

//清除操作的话，可以通过重新初始化Set来实现，如下即为实现过程：
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

//相等
//
//判断两个Set是否相等，可以通过循环遍历来实现，即将A中的每一个元素，查询在B中是否存在，只要有一个不存在，A和B就不相等
func (s *Set) Equal(other *Set) bool {
	// 如果两者Size不相等，就不用比较了
	if s.Size() != other.Size() {
		return false
	}

	// 迭代查询遍历
	for key := range s.m {
		// 只要有一个不存在就返回false
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//子集
func (s *Set) IsSubset(other *Set) bool {
	// s的size长于other，不用说了
	if s.Size() > other.Size() {
		return false
	}
	// 迭代遍历
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
