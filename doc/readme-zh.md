```
       ╭╮
╭─┬┬┬──┤└╮
╰┴┴─┴┴┴┴─╯
```
***
麻木的拇指是处理生活中常见的数字运算，单位计算与日期计算的工具
***
#### 举例如下
```

    $ numb
    Enter `q` to quit
    > 3735928559 >> 16 in hex
      0xdead
    > 1920/1080 in rat
      16/9
    > sin(pi/3)
      0.87
    > 3 millilightsecond to mile
      558.85 mile
    > 10 mile/hour to metre/sec
      4.47 metre/sec
    > total_grains = 2^64 - 1
      18446744073709551615
    > 3 GB / (15 Mbps) to minute
      26.67 minute
    > time in London
      13:04
    > today + 10 days
      05 Jun 2022
    > today - {1999-09-09} to months
      272.54 month
```
#### 安装
```
# build with Go >= 1.16
$ go install github.com/nkanaev/numb@latest
```
#### 更多帮助
```
   doc/help.txt     手册, 示例
    doc/spec.txt     命令行语法
    doc/dump.txt     各种设计笔记
    doc/todo.txt     待办事项
```



