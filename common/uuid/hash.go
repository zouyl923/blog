package uuid

import (
	"blog/common/helper"
	"strconv"
	"time"
)

// md5算法 32位uuid
func GenMd5Uuid(nonce string) string {
	t := time.Now().Unix()
	sTime := strconv.FormatInt(t, 10)
	str := nonce + sTime
	return helper.Md5(str)
}

// hash256算法 64为uuid
func GenHash256Uuid(nonce string) string {
	t := time.Now().Unix()
	sTime := strconv.FormatInt(t, 10)
	str := nonce + sTime
	return helper.Hash256(str)
}
