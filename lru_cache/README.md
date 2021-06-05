# intro
### LRU
LRU（Least Recently Used，最近最久未使用算法）是一種常見的緩存淘汰演算法，當緩存滿時，淘汰最近久未使用的元素。

思路是，如果一個數據在一段時間內沒有被訪問到，那麼可以認為，在將來被訪問的可能性也很小。因此，當緩存滿時，最久未被訪問的數據將優先淘汰。

具體作法是將最近使用的元素存放到靠近緩存頂部的位置。

當一個新條目被訪問時 LRU 將它放置到緩存的頂部，當緩存滿時，較早以前訪問的條目將從緩存底部被移除。

### groupcache LRU cache 介紹
在 Go 中，如果要使用 LRU cache，可以透過 Google 的 groupcache，透過 groupcache/lru/lru.go 實現一系列封裝過後的 LRU 緩存操作相關接口

```go
// 創建一個 LRU Cache
func New(maxEntries int) *Cache
 
// 向 Cache 中 插入一個 Key/ Value
func (c *Cache) Add(key Key, value interface{})

// 從 Cache 中獲取一個 Key 對應的 Value
func (c *Cache) Get(key Key) (value interface{}, ok bool)

// 從 Cache 中刪除一個 Key
func (c *Cache) Remove(key Key)

// 從 Cacahe 中刪除最久被訪問的數據
func (c *Cache) RemoveOldest()

// 獲取 Cache 中當前的元素個數
func (c *Cache) Len()

// 清空 Cache
func (c *Cache) Clear()
```
> 注意，groupcache 的 LRU Cache 並不是併發安全的，如果用於多個 Goroutine，需加鎖。


### 源碼解析
LRU Cache 基於`map`與`list`
- map 用於快速檢索
- list 實現 LRU

```go
package lru

import "container/list"

// Cache 資料結構，是一個 LRU Cache 但並不是併發安全的
type Cache struct {
	// MaxEntryies 是 Cache 中實體的最大數量，0表示不設限
	MaxEntries int
	
	// OnEvicted 定義一個可選的 callback function，當一個實體從 Cache 中被移除時執行
	OnEvicted func(key Key, value interface{})
	
	// ll 是一個雙向鏈表指針，執行一個 container/list package 中的雙向鏈表
	ll *list.List
	
	// cache 是一個 map，存放具體的 Key/ Value，對 value 是雙向鏈表中的具體元素，也就是 *Element
	cache map[interface{}]*list.Element
}

// Key 是 interface，可以為任意型態
type Key interface{}

// 一個 entry 包含一個 Key 和一個 value，皆為interface
type entry struct {
	key   Key
	value interface{}
}

// 創建一個 LRU Cache; maxEntries 為 0 表示緩存沒有大小限制
func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

// 向 Cache 中插入一個 Key/ Value
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

// 查詢一個 Key，並且返回對應的 Value
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

// 從 Cache 中刪除一個 Key/ Value
func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

// 從 Cache 中刪除最久未被訪問的數據
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

// 從 Cache 中刪除一個元素，供內部使用
func (c *Cache) removeElement(e *list.Element) {
	// 先從 list 中刪除
	c.ll.Remove(e)
	
	kv := e.Value.(*entry)

	// 再從 map 中刪除
	delete(c.cache, kv.key)
	
	// 如果callback函數不為空則調用
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

// 獲取 Cache 當前的元素個數
func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}

// 清空 Cache
func (c *Cache) Clear() {
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.ll = nil
	c.cache = nil
}
```



# refer:
- https://cloud.tencent.com/developer/article/1478020