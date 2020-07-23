#### Generate self-signed certificates for localhost test use
`openssl req  -nodes -new -x509  -keyout server.key -out server.cer`
- after generating: `export SHOP_TLSCERT=server.cert SHOP_TLSKEY=server.key`

#### Environment variables:
- SHOP_TLSCERT, SHOP_TLSKEY
- SHOP_BIND_ADDR
- LOG=* - print all logger info