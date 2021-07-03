# unsu
unsu - un-sudo your last command

### why?
When switching between unprivileged and priviliged users in linux I find myself either forgetting to use sudo or using sudo when I'm not able to.

It's quick and easy to to fix an un-sudo'd command:

`sudo !!`

But I've not found (or not googled) a way to do the reverse, remove sudo from the last command.

### how to use
 - Download unsu.py
 - Move it to somewhere in `echo $PATH`
 - Restart your shell session
 - And then:
```
sudo supervisorctl status
[sudo] password for unprivuser:
unsu !!
my-app                         RUNNING   pid 3300370, uptime 2 days, 21:20:07
``` 
