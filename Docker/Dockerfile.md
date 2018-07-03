### FROM

> FROM \<image> [AS \<name>]

> FROM \<image>[:\<tag>] [AS \<name>]

> FROM \<image>[@\<digest>] [AS \<name>]

### ENV (环境变量)
+ 常规用法  
    ```
    FROM busybox           # 父镜像
    ENV foo /bar           # 定义环境变量
    WORKDIR ${foo}         # WORKDIR /bar
    ADD . $foo             # ADD . /bar
    COPY \$foo /quux       # COPY $foo /quux
    ```      

+ 假如设定了环境变量
    ```
    ENV foo /bar           # 定义环境变量
    WORKDIR ${foo:-abc}    # WORKDIR /bar
    WORKDIR ${foo:+abc}    # WORKDIR abc
    ```

+ 假如没有设置环境变量 
    ```
    WORKDIR ${foo:-abc}    # WORKDIR abc
    WORKDIR ${foo:+abc}    # WORKDIR 
    ```
+ 变量赋值
    ```
    ENV abc=hello          # abc = hello
    ENV abc=bye def=$abc   # abc=bye def=hello
    ENV ghi=$abc           # ghi=bye
    ```

+ 可以使用环境变量的指令列表
    - ADD
    - COPY
    - ENV
    - EXPOSE
    - FROM
    - LABEL
    - STOPSIGNAL
    - USER
    - VOLUME
    - WORKDIR

### .dockerignore file
```
# comment
*/temp*
*/*/temp*
temp?
```

| 规则 | 行为 |
|-----|------|
| # comment | 忽略 |
| \*/temp\* | 排除所有root目录下以temp为开头的子目录或子文件，比如 /somedir/temporary.txt 或者/somedir/temp
| \*/\*/temp\* | 同上,比如/somedir/subdir/temporary.txt
| temp? | ?匹配任意单个字符，即排除root目录下以temp为开头，后面跟一个字符的文件或目录，比如 /tempa 或者 /tempb
| **/*.go | 排除所有以.go为结尾的文件，无论在哪个位置 |
| *.md !README.md | 排除root下所有.md文件除了README.md |

### ARG
```
ARG  CODE_VERSION=latest
FROM base:${CODE_VERSION}
CMD  /code/run-app

FROM extras:${CODE_VERSION}
CMD  /code/run-extras
```

```
ARG VERSION=latest
FROM busybox:$VERSION
ARG VERSION
RUN echo $VERSION > image_version
```

### RUN

13564960809 林
13962887855 孙