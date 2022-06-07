# Setup 
Setting up archgst is pretty straightforward. It requires three things to be done:
1. Install archgst.
1. Initialize the configuration. 
2. Start the parser. 

## Installing archgst
In order to install archgst you are required to have [Go 1.16+](https://golang.org/dl/) installed on your machine. Once you have it, the first thing to do is to clone the GitHub repository. To do this you can run

```shell
git clone https://github.com/giansalex/archway-gastracker.git
```

Then, you need to install the binary. To do this, run 

```shell
make install
```

This will put the `archgst` binary inside your `$GOPATH/bin` folder. You should now be able to run `archgst` to make sure it's installed: 

```shell
archgst version
```

## Initializing the configuration
In order to correctly parse and store the data based on your requirements, archgst allows you to customize its behavior via a YAML file called `config.yaml`. In order to create the first instance of the `config.yaml` file you can run

```shell
archgst init
```

This will create such file inside the `~/.archgst` folder.  
Note that if you want to change the folder used by archgst you can do this using the `--home` flag: 

```shell
archgst init --home /path/to/my/folder
```

Once the file is created, you are required to edit it and change the different values. To do this you can run 

```shell
nano ~/.archgst/config.yaml
```

For a better understanding of what each section and field refers to, please read the [config reference](https://github.com/forbole/juno/blob/v2/cosmos-stargate/.docs/config.md). 

## Running archgst 
Once the configuration file has been setup, you can run archgst using the following command: 

```shell
archgst parse
```

If you are using a custom folder for the configuration file, please specify it using the `--home` flag: 


```shell
archgst parse --home /path/to/my/config/folder
```

We highly suggest you running archgst as a system service so that it can be restarted automatically in the case it stops. To do this you can run: 

```shell
sudo tee /etc/systemd/system/archgst.service > /dev/null <<EOF
[Unit]
Description=archgst parser
After=network-online.target

[Service]
Type=simple
User=$USER
ExecStart=$GOPATH/bin/archgst start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF
```

Then you need to enable and start the service:

```shell
sudo systemctl enable archgst
sudo systemctl start archgst
```
