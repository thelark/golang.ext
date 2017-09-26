package safemap

import "sync"

type Map struct {
	*sync.RWMutex
	_map map[interface{}]interface{}
}

func New() *Map {
	return &Map{
		new(sync.RWMutex),
		make(map[interface{}]interface{}),
	}
}
func (m *Map) Get(key interface{}) interface{} {
	m.RLock()
	defer m.RUnlock()
	if val, ok := m._map[key]; ok {
		return val
	}
	return nil
}
func (m *Map) Set(key interface{}, val interface{}) bool {
	m.Lock()
	defer m.Unlock()
	if v, ok := m._map[key]; !ok {
		m._map[key] = val
	} else if v != val {
		m._map[key] = val
	} else {
		return false
	}
	return true
}
func (m *Map) Check(key interface{}) bool {
	m.RLock()
	defer m.RUnlock()
	if _, ok := m._map[key]; !ok {
		return false
	}
	return true
}
func (m *Map) Delete(key interface{}) {
	m.Lock()
	defer m.Unlock()
	delete(m._map, key)
}
