package mail

import (
	"crypto/tls"
	"io"
	"log"
	"net/smtp"
	"strings"

	"github.com/eynstudio/gobreak"
)

func NewClient(host, user, pwd string) *MailClient {
	return &MailClient{host: host, user: user, pwd: pwd}
}

type MailClient struct {
	host string
	user string
	pwd  string
}

func (p MailClient) SendWithSll(to []string, subject, msg string) error {
	conn, c, err := p.Conn()
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()
	s := &sendMail{c: c}
	s.send(p.user, to, subject, msg)
	return s.Err
}

func (p MailClient) Conn() (conn *tls.Conn, c *smtp.Client, err error) {
	conn, err = tls.Dial("tcp", p.host, nil)
	if err != nil {
		log.Println("Error Dialing", err)
		return
	}

	auth := smtp.PlainAuth("", p.user, p.pwd, p.host)
	c, err = smtp.NewClient(conn, p.host)
	if err != nil {
		log.Println("Error SMTP connection", err)
		return
	}

	if ok, _ := c.Extension("AUTH"); ok {
		if err = c.Auth(auth); err != nil {
			log.Println("Error during AUTH", err)
			return
		}
	}
	return
}

type sendMail struct {
	gobreak.Error
	c *smtp.Client
}

func (p *sendMail) send(user string, to []string, subject, msg string) {
	p.NoErrExec(func() { p.Err = p.c.Mail(user) })
	for _, it := range to {
		p.NoErrExec(func() { p.Err = p.c.Rcpt(it) })
	}

	data := []byte("To: " + strings.Join(to, ";") +
		"\r\nFrom: " + user +
		"\r\nSubject: " + subject +
		"\r\n\r\n" + msg)

	var w io.WriteCloser
	p.NoErrExec(func() { w, p.Err = p.c.Data() })
	p.NoErrExec(func() { _, p.Err = w.Write(data) })
	p.NoErrExec(func() { p.Err = w.Close() })
	p.c.Quit()
}
