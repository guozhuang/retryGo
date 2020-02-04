package mock

type Retrieve struct {
	Contents string
}

//这里细节处仅需要对应的结构体实现了被调用者的get方法
//只不过因为这个是测试用例，真正的项目中，肯定是先写当前的被调用者，
// 然后再调用者定义接口
func (r Retrieve) Get(url string) string {
	return r.Contents
}

/**
从这个角度来看，所以golang的net包只需要实现来listen以及处理的handle方法
然后在调用者声明来这些interface，即可正常使用这些方法来实现http的处理逻辑
【所以golang本身的http处理逻辑是非常清晰的】
*/
