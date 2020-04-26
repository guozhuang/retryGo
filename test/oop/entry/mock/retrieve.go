package mock

import "fmt"

//此处是被调用者
type Retrieve struct {
	Contents string
}

func (r *Retrieve) Get(url string) string {
	return r.Contents
}

//希望该结构体实现新的组合接口，所以扩充相应的方法
func (r *Retrieve) Post(url string, data map[string]string) string {
	if value, ok := data["user"]; ok {
		r.Contents = value //同时进行了值的变更，get时拿到的也是这里传入的值
		return value
	}

	return ""
}

//此时retrieve实现了stringer接口，所以在使用fmt的print方法时，就会自动调用这里的方法
func (r *Retrieve) String() string {
	return fmt.Sprintf("retrieve ins stringer: %s", r.Contents)
}
