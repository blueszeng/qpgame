package util

import (
	"crypto/md5"
	"encoding/hex"
	//"fmt"
	//. "qpgame/base"
	//	"strconv"
)

/**
 * 对一个字符串进行MD5加密,不可解密
 */
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s)) //使用zhifeiya名字做散列值，设定后不要变
	return hex.EncodeToString(h.Sum(nil))
}

func CheckPassword(src_pwd, dst_pwd string) bool {
	return GetMd5String(src_pwd) == dst_pwd
}
