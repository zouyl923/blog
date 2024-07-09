package regex

// 手机号验证
var Phone string = "^1([0-9]){10}$"

// 弱密码 至少包含字母、数字 6~32位
var WeakPassword string = "^(?=.*[a-zA-Z])(?=.*\\d).{6,32}$"

// 强密码 由字母、数字、特殊字符，任意2种组成，8-32位
var StrongPassword string = "^(?![a-zA-Z]+$)(?!\\d+$)(?![^\\da-zA-Z\\s]+$).{8,32}$"
