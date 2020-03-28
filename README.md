# 面试题
## 题目
有一个写满 url 大小为 100 GB的文件，求出现次数最多的 50 个 url。 可以用 golang 实现，代码可以写在 GitHub 上
## 解题思路
这道题属于经典的topK问题，解决方法就是`HashMap` + `最小堆`。先使用`HashMap`将文件切割100份，每份小文件大概1G左右。再使用50个文件的`最小堆`做一个出现次数的排序。

## 代码结构
```bash
top_url
├── BKDRHash  
├── CreateUrl  # 生成url
├── FileCut    
├── FileToHeap
├── MinHeap
├── README.md
├── ResultToFile
├── create_url  # 生成url即data.txt主程序，用于测试
├── data.txt
├── go.mod
├── main.go     
├── result.txt  # 结果文件
├── tmp  # 切割文件tmp目录
├── top_url   # 主程序
└── util
```

## 代码执行
```bash
cd CreateUrl 
go build -o ../create_url
cd ..
go build -o top_url
./create_url
./top_url
cat result.txt
```

