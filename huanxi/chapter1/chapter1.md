# Chapter1

## 变量

### 变量的类型

1. 简单数据类型 (值类型)

- 整型

  有符号: int, int8, int16, int32, int64

  无符号: uint, uint8, uint16, uint32, uint64, byte

- 浮点型

  float32, float64

- 字符型

  byte rune

- 字符串

  string

- 布尔型

  bool

- 复数

  complex64 complex128

2. 复杂类型

- 值类型

  1. struct

  2. array

- 引用类型

  1. slice

  2. map

  3. channel

### 常量

    const CONSTANT = VALUE

    const(
        CONSTANT1 = VALUE1
        CONSTANT2 = VALUE2
    )

### 枚举

```go
const(
    MONDAY      = iota + 1
    TUESDAY
    WEDNESDAY
    THURSDAY
    FRIDAY
    SATURDAY
    SUNDAY
)
```

### 变量的声明

1.  长声明
    `var variable TYPE`
    初始化
    `variable = value`

2.  长声明 + 初始化
    `var variable = value`

3.  短声明 + 初始化
    `variable := value`

4.  多个变量声明...

```go
var(
variable1 = value1
variable2 = value2
)
```

5.  多个常量声明...

```go
const(
VARIABLE1 = value1
VARIABLE2 = value2
)
```

6.  对于引用类型变量 map, slice, chan

```fo
m := make(map([string]sting), [cap])
var s = make([]sting, [len, cap])
var c chan int= make(chan int, [cap])
```

7.  new

## 流程语句

### if

- if

  ```go
  if mark >= 90 {
      fmt.Println("Excellent")
  } else if mark > 60 && mark < 90 {
      fmt.Println("Good")
  } else {
      fmt.Println("failed")
  }
  ```

### switch

- switch

  ```go
  swich {
  case mark >= 90:
      fmt.Println("Excellent")
  case mark >= 60 && mark < 90:
      fmt.Println("Good")
  case mark < 60:
      fmt.Println("failed")
  }
  ```

### for

- for

  ```go
  for {
      // do something
  }

  /* equivalent while true
  while (true) {
      // do something
  }
  */

  for i := 0; i < boundary; i++ {
      // do something
  }
  ```

## 函数

### 函数定义

```go
  func funcName(parametersList) (returnList) {
      // function body
  }
```

### 函数的注意点

1. 传值, 传引用问题

   1. 取决于实参类型, 实参为基本类型, struct 等时, 传值. 实参为引用类型, 如 map, slice, chan 等, 传引用

   2. 取决于 parametersList 的要求, go 语言会自动解引用与引用

2. 可以返回多的值

```go
  func MutiReturns(a int, b int) (sum int, diff int) {
      return a + b, a - b
  }
```

3. 变长函数, 可接受不定数量的参数

```go
  func MutiSummation(nums ...int) (sum int) {
      for _, v := range nums {
          sum += v
      }
      return
  }

  func main() {
      sum := MutiSummation(1, 2, 3)
      fmt.Println(sum)
      // output: 6
  }
```

4. 可存入变量中, 函数的本质是一个地址

```go
   f := MutiReturns
   fmt.Println(f(1, 2))
   // output: 3 -1
```

### 匿名函数

- 声明

  ```go
  func (parametersList) (returnList) {
      // function body
  }
  ```

- 调用

  ```go
  func (parametersList) (returnList) {
      // function body
  }()
  ```

- 闭包

  一个函数的返回值是一个函数, 返回的函数使用了非自己函数块定义的变量.

  ```go
  func makeSuffix(suffix string) func(str string) string {
      return func(str string) string {
          if !strings.HasSuffix(str, suffix) {
              return str + suffix
          }
          return str
      }
  }

  func main() {
      checkSuffix := makeSuffix(".txt")
      file := checkSuffix("rey")
      fmt.Println(file)
      // output: rey.txt
  }
  ```

## 指针

取引用(取地址): `&variable`
解引用(取值): `*variable`

## 数组

### 初始化数组

1.  指定大小+赋值 `variable := [3]int{1, 2, 3}`

2.  指定大小+不赋值 `variable := [3]int{}`

3.  推断大小 `variable := [...]int{1, 2, 3}`

### 遍历数组

1.  for

    ```go
    for i := 0; i < len(array); i++ {
      // do something
    }
    ```

2.  for range

    ```go
    for idx, value := range array {
      // do something
    }
    ```

## 二维数组

### 初始化二维数组

```go
a := [3][4]int{
 {0, 1, 2, 3} ,
 {4, 5, 6, 7} ,
 {8, 9, 10, 11},
}
```

### 遍历二维数组

```go
 for i := 0; i < len(array); i++ {
     for j := 0; j < len(array[i]); j++ {
         // do something
     }
 }
```

## 切片

### 初始化切片

1.  变量声明 `var sclice []int`

2.  字面量` var slice = []int{1, 2, 3}`

3.  从数组初始化 `slice := array[start:end[:max]]` cap = max - start

4.  make `slice := make([]int, [len, cap])`

### 遍历切片

1.  for

    ```go
    for i := 0; i < len(slice); i++ {
      // do something
    }
    ```

2.  for range

    ```go
    for idx, value := range slice {
      // do something
    }
    ```

### 如何复制切片

`copy(destnation, source)`

### 切片的增删改查

- 增 `slice := append(slice, interface...)`

- 删 `slice := append(start:targetIdx, targetIdx+1:end...)`

- 改 `slice[targetIdx] = value`

- 查 `slice[targetIdx]`

### 切片组成的二维数组

需要为每一个一维数组(切片)分配容量(make)

## map

### map 的特点

- 无序

- 不可比较

- k-v 对 不可寻址

- 需要使用 make 初始化

### map 的增删改查

- 增 `m[key] = value`

- 删 `delete(m, key)`

- 改 `m[key] = newValue`

- 查

  ```go
  if _, ok := m[key]; !ok {
      // not founded
      return
  } // founded
  ```

### map 的 key 的要求

可比较的类型可以做 map 的 key eg. 基本值类型, struct, array, channel etc.

## 结构体

### 结构体的定义

```go
type SturctName sturct {
    // body
}
```

结构体的继承 可以用 结构体的嵌套来实现 (ducky type)

### 方法的定义

```go
 func (receiver) funcName(parametersList) (returnList) {
     // function body
 }
```

### 面向对象的好处

高聚合, 低耦合, 高复用, 隐藏实现细节 user-friendly

## Coding By Rey

mail: 3065588496@qq.com
