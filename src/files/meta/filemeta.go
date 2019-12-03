package meta

//文件元数据结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta //包内全局的元数据信息，key使用sha1

func init() {
	//暂存在程序内存中
	fileMetas = make(map[string]FileMeta)
}

//进行文件元信息的追加
func UpdateFileMeta(fMeta FileMeta) {
	fileMetas[fMeta.FileSha1] = fMeta
}

//获取文件元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//文件元数据清理
func RemoveFileMeta(fileSha1 string) {
	//todo:需要注意协程同步以及验证是否有问题
	delete(fileMetas, fileSha1)
}
