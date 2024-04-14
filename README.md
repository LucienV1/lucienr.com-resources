# lucienr.com-resources
resources for lucienr.com

## usage
### server
clone

move to directory

build with go

move executable to /usr/bin/local/

move service file /etc/systemd/system/

start service

```
git clone https://github.com/LucienV1/lucienr.com-resources
cd lucienr.com-resources
go build
sudo mv randomcss /usr/bin/local/randomcss
sudo mv lucienr.com.service /etc/systemd/system/lucienr.com.service
sudo systemctl start lucienr.com.service
```
