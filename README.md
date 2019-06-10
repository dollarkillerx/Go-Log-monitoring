# Go-Log-monitoring
Go 生产级日志监控系统

### 简单概括
- 日志文件
- log_process 实时读取解析写入
- influxdb 存储
- grafana 展示

![](./README/bf.png)

### 读取模块实现
- 从日志文件末尾开始逐行读取
- 写入Read Channel
``` 
func (l *ReadFromFile) Read(rc chan []byte) {
	// 打开文件
	file, e := os.Open("./tmp/access.log")
	defer file.Close()
	if e != nil {
		panic(e.Error())
	}
	// 文件指针指向文件末尾 从末尾开始读
	file.Seek(0,2)

	// 从文件末尾开始卒行读取
	reader := bufio.NewReader(file)

	for  {
		bytes, e := reader.ReadBytes('\n')
		if e == io.EOF{
			time.Sleep(200 * time.Millisecond)
			continue
		}else if e != nil {
			log.Fatal(e.Error())
		}
		fmt.Printf("log: %s",bytes)
		rc <- bytes
	}
}
```

### 解析模块
- 从channel中读取每行日志数据
- 正则提取所需监控数据(path,status,method等)
- 写入write Channel
日志格式 
`127.0.0.1 -- [04/Mar/2019:13:49:52 +0000] http "Get /foo?query=t HTTP/1.0" 200 2133 "-" 
"KeepAliveCLient" "-" 1.005 1.854`