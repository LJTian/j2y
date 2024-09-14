# Info
json to yaml tool.

# Install
```shell
make install
```

# Use
```shell
j2y (json to yaml) is a tool for converting JSON data (from files or direct input) to YAML format.

Examples:
1.Convert from file: 
    j2y -f input.json -o output.yaml
2.Convert from command line JSON data: 
    j2y -i '{"key": "value"}' -o output.yaml
3.Convert from standard input: 
    echo '{"foo":1}' | j2y
4.Specify input file only, output to stdout: 
    j2y -f input.json
5.Specify JSON data only, output to stdout: 
    j2y -i '{"key": "value"}'

Usage:
  Usage: j2y [-f inputFile | -i jsonData] [output] [flags]

Flags:
  -f, --file string     Input JSON file path
  -h, --help            help for Usage:
  -i, --input string    Input JSON data
  -o, --output string   Output YAML file (optional)
```

# Use libraries
- github.com/spf13/cobra(https://github.com/spf13/cobra)
- gopkg.in/yaml.v2(https://gopkg.in/yaml.v2)

# Set the help language
Controlled by the environment variable J2Y_LANG. If not set, LANG is used. If no corresponding language is found, en is used.
