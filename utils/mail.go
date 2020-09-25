package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/gomail.v2"
)

func SendMail(email, reciver string) error {

	//设置随机码
	randCode := RandCode("0123456789", 5)

	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return err
	}
	defer r.Close()
	_, err = r.Do("SET", email, randCode, "EX", "300") //生存时间为5分钟
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	m := gomail.NewMessage()
	//发件人
	m.SetHeader("From", "邮箱号")

	//收件人
	m.SetHeader("To", reciver)

	//抄送人
	//m.SetAddressHeader("Cc", "test@126.com", "test")

	//邮件标题
	m.SetHeader("Subject", "blog验证码")

	//邮件内容
	m.SetBody("text/html", randCode)

	//邮件附件
	//m.Attach("C:\\Users\\User\\Pictures\\Saved Pictures\\1.jpg")
	d := gomail.NewDialer("smtp.qq.com", 465, "邮箱号", "授权码")


	//邮件发送服务器信息,使用授权码而非密码
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
