# dns-lookup-go

Build by running the below command in the current folder

`CGO_ENABLED=0 go build -o dns-check`

Run by:

`./dns-check divyendra.com`

```
 dns-lookup-go git:(main) ✗ ./dns-check divyendra.com
DNS:   divyendra.com 
CName: divyendra.com. 
IpsV4: [104.21.29.202 172.67.149.199] 
IpsV6: [2606:4700:3037::6815:1dca 2606:4700:3037::ac43:95c7]
```