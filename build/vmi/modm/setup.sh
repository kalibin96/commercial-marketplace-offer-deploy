#!/bin/bash

echo ""

# copy files to the correct location
echo "Copying caddy and docker compose files to $MODM_HOME"
sudo cp /tmp/Caddyfile $MODM_HOME/Caddyfile
sudo cp /tmp/docker-compose.yml $MODM_HOME/docker-compose.yml

echo "Performing git checkout on branch [$MODM_REPO_BRANCH]"
sudo rm -rf $MODM_HOME/source
sudo git clone --depth=1 --branch $MODM_REPO_BRANCH https://github.com/microsoft/commercial-marketplace-offer-deploy.git $MODM_HOME/source
sudo git config --global --add safe.directory $MODM_HOME/source

# build service host and install it
# ----------------------------------
cd $MODM_HOME/source
echo ""
echo "Building ServiceHost"
csproj=./src/ServiceHost/ServiceHost.csproj
out_path=./bin/servicehost
sudo mkdir -p $out_path

sudo dotnet restore $csproj
sudo dotnet build $csproj -c Release -o $out_path/build
sudo dotnet publish $csproj -c Release -o $out_path/publish

# setup daemon
echo "Installing ServiceHost as systemd service."
sudo cp $out_path/publish/modm-servicehost /usr/sbin/modm-servicehost
sudo cp /tmp/modm.service /etc/systemd/system/modm.service

# activate and start
sudo systemctl daemon-reload
sudo systemctl start modm

# support start on boot
sudo systemctl enable modm

# print out status
sudo systemctl status modm


# build final docker images that will represent MODM backend and its deployment engine (jenkins)
# ----------------------------------
echo ""
echo "Building container images"
sudo docker build ./src -t modm -f ./build/container/Dockerfile.modm  
sudo docker build . -t jenkins -f ./build/container/Dockerfile.jenkins