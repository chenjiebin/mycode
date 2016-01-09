#!/usr/bin/expect -f

set user root
set host 172.16.27.135
set timeout 120

spawn ssh $user@$host

send "tar -xf /root/go/gaceful-linux64.tar.bz -C /root/go\r"

send "chown -R root:root /root/go/gaceful-linux64\r"
send "chmod a+x /root/go/gaceful-linux64\r"
send "curl 'http://172.16.27.135:10000/goyaf_upgrade'\r"
send "exit\r"

expect eof