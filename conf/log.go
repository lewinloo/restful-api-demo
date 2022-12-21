package conf

type LogFormat string

const (
	// 文本格式
	TextFormat = LogFormat("text")
	// json 格式
	JSONFormat = LogFormat("json")
)

type LogTo string

const (
	// 保存到文件
	ToFile = LogTo("file")
	// 打印到标准输出
	ToStdout = LogTo("stdout")
)
