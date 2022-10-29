package command

var RunOptions struct {
	Patterns    []string `arg:"-p,env:GC_PATTERNS,required" help:"搜索路径，支持glob"`
	IgnoreError bool     `arg:"-i,env:GC_IGNORE_ERR" help:"是否忽略错误，默认是" default:"1"`
	TotalLines  int      `arg:"-l,env:GC_TOTAL_LINES" help:"最多读取多少行" default:"3000"`
	Output      string   `arg:"-o,env:GC_OUTPUT_PATH,required" help:"输出docx文件路径"`
	TrimLine    bool     `arg:"-t,env:GC_TRIM_LINE" help:"是否去除空行，默认是" default:"1"`
	Verbose     bool     `arg:"-v,env:GC_VERBOSE" help:"verbose模式" default:"0"`
}
