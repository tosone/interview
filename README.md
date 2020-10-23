# 面试

### 开放问题

- Golang 中的 interface struct 组合和 C++ 中的虚函数，虚函数，纯虚函数，继承，多态，重载，模板有哪些相同或者对应的东西，越多越好。
- top K 问题。
- Golang 有什么优势和缺点？这些特性适合做什么？
- 有两个超大的文件，里边都是电话号码，如何设计算法得到他们之间的差集？简述即可。
- 软件往往有很多依赖。Go 语言的处理方法是，将所有依赖打包进一个二进制文件，从而与环境无关。Docker 的方法是提供一个容器，在里面还原所有的依赖。如何看淡这个问题？

### 深入问题

- Golang 中如何会导致内存泄露以及如何排查？
- Goroutine 都说他节省内存，是调度的最小单元，那么他相比于 pthread 是为什么可以节省内存，他们分别占用多少内存？ 2K vs 16K
- C 中指针和 Golang 中指针有什么相同和不同的地方。了解过 CGO 么？他们是如何管理内存的？
- 内存对齐是什么？

``` go
package main

type S struct {
	A uint32
	B uint64
	C uint64
	D uint64
	E struct{}
}

func main() {
	fmt.Println(unsafe.Offsetof(S{}.E))
	fmt.Println(unsafe.Sizeof(S{}.E))
	fmt.Println(unsafe.Sizeof(S{}))
}
```

- Golang 内存管理的基本单位 8K。
- Slice 内存多大？
- 大端编码和小端编码是什么？
- 如何处理程序中的 err，如何去 check err？
- context 注意事项
- 我写的 Golang 的东西需要给别人做集成，但是我的其中算法的部分不想公开，这时候该怎么做？
- 请写一个关于 io.Reader 实现的程序。

### 容器

- Docker 网络模式以及他们之间的区别
- 你何时会使用像 ConfigMap 或 secret 这样的东西？他们的区别是什么？你都如何使用他们？
- 容器内暴露服务的方式以及访问方式？
- 简述 sidecar 容器，init 容器的区别？Pod 内容器启动顺序？
- 在构建和管理生产集群时遇到的主要问题是什么？
- 这块有遇见什么比较好用的工具么？他给你解决了什么问题？
- 私下自己有没有什么有趣的东西或者项目？
- namespace, systemd, runC, containerd, docker 这些都是什么东西，他们之间什么关系。

### 工具

- Git 中 branch tag commit 都是什么？rebase 和 merge 有什么区别？

### 协议

- Websocket 握手细节？
- TLS 握手细节？
- OAuth2，SSO 都是什么？
- restful 接口定义都包含哪些内容？

### 算法

- 全覆盖问题，如何判定给定的一个整数数组全覆盖到一个连续整数区间内的所有数字，要求内存最优时，如何设计算法？利用全覆盖的算法如何设计求解数独的算法，要求内存最优时，如何设计算法？

### 数据库

- MySQL 中有如下的两个表如何重新设计可以将其转换到 KV 的数据库中。

user

|id|username|password|card_id|email|
|:---:|:---:|:---:|:---:|:---:|
|1|tom|hashpwd1|1|tom@test.com|
|2|lice|hashpwd2|2|tom@test.com|

card

|id|name|card_id|valid_date|value|
|:---:|:---:|:---:|:---:|:---:|
|1|card1|id1|timestamp1|500|
|2|card2|id2|timestamp1|3000|

