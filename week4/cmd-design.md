[TOC]
# Agenda设计文档

## 文件内容说明
### user.txt
内部以JSON的格式存放着User结构的数据。
详情见`entity/useroper.go`

### cache.txt
存放着登录信息
未登录：`logout`
登录：`username(登录用户的用户名)`

### meeing.txt
内部以JSON的格式存放着Meeting结构的数据
详情见`entity/conference.go`

## 需求  



### 子命令 

- 大的方向上要实现`help、register、cm`三个子命令
- 用户信息存储在`curUser.txt`中，`User`和`Meeting`实体要用`json`进行存储
- 实现日志服务

---

 ### agenda help

寻求帮助使用子命令`agenda help`，列出命令说明。此外输入`agenda help register可以`列出regsiter命令的描述

---

### agenda user register

用户注册使用命令`agenda user register`

改指令可接受四个参数（必选）：

- --username/-u，用户名
- --password/-p，密码
- --email/-e，邮箱
- --telphone/-t，电话

**注**：

- 此外用户名是唯一的，要查重。
- 邮箱、电话要检查有效性。

**反馈**：

- 如果登记成功，返回成功注册的信息
- 如果登记失败，反馈错误信息

---

### agenda user login

用户登录使用子命令`agenda user login`

- --username/-u，用户名
- --password/-p，密码

**注**：

- **用户名**与**密码**这两个参数为必须

**反馈**：

- 用户名与密码都正确的时候返回一个成功登录的信息
- 登录失败，返回登录失败的信息

---

### agenda user logout

用户**登出**使用子命令`agenda user logout`

**注**：

用户登出后只能够使用用户注册和用户登录功能且这里命令**不接受参数**

**反馈**：

- 反馈登出信息

---

### agenda user lookup

在创建的用户中查找某用户使用子命令`agenda user lookup`

**注:**

只有在已经登录的状态才可以使用

**反馈**：

- 如果找到用户，返回已注册的所有用户的用户名、邮箱以及电话信息（打表）

---

### agenda user delete

删除用户使用子命令`agenda user delete`

**注**：

- 操作为删除自己的账号，自动注销
- 用户账号删除之后：
  - 以该用户为发起者的会议将会被删除
  - 以该用户为参与者的会议将会从参与者列表中移除该用户。如果此操作造成参与者数量为0，则会议也会被移除。

**反馈**:

- 返回成功注销的信息
- 失败要返回注销失败（有什么情况会导致注销失败吗？）

---

### agenda meeting create

创建会议使用子命令`agenda meeting create`

**参数：**

- --start/-s，开始时间，格式为(YYYY-MM-DD/HH:mm:ss)
- --end/-e，结束时间
- --title/-t，会议主题
- --participant/-p,会议参与者

**注：**

- 用户无法分身参与多个狐疑，如果用户已有的会议安排与将要创建的会议在时间上有重叠，则会无法创建这个会议。
- 会议主题是唯一的。

**反馈：**

- 成功返回创建成功提示信息
- 失败返回创建失败相关信息

---

### agenda meeting addUser

 添加会议参与者使用子命令`meeting addUser -p [Participator] -t [Title] `

**参数：**

- --title/-t，会议主题
- --participants/-p，要添加的用户名（可以为多个）

**反馈：**

- 成功返回相关信息
- 失败也返回相关信息

----

### agenda meeting deleteUser

删除会议参与者使用子命令`meeting deleteUser -p [Participator] -t [Title]` 

**参数：**

- --title/-t，会议主题
- --participants/-p，要添加的用户名（可以为多个）

**反馈：**

- 成功返回相关信息
- 失败也返回相关信息

----

### agenda meeting lookup

查找会议(仅登录可用)使用子命令`meeting lookup -s [StartTime] -e [EndTime] `

**参数：**

- --start/-s，开始时间
- --end/-e，结束时间

**注：**

- 已登录的用户可以查询自己的议程在某一时间段内的所有会议安排

----

### agenda meeting cancel

取消会议(已经登录的用户可以取消自己发起的某一会议安排)使用子命令` meeting cancel -t [title] `

**参数：**

- --title/-t，会议主题

----

### agenda meeting exit

退出会议(已经登录的用户可以退出自己参与的某一会议安排)使用子命令`meeting exit -t [title] `

**参数：**

- --title/-t，会议主题

 

**注：**

- 如果此操作造成会议参与者人数为0，则会议将会被删除

----

### agenda meeting clear

清空会议(已经登录的用户可以清空自己发起的所有会议安排)使用子命令`clear `

## 结构定义

### 用户(user)
位置：`entity/useroper.go`中
作用：定义用户所需要的结构，以及相关的操作函数
```go
type User struct{
    Username string
    Password string
    Email string
    Telphone string
}
```

### 会议(conference)
位置：`entity/conference.go`中
作用：定义会议所需要的结构，以及相关的操作函数
```go
type Meeting struct{
    StartTime string
    EndTime string
    Title string
    UserList []User
}
```

# 参考知识与备份

如果有可能经常用到的、比较重要的而且容易忘记的知识可以放在这里。

## JSON读写学习
### 编码
代码来自于网络
结构体转JSON代码：
```go
package main

import (
    "encoding/json"
    "fmt"
)

type DebugInfo struct {
    Level  string `json:"level,omitempty"` // Level解析为level,忽略空值
    Msg    string `json:"message"`         // Msg解析为message
    Author string `json:"-"`               // 忽略Author，或者设为未导出字段
}


func main() {

    dbgInfs := []DebugInfo{
        DebugInfo{"debug", `File: "test.txt" Not Found`, "Cynhard"},
        DebugInfo{"", "Logic error", "Gopher"},
    }

    if data, err := json.Marshal(dbgInfs); err == nil {
        fmt.Printf("%s\n", data)
    }
}
```

自定义结构对MarshalJSON()接口的实现
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Point struct{ X, Y int }

func (pt Point)MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf(`{"X":%d,"Y":%d}`, pt.X, pt.Y)), nil
}

func main() {
    if data, err := json.Marshal(Point{50, 50}); err == nil {
        fmt.Printf("%s\n", data)
    }
}
```

### 解码
解码主要使用的是JSON包的函数`func Unmarshal(data []byte,v interface{}) error`
```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    data := `[{"Level":"debug","Msg":"File: \"test.txt\" Not Found"},` +
        `{"Level":"","Msg":"Logic error"}]`

    var dbgInfos []map[string]string
    json.Unmarshal([]byte(data), &dbgInfos)

    fmt.Println(dbgInfos)
}
```

#### JSON转结构体
与编码一样，JSON是通过反射机制来实现解码的，所以结构必须导出所转换的字段，不导出的字段不会被JSON包解析。另外解析的时候不区分大小写。

```go
package main

import (
    "encoding/json"
    "fmt"
)

type DebugInfo struct {
    Level string
    Msg string
    author string  // 未导出字段不会被json解析
}

func (dbgInfo DebugInfo) String() string {
    return fmt.Sprintf("{Level: %s, Msg: %s}", dbgInfo.Level, dbgInfo.Msg)
}

func main() {
    data := `[{"level":"debug","msg":"File Not Found","author":"Cynhard"},` +
        `{"level":"","msg":"Logic error","author":"Gopher"}]`

    var dbgInfos []DebugInfo
    json.Unmarshal([]byte(data), &dbgInfos)

    fmt.Println(dbgInfos)
}
```

结构体也可以与编码的时候一样设置结构体字段标签

#### 转换接口
和编码的时候类似，解码使用接口Unmarshaler，需要实现`UnmarshalJSON([]byte) error`

下面这个接口没有实现解码算法，只是将参数打印出来
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Point struct{ X, Y int }

func (Point) UnmarshalJSON(data []byte) error {
    fmt.Println(string(data))
    return nil
}

func main() {
    data := `{"X":80,"Y":80}`
    var pt Point
    json.Unmarshal([]byte(data), &pt)
}
```

## 文件读写
使用ioutil的ReadFile和WriteFile，并使用stirng()等函数进行类型转换，例程：
```go
package main

import(
    "fmt"
    "io/ioutil"
)

func main() {
    //read
    b, err := ioutil.ReadFile("test.log")
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println(b)
    str := string(b)
    fmt.Println(str)
    //write
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("test.txt", d1, 0644)
    check(err)
}
```

## 实例
使用user进行的测试（已经通过）
```go
	var myuser []entity.User
	fmt.Println("JSON decode to structure test1")
	jsonStr := `[{"username":"wtysos11""password":"123456","email":"wtysos11"}{"username":"wtysos11","password":"123456""email":"wtysos11"}]`
	json.Unmarshal([]byte(jsonStr),&myuser)
	fmt.Println(myuser)
	
	if data,err:=json.Marshal(myuser);err==nil{
		fmt.Printf("%s\n",data)
	}
```

## Flag使用
### StringArray使用
使用`cmd.Flags().StringArray()`可以声明一个flag
使用`cmd.Flags().GetStringArray()`可以取得这个数组
对于数组而言只有重复使用标签才能够作为不同元素，比如`-a="ss" -a="tt"`，而同一个标签内都算作是一个字符窜

## 异常
```go
if 遇到错误
    return errors.New("遇到xx错误")
```