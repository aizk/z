# Outline

- zbytes 字符操作相关
- os.u 系统相关
- gin.u gin 框架相关封装
- random.u 随机生成数字、字符串
- crypto.u 加密相关工具封装
- slice.u 切片操作封装
- struct.u 结构体封装

# zslice

wrap of slice。

- `Clone()` 克隆一个 slice
- 删除某个位置的元素
- 查找某个元素的位置
- 复制一份 slice
- 泛型操作...
- map、reduce...
- filter...

# zos

- `Exist(path string)` 判断文件或目录是否存在
- `IsDir(path string)` 判断是否是目录
- `OpenFile(path string, perm 0777)` 打开文件
- `WorkDIR` 返回工作目录

## rand

- String() 随机生成字符串
- StringWithType() 可以选择生成的字符串中有哪些类型的字符

## crypto

- MD5()  md5 加密函数

## mysql

- GormOpen(options) 使用 gorm 打开数据库连接
