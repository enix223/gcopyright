package command

var RunOptions struct {
	Patterns       []string `arg:"-p,env:GC_PATTERNS,required" help:"搜索路径，支持glob"`
	AbortIfError   bool     `arg:"-e,env:GC_ABORT_IF_ERR" help:"当出现错误时是否中止程序，默认否"`
	TotalLines     int      `arg:"-l,env:GC_TOTAL_LINES" help:"最多读取多少行" default:"3000"`
	Output         string   `arg:"-o,env:GC_OUTPUT_PATH,required" help:"输出docx文件路径"`
	AllowEmptyLine bool     `arg:"-a,env:GC_ALLOW_EMPTY_LINE" help:"是允许空行，默认否"`
	Verbose        bool     `arg:"-v,env:GC_VERBOSE" help:"verbose模式"`
}
