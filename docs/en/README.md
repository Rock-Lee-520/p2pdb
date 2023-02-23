
<!-- [![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme) -->
<p align="center">
  <a href="https://opensource.org/licenses/Apache-2.0">
        <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"
            alt="License"></a>
    <a href="https://github.com/RichardLitt/standard-readme">
        <img src="https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square"
            alt="License"></a>
    <a href="https://godoc.org/github.com/Rock-liyi/p2pdb">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg"
            alt="GoDoc"></a>
                <a href="https://godoc.org/github.com/Rock-liyi/p2pdb">
        <img src="https://img.shields.io/badge/platform-Win32%20|%20GNU/Linux%20|%20macOS%20|%20FreeBSD%20-gold"
            alt="platform"></a>
</p>

# P2PDB

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)



### Introduction
P2PDB is a decentralized, distributed, peer-to-peer relational  database for web3. eventually consistent algorithm was implemented using  [Merker-CRDT](https://research.protocol.ai/blog/2019/a-new-lab-for-resilient-networks-research/PL-TechRep-merkleCRDT-v0.1-Dec30.pdf).  It uses IPFS Pubsub to automatically sync databases with peers and uses IPFS libp2p to build  a decentralized p2p  network, it's a relational  database that uses CRDT, POW, POS for conflict resolution. P2PDB vision is to build a decentralized database of industrial grade , making P2PDB an outstanding choice for local-first web applications, decentralized applications (DAPP) and blockchain applications,  it's developed based on [whitebook](docs/zh-cn/whitebook.md).

### Use the tutorial
[tutorial](https://rock-liyi.github.io/p2pdb/)


### Goal

The goal of this database is:

* A Dapp application data storage solution

* A decentralized distributed database solution

* An edge data storage solution

* An offline application data storage solution

### Features

1. Compatible with mysql general SQL syntax, can be connected with any mysql client

2. Decentralized, no central server

3. Once the historical data is generated, it cannot be tampered with

4. Data encryption transmission, public and private key encryption and decryption, authorization support to field level

5. Ultimate consistency, supporting strong ultimate consistency

6. Fault recovery, timed snapshot copy and automatic synchronization of data

7. Version proofing similar to blockchain algorithm to ensure that once the data is generated it will never be lost.    								   

### Use Cases

1. Offline web application storage  

2. Multi-player interactive games 

3. Multi-person collaborative documents 

4. Multi-person chat applications 

5. Edge cache storage 

6. Dapp application storage 

7. NFT application storage 

8. More.........


### Thanks
Thanks  to these excellent organizations 

- [libp2p](https://github.com/libp2p) 
- [ipfs](https://github.com/ipfs)
- [dolthub](https://github.com/dolthub)
- [berty](https://github.com/berty/go-ipfs-log)

Explanation: Due to the code used in this project, such as IPFS and Libp2p, some of the referenced code follows the relevant code protocols. Thanks to Protocol Labs for their contribution to Web3.0. The remaining codes referenced in independent modules or using codes are included in the LICENSE of independent modules with reference instructions.

###  Reference description

p2pdb is a project from 0 to 1, and also one of the earliest research projects in the field of decentralized databases. In the process of implementation, we have googled all articles and code implementations about decentralized databases around the world. It can be said that p2pdb stands on the shoulders of giants because many excellent predecessors paved the way for p2pdb. Here we sincerely thank all predecessors who have studied decentralized databases, as well as IPFS team and Protocol Labs for their contributions. If there are any references to codes or implementation ideas not mentioned in this project, please contact us and we will be very happy to acknowledge them with great honor.


### Maintenance instructions

Currently, the kkguan organization is maintaining the P2PDB project. Due to certain considerations, the repository is temporarily in private status. Note that the project is still in a rapid iteration phase, which means that most of the API are at risk of being refactored. It can currently be used for learning and research purposes, but it is not recommended for production use. The development team will do their best to push out version 1.0.0, but we cannot predict when a stable version will be released.



## Contributors ✨
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks  to these wonderful people :

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/her-cat"><img src="https://avatars.githubusercontent.com/u/18332628?v=4?s=100" width="100px;" alt=""/><br /><sub><b>她和她的猫</b></sub></a><br /><a href="#infra-her-cat" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=her-cat" title="Tests">⚠️</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=her-cat" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/Rock-liyi"><img src="https://avatars.githubusercontent.com/u/35133768?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Rock</b></sub></a><br /><a href="#infra-Rock-liyi" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=Rock-liyi" title="Tests">⚠️</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=Rock-liyi" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/PandaLIU-1111"><img src="https://avatars.githubusercontent.com/u/26201936?v=4?s=100" width="100px;" alt=""/><br /><sub><b>pandaLIU</b></sub></a><br /><a href="#infra-PandaLIU-1111" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=PandaLIU-1111" title="Tests">⚠️</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=PandaLIU-1111" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/CbYip"><img src="https://avatars.githubusercontent.com/u/31086265?v=4?s=100" width="100px;" alt=""/><br /><sub><b>CbYip</b></sub></a><br /><a href="#infra-CbYip" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=CbYip" title="Tests">⚠️</a> <a href="https://github.com/Rock-liyi/p2pdb/commits?author=CbYip" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
