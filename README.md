# Z!

code fast and reuse!

```go
import "github.com/aizk/z"
```
# zconv

type trans tools.

- Int64ToString
- IntToString

# zslice

wrap of slice。

- `Clone()` 克隆一个 slice
- 删除某个位置的元素
- 查找某个元素的位置
- 复制一份 slice
- 泛型操作...
- map、reduce...
- filter...

## Chunk(slice, num) slice

```go
s := []int{1, 2, 3, 4, 5}
n := Chunk(s, 2) // [[1, 2], [3, 4], [5]]
```

# zos

- `Exist(path string)` 判断文件或目录是否存在
- `IsDir(path string)` 判断是否是目录
- `OpenFile(path string, perm 0777)` 打开文件
- `WorkDIR` 返回工作目录

# rand

- String() 随机生成字符串
- StringWithType() 可以选择生成的字符串中有哪些类型的字符

# crypto

- MD5()  md5 加密函数
