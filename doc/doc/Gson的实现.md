第一版
NewDecodeF(file string) *Decoder
(dec *Decoder) RegisterHandlers(handler KeyFoundHandler, KeyDoneHandler, KeyDuplicateHandler) 
(dec *Decoder) Unmarshal() error

Scanner 负责扫描整个二进制流或者文档
  


Decode 负责将配置文件转为想要转的数据类型

type KeyFoundHandler func(string, any) bool 
type KeyDoneHandler func(string, any) bool
type KeyDuplicateHandler func(string, offset, line, lineoffset) bool


Valid(data []byte) bool
Unmarshal(data []byte, v any) error

NewDecoderR(r io.Reader) *Decoder
NewDecoderF(file string) *Decoder
NewDecoderB(data []byte) *Decoder
NewDecoderS(str string) *Decoder
(dec *Decoder) RegisterHandlers(handler KeyFoundHandler, KeyDoneHandler, KeyDuplicateHandler) 
(dec *Decoder) Decode(v any) error
(dec *Decoder) Unmarshal() error


Gson 
map[string] string VarMap
map[string] Gson
[]Gson
parent Gson
origin Gson
value interface{}






Encode 负责将相关的数据类型转为gson的二进制文件

NewEncoder


扫描键的时候，如果发现有!，那么就要记录下这个键
除非遇到转义，否则不需要将字符串特地移动到另一个字符串去，遇到转义才需要转到临时变量去


所有基础的数据类型全部都转为自定义的数据类型，并且像easyjson一样，为所有的自定义基础数据类型实现一个opt，为所有基础数据类型实现一个转为Gson的接口
如果需要转为基础数据类型时，则将其通过自转为自定义的数据类型来实现转换


gson -all path   //将某个路径下的所有添加了//:gson注释的struct或者其余类型转为implement了解析接口的数据

gengson  转换的入口
gen 解析数据转换接口
   generator.go
decode   解析接口
decodetool   转换的通用解析，给gen和decode共用
   实现decoder.go具体功能 
opt   实现的通用的数据类型转换
gson   
   helpers.go 
	    Unmarshaler UnmarshalGSON   接口，将gson字符进行解析
			Optional IsDefined   接口，是否定义
			         IsDone      接口，是否完成
			UnknownsUnmarshaler UnmarshalUnknown   接口，处理未知得键
			isNilInterface   函数，是否未空值
			Unmarshal     函数，转换byte为v
			UnmarshalFromReader   函数，转换io.Reader为v
	 gson.go
	    struct varmap存放变量，只有初始值有；obj存放被标记了的对象，只有初始值有，它直接指向对应的子gson的指针；valuemap，存放了所有的值，可以存放任何类型的值，包括gson、varstring、objstring；prefather，前置父gson的标签；afterfather，后置父gson的标签；originGson，原始的gson，直接指向指针
	    Gson2Json  函数，将gson转为json字符串
			实现Unmarshaler和Optional、UnknowsUnmarshaler接口
			MustString
			MustInt
			Must***
			SetVar
			GetVar

	 varstring.go
	    struct varKey占位符；value值字符；origin原始的gson（gson指针）
	    MustString
	 objstring.go
	    struct key替换标签；value值字符；origin原始的gson（gson指针）
	 	  MustString 

jwriter.Writer 
    实现一个转换为具体数据的接口




scanner  扫描数据文件的实现
    scanner.go 实现了扫描文件的功能
buffer   缓存数据结构


gson的实现需要将字节从原始的位置移动到临时位置，这个移动仅需一次，使用lexer.go里面相关的内容进行存储临时的字节，然后将数据保存到一个类似buffer的东西，把这个buffer作为参数传递给每个变量的回调，每个变量还需要添加一个回调接口，这个接口有一个键、下标、值（即刚刚说到的buffer），然后根据键、下标做出不同的处理，以便将值赋值给目标变量。另外如果没有对应的回调接口，则通过反射找到对应的下标、键的值来赋值。

遇到一个空白字符，结尾
遇到一个{、}结尾
遇到一个!结尾
遇到一个@结尾
遇到一个$结尾
遇到一个:结尾
遇到一个=结尾
遇到与之匹配的``结尾
遇到一个#结尾
注释中遇到换行结尾

遇到一个#按照注释处理
遇到一个：按照键处理
遇到一个=按照变量处理

!前面必须有键，所以遇到!直接按照值来处理
@前面必须的空格之前的数据都按照值来处理，$也是
遇到一个!、@、$暂时不处理，等待后面的值取得，!等待一个:来处理，$和@等待一个空白符来处理
遇到：先确定是否在标签中，然后再遇到：的处理，再根据是否在标签中决定分割处理
需要一个标志是否在标签中

遇到一个{和}直接按照值处理
遇到一个@按照替换符处理（前面可以无需有键），暂不处理，等待后面的值取得，




