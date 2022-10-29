# gcopyright 软著代码生成利器

## 使用方法:

```bash
gcopyright \
    -o "/tmp/code.docx" \
    -p "./example/**/*.js" "./example/**/*.css"
```

## 参数

| 参数 | 描述 |
|-|-|
| `--patterns PATTERNS`, `-p PATTERNS` | 搜索路径，支持glob |
| `--ignoreerror`, `-i`  | 是否忽略错误，默认是 |
| `--totallines TOTALLINES`, `-l TOTALLINES` | 最多读取多少行，默认3000  |
| `--output OUTPUT`, `-o OUTPUT` | 输出docx文件路径|
| `--trimline`, `-t`  | 是否去除空行，默认是|
| `--verbose`, `-v` | verbose模式 |