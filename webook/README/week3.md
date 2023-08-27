# 作业：修改已有的部署方案
* 将 webook 的启动端口修改 8081。
* 将 webook 修改为部署 2 个 Pod。
* 将 webook 访问 Redis 的端口修改为 6380。
* 将 webook 访问 MySQL 的端口修改为 3308。

你需要提供：

* kubectl get services 的截图。
* kubectl get pods 的截图。
* 你通过浏览器访问 webook，能够正确得到响应的截图。

---

## 流程：

![1](https://github.com/LiAyasan/eimy-geektime-basic-go-homework/blob/master/webook/image/week3/%E4%BF%AE%E6%94%B9%E5%90%AF%E5%8A%A8%E7%AB%AF%E5%8F%A3.png?raw=true)

![2](https://github.com/LiAyasan/eimy-geektime-basic-go-homework/blob/master/webook/image/week3/%E4%BF%AE%E6%94%B9%E4%B8%BA2%E4%B8%AApod.png?raw=true)

![3](https://github.com/LiAyasan/eimy-geektime-basic-go-homework/blob/master/webook/image/week3/%E4%BF%AE%E6%94%B9redis%E6%8E%A5%E5%85%A5%E7%AB%AF%E5%8F%A3.png?raw=true)

![4](https://github.com/LiAyasan/eimy-geektime-basic-go-homework/blob/master/webook/image/week3/%E4%BF%AE%E6%94%B9mysql%E6%8E%A5%E5%85%A5%E7%AB%AF%E5%8F%A3.png?raw=true)

![service](https://github.com/LiAyasan/eimy-geektime-basic-go-homework/blob/master/webook/image/week3/get%20service.png?raw=true)

![pod](https://github.com/LiAyasan/eimy-geektime-basic-go-homework/blob/master/webook/image/week3/get%20deployment.png?raw=true)

最后测试项目能否正常运行

![postman](https://github.com/LiAyasan/eimy-geektime-basic-go-homework/blob/master/webook/image/week3/postman.png?raw=true)

## 心得

学会如何把项目部署到k8s，并使其与mysql、redis联通。