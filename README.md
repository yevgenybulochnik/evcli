# Evcli

Go command line utility to manage libvirt resources.

### Basic Usage

Please make sure `LIBVIRT_DEFAULT_URI` env var is set to `qemu:///system`. 
In order to run commands one can either build the binary for this project with `go build .` at the project root. Or you can execute the commands wih `go run . <your command>`. Executing `go run .` at the project root, for example, will give you the cli help message. Get further help by typing in specific commands, ie `go run . download`.
1. Generate 2 storage pools, one for images and one for vms. I typically do this in my home directory
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
