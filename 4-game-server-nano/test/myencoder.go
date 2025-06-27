package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

// 定义数据包类型，参考 nano 框架
const (
	Handshake    = 0x01
	HandshakeAck = 0x02
	Heartbeat    = 0x03
	Data         = 0x04
	Kick         = 0x05
)

// Encode 手动编码消息
func Encode(typ byte, data []byte) ([]byte, error) {
	length := len(data)
	// 头部长度为 4 字节（1 字节类型 + 3 字节长度）
	header := make([]byte, 4)
	header[0] = typ
	// 创建一个临时 4 字节切片来存储 uint32 类型的长度
	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, uint32(length))
	// 复制后 3 字节到 header[1:]
	copy(header[1:], tmp[1:])
	// 组合头部和数据
	encoded := make([]byte, 0, 4+length)
	encoded = append(encoded, header...)
	encoded = append(encoded, data...)
	return encoded, nil
}

func main() {
	// 示例消息数据
	messageData := "{\n    \"route\": \"room.handlejoinroom\",\n    \"data\": {\n        \"roomID\": \"123456\",\n        \"playerName\": \"TestPlayer\",\n        \"playerID\": \"p001\"\n    }\n}"
	// 将消息数据序列化为 JSON 字节切片
	jsonData, err := json.Marshal(messageData)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}
	// 使用手动编码方法进行封装
	encodedData, err := Encode(Data, jsonData)
	if err != nil {
		fmt.Println("Encode error:", err)
		return
	}
	// 打印十六进制编码的数据，方便在 Postman 中发送
	fmt.Printf("Encoded data (hex): %x\n", encodedData)
}
