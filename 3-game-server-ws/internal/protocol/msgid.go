package protocol

// type 创建一个新的类型，MsgID 是这个新类型的名称。
// int 是 MsgID 类型的基础类型，它指定了 MsgID 可以存储的整数值的范围
// 它可以帮助开发者明确变量的用途和类型，从而避免类型错误和混淆，增加代码的可读性和可维护性，特别是在大型项目中
type MsgID int

// iota 是 Go 语言里的预定义标识符，在 cst 常量组中从 0 开始递增，每遇到一个新的 cst 常量组，iota 就会重置为 0。
const (
	MsgHeartbeat MsgID = iota + 1
	MsgChat
	MsgError
	MsgSystem
)
