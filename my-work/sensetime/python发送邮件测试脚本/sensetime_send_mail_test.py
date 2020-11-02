#!/usr/bin/env python
# -*- coding: utf-8 -*-

import smtplib
from email.header import Header
from email.mime.text import MIMEText

# mail_host = "smtp.163.com"
mail_host = "notice.sensetime.com"
# mail_host = "smtp.partner.outlook.cn"
# mail_user = "wangyecheng465@163.com"
mail_user = "idea.notify@notice.sensetime.com"
# mail_user = "wangyecheng_vendor@sensetime.com"
mail_pass = "99gUWS96dz"
# mail_pass = "1xxlxx^"

sender = 'idea.notify@notice.sensetime.com'
# sender = 'wangyecheng_vendor@sensetime.com'
receivers = ['wangyecheng_vendor@sensetime.com']

content = 'Python test to send mail'
title = 'Test'


def sendEmail():
    message = MIMEText(content, 'plain', 'utf-8')
    message['From'] = "{}".format(sender)
    message['To'] = ",".join(receivers)
    message['Subject'] = title

    try:
        # smtpObj = smtplib.SMTP_SSL(mail_host, 465)
        smtpObj = smtplib.SMTP_SSL(mail_host, 587)
        smtpObj.login(mail_user, mail_pass)
        smtpObj.sendmail(sender, receivers, message.as_string())
        print("mail has been send successfully.")
    except smtplib.SMTPException as e:
        print(e)


def send_email2(SMTP_host, from_account, from_passwd, to_account, subject, content):
    # email_client = smtplib.SMTP(SMTP_host)
    email_client = smtplib.SMTP(SMTP_host, 587)

    email_client.ehlo()  # 向Gamil发送SMTP 'ehlo' 命令
    email_client.starttls()

    email_client.login(from_account, from_passwd)
    # create msg
    msg = MIMEText(content, 'plain', 'utf-8')
    msg['Subject'] = Header(subject, 'utf-8')  # subject
    msg['From'] = from_account
    msg['To'] = to_account
    # email_client.sendmail(from_account, to_account, msg.as_string())
    email_client.sendmail(from_account, to_account, "hi")
    email_client.quit()

    print("successful send mail")


if __name__ == '__main__':
    # sendEmail()
    send_email2(mail_host, mail_user, mail_pass, receivers, title, content)
