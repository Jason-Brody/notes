### Elasticsearch基础

+ 查询
``` json
{
    "query": {
        "match_all": {}
    }
}
```

curl -XPUT 'http://localhost:9200//employees/employee/1' -d '{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}'
+ 索引文档
    1. 索引一条数据
    ``` json
    curl -XPUT 'http://localhost:9200//employees/employee/1' -d '
    {
        "first_name" : "John",
        "last_name" :  "Smith",
        "age" :        25,
        "about" :      "I love to go rock climbing",
        "interests": [ "sports", "music" ]
    }
    '
    ```

    megacorp -- 索引名称  
    employee -- 类型名称  
    1 -- 特定id

+ 分片
    1. 创建Blog的分片
    ``` json
    curl -XPUT 'http://localhost:9200/blogs' -d '
    {
        "settings":{
            "number_of_shards":3,
            "number_of_replicas":1
        }
    }
    '
    ```

+ 分析器
    1. 测试分析器
    ``` json
    curl -XGET 'http://localhost:9200/_analyze' -d '
    {
        "analyzer": "standard",
        "text": "Text to analyze"
    }
    '
    ```

+ Index
    1. analyzed  
    首先分析字符串，然后索引它
    2. not analyzed
    索引这个域，所以可以搜索到它，但索引指定的精确值。不对它进行分析
    3. no
    Don’t index this field at all不索引这个域。这个域不会被搜索到
