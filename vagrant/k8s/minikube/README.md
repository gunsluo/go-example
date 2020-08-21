# Command Line


```
vagrant plugin install vagrant-vbguest

vagrant init centos/7
```

modify Vagrantfile
```
  config.vm.define "minikube" do |node|
    node.vm.box = "centos/7"
    node.vm.hostname = "minikube"

    node.vm.provider "virtualbox" do |vb|
        vb.memory = "2048"
        vb.cpus = 2
    end
    node.vm.network "private_network", ip: "192.168.10.2"
    node.vm.synced_folder "./share", "/share"
  end
```

```
mkdir share

vagrant up
```

```
vagrant status

Current machine states:

minikube                  running (virtualbox)
```

```
vagrant ssh

sudo su

yum install -y net-tools

yum install -y telnet
```

```
vagrant suspend
vagrant resume

vagrant halt

vagrant destroy
```
