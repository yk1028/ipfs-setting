# docker install
cd ~
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# go install
sudo apt-get install wget
cd /tmp
wget https://go.dev/dl/go1.18.2.linux-amd64.tar.gz
sudo tar -xvf go1.18.2.linux-amd64.tar.gz
sudo mv go /usr/local
cd ~

# set go root, path (ubuntu)
echo 'GOROOT="/usr/local/go"' >> ~/.profile
echo 'GOPATH="$HOME/go"' >> ~/.profile
echo 'PATH="$GOPATH/bin:$GOROOT/bin:$PATH"' >> ~/.profile
source ~/.profile

# docker-compose install
sudo curl -L "https://github.com/docker/compose/releases/download/1.28.5/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
