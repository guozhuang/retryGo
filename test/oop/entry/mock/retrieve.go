package mock

//此处是被调用者
type Retrieve struct {
	Contents string
}

func (r *Retrieve) Get(url string) string {
	return r.Contents
}
