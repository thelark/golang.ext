<style>* {font-family:'微软雅黑';}</style>
# Golang 包扩展
 ## base 包
  - [x] Struct => Map
  - [x] Map => Struct
  - [x] 判断字符串是否为空

 ## color 包
  #### cmd 包
   - [x] Black 前景色
   - [x] Red 前景色
   - [x] Green 前景色
   - [x] Yellow 前景色
   - [x] Blue 前景色
   - [x] Magenta 前景色
   - [x] Cyan 前景色
   - [x] White 前景色
   - [x] Gray 前景色
   - [ ] Black 背景色
   - [ ] Red 背景色
   - [ ] Green 背景色
   - [ ] Yellow 背景色
   - [ ] Blue 背景色
   - [ ] Magenta 背景色
   - [ ] Cyan 背景色
   - [ ] White 背景色
   - [ ] Gray 背景色
 ## file 包
  - [x] 判断路径是否存在（文件|文件夹）
  - [x] 创建目录（多级创建）
 ## path 包
  - [x] 获取文件路径的相对路径 (相对GOPATH)
  - [ ] 获取以当前执行目录为相对参照的文件的绝对路径
 ## log 包
  - [x] FATAL 输出
  - [x] ERROR 输出
  - [x] WARN 输出
  - [x] INFO 输出
  - [x] DEBUG 输出
  - [x] 初始化 log 文件存储
  - [x] Open 开启日志记录
  - [x] Close 关闭日志记录
  - [x] SetFlags 设置记录项标识
  - [x] SetFile 设置记录文件
  - [x] SetLevel 设置日志级别

  > 日志级别
  > - OFF 不记录任何
  > - FATAL 记录致命错误
  > - ERROR 记录普通错误及以上
  > - WARN 记录警告及以上
  > - INFO 记录普通输出及以上
  > - DEBUG 记录调试输出及以上
  > - ALL 记录所有日志

  > 日志标识
  > - Ltime 时间标识
  > - Ldate 日期标识
  > - Lmicroseconds 微秒标识
  > - Llongfile 长文件名标识
  > - Lshortfile 短文件名标识
  > - LUTC UTC时间标识

 ## safe 包
  #### safemap 包
   - [x] New 返回 SafeMap
   - [x] Get 获取 Map 的值
   - [x] Set 设置 Map 的值
   - [x] Check 校验 Map 的值
   - [x] Delete 删除 Map 的值
  #### safefile 包
   - [ ] 安全读
   - [ ] 安全写
   - [ ] 安全复制
   - [ ] 安全删除
 ## unsafe 包
  - [x] string => int
  - [x] string => int8
  - [x] string => int16
  - [x] string => int32
  - [x] string => int64
  - [x] string => uint8
  - [x] string => uint16
  - [x] string => uint32
  - [x] string => uint64
  - [x] string => float64
  - [x] string => float32
  - [x] string => bool
 
 ## time 包
  - [x] time.Time => string
  - [x] 时间戳 => string
  - [x] string => (UTC)time.Time (不安全)
  - [x] string => (本地)time.Time (不安全)
  - [x] string => (UTC)时间戳 (不安全)
  - [x] string => (本地)时间戳 (不安全)
 
 ## wrong 包
  - [x] Println 错误按行打印
  - [x] Panicln 错误恐慌打印 (panic)
  - [x] Fatalln 错误致命打印 (fatal)

 ## crypto 包
  - [ ] MD5 加密
  - [ ] RSA 加密
  - [ ] AES 加密


> ### 注意
> * 所谓不安全即不考虑 error



