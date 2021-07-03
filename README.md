# unsu
unsu - un-sudo your last command

### why?
When switching between unprivileged and priviliged users in linux I find myself either forgetting to use sudo or using sudo when I'm not able to.

It's quick and easy to to fix an un-sudo'd command:

`sudo !!`

But I've not found (or not googled) a way to do the reverse, remove sudo from the last command.

### how to use
 - Download unsu.py as `unsu`
 - Move it to somewhere in `echo $PATH`
 - `chmod +x unsu # make it executable`
 - Restart your shell session
 - And then:
```
sudo supervisorctl status
[sudo] password for unprivuser:
unsu !!
my-app                         RUNNING   pid 3300370, uptime 2 days, 21:20:07
``` 

### Enjoy
If you use this, I hope it brings you some relief to a semi-regular issue like mine!
Also I plan to repeat this exercise in other languages, python's just easiest!
