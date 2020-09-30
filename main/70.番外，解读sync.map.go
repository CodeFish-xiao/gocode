package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

//Map类似于Go map [interface {} {interface] interface {}，
//但可以安全地供多个goroutine并发使用，而无需额外的锁定或协调。
//加载，存储和删除均以摊销的固定时间运行。

//映射类型是专用的。 大多数代码应改用带有单独锁定或协调功能的普通Go映射，
//以提高类型安全性，并使其更易于维护其他不变式以及映射内容。
//
//
//Map类型针对两种常见用例进行了优化：
//（1）给定键的条目仅写入一次但多次读取（例如仅在增长的高速缓存中），
//或者（2）当多个goroutine进行读取，写入和； 覆盖不相交的键集的条目。
//在这两种情况下，与与单独的Mutex或RWMutex配对的Go映射相比，使用Map可以显着减少锁争用。
type Map struct {
	mu sync.Mutex

	//read包含映射内容中对当前并发访问安全的部分（无论是否保留mu）。
	//read 字段本身始终是可以安全加载的，但只能与 mu 一起保存。
	//
	//存储在read中的条目可以不使用mu并发更新，但是可以更新
	//先前删除的条目要求将该条目复制到脏条目
	//映射，并保留mu。
	//

	//存储在read中的条目可以不使用mu并发更新，但是更新以前删除的条目要求将该条目复制到脏映射，并且在保留mu的情况下不删除它。
	read atomic.Value // 只读

	//脏映射包含映射内容中需要保留mu的部分。 为了确保可以将脏映射快速提升到读取映射，它还包括读取映射中的所有未删除条目。
	//
	// 删除的条目不会存储在脏映射中。 必须先清除干净映射中的已删除条目，然后将其添加到脏映射中，然后才能向其存储新值。
	//
	// 如果脏映射为nil，则对映射的下一次写入将通过制作干净映射的浅表副本（省略陈旧的条目）来初始化它。
	dirty map[interface{}]*entry

	// 未命中数最后更新，因为读取映射加载的数量需要锁定mu，以确定键是否存在。
	//
	// 一旦发生足够的未命中以支付复制脏映射的成本，
	//该脏映射将被提升为已读映射（处于未修改状态），
	//并且该映射的下一个存储区将进行新的脏复制。
	misses int
}

// readOnly是原子形式存储在Map.read字段中的不变结构。
type readOnly struct {
	m       map[interface{}]*entry
	amended bool //如果脏映射包含不在m中的某个键，则为true。
}

//
//expunged是一个任意指针，用于标记已从脏映射中删除的条目。
var expunged = unsafe.Pointer(new(interface{}))

//条目是映射中对应于特定键的插槽。
type entry struct {
	// p points to the interface{} value stored for the entry.
	// p指向为 entry 存储的interface {}值。
	// If p == nil, the entry has been deleted and m.dirty == nil.
	//
	// If p == expunged, the entry has been deleted, m.dirty != nil, and the entry
	// is missing from m.dirty.
	//
	// 除此以外, dirty有效，并记录在m.read.m [key] and, if m.dirty
	// != nil, in m.dirty[key].
	//
	// 可以通过原子替换为nil来删除 entry :下次创建m.dirty时，它将自动用expunged替换nil，并保持m.dirty [key]不变。
	// 条目的关联值可以通过原子替换来更新, provided
	//提供 p != expunged. If p == expunged,
	//只有在首先设置m.dirty [key] = e之后，
	//才可以更新entry的关联值，以便使用脏映射的查找找到该entry 。
	p unsafe.Pointer // *interface{}
}

func newEntry(i interface{}) *entry {
	return &entry{p: unsafe.Pointer(&i)}
}

// Load返回存储在映射中的键值，如果没有值，则返回nil。
//确定结果表明是否在地图中找到了值。
// The ok 结果表明是否在映射中找到了值。
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		// 如果在我们被m.mu封锁时提升了m.dirty，请避免报告虚假的遗漏。
		//(如果相同的密钥的进一步负荷不会错过, 则不值得复制该密钥的脏映射。.)
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			// 不管entry是否存在，都要记录miss:
			// 该键将采用慢速路径，直到将脏映射提升为已读映射为止。
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if !ok {
		return nil, false
	}
	return e.load()
}

func (e *entry) load() (value interface{}, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == nil || p == expunged {
		return nil, false
	}
	return *(*interface{})(p), true
}

// Store 设置键的值。
func (m *Map) Store(key, value interface{}) {
	read, _ := m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok && e.tryStore(&value) {
		return
	}

	m.mu.Lock()
	read, _ = m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			// 该条目先前已删除，这意味着存在一个非零的脏映射，并且该条目不在其中。
			m.dirty[key] = e
		}
		e.storeLocked(&value)
	} else if e, ok := m.dirty[key]; ok {
		e.storeLocked(&value)
	} else {
		if !read.amended {
			// 我们正在向脏映射添加第一个新key。
			// Make sure it is allocated and mark the read-only map as incomplete.
			//确保已分配它，并将只读映射标记为不完整。
			m.dirtyLocked()
			m.read.Store(readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
	}
	m.mu.Unlock()
}

// tryStore 如果entry未删除，则存储一个值。
//
// 如果删除该条目，则tryStore返回false并使该条目保持不变。
func (e *entry) tryStore(i *interface{}) bool {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == expunged {
			return false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
			return true
		}
	}
}

// unexpungeLocked 确保该条目未标记为清除。
//
// 如果该条目先前已删除，则必须在解锁m.mu之前将其添加到脏映射中。
func (e *entry) unexpungeLocked() (wasExpunged bool) {
	return atomic.CompareAndSwapPointer(&e.p, expunged, nil)
}

// storeLocked 无条件地将值存储到条目。
//
// 必须知道该条目不会被清除。
func (e *entry) storeLocked(i *interface{}) {
	atomic.StorePointer(&e.p, unsafe.Pointer(i))
}

// LoadOrStore 返回键的现有值（如果存在）。
// 否则，它将存储并返回给定的值。
// 如果已加载该值，则加载的结果为true；如果已存储，则为false。
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	// 避免锁定，如果它是一个干净的命中。
	read, _ := m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		actual, loaded, ok := e.tryLoadOrStore(value)
		if ok {
			return actual, loaded
		}
	}

	m.mu.Lock()
	read, _ = m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			m.dirty[key] = e
		}
		actual, loaded, _ = e.tryLoadOrStore(value)
	} else if e, ok := m.dirty[key]; ok {
		actual, loaded, _ = e.tryLoadOrStore(value)
		m.missLocked()
	} else {
		if !read.amended {
			// 我们正在向脏映射添加第一个新key。
			// Make sure it is allocated and mark the read-only map as incomplete.
			//确保已分配它，并将只读映射标记为不完整。
			m.dirtyLocked()
			m.read.Store(readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
		actual, loaded = value, false
	}
	m.mu.Unlock()

	return actual, loaded
}

// tryLoadOrStore 如果条目没有删除，则自动加载或存储一个值。
//
// 如果该条目被删除，tryLoadOrStore将使该条目保持不变，并以ok == false返回。
func (e *entry) tryLoadOrStore(i interface{}) (actual interface{}, loaded, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == expunged {
		return nil, false, false
	}
	if p != nil {
		return *(*interface{})(p), true, true
	}

	//在第一次加载后复制接口，以使此方法更适合进行分析：如果我们单击“加载”路径或删除了条目，则不应理会堆分配。
	ic := i
	for {
		if atomic.CompareAndSwapPointer(&e.p, nil, unsafe.Pointer(&ic)) {
			return i, false, true
		}
		p = atomic.LoadPointer(&e.p)
		if p == expunged {
			return nil, false, false
		}
		if p != nil {
			return *(*interface{})(p), true, true
		}
	}
}

// LoadAndDelete 删除键的值，如果有则返回前一个值。
// 加载的结果报告密钥是否存在。
func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			delete(m.dirty, key)
			//无论是否存在该条目，都记录一个未命中：该键将采用慢速路径，直到将脏映射提升为已读映射为止。
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if ok {
		return e.delete()
	}
	return nil, false
}

// Delete 删除键的值。
func (m *Map) Delete(key interface{}) {
	m.LoadAndDelete(key)
}

func (e *entry) delete() (value interface{}, ok bool) {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == nil || p == expunged {
			return nil, false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, nil) {
			return *(*interface{})(p), true
		}
	}
}

// Range 对映射中存在的每个键和值依次调用f。 如果f返回false，则range停止迭代。
//
//Range不一定与Map内容的任何一致快照相对应：
//不会多次访问任何键，但是如果同时存储或删除任何键的值，则Range可能会在Range调用期间从任何点反映该键的任何映射。
// Range 可能是O（N），且映射中的元素数即使在恒定数量的调用后f返回false也是如此。
func (m *Map) Range(f func(key, value interface{}) bool) {
	// 我们需要能够遍历对Range的调用开始时已经存在的所有键。
	// 如果read.amended为false，则read.m满足该属性，而无需我们长时间持有m.mu。
	read, _ := m.read.Load().(readOnly)
	if read.amended {
		// m.dirty包含不在read.m中的键。
		//幸运的是，Range已经是O（N）（假设调用者没有提前中断），
		//因此对Range的调用将摊销映射的整个副本：我们可以立即升级脏副本！
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly)
		if read.amended {
			read = readOnly{m: m.dirty}
			m.read.Store(read)
			m.dirty = nil
			m.misses = 0
		}
		m.mu.Unlock()
	}

	for k, e := range read.m {
		v, ok := e.load()
		if !ok {
			continue
		}
		if !f(k, v) {
			break
		}
	}
}

func (m *Map) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}
	m.read.Store(readOnly{m: m.dirty})
	m.dirty = nil
	m.misses = 0
}

func (m *Map) dirtyLocked() {
	if m.dirty != nil {
		return
	}

	read, _ := m.read.Load().(readOnly)
	m.dirty = make(map[interface{}]*entry, len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() {
			m.dirty[k] = e
		}
	}
}

func (e *entry) tryExpungeLocked() (isExpunged bool) {
	p := atomic.LoadPointer(&e.p)
	for p == nil {
		if atomic.CompareAndSwapPointer(&e.p, nil, expunged) {
			return true
		}
		p = atomic.LoadPointer(&e.p)
	}
	return p == expunged
}
