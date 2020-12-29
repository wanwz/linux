## Go语言特征
### 声明
1. 可见性
    - 声明在函数内部，是函数的本地值，只能在函数内部使用
    - 声明在函数外部，是对当前包可见的全局值，只要是.go文件用了这个包，就都可以使用，待验证2020.12.29
    - 声明在函数外部且首字母大写是所有包可见的全局值
2. 声明方式
    - var 声明变量
    - const 声明常量
    - type 声明类型
    - func 声明函数
3. 声明顺序
    1. 第一行 package xxx，声明该文件属于哪个包
    2. import xxx
    3. 类型 常量 变量 函数的声明
### 项目构建及编译
1. 一个Go工程主要包含三个目录：src(源代码)、pkg(包文件)、bin(相关bin文件)
2. 编译：
    - go build 用于编译包生成可执行文件放在当前目录下，必须有main包才可以
    - go install 主要用来生成库和工具，如果有main包编译后生成的可执行工具保存在bin目录下，编译后的库文件保存在pkg目录下（项目比较大建议使用）
    - go run xx.go 直接编译执行.go文件
### 内置类型和函数
1. 值类型
    1. 布尔 bool
    2. 整数 int(32 or 64) int8 int16 int32 int64
    3. 浮点 float32 float64
    4. 字符串 string
    5. 复数 complex64 complex128
    6. 数组 array
2. 引用类型
    1. slice 序列数组
    2. map 映射
    3. chan 管道
### 下划线
'_'是特殊标识符，用来忽略结果
1. 在import中
    1. import _ 包路径 只调用包中init函数
2. 在代码中
    1. 忽略这个变量，例
        ```go
        os.Open返回值为*os.File,error
        普通写法f, err := os.Open("xxx")
        如不想要返回的错误值，则
        f, _ := os.Open("xxx")
        ```
### 变量
1. 原理
    1. 程序运行过程中的数据是保存在内存中的
    2. 操作数据时，首先要去内存中找到这个变量（内存地址）
    3. 如果直接在代码中通过内存地址来操作变量时，代码可读性变差
    4. 因此需要一个变量将这个数据的内存地址保存起来，通过这个变量来找内存上对应的数据
2. 声明
    1. 同一个作用域内不支持重新声明
    2. 变量声明后必须使用，不使用报错
    3. 格式
        - var 变量名 变量类型 var name string|var age int |var isOk bool
        - 批量声明
            ```go
            var (
                a string
                b int
                c bool
                d float32
            )
            ```
    4. 初始化：声明变量时，会自动对变量对应的内存区域进行初始化操作，每个变量会被初始化成其类型的默认值
        - 整数和浮点型 0
        - 字符串 空字符串
        - 布尔 false
        - 切片、函数、指针 nil
        - 也可在初始化时指定初始值
            1. 格式 var 变量名 类型 = 表达式 var name string = "zhangsan"|var name, age = "zhangsan", 100|var name = "zhangsan"
    5. 短变量
        1. 只能在函数内部使用':='声明并初始化变量
            ```go
            package main
            import "fmt"
            //全局变量
            var name1 = "zhangsan"
            func main() {
                name2 := "lisi"
                //局部变量
                name1 := "wangwu"
                fmt.Println(name1, name2)
            }
            //输出 wangwu lisi
            ```
### 常量
1. 常量是恒定不变的值
2. 在定义时必须赋值
    ```go
    const pi = 3.1415
    const e = 2.7182
    const (
        pi = 3.1415
        e = 2.7182
        a // a = 2.7182
        b // b = 2.7182
        c = 10
        d // d = 10
    )
    ```
3. iota
    ```go
    const (
        n1 = iota //n1 =0
        n2 //n1 =1
        n3 //n1 =2
        _
        n4 //n4 =4
        n5 = 10
        n6 = iota //n6 = 5
        n7 //n7 = 6
    )
    const (
        a, b = iota + 1, iota + 2 //1,2
        c, d //2,3
    )
    ```
