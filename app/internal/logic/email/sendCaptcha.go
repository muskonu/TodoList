package email

import (
	"TodoList/common/globalkey"
	"context"
	"github.com/jordan-wright/email"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func NewSendCaptcha(redis *redis.Client) func(receiver string) error {
	return func(receiver string) error {
		em := email.NewEmail()
		// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱

		em.From = "TodoList <1706459198@qq.com>"

		// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
		em.To = []string{receiver}

		// 设置主题

		em.Subject = "您的验证码为："

		// 简单设置文件发送的内容，暂时设置成纯文本
		num := rand.Intn(900000) + 100000
		n := strconv.Itoa(num)
		success, err := redis.SetNX(context.Background(), globalkey.Email(receiver), num, 1*time.Minute).Result()
		if err != nil {
			return err
		}
		if !success {
			return errors.New("验证码未过期")
		}
		em.HTML = []byte("您的验证码为：<h1>" + n + "</h1>")

		//设置服务器相关的配置
		err = em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1706459198@qq.com", "ywpbbkfisdbxdibi", "smtp.qq.com"))
		if err != nil {
			logx.Error()
			return err
		}
		return nil
	}
}
