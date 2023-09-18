package common

// 配置文件前缀
const (
	KeyConfigPrefix  string = "key."  // key配置前缀
	DescConfigPrefix string = "desc." // 描述配置前缀
	TypeConfigPrefix string = "type." // 值类型配置前缀
)

// 指定配置 key
const (
	HttpBindConfig string = KeyConfigPrefix + "bind"
	HttpPortConfig string = KeyConfigPrefix + "port"
	LogLevelConfig string = KeyConfigPrefix + "log.level"
)
