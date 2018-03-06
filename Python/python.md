### Python学习总结

#### 注释
``` python
# 行级注释

"""
块级注释
"""

```




### Map
``` python
x = {'test':1,'test2':2}

# 遍历
for item in x
    print(item,x[item])
```

### 函数
1. 定义
``` python
def banner(msg,border);
    line = border * len(msg)
    print(line)
    print(msg)
    print(line)

# 函数默认值
def banner(msg,border='-')
    line = border * len(msg)
    print(line)
    print(msg)
    print(line)
```

2. 函数调用
``` python
banner('abc') #border默认值为-
banner('abc','*')
banner(border='*',msg='abc')
```

### 集合类型
1. str
``` python
# 字符串连接
','.join(["Hello","World"]) # Hello,World
# 字符串拆分
hello,split,world = "Hello,World".partition(',')
# 格式化
"My name is {0} with {1} years old".format('test',32)
"Current position {latitude} {longtitude}".format(longtitude="5E",latitude="60N")

pos = (1,2,3)
"Position x={p[0]} y={p[1]} z={p[2]}".format(p=pos)
```
2. list
``` python
x = [1,2,3,4]
x = list("abcdefg")

# 子串
x[2:] # [3,4]
x[:2] # [1,2]

# 遍历
for item in x
    print(item)

```
3. dict
4. tuple
``` python
# 初始化
x = ("ABC",1,3.454,True)

# 遍历
for item in x:
    print(item)
```
5. range
6. set