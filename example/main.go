package main

import (
	"flag"
	"fmt"
	"github.com/mm-ooto/awsEmailSmsSend/config"
	"github.com/mm-ooto/awsEmailSmsSend/utils"
)

var (
	runMode string
)

func init() {
	flag.StringVar(&runMode, "runMode", "dev", "managed environment")
	flag.Parse()
}

func main() {
	config.LoadConfig(runMode)
	confData := config.ConfigData
	if err := sendEmail(confData); err != nil {
		fmt.Printf("sendEmail err:%s\n", err.Error())
		return
	}
	if err := sendSns(confData); err != nil {
		fmt.Printf("sendSns err:%s\n", err.Error())
		return
	}
}

func sendEmail(confData *config.Config) (err error) {
	newEService := utils.NewEmailService(&utils.AwsConfig{
		Region:      confData.Email.Region,
		AccessKeyId: confData.Email.AccessKeyId,
		SecretKey:   confData.Email.SecretKey,
	})
	if err = newEService.SendEmail(&utils.EmailData{
		Recipient: "456789@163.com",
		Body:      "xxxx_body_xxxx",
		Subject:   "xxxx_subject_xxxx",
		Sender:    "123456@163.com",
		CharSet:   "utf-8",
	}); err != nil {
		return
	}
	return
}

func sendSns(confData *config.Config) (err error) {
	newSnsService := utils.NewSnsService(&utils.AwsConfig{
		Region:      confData.Email.Region,
		AccessKeyId: confData.Email.AccessKeyId,
		SecretKey:   confData.Email.SecretKey,
	})
	if err = newSnsService.SendSns(&utils.SnsData{
		Recipient: "456789@163.com",
		Body:      "xxxx_body_xxxx",
	}); err != nil {
		return
	}
	return
}
