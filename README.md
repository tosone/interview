# 面试

## 开放问题

- top K 问题。
- Golang 有什么优势和缺点？这些特性适合做什么？
- 有一个超大的文件，里边全部是电话号码，并且号码没有重复，如何设计算法快速找出其中一个？简述即可。
- 软件往往有很多依赖。Go 语言的处理方法是，将所有依赖打包进一个二进制文件，从而与环境无关。Docker 的方法是提供一个容器，在里面还原所有的依赖。如何看待这个问题？

## 数据结构

- HashMap 底层原理是什么？
- 为什么设计的时候需要设置负载因子？初始大小？
- 为什么它是一种不安全的数据结构？怎么实现线程安全的 HashMap？

## 通讯协议

- UDP 和 TCP 之间的区别是什么？
- TCP 粘包的主要原因是什么？UDP是否会出现？如何解决粘包问题？[参考](https://draveness.me/whys-the-design-tcp-message-frame/)
- UDP 变 TCP 一样的可靠通信需要做什么？
- 请描述一种你知道的共识算法，以及其中的原理？
- TCP 的多路复用是如何实现的？
- 我们再集群内创建负载均衡的时候，需要设置一个 VIP，你知道这个 VIP 到底是什么东西么？这个 VIP 有什么限制？我们可以通过哪些协议探测 VIP？

## 原理

- 负载均衡的原理是什么？常见的负载均衡的算法都有哪些？
  > HTTP 重定向，DNS 负载均衡，反向代理负载均衡，IP负载均衡(LVS-NAT)，直接路由(LVS-DR)

## 操作系统

- 为什么 Linux 需要虚拟内存？[参考](https://draveness.me/whys-the-design-os-virtual-memory/)
- 进程和线程以及它们的区别？僵尸进程是如何产生的？
- 进程间通讯都有哪些方式？mmap pipe fifo 都分别是什么？

## 数据库

- 数据库的乐观锁和悲观锁是什么？有什么区别？使用场景？
- 索引是什么？有什么作用以及优缺点？使用索引查询一定能提高查询的性能吗？

## 场景设计

- 如果基于 Kubernetes 实现 serverless 需要做什么？快速启动？调度？超卖？
- 一个任务有多个计算单元，每个计算单元相互之间有依赖关系，这样的系统应该如何设计？计算单元的重试应该如何设计？
- 一个耗时任务，任务本身的状态会有等待执行，执行中，执行完成，执行失败等等状态，用户在前端创建完了这个任务之后，后端后续的逻辑应该是什么样的？整个服务的架构应该是什么样的？如果你的服务要做高可用，应该做什么样的处理？

## Golang

- init 和 defer 的执行顺序是什么？
- Golang 中如何会导致内存泄露以及如何排查？
- Golang 中线程安全的数据结构都有哪些？
- Goroutine 都说他节省内存，是调度的最小单元，那么他相比于 pthread 是为什么可以节省内存，他们分别占用多少内存？ 2K vs 16K
- C 中指针和 Golang 中指针有什么相同和不同的地方。了解过 CGO 么？他们是如何管理内存的？
- 内存对齐是什么？
- Golang 中的 string 数据结构是什么样的？获取长度的时候时间复杂度是什么样的？与 Redis 中字符串的相比有什么异同？
- slice 的原理，请简述一下？

## C++

- C++ 中的字符串和 C 中的字符串有什么区别？
  - C 中的字符串操作的时候复杂度是什么样的？
  - C++ 字符串容量自动扩容是怎么做的有什么好处？
  - Redis 的 sds 有了解过么？

## 容器

- 从认证到拉取镜像整个的流程是什么样的？
- 外部的请求是是如何一步步落到 Pod 上的？
- 简述 init 容器是做什么用的？如何保证 Pod 内容器启动顺序？
- CGroups，namespace 在 Docker 中扮演什么角色，都有什么作用？
- 有状态服务和无状态服务之间的区别？
- 应用访问 service 发现网络不通的时候该如何排查？
- K8s 集群中有几种 IP，node ip、service ip、pod ip、cluster ip 分别都是由谁来分配的？他们的作用都是什么？
- K8s 如何实现 request 和 limit 的？
- k8s 的 pause 容器有什么用，是否可以去掉？
- pod 内的容器如何进行通信？
- 请你说一下kubernetes针对pod资源对象的健康监测机制？几种探针的区别？

## 工具

- Git 中 rebase 和 merge 有什么区别？

## 协议

- restful 接口定义都包含哪些内容？
- Websocket 握手细节？
- OAuth2，SSO 都是什么？
- TLS 握手细节？

## 算法

- 全覆盖问题，如何判定给定的一个整数数组全覆盖到一个连续整数区间内的所有数字，要求内存最优时，如何设计算法？利用全覆盖的算法如何设计求解数独的算法，要求内存最优时，如何设计算法？
- 买卖股票的最佳时机 - [121](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)
- 岛屿数量 - [200](https://leetcode-cn.com/problems/number-of-islands/)
