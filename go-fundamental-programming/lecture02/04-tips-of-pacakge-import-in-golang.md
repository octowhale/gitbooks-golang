# 包引入的规则和方法

引入包的时候使用 `import` 关键字。
如果要引入多个包的时候，可以使用多个 `import` 进行引入。

```go
import "fmt"
import "io"
import "os"
import "strings"
import "time"
import "math"
```

当然，如果不想写那么多个 `import` 的话，可以使用使用『圆括号』的方式简写。

```go
import (
    "fmt"
    "io"
    "os"
    "strings"
    "time"
    "math"
)
```

## 包引入的命令格式

在 golang 中，引入包的命令格式为 

```go
import [别名] 包名
```

其中，别名可以省略。

### 标准引入

例如我们最常见的  `import "fmt"`。

```go
// 标准引入
import "fmt"
func main(){
	fmt.Println("标准引入 : import \"fmt\"")
}
```

### 别名引入

在使用别名引入的时候，程序中调用包的函数时，就不能再使用包名了，而使用『别名』进行替代。

```go
// 别名引入
import std "fmt"
func main(){
	std.Println("别名引入： import std \"fmt\"")
}
```

### 隐藏引入

当别名为『逗点 (.)』的时候，这种引入被称为隐藏引入。在程序中使用包中的函数，可以直接使用函数名，而不再使用报名。

当引入包过多的时候不推荐使用这种方法，原因是不容易区分调用的函数来自哪个包。

```go
// 隐藏引入
import . "fmt"
func main() {
	Println("隐藏引入: import .\"fmt\"")
}
```


## 包的调用

导入包之后，可以使用 `pakcage_name.func_name` 的格式（包名与函数之间用『逗点』分隔），对包中的函数进行调用。
如果导入包之后 **未调用** 包中的函数或类型，程序在编译的时候会报错。
