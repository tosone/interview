# 架构师面试

## 开放问题

- 为什么需要微服务？解决了什么问题？

## 设计

- HashMap 底层原理，为什么它是一种不安全的数据结构？怎么实现线程安全的 HashMap？
- 如何设计用户注册系统的时候发送给手机验证码的整个流程？

## 协议

- TCP 粘包的主要原因是什么？UDP是否会出现？如何解决粘包问题？[参考](https://draveness.me/whys-the-design-tcp-message-frame/)
- TCP 中为什么会有 TIME_WAIT 这种状态？用来做什么？
- 请描述一种你知道的共识算法，以及其中的原理？
- 如何实现 IO 的多路复用？

## 原理

- 负载均衡的原理是什么？常见的负载均衡的算法都有哪些？
  > HTTP 重定向，DNS 负载均衡，反向代理负载均衡，IP负载均衡(LVS-NAT)，直接路由(LVS-DR)
- CDN 为什么这么快，原理是什么？

## 操作系统

- 为什么 Linux 需要虚拟内存？[参考](https://draveness.me/whys-the-design-os-virtual-memory/)
- 进程和线程以及它们的区别？僵尸进程是如何产生的？
- 进程间通讯都有哪些方式？mmap pipe fifo 都分别是什么？

## 数据库

- 数据库的乐观锁和悲观锁是什么？有什么区别？使用场景？
- 索引是什么？有什么作用以及优缺点？使用索引查询一定能提高查询的性能吗？

## 场景设计

- 我们再集群内创建负载均衡的时候，需要设置一个 VIP，你知道这个 VIP 到底是什么东西么？这个 VIP 有什么限制？我们可以通过哪些协议探测 VIP？
- 一个任务有多个计算单元，每个计算单元相互之间有依赖关系，请设计一种架构使其满足其中一个计算单元挂掉了不影响整体任务，若其中一个计算单元非常耗时，如何处理？

## 算法题目

- 买卖股票的最佳时机。[121](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)
- 缺失的第一个正数。[41](https://leetcode-cn.com/problems/first-missing-positive/)
