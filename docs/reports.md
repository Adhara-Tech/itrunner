# Console report
+-----------------+---------+---------+
|      GROUP      | VERSION | RESULT  |
+-----------------+---------+---------+
| postgres        |    10.9 | Failure |
+-----------------+---------+---------+
| Postgresql demo |    10.9 | Failure |
+                 +---------+---------+
|                 |   10.10 | Failure |
+                 +---------+---------+
|                 |   10.11 | Failure |
+-----------------+---------+---------+
| Rabbit demo     |     1.4 | Failure |
+                 +---------+---------+
|                 |     2.4 | Failure |
+                 +---------+---------+
|                 |     3.7 | Failure |
+-----------------+---------+---------+

# Json report format (proposal) 
```
{

    "suite":{
        "group":[
            {
                "name":"",
                "versions":[
                    {
                        "version":"",
                        "result":"",
                    },
                    {
                        "version":"",
                        "result":"",
                    }
                ]
            }
        ]
    }
}
```