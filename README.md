# gcopyright 软著代码生成利器

`gcopyright`是针对软著申请过程中，需要提供一个3000行代码的word文件而开发的软件。

使用`gcopyright`直接从代码库种提取3000行代码，并生成word文件，无需安装任何环境，下载可执行文件即可立刻使用。

## 1. 下载binary

本repo已预编译了`linux`, `darwin`, `windows`系统的binary，如果不想自行编译，可直接[下载最新版本]((https://github.com/enix223/gcopyright/releases/))

## 2. 使用方法:

```bash
gcopyright \
    -o "/tmp/code.docx" \
    -p "./example/**/*.js" "./example/**/*.css"
```

## 3. 参数

| 参数 | 描述 |
|-|-|
| `--patterns PATTERNS`, `-p PATTERNS` | 搜索路径，支持glob |
| `--ignoreerror`, `-i`  | 是否忽略错误，默认是 |
| `--totallines TOTALLINES`, `-l TOTALLINES` | 最多读取多少行，默认3000  |
| `--output OUTPUT`, `-o OUTPUT` | 输出docx文件路径|
| `--trimline`, `-t`  | 是否去除空行，默认是|
| `--verbose`, `-v` | verbose模式 |