Goson
===

---
## 保留字符
= 表示常量定义
: 键值对
{ 字典开头
} 字典结尾
$ 常量替换
# 注释
% 转义
` 强制字符串
@ 表示使用键值替换
%{ == {
%} == }
%# == #
%: == :
%= == =
%% == %
%$ == $
%` == `
%@ == @

---
## 样例
aa=333
bb=333
cc=333
dd=333
a1: a1 
a2: a2
a3: a3_1 a3_2 a3_3
a4: a4_1 a4_2 a4_5
a5: {
	a51: {
	  a52: a53
		a53: a54 
	}
	a52: {
		a55: a55 
		a56: a56
	}
}

---
## 设计
type Goson struct {
	data []byte   //记录当前读入内存的数据
	dataBegin int //本次读取数据的起始点
	DataNone int  //上次未读取到数据的起始点
	dataLen int   //内存数据长度
	dataOff int   //记录当前内存数据的off
	scan scanner  //数据浏览器
	conf string   //配置文件路径
	confBegin int     //配置文件本次读取的起始位置
	confOff int       //配置文件的读取off
	confNone int      //配置文件上次未读取到数据的起始点
	maxLen int        //一次从文件读取多少字符
	done bool      //是否读取完成
	line int       //读取到第几行
	column int     //读取到第几个字符串
	errorContext *errorContext
	useNumber bool
	disallowUnkonwnFields bool
}

func GetGosonByFile(file string)
func GetGosonByByte(bytes []byte)

//将data转为any，如果没有指定类型，将转为map[string]interface{}
func Unmarshal(data []byte, v any) error {
}

//将某个键值转为any
func UnmarshalKey(key string, v any) error {

}

func GetBool(key string) 
func GetInt(key string)
func GetBool(key string) 
func GetInt(key string)
func GetBool(key string) 
func GetInt(key string)
func GetBool(key string) 
func GetInt(key string)
func GetBool(key string) 
func GetInt(key string)
func GetBool(key string) 
func GetInt(key string)
func GetBool(key string) 
func GetInt(key string)

func Close() {

}
