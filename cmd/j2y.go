package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"

    "gopkg.in/yaml.v2"
    "github.com/spf13/cobra"
)

func main() {
    var inputFile string
    var outputFile string

    // 创建根命令
    var rootCmd = &cobra.Command{
        Use:   "json2yaml",
        Short: "将 JSON 文件转换为 YAML 文件",
        Run: func(cmd *cobra.Command, args []string) {
            // 读取 JSON 文件
            jsonFile, err := ioutil.ReadFile(inputFile)
            if err != nil {
                fmt.Printf("读取文件失败: %v\n", err)
                return
            }

            // 解析 JSON
            var data interface{}
            if err := json.Unmarshal(jsonFile, &data); err != nil {
                fmt.Printf("解析 JSON 失败: %v\n", err)
                return
            }

            // 转换为 YAML
            yamlData, err := yaml.Marshal(data)
            if err != nil {
                fmt.Printf("转换为 YAML 失败: %v\n", err)
                return
            }

			if outputFile == "" {
				fmt.Printf("%v\n", string(yamlData))
				return
			}

            // 写入 YAML 文件
            if err := ioutil.WriteFile(outputFile, yamlData, 0644); err != nil {
                fmt.Printf("写入文件失败: %v\n", err)
                return
            }

            fmt.Println("转换成功!")
        },
    }

    // 添加命令行参数
    rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "输入 JSON 文件")
    rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "输出 YAML 文件")
    rootCmd.MarkFlagRequired("input")
	
	// 可以没有输出文件，没有设置输出文件将打印到标准输出 
    //rootCmd.MarkFlagRequired("output")

    // 运行命令
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
