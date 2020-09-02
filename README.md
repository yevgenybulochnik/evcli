你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Evcli

Go command line utility to manage libvirt resources.

### Basic Usage

Please make sure `LIBVIRT_DEFAULT_URI` env var is set to `qemu:///system`. Also for ubuntu install `libvirt-bin`, `libvirt-dev`, `qemu-kvm`. 
In order to run commands one can either build the binary for this project with `go build .` at the project root. Or you can execute the commands with `go run . <your command>`. Executing `go run .` at the project root, for example, will give you the cli help message. Get further help by typing in specific commands, ie `go run . download`.
1. Generate 2 storage pools, one for images and one for vms. I typically do this in my home directory
    - Create an images and vms dir at `~`
    - Execute `evcli create pool images -p /home/<your-username>/images `
    - Execute `evcli create pool vms -p /home/<your-username>`
2. Download a cloud image, currently only ubuntu and centos images are available.
    - Execute `evcli download ubuntu1804`
    - This command will prompt you to create a `.evcli` config dir at `~`. When an image downloads a profile will automatially be generated at `~/.evcli/profiles.yaml`
3. Create a new vm
    - Execute `evcli create vm <vm-name> --profile ubuntu1804`
4. List all vms
    - Execute `evcli list vms`
    - Wait for your newly created vm to have an ip address
5. Once an ip has been generated ssh into your newly created vm with `evcli ssh <your-vm-name>`
