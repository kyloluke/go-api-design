package mail

// Driver interface 为我们后续扩展使用其他发送邮件的渠道提供了方便。
// 发送邮件可不止 SMTP 一种方式。常见的还有 SAAS 服务提供的 HTTP Mail API，如 sendcloud.sohu.com/home ，sendgrid.com/ 。

// 本课程使用 SMTP Driver

type Driver interface {
	Send(email Email, config map[string]string) bool
}
