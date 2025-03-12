# GBM
<p style="text-align: center">一个简单的GB28181流媒体模拟工具</p>
<p style="text-align: center"><img src="https://camo.githubusercontent.com/968fd36a2415616f033b8588ef156e8f23d6e4c1e864f7908e252ad73e10c23d/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d4d49542d677265656e3f7374796c653d666c61742d737175617265"  alt="license"/></p>

## 简介
本项目使用GO语言编写，采用宽松的**MIT**协议，用于模拟`GB/T28181`摄像头推流，数据包来自于Wireshark抓包的海康摄像头。

## 使用
```shell
# Windows 推流命令
.\GBM.exe -host 192.168.1.99 -port 5060
# MacOS/Linux
./GBM -host 192.168.1.99 -port 5060
```
