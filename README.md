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

example: curl 'http://192.168.2.172:8888/query?chainId=11155111&&addr=0x0000000000000000000000000000000000000000&&page=1&&pageSize=12'

#### httpResponse
{"status":0,"result":[{"height":"5378749","blockhash":"0xcced0a4b26e1d7dd5c547a2e33113dc5ad2f33b69f4a493326b4fb99b36f1c52","timestamp":"1709098656","txhash":"0xc4cde5a2b122b62a96421a12a6ef79be79027fdefbee407c90ce6e348f4980c0","toaddr":"0xab420Eb61b0eaf0Bf11ad7C487CA7321D3F49dCA","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378749","blockhash":"0xcced0a4b26e1d7dd5c547a2e33113dc5ad2f33b69f4a493326b4fb99b36f1c52","timestamp":"1709098656","txhash":"0x58dd1c0cdd85123a6fb544e3e315fef45da20f0cc09fa041a450293af54598aa","toaddr":"0xC37C8122b81620f0f4E2ED0b250413702e4F4aE0","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378749","blockhash":"0xcced0a4b26e1d7dd5c547a2e33113dc5ad2f33b69f4a493326b4fb99b36f1c52","timestamp":"1709098656","txhash":"0xc87eb7ea3ffc1011817e2db26c311ce9cced802015dc8a7756c326b5bc42c5a1","toaddr":"0x38a0F9D08dc5Ed7379167F6065a6F20765e4e865","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378749","blockhash":"0xcced0a4b26e1d7dd5c547a2e33113dc5ad2f33b69f4a493326b4fb99b36f1c52","timestamp":"1709098656","txhash":"0x52a826aed3f06b70c07555b8d6419aff35d476cbade3f83c0f71f2f0d12c970d","toaddr":"0x271142db55dD5EAEb2ed18433F8d8672C246aA92","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378749","blockhash":"0xcced0a4b26e1d7dd5c547a2e33113dc5ad2f33b69f4a493326b4fb99b36f1c52","timestamp":"1709098656","txhash":"0x47551d3da915d814aac3c24bc32ce997ed33b84802e0410423035cd44b8da84e","toaddr":"0x33c86a2a5638cBdF299052260c3b15827Fb35f82","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378750","blockhash":"0x13e23795c550de257d59abee589fbb619cf34843d472ec9fad9bb7803adfe377","timestamp":"1709098668","txhash":"0x5fc004e173a27aa20581faf90bab39b2fcd13d27b68efb992d5d18989bdd3ae3","toaddr":"0xdB1ff830071DcFafd3dc55792Fd52C2bD95b7e12","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378750","blockhash":"0x13e23795c550de257d59abee589fbb619cf34843d472ec9fad9bb7803adfe377","timestamp":"1709098668","txhash":"0xff45de8ab8e27e36eef2078c113cb6dd64115baef787630d05f0c0f7aac5e36c","toaddr":"0xdB1ff830071DcFafd3dc55792Fd52C2bD95b7e12","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378750","blockhash":"0x13e23795c550de257d59abee589fbb619cf34843d472ec9fad9bb7803adfe377","timestamp":"1709098668","txhash":"0x50f52136079f6d7d768e08881dc72eb455b80fca2f6755166a8c79f714f8840f","toaddr":"0xdB1ff830071DcFafd3dc55792Fd52C2bD95b7e12","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378750","blockhash":"0x13e23795c550de257d59abee589fbb619cf34843d472ec9fad9bb7803adfe377","timestamp":"1709098668","txhash":"0xee1e1e51936c0e0ec868eeee750394c973c32096d607d7ce513794d55b6e8b9e","toaddr":"0xdB1ff830071DcFafd3dc55792Fd52C2bD95b7e12","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378751","blockhash":"0x28a0bbfbe884902946b6972d23991a001fb91ec46cb2c8feecaa68451e800339","timestamp":"1709098680","txhash":"0xff99c91f6c2d3644b1e4ad152844d9bda172d35be94fbe0725cc2299938b1e5d","toaddr":"0xdB1ff830071DcFafd3dc55792Fd52C2bD95b7e12","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378751","blockhash":"0x28a0bbfbe884902946b6972d23991a001fb91ec46cb2c8feecaa68451e800339","timestamp":"1709098680","txhash":"0x0b3ffc16cd99622379648262c152acf86fac31846da7df4c9c98e3037b05b909","toaddr":"0xdB1ff830071DcFafd3dc55792Fd52C2bD95b7e12","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"},{"height":"5378751","blockhash":"0x28a0bbfbe884902946b6972d23991a001fb91ec46cb2c8feecaa68451e800339","timestamp":"1709098680","txhash":"0x238fc437459fa132242acb6267f79762ac742ac5197bbf04403e4b49e131c237","toaddr":"0xdB1ff830071DcFafd3dc55792Fd52C2bD95b7e12","value":"10000000000000000000000","contract":"0x7f11f79DEA8CE904ed0249a23930f2e59b43a385"}]}