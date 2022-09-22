#!/bin/bash
goPATH="/home/vagrant/.go"
goROOT="/usr/local/go"
export DEBIAN_FRONTEND=noninteractive
workDir="/temp/work"
Cyan='\033[0;36m'
Color_Off='\033[0m'
sudo mkdir -p /temp
sudo mkdir -p /home/vagrant/.go
installNode () {
  echo "...$(Cyan)Installing Nodejs..."
  curl -fsSL https://deb.nodesource.com/setup_16.x -o /tmp/nodesource_setup.sh
  sudo bash /tmp/nodesource_setup.sh
  apt-get install -y nodejs
  npm install -g yarn
  echo "...Install Nodejs successfully!..."
}

installGo () {
  echo "...Installing Golang..."
  wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz
  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz
  echo "...writing env golang to bashrc..."
  sudo cat << EOF >> /home/vagrant/.bashrc
export GOPATH=$goPATH
export GOROOT=$goROOT
export PATH=$PATH:$goROOT/bin:$goPATH/bin
EOF
  cat /home/vagrant/.bashrc
  source /home/vagrant/.bashrc
  go env
  echo "....Install Golang successfully!..."
}

installMysql () {
  echo "Installing mysql"

  wget -c https://dev.mysql.com/get/${MYSQL_8_file}
  sudo -E dpkg -i ${MYSQL_8_file}
  sudo apt update
  sudo -E apt install -y mysql-server
  echo "set up user and password and database default"
  sudo bash -c "mysql << EOF
DROP DATABASE IF EXISTS mydb;
CREATE DATABASE mydb;
DROP USER IF EXISTS 'thaianh'@'localhost';
CREATE USER 'thaianh'@'localhost' IDENTIFIED WITH mysql_native_password BY 'thaianh1711';
GRANT ALL PRIVILEGES ON mydb.* TO 'thaianh'@'localhost' WITH GRANT OPTION;
FLUSH PRIVILEGES;
EOF"
}
installNode
installGo
installMysql
