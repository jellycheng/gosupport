package gosupport

import (
	"sync"
	"sync/atomic"
)

type DataManage struct {
	
	DataMutex *sync.RWMutex
	Data map[string]interface{}
	
}

func (dm *DataManage) Set(key string, value interface{}) {
	if dm.DataMutex == nil {
		dm.DataMutex = &sync.RWMutex{}
	}

	dm.DataMutex.Lock()
	if dm.Data == nil {
		dm.Data = make(map[string]interface{})
	}

	dm.Data[key] = value
	dm.DataMutex.Unlock()
}


func (dm *DataManage) Get(key string) (value interface{}, exists bool) {
	if dm.DataMutex == nil {
		dm.DataMutex = &sync.RWMutex{}
	}

	dm.DataMutex.RLock()
	value, exists = dm.Data[key]
	dm.DataMutex.RUnlock()
	return
}

func (dm *DataManage) MustGet(key string) interface{} {
	if value, exists := dm.Get(key); exists {
		return value
	}
	panic("DataManage struct Data Key \"" + key + "\" does not exist")
}

func (dm *DataManage) GetString(key string) (s string) {
	if val, ok := dm.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

func (dm *DataManage) GetBool(key string) (b bool) {
	if val, ok := dm.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

func (dm *DataManage) GetInt(key string) (i int) {
	if val, ok := dm.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

func (dm *DataManage) GetInt64(key string) (i64 int64) {
	if val, ok := dm.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

func (dm *DataManage) GetFloat64(key string) (f64 float64) {
	if val, ok := dm.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

//	globalCfg := gosupport.NewGlobalCfgSingleton()
//	globalCfg.Set("host", "127.0.0.1")
var muGlobalcfgDM_1 sync.Mutex
var globalcfgDM_1 *DataManage
var globalcfgInit_1 uint32
func NewGlobalCfgSingleton() *DataManage {
	if atomic.LoadUint32(&globalcfgInit_1) == 1 {//确保原子性
		return globalcfgDM_1
	}
	muGlobalcfgDM_1.Lock()
	defer muGlobalcfgDM_1.Unlock()
	if globalcfgInit_1 == 0 { //未初始化
		globalcfgDM_1 = new(DataManage)
		atomic.StoreUint32(&globalcfgInit_1, 1)
	}
	return globalcfgDM_1
}

