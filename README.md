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

- For ease of finding environment variable, path..., locate your Apache main folder and run a command on all files inside of directory to find the desired variable:
```sh
grep -iR <variable_to_search> <dir_name> [| wc -l] # include wc -l if want to count number of words
# The Regex syntax is same as using Regex in Javascript
```

### Creating first website

- The site that we saw earlier pointing to our localhost is an basic site, and it also locate in DocumentRoot path.

- Let create a new directory ```cgi```(any name will works) to store out new site: ```/var/html/gci``` and a new html file inside it, ```mysite.html``` including in this repository.

- Now, lets look at Virtual Host, which help Apache handles incoming requests and have multiple sites running on the same server. We will start by setting up Virtual Host configuration files.

- The ```/etc/apache2/sites-available/000-default.conf``` is the default configuration file for VirtualHost, and we can use that as a base for our new directory hosting.
```sh
sudo cp 000-default.conf gci.conf
# edit conf file with your favorite editor: vim, nano, gedit...
nano gci.conf
```

- Respectively modfied the ```ServerAdmin```(Should be Your Email), ```DocumentRoot``` point to where our site files hosted, and the ```ServerName```. There maybe more to configurate... but those alone should be enough.

```
# .... #
ServerAdmin admin.email@name.org
# .... #
DocumentRoot /var/www/gci/
# .... #
ServerName gci.mysite.com # routing of sitename based on this
# .... #
```

- Next, activate the previously configured VirtualHost file by running these command inside the directory where your config file located.
```sh
sudo a2ensite gci.conf
# output ...
sudo systemctl restart apache2
```

- Further Virtual Host Configuration can be found on [apache.org](https://httpd.apache.org/docs/2.4/vhosts/examples.html)

### References
1. https://ubuntu.com/tutorials/install-and-configure-apache
