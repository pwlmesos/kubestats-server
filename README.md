# kubestats-server

Server to get stats on Kubenetes cluster


## GO Help for newbs

Hit `Crtl+C` not `Crtl+x` else
to kill a process on Mac

```bash
$> lsof -i tcp:8080
COMMAND   PID         USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
server  49011 patricklogan    6u  IPv6 0xf148eca9f4ef31a9      0t0  TCP *:http-alt (LISTEN)

$> sudo kill -9 49011
fish: Job 1, './server' terminated by signal SIGKILL (Forced quit)
```
