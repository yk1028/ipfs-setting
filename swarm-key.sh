
# set swarm.key
git clone https://github.com/Kubuxu/go-ipfs-swarm-key-gen.git
cd go-ipfs-swarm-key-gen
go build ipfs-swarm-key-gen/main.go
mkdir ~/data
mkdir ~/data/ipfs_data
./main > ~/data/ipfs_data/swarm.key