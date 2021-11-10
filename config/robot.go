/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmdutils

import (
	"log"
	"strings"

 	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	MSGTYPE_TEXT     = "text"
	MSGTYPE_MARKDOWM = "markdown"
)

type RobotConf struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content       string   `json:"content"`
		MentionedList []string `json:"mentioned_list"`
		//MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
	Webhook string `json:"webhook"`
	Member  []struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"member"`
}

// robotCmd represents the robot command
var robotCmd = &cobra.Command{
	Use:   "robot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置文件路径
		conf, err := cmd.Flags().GetString("conf")
		if err != nil {
			log.Fatalf("get conf err:%v", err)
		}

		// 读取配置
		var robotConf RobotConf
		viper.SetConfigFile(conf)
		err = viper.ReadInConfig()
		if err != nil {
			log.Fatalf("read config failed: %v", err)
		}
		err = viper.Unmarshal(&robotConf)
		if err != nil {
			log.Fatalf("Unmarshal config failed: %v", err)
		}

		// viper无法解析带下划线的key, 手动赋值处理一下
		robotConf.Text.MentionedList = viper.GetStringSlice("text.mentioned_list")

		// 发送内容
		// 获取所有成员信息，筛出值日生
		var name, code string
		var member []string
		for _, user := range robotConf.Member {
			userName := strings.TrimPrefix(strings.TrimSuffix(user.Name, "】"), "【")
			if len(userName) != len(user.Name) {
				name = user.Name
				code = user.Code
			}
			member = append(member, user.Name)
		}

		var sendMsg RobotConf
		switch robotConf.Msgtype {
		case MSGTYPE_TEXT:
			sendMsg = RobotConf{
				Msgtype: MSGTYPE_TEXT,
				Text:    robotConf.Text,
			}

			sendMsg.Text.Content = strings.ReplaceAll(sendMsg.Text.Content, "{member}", strings.Join(member, "，"))
			sendMsg.Text.Content = strings.ReplaceAll(sendMsg.Text.Content, "{name}", name)
			sendMsg.Text.Content = strings.ReplaceAll(sendMsg.Text.Content, "{code}", code)

			for i, s := range sendMsg.Text.MentionedList {
				sendMsg.Text.MentionedList[i] = strings.ReplaceAll(s, "{code}", code)
			}
		case MSGTYPE_MARKDOWM:
			sendMsg = RobotConf{
				Msgtype:  MSGTYPE_MARKDOWM,
				Markdown: robotConf.Markdown,
			}

			sendMsg.Markdown.Content = strings.ReplaceAll(sendMsg.Markdown.Content, "{member}", strings.Join(member, "，"))
			sendMsg.Markdown.Content = strings.ReplaceAll(sendMsg.Markdown.Content, "{name}", name)
			sendMsg.Markdown.Content = strings.ReplaceAll(sendMsg.Markdown.Content, "{code}", code)
		default:
			return
		}

		err = pkg.PostJson(robotConf.Webhook, sendMsg, nil)
		if err != nil {
			log.Fatalf("PostJson failed: %v", err)
		}

		// 重置配置文件
		memberLen := len(robotConf.Member)
		if memberLen > 0 {
			for i, user := range robotConf.Member {
				userName := strings.TrimPrefix(strings.TrimSuffix(user.Name, "】"), "【")
				if len(userName) != len(user.Name) {
					next := i + 1
					if next >= memberLen {
						next = 0
					}

					robotConf.Member[i].Name = userName
					robotConf.Member[next].Name = "【" + robotConf.Member[next].Name + "】"

					break
				}
			}

			viper.Set("member", robotConf.Member)
			err = viper.WriteConfigAs(conf)
			if err != nil {
				log.Fatalf("WriteConfigAs failed: %v", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(robotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// robotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// robotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	robotCmd.Flags().StringP("conf", "c", "", "config file")
	err := robotCmd.MarkFlagRequired("conf")
	if err != nil {
		log.Fatalf("robot init err: %v", err)
	}
}
