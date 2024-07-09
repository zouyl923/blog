package helper

import "encoding/json"

/**
 * 把m转换成c
 */
func Swap(m interface{}, c interface{}) {
	j, _ := json.Marshal(m)
	_ = json.Unmarshal(j, c)
}
