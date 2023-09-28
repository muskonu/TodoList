package email

import (
	"TodoList/model"
	"bytes"
	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
	"html/template"
	"net/smtp"
)

func SendWarning(receiver string, todo *model.TodoList) error {
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱

	em.From = "TodoList <1706459198@qq.com>"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{receiver}

	// 设置主题

	em.Subject = "你有待办事项未完成！！！！！"

	data := "您在 " + todo.DueDate.Format("2006.01.02 15:04:05") + " 还存在代办事项：" + todo.Content
	a := template.New("email")
	parse, err := a.Parse(tmp)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = parse.Execute(&b, data)
	if err != nil {
		return err
	}

	em.HTML = b.Bytes()

	//设置服务器相关的配置
	err = em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1706459198@qq.com", "ywpbbkfisdbxdibi", "smtp.qq.com"))
	if err != nil {
		logx.Error()
		return err
	}
	return nil
}

var tmp = `<meta http-equiv="Content-Type" content="text/html;charset=utf-8">
<div>
    <includetail>
        <div align="center">
            <div class="open_email" style="margin-left: 8px; margin-top: 8px; margin-bottom: 8px; margin-right: 8px;">
                <div>
                    <br>
                    <span class="genEmailContent">
                        <div id="cTMail-Wrap"
                             style="word-break: break-all;box-sizing:border-box;text-align:center;min-width:320px; max-width:660px; border:1px solid #f6f6f6; background-color:#f7f8fa; margin:auto; padding:20px 0 30px; font-family:'helvetica neue',PingFangSC-Light,arial,'hiragino sans gb','microsoft yahei ui','microsoft yahei',simsun,sans-serif">
                            <div class="main-content" style="">
                                <table style="width:100%;font-weight:300;margin-bottom:10px;border-collapse:collapse">
                                    <tbody>
                                    <tr style="font-weight:300">
                                        <td style="width:3%;max-width:30px;"></td>
                                        <td style="max-width:600px;">
                                            <div id="cTMail-logo" style="width:92px; height:25px;">
                                            </div>
                                            <p style="height:2px;background-color: #00a4ff;border: 0;font-size:0;padding:0;width:100%;margin-top:20px;"></p>

                                            <div id="cTMail-inner" style="background-color:#fff; padding:23px 0 20px;box-shadow: 0px 1px 1px 0px rgba(122, 55, 55, 0.2);text-align:left;">
                                                <table style="width:100%;font-weight:300;margin-bottom:10px;border-collapse:collapse;text-align:left;">
                                                    <tbody>
                                                    <tr style="font-weight:300">
                                                        <td style="width:3.2%;max-width:30px;"></td>
                                                        <td style="max-width:480px;text-align:left;">
                                                            <h1 id="cTMail-title" style="font-size: 20px; line-height: 36px; margin: 0px 0px 22px;">
                                                                您有待办事项未完成！
                                                            </h1>

                                                            <p class="cTMail-content" style="line-height: 24px; margin: 6px 0px 0px; overflow-wrap: break-word; word-break: break-all;">
                                                                <span style="color: rgb(51, 51, 51); font-size: 14px;">
                                                                    <span style="font-weight: bold;">{{.}}</span>
                                                                </span>
                                                            </p>

                                                            <dl style="font-size: 14px; color: rgb(51, 51, 51); line-height: 18px;">
                                                                <dd style="margin: 0px 0px 6px; padding: 0px; font-size: 12px; line-height: 22px;">
                                                                    <p id="cTMail-sender" style="font-size: 14px; line-height: 26px; word-wrap: break-word; word-break: break-all; margin-top: 32px;">

                                                                        <br>
                                                                        <strong>TodoList</strong>
                                                                    </p>
                                                                </dd>
                                                            </dl>
                                                        </td>
                                                        <td style="width:3.2%;max-width:30px;"></td>
                                                    </tr>
                                                    </tbody>
                                                </table>
                                            </div>

                                        </td>
                                        <td style="width:3%;max-width:30px;"></td>
                                    </tr>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </span>
                    <br>
                </div>
            </div>
        </div>
    </includetail>
</div>
`
