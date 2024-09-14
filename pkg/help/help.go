package help

import (
    "fmt"
)

type Help struct {
    Usage     string    `json:"usage"`
    ShortDesc string    `json:"short_desc"`
    LongDesc  string    `json:"long_desc"`
    Examples  []Example `json:"examples"`
}

type Example struct {
    Description string `json:"description"`
    Command     string `json:"command"`
}

// Help information in various languages
var HelpData = map[string]*Help{
    "en": {
        Usage:     "Usage: j2y [-f inputFile | -i jsonData] [output]",
        ShortDesc: "j2y (json to yaml) is a tool for converting JSON data (from files or direct input) to YAML format.",
        Examples: []Example{
            {"Convert from file:", "j2y -f input.json output.yaml"},
            {"Convert from command line JSON data:", "j2y -i '{\"key\": \"value\"}' output.yaml"},
            {"Convert from standard input:", "echo '{\"foo\":1}' | j2y"},
            {"Specify input file only, output to stdout:", "j2y -f input.json"},
            {"Specify JSON data only, output to stdout:", "j2y -i '{\"key\": \"value\"}'"},
        },
    },
    "zh_CN.UTF-8": {
        Usage:     "用法: j2y [-f 输入文件 | -i JSON 数据] [输出]",
        ShortDesc: "j2y（json to yaml）是一个工具，用于将 JSON 数据（从文件或直接输入）转换为 YAML 格式。",
        Examples: []Example{
            {"从文件转换:", "j2y -f input.json output.yaml"},
            {"从命令行输入 JSON 数据:", "j2y -i '{\"key\": \"value\"}' output.yaml"},
            {"从标准输入转换:", "echo '{\"foo\":1}' | j2y"},
            {"只指定输入文件，输出到标准输出:", "j2y -f input.json"},
            {"只指定 JSON 数据，输出到标准输出:", "j2y -i '{\"key\": \"value\"}'"},
        },
    },
}

// Initialize LongDesc for each Help entry
func init() {
    for lang, help := range HelpData {
        help.LongDesc = generateLongDesc(lang)
    }
}

// Generate LongDesc with formatted examples
func generateLongDesc(lang string) string {
    var examples []string
    num := 1
    examples = append(examples,"\n");
    for _, example := range HelpData[lang].Examples {
        
        examples = append(examples, fmt.Sprintf("%d. %s \n\t %s \n",num, example.Description, example.Command))
        num++
    }
    return fmt.Sprintf("%s\n\nExamples:\n%s", HelpData[lang].ShortDesc,examples)
}