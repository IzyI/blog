---
title: "Китайский SOCKS5 (shadowsocks)"
date: 2020-04-09T18:05:01+03:00
type: it
draft: false
linkhabr: false
description: shadowsocks socks5 proxy go
image: default_open_graph.jpg
tags: [proxy]
---
Писал я как то одну программу благодаря ей, я  познакомился с очень удобной системой проксей под названием Shadowsocks.

И был я очень удивлен его открытостью, скоростью , простотой работы и большим количеством языковых
реализаций под всевозможные платформы.  Shadowsocks уже несколько лет работает в Китае и не плохо показал себя в борьбе за
контент против великого китайского фаервола.

<!--more-->
Для ознакомления с проектом можно пройти по ссылке: <a href="https://shadowsocks.org/en/index.html" target="_blank">ТыЦ</a>

Сейчас же мне резко понадобился прокси для моих мелочей. И я решил что это отличный повод рассказать как его поднимать.


<br><br>
<h3 class="text-center">Сервер</h3>


Я выбрал  реалзцию Outline-ss-server сервера shadowsocks:
 <a href="https://github.com/Jigsaw-Code/outline-ss-server" target="_blank">ТыЦ</a> 
написанный на go.

Выбрал я его так как у него есть уже настроенная реализация сбора метрики <em>Prometheus</em>. Да и в целом почему бы и нет.
Устанавливаем командой: 
{{< highlight bash  >}}
GO111MODULE=off go get github.com/shadowsocks/go-shadowsocks2 github.com/prometheus/prometheus/cmd/...
{{< /highlight >}}

Если же вы <em>хотеть</em> проще и быстрее то можете установить сишную реализацию shadowsocks-libev
<a href="https://github.com/Jigsaw-Code/shadowsocks-libev" target="_blank">ТыЦ</a>

Теперь нам надо собрать outline-ss-server . для этого надо пройти в папку где у вас установлен сервер в моем случае это
<em class="redic">~/work/src/github.com/Jigsaw-code/outline-ss-server</em> и вызвать <em class="redic">go build .</em>

После надо перейти в папку где у вас происходит сборка программы и сделать символьную ссылку
<em class="redic">ln -s /root/work/bin/outline-ss-server /usr/local/bin/</em>.

Также нам понадобится конфиг файл config.yml:

{{< highlight  yaml >}}
keys:
    - id: user-0
      port: 8388
      cipher: chacha20-ietf-poly1305
      secret: 45hgwtregy353h

{{< /highlight >}}
<br/>

Пишем service чтобы автоматически поднимался наш сервис:
 
{{< highlight service  >}}
[Unit]
Description=shadowsocks_run
After=network.target
After=syslog.target

[Service]
WorkingDirectory=/root/work/bin
ExecStart=outline-ss-server -config /root/work/bin/config.yml
TimeoutStopSec=5
PrivateTmp=true
Restart=always

[Install]
WantedBy=multi-user.target
{{< /highlight >}}
<br/>

Затем надо вызвать <em class="redic">systemctl daemon-reload</em> и можно запускать <em class="redic">service shadowsockrun start</em>.

<br><br>
<h3 class="text-center">Клиент</h3>

Теперь мы можем подключаться свой laptop через клиент. В качестве клиента я взял shadowsocks-libev.
Устанавливаем и правим конфиг
<em class="redic">sudo nano /etc/shadowsocks-libev/config.json </em>


{{< highlight json  >}}
{
    "server": "IP_удаленного_сервера",
    "mode":"tcp_and_udp",
    "server_port":8388,
    "local_port":1081,
    "password":"45hgwtregy353h",
    "timeout":60,
    "method":"chacha20-ietf-poly1305"
}
{{< /highlight >}}
<br/>

И вызываем ss-local. вывод должен быть такой:
{{< highlight bash  >}}
2020-04-09 20:45:58 INFO: initializing ciphers... chacha20-ietf-poly1305
2020-04-09 20:45:58 INFO: listening at 127.0.0.1:1081
2020-04-09 20:45:58 INFO: udprelay enabled
{{< /highlight >}}
<br/>

Туннель поднялся теперь мы можем проверить его нашим тестовым скриптом на python (не забываем установить pysocks)

{{< highlight python  >}}
import requests

ip = '127.0.0.1'
port = "1081"
auth = ""
proxies = {
    "http": f"socks5://{auth}{ip}:{port}",
    "https": f"socks5://{auth}{ip}:{port}",
}
print(requests.get("http://ifconfig.me").text)
print(requests.get("http://ifconfig.me", proxies=proxies,timeout=50).text)
{{< /highlight >}}
