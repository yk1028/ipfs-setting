# IPFS-setting
### OS
- ubuntu 18.04 x86/64

### Init 
- init.sh
- source ~/.profile

### Key setting 
- swarm-key.sh -> ~/data/ipfs_data/swarm.key (하나의 노드에서 실행 후 나머지 복사)
- cluster-secret.sh -> docker-compose.yml - ipfs-cluster - environment - CLUSTER_SECRET
- directory 분리시 key setting 주의

## Docker command
### ipfs-node
- `sudo docker-compose up -d ipfs-node`
  - docker-compse.yml이 있는 경로에서 실행
- 아래 명령어에서 ipfs-node == container name (ex. `ubuntu_ipfs-node_1`)
- 맨 앞에 붙는 이름(ex. `ubuntu`)은 diretory 이름
  
* boot node 초기화
```shell
$ docker exec ipfs-node ipfs bootstrap rm --all
```

* peer 확인
```shell
$ docker exec ipfs-node ipfs id
{
	"ID": "12D3KooWSBgUZ3oTCCA3UavCrhjsbMK53zvhWFkvFX4UVP3Dxk1N",
	"PublicKey": "CAESIPMw4z2xUZIoyQeJ86hXvBccKAobVLjYjZQueEIK/mw5",
	"Addresses": [
		"/ip4/127.0.0.1/tcp/4001/p2p/12D3KooWSBgUZ3oTCCA3UavCrhjsbMK53zvhWFkvFX4UVP3Dxk1N" # peer address
	],
	"AgentVersion": "go-ipfs/0.10.0/64b532f",
	"ProtocolVersion": "ipfs/0.1.0",
	"Protocols": [
		"/ipfs/bitswap",
		"/ipfs/bitswap/1.0.0",
		"/ipfs/bitswap/1.1.0",
		"/ipfs/bitswap/1.2.0",
		"/ipfs/id/1.0.0",
		"/ipfs/id/push/1.0.0",
		"/ipfs/lan/kad/1.0.0",
		"/ipfs/ping/1.0.0",
		"/libp2p/autonat/1.0.0",
		"/libp2p/circuit/relay/0.1.0",
		"/p2p/id/delta/1.0.0",
		"/x/"
	]
}
```

* bootstrap 등록 (안해도 동작)
```shell
$ docker exec ipfs-node ipfs bootstrap add /ip4/{연결할 peer의 IP}/tcp/4001/p2p/{연결할 peer의 id}


# 예제
$ docker exec ipfs-node ipfs bootstrap add /ip4/172.31.28.135/tcp/4001/p2p/12D3KooWSBgUZ3oTCCA3UavCrhjsbMK53zvhWFkvFX4UVP3Dxk1N                                      
```

* boot node 확인
```shell
$ docker exec ipfs-node ipfs bootstrap
```

* swarm peer 연결
  * swarm.key 등록 필요
  * 주의 : `... no good addresses` 이런 error가 발생하면 swarm filters 확인 및 삭제
```shell
$ docker exec ipfs-node ipfs swarm connect /ip4/{연결할 peer의 IP}/tcp/4001/p2p/{연결할 peer의 id}
# 예제
$ docker exec ipfs-node ipfs swarm connect /ip4/172.31.28.135/tcp/4001/p2p/12D3KooWSBgUZ3oTCCA3UavCrhjsbMK53zvhWFkvFX4UVP3Dxk1N                                      
```

* 연결된 peer 확인
```shell
$ docker exec ipfs-node ipfs swarm peers
```

* cli로 테스트 파일 추가
  - data/ipfs_stage directory에 ipfs에 추가하고 싶은 test 파일 추가
  - `ipfs add export/{file_path}`로 ipfs에 등록
  - `ipfs cat {cid}`로 확인
  
### ipfs-cluster
- `sudo docker-compose up -d ipfs-cluster`
  - docker-compse.yml이 있는 경로에서 실행
  - identity.json, service.json 없으면 자동 생성
- cluster peer 연결
  - /data/ipfs_cluster/service.json peer_addresses 추가
  - 하나의 ipfs-cluster에서 자기 자신을 포함한 모든 cluster의 address 추가
  - 등록 후 한 쪽 재시작 (docker process 제거 후 다시 시작)
  ```json
    {
        "cluster": {
        ...
            "peer_addresses": [
            "/ip4/3.35.210.47/tcp/9096/ipfs/12D3KooWCikJL8599zzvGEt1cujK5CbBoei8vHRThDWE6gJqJB6Z",
            "/ip4/13.125.54.137/tcp/9096/ipfs/12D3KooWJ3LndWyqYjMst2hHauY1NjcT1KEzDhHsreWx32mhtXSo"
            ]
        },
        ...
    }
  ```
