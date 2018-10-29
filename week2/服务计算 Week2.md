`服务计算`——安装 go 语言开发环境

[TOC]

# 1. 安装VScode编辑器

步骤：

​	① 进入给定[链接](https://code.visualstudio.com/docs/setup/linux) 

​	② 下载地址里废话不用多看，直接ctrl + f 定位 `centos `所在的版本

​	③ 打开命令行窗口，使用 `yum `指令完成安装

```python
		yum check-update
		sudo yum install code
```

​	④  安装完成，在命令行窗口输入 code , 若出现 VScode 终端窗口，则证明安装完成

![1](C:\Users\WYX\Desktop\服务计算\week2\1.png)

# 2. Golang

## 2.1 安装Golang

步骤：

​		①  终端输入以下，输入管理员密码，等待安装完成 

​		 	 ` sudo yum install golang  `

​		②  查看安装目录

​		         `rpm -ql golang |more `

![2](C:\Users\WYX\Desktop\服务计算\week2\2.png)

​		③  测试安装

​			 `go version ` 

![3](C:\Users\WYX\Desktop\服务计算\week2\3.png)

## 2.2 Golang 环境变量配置 

注：进行这一步我参考了`如何使用Go编程`这个教程

### 2.2.1 GOPATH 

> GOPATH 环境变量:指定工作空间位置

① 首先创建一个工作空间目录，并设置相应的GOPATH

​	指令：	 `mkdir $HOME/work `    - >    

​         ` export GOPATH=$HOME/work ` 

② 将此工作空间的`bin子目录`添加到` PATH`中

​	指令：    ` export PATH=$PATH:$GOPATH/bin `

### 2.2.2 包路径

​	标准库中的包有给定的短路径，比如 `"fmt"` 和 `"net/http"`。 对于你自己的包，你必须选择一个基本路径，来保证它不会与将来添加到标准库， 或其它扩展库中的包相冲突 



​	这里我们使用 `github.com/user `作为我们的基本路径，我们在工作空间中创建一个目录，并将源码存放进去：

​	指令：` mkdir -p $GOPATH/src/github.com/user ` 

## 2.3 运行简单Golang程序

### 2.3.1 设置包路径

​	指令：`mkdir $GOPATH/src/github.com/user/hello `

### 2.3.2 创建 hello.go 文件

​	在包目录下创建`hello.go`文件，并输入以下代码用做测试：

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
```

### 2.3.3 编译执行

编译：

​	指令： ` go install github.com/user/hello ` 

注：该指令完成后，会产生一个可执行的二进制文件，接着会将二进制文件生成 `hello..exe` 文件并安装到` bin ` 目录中



执行：

​	指令：`$GOPATH/bin/hello ` 	

注：此时命令窗口应该会打印 `Hello,world`字样

---



![7](C:\Users\WYX\Desktop\服务计算\week2\7.png)



### 2.3.4 提交代码到远程github库

​	[附Git个人学习教程含提交方法](https://blog.csdn.net/wyxwyx469410930/article/details/82826608)

## 2.4 运行第一个库

注：这里自主编写一个库，并让 hello 程序来使用它

### 2.4.1 创建包目录

​	指令：` mkdir $GOPATH/src/github.com/user/stringutil`

### 2.4.2 创建库文件

​	在该目录下创建名为 ` reverse.go`，并将内容设置如下：

```go
// stringutil 包含有用于处理字符串的工具函数。
package stringutil

// Reverse 将其实参字符串以符文为单位左右反转。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

### 2.4.3 编译库文件

​	指令：` go build github.com/user/stringutil` 

### 2.4.4 更改hello.go文件内容 

```go
package main

import (
	"fmt"

	"github.com/user/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
```

### 2.4.5 工作空间

​	做完上述工作后，工作空间如下：

![4](C:\Users\WYX\Desktop\服务计算\week2\4.png)

​	 `go install`   会将  `stringutil.a`对象放到 `pkg/linux_amd64` 目录中，它会反映出其源码目录。 这就是在此之后调用 `go` 工具，能找到包对象并避免不必要的重新编译的原因。 `linux_amd64` 这部分能帮助跨平台编译，并反映出你的操作系统和架构。 

### 2.4.6 运行

![5](C:\Users\WYX\Desktop\服务计算\week2\5.png)

## 2.5 包

​	Go源文件中的第一个语句必须是

```
package 名称
```

​	① 这里的 **名称 **即为导入该包时使用的默认名称。 （一个包中的所有文件都必须使用相同的 **名称**。）

​	② Go的约定是包名为导入路径的最后一个元素：作为 “`crypto/rot13`” 导入的包应命名为 `rot13`。

​	③ 可执行命令必须使用 `package main`。

​	④ 链接成单个二进制文件的所有包，其包名无需是唯一的，只有导入路径（它们的完整文件名） 才是唯一的。

## 2.6 Go_test

​	Go拥有一个轻量级的测试框架，它由 `go test` 命令和 `testing` 包构成。

​	你可以通过创建一个名字以 `_test.go` 结尾的，包含名为 `TestXXX` 且签名为 `func (t *testing.T)` 函数的文件来编写测试。 测试框架会运行每一个这样的函数；若该函数调用了像 `t.Error` 或 `t.Fail` 这样表示失败的函数，此测试即表示失败。

​	我们可通过创建文件 `$GOPATH/src/github.com/user/stringutil/reverse_test.go` 来为 `stringutil` 添加测试，其内容如下：

```go
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

​	接着使用 `go test` 运行该测试：

​		指令：` go test github.com/user/stringutil`

​		运行完成：`  ok  	github.com/user/stringutil 0.001s `

​	如图：

![6](C:\Users\WYX\Desktop\服务计算\week2\6.png)

## 2.7 远程包

​	 `go `工具可**通过根据导入路径的描述来获取包源代码的特性**从远程代码库自动获取包。



​		① 	若你在包的导入路径中包含了代码仓库的URL，`go get` 就会自动地获取、 构建并安装它 。



​		②    若指定的包不在工作空间中，`go get` 就会将会将它放到 `GOPATH` 指定的第一个工作空间内。 



![8](C:\Users\WYX\Desktop\服务计算\week2\8.png)

# 3. 必要的工具和插件的安装 

## 3.1 安装 Git 客户端

> 这一步安装的客户端是版本比较落后的，但是已经足够满足我们使用。

## 3.2 安装 Go 工具

安装指令：

​		①  创建文件夹 

​				 `mkdir $GOPATH/src/golang.org/x/  `  

​		②  下载源码 (**需要等待一段时间**)

​				 ` go get -d github.com/golang/tools `

​		③   copy源码  

​				 ` cp $GOPATH/src/github.com/golang/tools $GOPATH/src/golang.org/x/ -rf `

​		④ 安装工具包

​				 `go install golang.org/x/tools/go/buildutil`

随后进入`Vscode`，此时 `Vscode `系统右下角会提示你安装插件，点击` install all `按钮即可。	 

## 3.3 编译运行

​		指令：    ` go install github.com/github-user/hello ` ->

​				 `hello` 

![10](C:\Users\WYX\Desktop\服务计算\week2\10.png)

# 4. 安装运行Go tour

指令：

​		①  ` go install github.com/github-user/hello`

​		②   `go tour `  

![9](C:\Users\WYX\Desktop\服务计算\week2\9.png)