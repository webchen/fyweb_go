package config

// CRYPTOKEY ，加密时的key
const CRYPTOKEY string = "1111111100000000"

// CRYPTOIV ，加密时的偏移量
const CRYPTOIV string = "1111111100000000"

// LOGDIR ， 日志的路径
const LOGDIR string = "/Logs/"

// ACCESSLOGCHECKMINUTE , 日志检查更新时间（秒)
const ACCESSLOGCHECKMINUTE = 1800
// ACCESSLOGCHECKSIZE , 日志检查更新的SIZE（以B为单位）
const ACCESSLOGCHECKSIZE = 1024 * 1024 * 100