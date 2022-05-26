# Zenon Dev-net Node

Dev-net version of the official zenon implementation. Intended to be used by builders developing on top of the zenon ecosystem, rather than developing the core zenon-node.  

## Building from source

```shell
make znnd
```

## Changes from main-net

- different **genesis**
  - there is only 1 initial pillar
  - pre-mint enough funds for 200 pillars, controlled by the pillar owner 
- different **chain identifier**; transactions made on main-net can't be published here, not the other way around
  - `ChainId=2` 
- different **location**; the different ledgers don't collide with one another
  - `znn-devnet` vs `znn`
  - for linux, the default location is `~/.znn-devnet`
- different **seeds**; all seeds were removed, the only available seed being a gateway provided by default  
- different default **ports**; tries to make connecting to the `devnet` an intentional process 
  - `DefaultListenPort = 35595` vs 35**9**95
  - `DefaultHTTPPort   = 35597` vs 35**9**97
  - `DefaultWSPort     = 35598` vs 35**9**98
- enable accelerator by default 

## How to interact

**Host a Node** by running the `znnd-devnet` binary. The process should work out of the box. You can build it yourself or download a build one from the release.

**Explore** the network using [explorer.zenon.network](http://explorer.zenon.network/) and connect to the node `http://172.104.203.145:35597`

**S Y R I U S** support is not possible yet, due to limitations on changing the `Client network identitifer`
