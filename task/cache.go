package task

type cache struct {
	cacheMap map[string]string
}

var (
	Cache cache
)

func init() {
	Cache.cacheMap = make(map[string]string)
}

//insert the record to cache
func (c cache) insert(key, value string) {
	c.cacheMap[key] = value
}

//get the record from cache
func (c cache) get(key string) string {
	if val, ok := c.cacheMap[key]; ok {
		return val
	}
	return ""
}
