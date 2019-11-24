package goutil

// 获取一个以 \0 结尾的 c 字符串,如果没有则以长度不变
func GetGStrFromCStr(data []byte) string {
	vaildByte := make([]byte, 0)
	for _, c := range data {
		// 遇到0 终止
		if c == 0x00 {
			break
		}
		vaildByte = append(vaildByte, c)
	}
	return string(vaildByte)
}

// 在 go string 后面补充 \0。
func GetCStrFromGStr(gstr string) []byte {
	cstr := []byte(gstr)[0:]
	cstr = append(cstr, 0x00)
	return cstr
}
