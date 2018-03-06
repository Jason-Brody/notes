# Groovy 学习资料

### 安装
+ HomeBrew 安装
    > brew install groovy

### 数组
``` groovy
int[] array = [1,2,3]
```

### 闭包
``` groovy
import org.testng.Assert
import org.testng.annotations.Test

class ClosureTest {


    @Test
    void withOutParameter(){
        def print = { print 'HelloWorld'}
        print.call()
    }

    @Test
    void withParameter(){
        def print = { msg -> print msg}
        print.call("HOHOHO")
    }

    @Test
    void withReturnValue(){
        def print = {
            String msg ->
                if (msg.length() > 10){
                    return "A"
                }else{
                    return "b"
                }
        }
        def result = print("1l2k3jkl2j3lk322")
        Assert.assertEquals(result,"A")
    }
}
```

### DSL
+ 参考资料
    1. https://www.cnblogs.com/xingzc/p/6347673.html
    2. http://blog.csdn.net/hivon/article/details/2335248
    3. https://www.w3cschool.cn/groovy/groovy_meta_object_programming.html

### <<操作符
+ 参考资料
    1. http://blog.csdn.net/hivon/article/details/2341312
