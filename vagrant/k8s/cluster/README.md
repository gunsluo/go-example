# Command Line


```
vagrant plugin install vagrant-vbguest

vagrant init centos/7
```

modify Vagrantfile
```
  config.vm.define "master" do |node|
    node.vm.box = "centos/7"
    node.vm.hostname = "master"

    node.vm.provider "virtualbox" do |vb|
        vb.memory = "2048"
        vb.cpus = 2
    end
    node.vm.network "private_network", ip: "192.168.10.2"
    node.vm.synced_folder "./share", "/share"
  end

  config.vm.define "node01" do |node|
    node.vm.box = "centos/7"
    node.vm.hostname = "node01"

    node.vm.provider "virtualbox" do |vb|
        vb.memory = "2048"
        vb.cpus = 2
    end
    node.vm.network "private_network", ip: "192.168.10.3"
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
