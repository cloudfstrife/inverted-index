package inverted

import "sync"

// IDContainer document id container
type IDContainer interface {
	// Push document ID into container
	Push(int64)
	// Pop  remove document ID from  container
	Pop(int64)
	// Array return all document ID
	Array() []int64
}

// -----------------------------------------------------------------------

// Container default id container
type Container struct {
	A []int64
}

// NewIDContainer create document id container
func NewIDContainer() IDContainer {
	return &Container{
		A: make([]int64, 0, 10000),
	}
}

// Push document ID into container
func (c *Container) Push(id int64) {
	i := 0
	for ; i < len(c.A); i++ {
		if c.A[i] == id {
			break
		}
	}
	if i != len(c.A) {
		return
	}
	c.A = append(c.A, id)
}

// Pop  remove document ID from container
func (c *Container) Pop(id int64) {
	for i := 0; i < len(c.A); i++ {
		if c.A[i] == id {
			c.A = append(c.A[:i], c.A[i+1:]...)
			i--
		}
	}
}

// Array return all document ID
func (c *Container) Array() []int64 {
	r := make([]int64, len(c.A))
	copy(r, c.A)
	return r
}

// -----------------------------------------------------------------------

// Index implement inverted index
type Index struct {
	lock *sync.RWMutex
	M    map[string]IDContainer
}

//NewIndex create inverted index
func NewIndex() Index {
	return Index{
		lock: &sync.RWMutex{},
		M:    make(map[string]IDContainer),
	}
}

// Push push keyword and document id into index
func (i Index) Push(k string, id int64) {
	i.lock.Lock()
	defer i.lock.Unlock()
	if v, ok := i.M[k]; ok {
		v.Push(id)
		return
	}
	c := NewIDContainer()
	c.Push(id)
	i.M[k] = c
}

// Pop remove document id from keyword
func (i Index) Pop(k string, id int64) {
	i.lock.Lock()
	defer i.lock.Unlock()
	if v, ok := i.M[k]; ok {
		v.Pop(id)
	}
}

// GetAllID get all document id
func (i Index) GetAllID(k string) []int64 {
	i.lock.RLock()
	defer i.lock.RUnlock()
	if v, ok := i.M[k]; ok {
		return v.Array()
	}
	return nil
}
