package main

import (
	"airiacloud/utils"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
	"strings"
	"time"
)

type DDApp struct {
	AccessToken string
	Secret string
	Command string
	Config string
}

type DDNATItem struct {
	Name string `json:"name"`
	Url string `json:"url"`
}
type DDNAT struct {
	Title string `json:"title"`
	Items []DDNATItem `json:"items"`
	Token string `json:"token"`
	Secret string `json:"secret"`
}

func (dd *DDNAT)Load(conf string)error {
	fd, err := os.Open(conf)
	if err != nil {
		fmt.Println("dinggroup [path] ", conf)
		return err
	}
	defer fd.Close()
	// 创建json解码器
	decoder := json.NewDecoder(fd)
	err = decoder.Decode(dd)
	if err != nil {
		fmt.Println("invalid ", err.Error())
		return err
	}
	return  nil
}

func (dd *DDNAT) ToString() string {

	markText := ""
	markText += fmt.Sprintf("**项目** : %s \n\n", dd.Title)
	markText += fmt.Sprintf("**日期** : %s \n\n", time.Now().Format("2006-01-02"))
	if len(dd.Items) > 0 {
		for _, item := range dd.Items {
			markText += fmt.Sprintf("- [%s](%s)\n\n", item.Name, item.Url)
		}
	}
	return markText
}

type DDMessage struct {
	Title string `json:"title"`
	Describe string `json:"describe"`
}

func (dd *DDMessage) ToString() string {
	markText := ""
	markText += fmt.Sprintf("**标题** : %s \n\n", dd.Title)
	markText += fmt.Sprintf("**日期** : %s \n\n", time.Now().Format("2006-01-02 15:04:05"))
	markText += fmt.Sprintf("**描述** : %s \n\n", dd.Describe)
	return markText

}

type DDMonitor struct {
	Title string `json:"title"`
	Describe string `json:"describe"`
	CPU float64
	Memory float64
	DiskUsage []float64
	Disks []string
	Paths string
	NeedReport bool

}

func (dd *DDMonitor) DO() bool {

	dd.Disks=strings.Split(dd.Paths, ",")
	dd.NeedReport = false
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Percent(0, false)

	dd.Memory = v.UsedPercent
	dd.CPU = c[0]

	if dd.Memory > 80 || dd.CPU > 85 {
		dd.NeedReport = true
	}
	for _, p := range dd.Disks {
		d, _ := disk.Usage(p)
		if d.UsedPercent > 85 {
			dd.NeedReport = true
		}
		dd.DiskUsage= append(dd.DiskUsage, d.UsedPercent)
	}
	return dd.NeedReport
}

func (dd *DDMonitor) ToString() string {
	markText := ""
	if dd.NeedReport {
		markText += fmt.Sprintf("**标题** : %s \n\n", dd.Title)
		markText += fmt.Sprintf("**日期** : %s \n\n", time.Now().Format("2006-01-02 15:04:05"))
		markText += fmt.Sprintf("**简述** : %s \n\n", dd.Describe)
		markText += fmt.Sprintf("**CPU** : %f \n\n", dd.CPU)
		markText += fmt.Sprintf("**MEM** : %f \n\n", dd.Memory)
		markText += fmt.Sprintf("**DISK**  \n\n")
		for i, u := range dd.DiskUsage {
			markText += fmt.Sprintf("**%s** : %f \n\n", dd.Disks[i], u)
		}
	}
	return markText

}
func main() {

	//sn := utils.SimpleNotice{}
	dapp := DDApp{}
	//items := ""
	//var receiver Robot
	// nat
	natCmd := flag.NewFlagSet("nat", flag.ExitOnError)
	natCmd.StringVar(&dapp.AccessToken, "token", "6587d40230371eb38fff496113ffdc4500b0100dd7208ef1e779313573f3c430", "token" )
	natCmd.StringVar(&dapp.Secret, "secret", "SEC770b9531b28ba60150c930e964865fbd8d8649e1e9409e65aeb6c4e00aa06bf8", "secret" )
	natCmd.StringVar(&dapp.Config, "config", "config", "path of config")
	// message

	msgCmd := flag.NewFlagSet("message", flag.ExitOnError)
	msgData := DDMessage{}
	msgCmd.StringVar(&dapp.AccessToken, "token", "6587d40230371eb38fff496113ffdc4500b0100dd7208ef1e779313573f3c430", "token" )
	msgCmd.StringVar(&dapp.Secret, "secret", "SEC770b9531b28ba60150c930e964865fbd8d8649e1e9409e65aeb6c4e00aa06bf8", "secret" )
	msgCmd.StringVar(&msgData.Title, "title", "title", "title" )
	msgCmd.StringVar(&msgData.Describe, "describe", "desc", "desc" )

	monitorCmd := flag.NewFlagSet("monitor", flag.ExitOnError)
	monitorData := DDMonitor{}
	monitorCmd.StringVar(&dapp.AccessToken, "token", "6587d40230371eb38fff496113ffdc4500b0100dd7208ef1e779313573f3c430", "token" )
	monitorCmd.StringVar(&dapp.Secret, "secret", "SEC770b9531b28ba60150c930e964865fbd8d8649e1e9409e65aeb6c4e00aa06bf8", "secret" )
	monitorCmd.StringVar(&monitorData.Title, "title", "title", "title" )
	monitorCmd.StringVar(&monitorData.Describe, "describe", "desc", "desc" )
	monitorCmd.StringVar(&monitorData.Paths, "paths", "/", "desc" )

	subCommand := os.Args[1]

	switch subCommand {
	case "nat":
		natCmd.Parse(os.Args[2:])
		data := &DDNAT{}
		if e := data.Load(dapp.Config); e == nil {
			mkdn := data.ToString()
			utils.HookSendMarkdown(data.Title, mkdn, data.Token, data.Secret)
		} else {
			fmt.Println(e.Error())
			os.Exit(1)
		}

	case "monitor":
		monitorCmd.Parse(os.Args[2:])
		if monitorData.DO() {
			utils.HookSendMarkdown(monitorData.Title, monitorData.ToString(), dapp.AccessToken, dapp.Secret)
		}
	case "message":
		msgCmd.Parse(os.Args[2:])
		utils.HookSendMarkdown(msgData.Title, msgData.ToString(), dapp.AccessToken, dapp.Secret)
	default:
		fmt.Println("invalid command ", subCommand)
		os.Exit(1)
	}
}
