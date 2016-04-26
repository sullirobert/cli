package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	Topics = append(Topics, TopicSet{
		{
			Name:   "debug",
			Hidden: true,
			Commands: CommandSet{
				{
					Topic:   "debug",
					Command: "errlog",
					Flags:   []Flag{{Name: "num", Char: "n", HasValue: true}},
					Run: func(ctx *Context) {
						numS, _ := ctx.Flags["num"].(string)
						if numS == "" {
							numS = "30"
						}
						num, err := strconv.Atoi(numS)
						must(err)
						body, err := ioutil.ReadFile(ErrLogPath)
						must(err)
						lines := strings.Split(string(body), "\n")
						start := len(lines) - num - 1
						if start < 0 {
							start = 0
						}
						end := len(lines) - 1
						lines = lines[start:end]
						fmt.Println(strings.Join(lines, "\n"))
					},
				},
			},
		},
	}...)
}
