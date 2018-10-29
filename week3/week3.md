服务计算——selpg

[TOC]

# 1. 项目需求

​	作业要求中的千言万语汇成一句话：使用 golang 开发 [开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html) 中的 **selpg** 

# 2. 分析设计 

​	刚开始做这个东西的时候，看完了老师给的参考资料，自己还仍旧完完全全是一头雾水，丝毫不知道从何处下手，随后就**github**以及**csdn**了一波，参考了前辈已经搞过的东西，理清了思路，然后再开始着手做就变得简单多了，接下来我会将自己做这个项目的过程大概梳理一下，给后人减轻压力。

## 2.1 项目所涉及的Go语言知识点

​	做这次作业需要使用的Go语言的知识点大概有以下几个方面

​		① 函数  

​		② 结构体

​		③ 参数的绑定

​		④ 异常处理

​		⑤ 文件读写

​		⑥ 命令行参数的解析

## 2.2 理解selpg 

​	[selpg](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html) 是一个自定义命令行程序，全称select page，即从源（标准输入流或文件）读取指定页数的内容到目的地（标准输出流或给给打印机打印） 

## 2.3 Go使用flag包解释命令行参数

flag包能解析的参数有如下四种形式，使用的时候分三种类型：

类型一 
   `cmd  -flag=x // 支持所有类型` 
   `cmd  -flag x // 只支持非bool类型 `

类型二 
   `cmd  -flag // 只支持bool类型` 

类型三 
   `cmd  abc // 没有flag的参数`

### 2.3.1 类型一 

类型一定义参数有两种形式，下方Xxx代表类型 

#### 2.3.1.1 类型一定义参数

1. **flag.Xxx()** 

   `flag.String()`, `Bool()`, `Int()` 这种形式返回一个指向该参数的指针

   例子

   ```go
   var date = flag.Int("d", 20171107, "help message for date") 
   
   flag.Parse()
   ```

   (**注意: date 是指向 -d 参数值的指针**)

   ​	若在命令行输入 `-d=20121212 -a=20` 或者 `-d 20121212 -a 20`，则我们在程序中即可用 *date 取得 -d 的值20121212

   ​	若只输入 `-a=20` 或者 `-a 20`，因为缺省标志 `-d` ，则 *date 取到的是 -d 定义中的默认值 20171107 

2. **flag.XxxVar()** 

   `flag.StringVar()`, `BoolVar()`, `IntVar()` 这种形式把参数绑定到一个变量

   例子  

   ```go
   var age int
   
   flag.IntVar(&age, "a", 18, "help message for age")
   
   flag.Parse()
   ```

   ​	命令行输入`-d=20121212 -a=20` 或者`-d=20121212 -a 20` ，在程序里 变量age 则获取到 -a 的值 20

   ​	若只输入`-d=20121212` 或者`-d=20121212`，缺省标志-a， 则 变量age 获取定义中的默认值 18

#### 2.3.1.2 类型一解析参数

​	注意到上面例子最后都带有一行 flag.Parse()

​	因为定义好参数后，只有调用方法 flag.Parse() 解析命令行参数到定义的flag，这样我们才能使用上面两个例子的 *date 和 age 取得对应flag的参数值

### 2.3.2 类型二

`cmd  -flag //该形式只支持bool类型，对应的值是1, 0, t, f, true, false, TRUE, FALSE, True, False`

​	默认的，如果我们在命令行里提供了-flag，则其对应的值为true，否则为flag.Bool/BoolVar中指定的默认值；如果希望显示设置为false则使用-flag=false。

例子

```go
var exist_f = flag.Bool("f", false, "help message for format")

flag.Parse()
```

​	当在命令行输入 `-d=123 -f` 时，程序里 *exist_f 的值就为 true 了

​	若命令行只输入 `-d=123` 时，因为缺省-f，*exist_f 的值为事先定义中默认的 false

### 2.3.3 类型三

`cmd  abc //没有flag的参数`

​	1.通过 flag.Args() 获取非flag参数列表

​	2.通过 flag.Arg(i) 来获取非flag命令行第i个参数，i 从 0 开始

​	3.通过flag.NArg() 获得非flag参数个数

例子

​	当命令行敲入 `-d 20121212 -a 20 Wang yx` ，最后两个值即为不带标志的参数

```go
flag.Parse()

var fullName = flag.Args() // fullName = ['Wang', 'yx']

var firstName = flag.Arg(0) // firstName = Wang

var lastName = flag.Arg(1) // lastName = yx

var num = flag.NArg() // num = 2
```

## 2.4 命令行参数设计

`selpg -s startPage  -e endPage [-l linePerPage | -f ][-d dest] filename`

必需标志以及参数：

- -s，后面接开始读取的页号 int
- -e，后面接结束读取的页号 int 
   s和e都要大于1，并且s <= e，否则提示错误

可选参数：

- -l，后面跟行数 int，代表多少行分为一页，不指定 -l 又缺少 -f 则默认按照72行分一页
- -f，该标志无参数，代表按照分页符’\f’ 分页
- -d，~~后面接打印机标号，用于将内容传送给打印机打印~~ 我没有打印机用于测试，所以当使用 -d destination(随便一个字符串作参数)时，就会通过管道把内容发送给 grep命令，并把grep处理结果显示到屏幕
- filename，唯一一个无标识参数，代表选择读取的文件名
- [] 中内容为可选择性输入的

## 2.5 读写函数

​	官网解释：bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。 

### 2.5.1 bufio.NewReader

​	os.Stdlin和打开文件都能得到输入数据流，而bufio.NewReader就是一个对象，该对象带有一个缓冲区，与一个数据流绑定。通过

```go
rd := bufio.NewReader(fin)
if psa.page_type == false {
    line_ctr = 0
    page_ctr = 1
    for true {
        line, err2 = rd.ReadString('\n')
        if err2 != nil { /* error or EOF */
            break
        }
        line_ctr++
        if line_ctr > psa.page_len {
            page_ctr++
            line_ctr = 1
        }
        if page_ctr >= psa.start_page && page_ctr <= psa.end_page {
            fmt.Fprintf(fout, "%s", line)
        }
    }
} 
	
```

​	我们创建了一个newreader，是bufio.newReader类型，它绑定了一个文件流，用bufio.NewReader对象的好处是我们可以利用这个对象的函数方法，如

```go
line, err := reader.ReadBytes('\n')1
```

​	这样就能按行读取缓冲区中的数据。

### 2.5.2 bufio.Writer 

​	这个对象和bufio.Reader很像，不过它是绑定了一个输出数据流。 
 	可以利用

```go
writer := bufio.NewWriter(os.Stdout)
errW := writer.Write(byte[]("line 1"))
```

​	导出到标准输出，同样也可以绑定到文件数据流中 

## 2.6 os/exe

​	exec包执行外部命令。它包装了os.StartProcess函数以便更容易的修正输入和输出，使用管道连接I/O，以及作其它的一些调整。 
 	这个包可以帮助我们在程序中启动子进程并使用管道连接I/0。具体用法是先建立一个子进程的输入管道，启动子进程，从父进程接收标准输出，关闭管道，关闭子进程。  

 ### 2.6.1 exec.Command(filePath)

```go
cmd_grep := exec.Command("./" + args.printDestination)
```

​	创建一个命令对象，参数为子进程路径和子进程参数（可选）

### 2.6.2 func (c \*Cmd) StdinPipe() (io.WriteCloser, error) 

 	它构建了到子进程的输入管道，返回一个io.WriterClose对象，这个对象绑定了一个数据流，该数据流是子进程的数据输入流。所以可以直接调用对象的方法writer.Write([]byte)将数据转化为到子进程的标准输入，如：

```
_, errW := writer.Write([]byte("line 3"))
```

### 2.6.3 func (c \*Cmd) Start() error 

​	启动子进程

### 2.6.4 func (\*Cmd) Wait 

​	在命令执行完成后调用，返回所调用命令的执行情况，同时释放资源

# 3. 代码设计

大体框架是：

​	①  读取命令行参数

​	②  打开需要处理的文件

​	③  根据用户给定参数进行操作

​	④  结束读写

## 3.1 引入所需要的包

```go
import (
    "fmt"
	"os"
	"bufio"
	"github.com/spf13/pflag"
	"os/exec"
	"io"
)
```

注：由于这里老师要求要使用[Unix标准](http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html)，所以引入`pflag`包代替`flag`，但是两个包的函数在使用上面基本是相同的。

## 3.2 创立记录数据的结构体

```go
type selpg_args struct {
	start_page  int  //开始页码  
	end_page    int  //结束页码
	in_filename string  // 文件名
	page_len    int    // 每一页的大小
	page_type   bool   // 页类型
	print_dest  string  // 打印地
}
```

## 3.3 main 函数

```go
func main() {
    //创建结构体数据
	sa := selpg_args{0, 0, "", 5, false, ""}
	//拿到生成的可执行文件的名字
    progname = os.Args[0]
	//初始化结构体数据
	Parser(&sa)
    //处理参数
	processArgs(len(os.Args), &sa)
	//传递参数，完成对应操作
    processInput(&sa)
}
```

## 3.4 Parser函数

```go
func Parser(p *selpg_args) {
	pflag.Usage = usage
	pflag.IntVarP(&p.start_page,"start", "s", 0, "首页")
	pflag.IntVarP(&p.end_page,"end","e", 0, "尾页")
	pflag.IntVarP(&p.page_len,"linenum", "l", 5, "打印的每页行数")
	pflag.BoolVarP(&p.page_type,"printdes","f", false, "是否用换页符换页")
	pflag.StringVarP(&p.print_dest, "othertype","d", "", "打印目的地")
	pflag.Parse()
}
```

注：这里使用`pflag`包的函数，完成对命令行数据的操作。

## 3.5 processArgs函数 

​	 ` processArgs函数 `主要是用于处理输入时候的各种错误，比如起始页码是负数，终止页码小于起始页码等情况，增加程序的健壮性，这里给出部分情况。

```go
func processArgs(psa *selpg_args) {
//处理-s 这种情况
if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "%s: 1st arg should be -s=start_page\n", progname)
		pflag.Usage()
		os.Exit(2)
}
//处理起始的页码小于1的情况
if psa.start_page < 1  {
    fmt.Fprintf(os.Stderr, "%s: invalid start page %s\n", progname, psa.start_page)
    pflag.Usage()
    os.Exit(3)
}
....
....
//检查要操作的文件是否存在
if pflag.NArg() > 0 {
		psa.in_filename = pflag.Arg(0)
		/* check if file exists */
		file, err := os.Open(psa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n", progname, psa.in_filename)
			os.Exit(7)
}
		/* check if file is readable */
		file, err = os.OpenFile(psa.in_filename, os.O_RDONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				fmt.Fprintf(os.Stderr, "%s: input file \"%s\" exists but cannot be read\n", progname, psa.in_filename)
				os.Exit(8)
			}
		}
		file.Close()
}
```

## 3.6 processInput函数 

​	 `processInput函数`用来对文件的内容进行处理，根据后缀参数的类型完成操作，这里只给出函数的一部分。

```go
func processInput(psa *selpg_args) {
	fin := os.Stdin
	fout := os.Stdout
	var (
		 page_ctr int
		 line_ctr int
		 err error
		 err1 error
		 err2 error
		 line string
		 cmd *exec.Cmd
		 stdin io.WriteCloser
	)
.....
.....
```

# 4. 程序测试

​	程序测试部分按照参考文献的测试部分一步一步进行。

​	注：in.txt为已含数据文件，out.txt为空白无数据文件，error.txt为空白无数据文件。

## 4.1 selpg -s1 -e1 in.txt  

显示第一页中的所有内容

![1](https://github.com/wyxwyx46941930/ServiceComputing/blob/master/Desktop/%E6%9C%8D%E5%8A%A1%E8%AE%A1%E7%AE%97/week3/1.png)

## 4.2 selpg -s1 -e1 < in.txt 

正常显示第一页中的所有内容

![2](https://github.com/wyxwyx46941930/ServiceComputing/blob/master/Desktop/%E6%9C%8D%E5%8A%A1%E8%AE%A1%E7%AE%97/week3/2.png)

## 4.3 selpg -s1 -e1 in.txt >out.txt

将in.txt文件中的内容写到out.txt文件中

![3](https://github.com/wyxwyx46941930/ServiceComputing/blob/master/Desktop/%E6%9C%8D%E5%8A%A1%E8%AE%A1%E7%AE%97/week3/3.png)

## 4.4 selpg -s1 -e1 in.txt 2>error.txt 

把运行的错误信息写入error.txt中，这里由于运行正常，所以没有内容写入

![4](https://github.com/wyxwyx46941930/ServiceComputing/blob/master/Desktop/%E6%9C%8D%E5%8A%A1%E8%AE%A1%E7%AE%97/week3/4.png)

## 4.5 selpg -s1 -e1 in.txt  2>dev/null丢弃无用的错误输出

![5](https://github.com/wyxwyx46941930/ServiceComputing/blob/master/Desktop/%E6%9C%8D%E5%8A%A1%E8%AE%A1%E7%AE%97/week3/5.png)

## 4.6 使用linux的ps指令

![6](https://github.com/wyxwyx46941930/ServiceComputing/blob/master/Desktop/%E6%9C%8D%E5%8A%A1%E8%AE%A1%E7%AE%97/week3/6.png)

## 4.7 使用cat指令

![7](https://github.com/wyxwyx46941930/ServiceComputing/blob/master/Desktop/%E6%9C%8D%E5%8A%A1%E8%AE%A1%E7%AE%97/week3/7.png)

## 4.8 更改可显示的行数l

![8](https://github.com/wyxwyx46941930/ServiceComputing/blob/master/Desktop/%E6%9C%8D%E5%8A%A1%E8%AE%A1%E7%AE%97/week3/8.png)

## 4.9 报错信息

![9](C:\Users\WYX\Desktop\服务计算\week3\9.png)

# 5. 资料参考

① 舍友梓溢大佬的[博客](https://z1wu.github.io/post/service_computing_3/)，前排奶一波，梓溢大佬牛逼!!!

② 如何[使用flag与goflag](https://o-my-chenjian.com/2017/09/20/Using-Flag-And-Pflag-With-Golang/)。

③ [宇翔师兄的github](https://github.com/Mensu/selpg)，宇翔师兄写出来的代码我所不能及也！

④ [GO语言实现selpg](https://blog.csdn.net/kunailin/article/details/78262456)

---

最后附上[自己的github](https://github.com/wyxwyx46941930/ServiceComputing/tree/master/Desktop/%E6%9C%8D%E5%8A%A1%E8%AE%A1%E7%AE%97)