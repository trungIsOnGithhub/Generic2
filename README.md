##### Home Lab practicing Apache Web Server on Linux

> This lab was operated on Ubuntu Desktop 22.4

### Installation and Download:

- As usual, update and upgrade all packages
```sh
sudo apt update
sudo apt upgrade
```
-  Install Apache Server
```sh
sudo apt install apache2
```
- To check if Apache Web systemd has been setup and running:
```sh
sudo systemctl status apache2
# if not already started
# sudo systemctl restart apache2
```

![Ubuntu localhost is now Apache Web Server](/assets/apache-running-local.png)

You can check your IP Adress with terminal command, looking for:
- LOOPBACK(*lo: tag*) inet(IPv4)
- BOARDCAST(*wl: tag*) inet(IPv4)
then type it to brower
```sh
ifconfig
# or
ip a
```

### Finding and Configurating

- Default configuration files directory of Apache in Linux distributions is ```/etc/apache2``` for Ubuntu/Debian or ```/etc/httpd``` for CentOS/RHEL/Fedora, there are some useful configuration files with ```.conf``` extension by useful directories here.

> ```etc``` is folder contain all system configuration files, full-name etcetera back to the early day of UNIX

- With CentOS/RHEL/Fedora, finding and modify your RootDocument variable in ```/etc/httpd/httpd.conf | etc/httpd/ssl.conf``` use the command below
```sh
grep -i '*Root' /etc/httpd/conf/httpd.conf
grep -i '*Root' /etc/apache2/*.conf
```

- With Ubuntu/Debian, suddenly there are none of the above! The default document folder is ```/var/www/``` and the main cofiguration file is in ```apache2.conf | sites-enabled/000-default```