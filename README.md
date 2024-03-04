# chaindata
## 0. Compile
make

## 1. Start
### using default(chain.ini and port 8888)
./chaindata start

## 2. Service
### 2.1 Count
#### httpRequest
two parameters needed here:
- chainId, chain for querying
- addr, address for querying

example: curl 'http://192.168.2.172:8888/count?chainId=11155111&&addr=0x0000000000000000000000000000000000000000'
#### httpResponse
{"status":0,"result":5221}

### 2.2 Query
#### httpRequest
four parameters needed here:
- chainId, chain for querying
- addr, address for querying
- page, starts with 0
- pageSize, numbers per page

example: curl 'https://robin-chaindata.rangersprotocol.com/query?chainId=11155111&&addr=0xFB18E6FF5F94Bdf0115Ed4c61F9Cf49041245dEd&&page=0&&pageSize=12

#### httpResponse
{"status":0,"result":[{"height":"5413149","blockhash":"0xe6765e1d0ca9091b944eaf0affa8ef1c5f7114d5619630066f0baf10ca2f5f89","timestamp":"1709535168","txhash":"0x4bad2fda95db889f1a3bb4bfd4fe59e190ae24b25fa13591eaf94037127201db","toaddr":"0xB48Cb9714e30b6441d22492883a622E08A12a902","value":"0","contract":"","gas":"3000000","gasprice":"1638967981"},{"height":"5413157","blockhash":"0x3dbf398275f6dc59c865c9f2a6828f0017639ebe2db773de1e9e82cab607be08","timestamp":"1709535264","txhash":"0xc1cea72dc665b1ee42d94728e9d2f87d4a636afaa00022b6b915b0e4685cc093","toaddr":"0xB48Cb9714e30b6441d22492883a622E08A12a902","value":"0","contract":"","gas":"3000000","gasprice":"1615032380"},{"height":"5413174","blockhash":"0x3640f582f2204e250a58c7288250141fd83e4f4382a3c258a6d06c59565a89ae","timestamp":"1709535480","txhash":"0x6b46137824d7b3b382148aa512159f458158d80b2545fece755c32e72665a7fb","toaddr":"0xfE8F287209904B5DE906bEEE23FBba6a54b94C66","value":"0","contract":"","gas":"70316","gasprice":"1621628916"},{"height":"5413175","blockhash":"0x58eb8352bf4e249d052f67ba36838d2a57b1bf7f322a3db3e84a67d683d613fa","timestamp":"1709535492","txhash":"0x2d92f60a73fe607a784105da8065756f9b25d79c774c333eb8bdc2d7d7dca34a","toaddr":"0xfE8F287209904B5DE906bEEE23FBba6a54b94C66","value":"0","contract":"","gas":"53160","gasprice":"1632440811"},{"height":"5413177","blockhash":"0xbde1f62dbfb6631f0840fcc77a1b5ef8996c3c758a60f90974bb881132774d52","timestamp":"1709535516","txhash":"0x8164c21d3923e29c6f9dbfe93a89d2c93af6298e90d361b8ce9326abbae6162d","toaddr":"0xfE8F287209904B5DE906bEEE23FBba6a54b94C66","value":"0","contract":"","gas":"53160","gasprice":"1621023306"},{"height":"5413216","blockhash":"0x748d77d3b26b99b2a75f6427677084915a11f178846be6e4438e656cebb68bf7","timestamp":"1709535984","txhash":"0x35e1ab79876b354a63c3df554a2e23247ebaf2ced95a1808ba3feb436f272ee7","toaddr":"0xfE8F287209904B5DE906bEEE23FBba6a54b94C66","value":"0","contract":"","gas":"70328","gasprice":"1622844301"},{"height":"5413217","blockhash":"0x0c6cdcf94eac9f5fe4bf6ed769ec13d05491ea314d269dd880db9eba856c4d25","timestamp":"1709535996","txhash":"0x9d8e539caea8aae18a67467b08b5ea65cabafd865b62792e3b08f76c61aaf1b4","toaddr":"0xfE8F287209904B5DE906bEEE23FBba6a54b94C66","value":"0","contract":"","gas":"53160","gasprice":"1628848492"},{"height":"5413237","blockhash":"0x533e38dc5e52f6f645bc6523550091c5323115de293f0ddc80b391e8a99cfd6c","timestamp":"1709536248","txhash":"0xfa7565d58014f14bea976b90a4b6f78fcd2918c72419e388b39dee2cda0d05e7","toaddr":"0xF362fdD4ba74C339db38D2a12CB2a0F95a61BEdf","value":"0","contract":"","gas":"502493","gasprice":"147843743"},{"height":"5413239","blockhash":"0x3c7f0b2fcf45d0585368b1a6f26cd01f4d8b7f04346b9f669575677c313f6d23","timestamp":"1709536272","txhash":"0xb1cd344b3da19c7ccd4d525453a05a617fa5a4b9f7324a67fb7d6eeab18ed23e","toaddr":"0xF362fdD4ba74C339db38D2a12CB2a0F95a61BEdf","value":"0","contract":"","gas":"800000","gasprice":"1645321512"},{"height":"5413240","blockhash":"0x929e544ed1fad6f073e8263a2c1725573dc8d3336068b2c359306da7f7d09a57","timestamp":"1709536284","txhash":"0x53455afe6b378cee0d3500c62102d88e69f00ba2dca1b45b73d7a128b09dad38","toaddr":"0xF362fdD4ba74C339db38D2a12CB2a0F95a61BEdf","value":"0","contract":"","gas":"128365","gasprice":"1142309003"}]}
