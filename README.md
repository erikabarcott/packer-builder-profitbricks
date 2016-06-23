# Profitbricks packer builder

This builder plugin extends packer.io to support building images for ProfitBricks. 

You can check out packer [here](https://packer.io).


## Dependencies
* Packer >= v0.10.1 (https://packer.io)
* Golang (tested with 1.6) 
* Godep >= v62


## Install Go

Follow these instructions and install golang on your system:
* https://golang.org/doc/install

## Install Packer

Follow instructions https://www.packer.io/intro/getting-started/setup.html

## Compile the plugin

Once you have installed Packer, you must compile this plugin and install the
resulting binary.

```shell
go get https://github.com/profitbricks/packer-builder-profitbricks
cd $GOPATH/src/github.com/profitbricks/packer-builder-profitbricks
make install
```

If the build is successful, you should now have `packer-builder-profitbricks`
binary in your `~/.packer.d/plugins` if you are using Linux or Mac or on Windows `%APPDATA%/packer.d/plugins` directory and you are
ready to get going with packer.

## Download the plugin

Alternatively you can download prebuilt binaries from https://github.com/profitbricks/packer-builder-profitbricks/releases/tag/v1.0.0. Once you have downloaded binary for your distribution. Place the binary to  `~/.packer.d/plugins` if you are using Linux or Mac or on Windows `%APPDATA%/packer.d/plugins` directory.

## Example

Once you've setup the above, you are good to go with an example.
To get a quick start run:

```shell
cd $GOPATH/src/github.com/profitbricks/packer-builder-profitbricks
```

There you will find example `config.json`

```
{
  "builders": [
    {
      "image": "Ubuntu-16.04",
      "pbusername": "pb_username",
      "pbpassword": "pb_password",
      "servername": "packer",
      "type": "profitbricks"
    }
}
```

To validate `config.json` run:

```shell
packer validate config.json
```

or if you want to get suggestions on how to fix your config run:

```shell
packer fix config.json
```

To build a Profitbricks packer image run: 

```shell
packer build config.json

==> profitbricks: Creating temporary SSH key for instance...
==> profitbricks: Creating Virutal datacenter...
==> profitbricks: Creating Profitbricks server...
==> profitbricks: Creating a volume
==> profitbricks: Creating a LAN
==> profitbricks: Creating a NIC
==> profitbricks: Done!
==> profitbricks:
==> profitbricks: Waiting for SSH to become available...
==> profitbricks: Connected to SSH!
==> profitbricks: Creating Profitbricks snapshot...
==> profitbricks: Removing Virtual Data Center
Build 'profitbricks' finished.


==> Builds finished. The artifacts of successful builds are:
--> profitbricks: A snapshot was created: 'packerSnapshot'

```

## Available config parameters

Requied parameters:

```shell
Builder type - "type" 
Profitbricks username - "pbusername"
Profitibricks passwrod - "pbpassword"
Profitbricks REST Url - "pburl"
Server name - "servername"
Profitbricks volume image - "image"
```

Optional parameters:

```shell
Profitbricks region default value "us/las" - "region"
Desired disk size default value 50gb - "disksize"
Desired disk type default value "HDD" - "disktype"
Number of server cores default value 4 - "cores"
RAM size for the server default value "2048" - "ram"
```

## Support

You are welcome to contact us with questions or comments at [ProfitBricks DevOps Central](https://devops.profitbricks.com/). Please report any issues via [GitHub's issue tracker](https://github.com/profitbricks/docker-machine-driver-profitbricks/issues).