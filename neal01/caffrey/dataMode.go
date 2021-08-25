package caffrey
import "sync"

//声明一个保存数据的结构体
type SafeData struct {
	data []string
	m sync.Mutex
}

//获取存储数据的方法
func (s *SafeData) Get() []string {
	s.m.Lock()
	defer s.m.Unlock()
	return s.data
}

//增加存储数据的方法
func (s *SafeData) Add(slice []string) {
	s.m.Lock()
	defer s.m.Unlock()
	//将数据放入存储数据中
	s.data = append(s.data, slice...)
	//对存储数据进行去重
	res := make([]string,0,len(s.data))
	temp := map[string]struct{}{}
	for _,val := range s.data{
		if _,ok := temp[val];!ok {
			temp[val] = struct{}{}
			res = append(res, val)
		}
	}
	s.data = append(s.data[:0],res... )
}


//返回数据结构
type ResponeData struct {
	Data []bool `json:"data"`
}