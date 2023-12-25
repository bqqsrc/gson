# Gson

Go Script Object Notation

Copyright (C) 晓白齐齐,版权所有.

---
## gson的语法
### 构造字符
gson保留以下6个字符作为构造字符，以下是gson的构造字符一览表：

|构造字符|名称|作用|使用范例|
|---|---|---|---|
|:|键定义构造字符|定义一个键|```root: /User/testpath/test.gs```|
|{|对象起始构造字符|标志一个对象的起始构造字符|```Obj: { ele1: 33.5 }```|
|}|对象结尾构造字符|标志一个对象的结尾构造字符|同上|
|\`|强制字符串构造字符|标志一个字符串的起始字符|numstr: \`35.898\`|
|%|转义字符|对构造字符进行转义|mapstr: %{aa%:3,bb%:4%}|
|#|注释符|添加一段注释|```prefix: /Users/goper/conf/goper.gs #goper的配置文件路径```|


更多详解，可查看[这篇文档](https://github.com/bqqsrc/documents/blob/main/2.gson/1.Gson%E4%BD%BF%E7%94%A8%E8%A7%84%E8%8C%83.md)

---