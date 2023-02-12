package utils

// 发送验证码主邮箱
const ServerEmail = ""

// 发送邮箱授权码
const EmailAuthCode = ""

// 邮箱Smtp地址
const EmailSmtpAddr = ""

// 邮箱SmtpHost
const EmailSmtpHost = ""

const VerificationCodeLength = 6

const DefaultPageSize = 8

// 存储桶名称
const BucketNameWithAPPID = ""

// 存储桶地域
const CosRegion = ""
const CosUrl = "https://" + BucketNameWithAPPID + "" + CosRegion + ""
const SecretID = ""
const SecretKey = ""

// 块大小，1MB
const ChunkSize = 1024 * 1024

// 缓存标识key
const CacheRepositoryKey = "cache:repository:"
const CacheShareKey = "cache:share:"
const CacheUserKey = "cache:user:"
const CacheUserRepositoryKey = "cache:userRepository:"
const CacheEmailCodeKey = "cache:email:code:"

// 缓存过期
const RedisLockExpireSeconds = 10

const EmailCodeExpireSeconds = 300

// 当前时间戳
const BeginTimeStamp = 1675580392

// Id序列号部分的位长
const IdCountBit = 32
