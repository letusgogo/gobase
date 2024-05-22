package util

// GetGStrFromCStr 获取一个以 \0 结尾的 c 字符串,如果没有则以长度不变
func GetGStrFromCStr(data []byte) string {
	validByte := make([]byte, 0)
	for _, c := range data {
		// 遇到0 终止
		if c == 0x00 {
			break
		}
		validByte = append(validByte, c)
	}
	return string(validByte)
}

// GetCStrFromGStr 在 go string 后面补充 \0。
func GetCStrFromGStr(gStr string) []byte {
	cStr := []byte(gStr)[0:]
	cStr = append(cStr, 0x00)
	return cStr
}
