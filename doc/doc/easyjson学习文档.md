easyjson学习文档
===

- 2022.11.19

---
## 核心代码
* benchmark 测试代码
* bootstrap  ??
* buffer 内存池
* easyjson easyjson的工具代码，把对应的结构体自动实现implement相关接口的夫夫
* gen 核心代码
   * decoder.go 解析代码
	 * encoder.go 反解析代码
	 * generator.go 这是把对应的struct转为实现了接口的方法

* opt 与转换为何种数据有关？？？
* parser 转换数据？？？
这2个是easyjson的直接包
* helpers.go  ????
   * 这个文件定义了很多json解析的接口定义，以及直接可以调用入口，可以转换为各种想要的格式
   * Marshal(v Marshaler)其中的v是一个Marshaler类型，每个要转换为json字符串或者字节的结构体都要实现Marshaler的接口，实现MarshalEasyJSON
	 同理Unmarshaler也是如此
* raw.go 
   * 这个文件它定义了将一个片段的byte（可能是字符串、数组、bool、对象等）和json数据类型的互相转换，主要是转换为了jwriter.Writer、jlexer.Lexer
* unknown_fields.go 
   * 这个文件定义了对未知的域的处理接口
* jlexer和jwriter是主要的东西，很多数据都是转换为这两个东西
* jlexer定义了两个bytesToStr，一个unsafe
   * bytestostr定义了两个bytesToStr，一个unsafe
   * error.go定义了错误结构体
	 * lexer.go定义了解析的具体逻辑，从数据中一个一个扫描字符来处理的操作
* jwriter是将具体类型写为字符和二进制流的操作
* opt实现了每种数据类型实现的json转换的接口
* buffer定义了一个内存池空间

easyjson是定义了一组接口，在运行时通过判断对应的结构体是否实现了这些接口，如果实现就会调用这些接口来转换，如果没有实现就会报错，因此在运行时是直接调用对应的接口实现的，而不是通过反射实现的。
而实现这些接口并不是手动写的，easyjson写了一个自动转换的脚本，运行这个工具就可以将其自动写好了。这个脚本的转换过程也是通过反射来实现的，但是它和go自带的json工具的不同在于，easyjson是在运行前一次转换，而go则是在运行时转换。

#### intern
* intern.go 只与内存池有关
   String(s) 将一个字符数组转成一个字符串
	 Bytes(s)将一个byte数组转为string
	 sync.Pool是一个内存池，这里面的东西随时可能会被回收，如果出现从map[string]string中拿到的里面就已经有对应的字符串的情况，那么大概率是因为刚刚把这个map[string]string对应的内存释放，里面的内容还未释放，因此再次获取到的还是原来那个，所以会先看看拿到的map有没有想要的string

---
## easyjson的核心要点
* 运行时核心仅easyjson和jlexer和jwriter几个
* opt是所有数据类型的一个转换
* gen和parser是运行前将一个数据类型解析后得到的implement接口的实现过程，很重要，但是不是在运行时使用的，而是运行之前
* 所有数据类型都要实现helpers的几个接口，才可以使用
* 要添加注释//**.json才会转换为相关代码
* easyjson是转换的入口代码，bootstrap、gen和parser是转换的过程
* benchmark和tests是测试代码
* intern仅仅实现了一个字符数组转字符串和字节数组转字符串的接口，它使用内存池，减少gc开销


---
## go自带的json包的核心要点
* scanner.go实现扫描的功能
* stream.go实现从io.Reader或者io.Writer读取数据或者写入数据的功能
* tags.go实现了获取标签的作用
* encode.go和decode.go实现了解析到具体数据类型，或者从数据类型解析到json的功能，它是通过反射的方式，而不是像easyjson一样为每种数据类型实现一个解析接口的方式

---
## jlexer
bytestostr.go和bytestostr.go实现了2个接口bytesToStr，将byte数组转为字符串
error.go定义了一种错误
lexer.go基本都是在读取和解析json的一个过程实现

---
## 常量的定义可以使用byte重新定义，参考lexer.go的tokenKind

