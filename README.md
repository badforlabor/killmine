# killmine
额，linux干掉挖矿进程的golang小程序


一个是模拟挖矿的工具：runcpu
另一个是杀进程的工具：top_cpu

很简单，就是每秒检查一下cpu占用率最高的进程，如果频繁出现，那么就杀掉该进程

