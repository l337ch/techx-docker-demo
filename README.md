#Vagrant and Docker TechX Demo
Using vagrant and docker in your development environment.

__NOTE__ The following source material is target at Windows users. YMMV if you attempt to try this with any other system.

## Installing
###Prerequisites
To run the examples in this demo, you will to install the following:
- Ruby - http://rubyinstaller.org/
- Git with git-bash - https://git-scm.com/download/win
- VirtualBox - https://www.virtualbox.org/wiki/Downloads
- Text editor such as VIM or EMACS or Notepad ++

### Install Vagrant
The Vagrant installation package can be found at https://www.vagrantup.com/downloads.html. Run the install package after it is downloaded, accepting the default in the process.

### Install Docker Machine
Run the following to install Docker Machine on a Windows machine.

```bash
if [[ ! -d "$HOME/bin" ]]; then mkdir -p "$HOME/bin"; fi && \
curl -L https://github.com/docker/machine/releases/download/v0.6.0/docker-machine-Windows-x86_64.exe > "$HOME/bin/docker-machine.exe" && \
chmod +x "$HOME/bin/docker-machine.exe"
```

##Vagrant
Let's get up and running with vagrant with including some demos.

####TL;DR
For the impatient, you can get a single machine running by running:

```bash
vagrant init ubuntu/trusty64
vagrant up
```

This will create a default Vagantfile using an Ubuntu based `box` from Altas.
Running `vagrant up` in the same location as the Vagrantfile will create a local VirtualBox Machine.
Login into the machine to start driving your new machine with `vagrant ssh`.


###Running with Vagrant

####Sharing folders
Sharing folders between the local host and guest machine can be accomplished by enabling _synced_folders_ .
To enable `synced_folders`, modify the 'config.vm.synced_folder' of the Vagrantfile with:

```bash
config.vm.synced_folder "../../data", "/vagrant_data"
```

Run `vagrant reload` to see the synced folder added to the machine. This will restart the machine and pick up any new Vagrantfile changes.

####Networking
If services need to be reached over the network, use _config.vm.network_ to add an IP address to the machine
Add `config.vm.network "private_network", ip: "172.16.1.100"` under the `config.vm.box` line.
Reload the machine with `vagrant reload`

####Provisioning
Provisioning with _config.vm.provision_ can be used to install software software, alter configurations, and to run arbitrary scripts.
For demo purposes, edit the Vagrantfile and uncomment the following lines.

```bash
config.vm.provision "shell", inline: <<-SHELL
  sudo apt-get update
  sudo apt-get install -y apache2
SHELL
```

Run `vagrant provision` to have the changes reflected on the machine

###Multi Machine Vagrant

##Docker Machine

###Provisioning Docker Engine
Using Docker Machine, install Docker Engine on the machine that we built previously.
Run

```bash
docker-machine create -d generic \
--generic-ssh-user vagrant \
--generic-ssh-key .vagrant/machines/default/virtualbox/private_key \
--generic-ip-address 172.16.1.100 \
--engine-install-url "https://get.docker.com" \
default
```

##License

Copyright (c) 2016 Lee Chang


Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
