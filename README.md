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
Let's get up and running with vagrant with some demos.

####TL;DR
For the impatient, you can get a single machine running by running:

```bash
vagrant init hashicorp/precise64
vagrant up
```

This will create a default Vagantfile using an Ubuntu based `box` from Altas.
Running `vagrant up` in the same location as the Vagrantfile will create a local VirtualBox Machine.

##License

Copyright (c) <2016> Lee Chang


Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
