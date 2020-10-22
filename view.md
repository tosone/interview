### 面试

1. Go 中的 interface struct 组合和 C++ 中的虚函数，纯虚函数，友元函数，继承，多态，重载，模板有哪些相同或者对应的东西，可以讲一下，越多越好。
2. C 中指针和 Golang 中指针有什么相同和不同的地方。了解过 CGO 么？他们是如何管理内存的？
3. 内存对齐是什么？

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

4. Golang 内存管理的基本单位 8K。
5. Slice 内存多大？
6. 大端编码和小端编码是什么？
7. 有两个超大的文件，里边都是电话号码，如何设计算法得到他们之间的差集？简述即可。
8. err check
9. Go 的优势，有什么特性是你特别喜欢的？这些特性适合做什么？
10. context 注意事项
11. Docker 网络模式以及他们之间的区别
12. 一个简单的基于容器的 DevOps 系统需要包含什么样的功能？
13. 你何时会使用像 ConfigMap 或 secret 这样的东西？他们的区别是什么？你都如何使用他们？
14. 容器内暴露服务的方式以及访问方式？
15. 简述 sidecar 容器，init 容器的区别？Pod 内容器启动顺序？
16. 在构建和管理生产集群时遇到的主要问题是什么？
17. 全覆盖问题，如何判定给定的一个整数数组全覆盖到一个连续整数区间内的所有数字，要求内存最优时，如何设计算法？利用全覆盖的算法如何设计求解数独的算法，要求内存最优时，如何设计算法？
18. 这块有遇见什么比较好用的工具么？他给你解决了什么问题？
19. 私下自己有没有什么有趣的东西或者项目？
20. namespace, systemd, runC, containerd, docker 这些都是什么东西，他们之间什么关系。
