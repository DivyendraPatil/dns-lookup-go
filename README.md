# dns-lookup-go

This is a simple command-line tool written in Go that performs a DNS lookup and displays the IP addresses and CNAME for a given domain name.

## Usage

To use this tool, open a terminal and navigate to the directory where the program is saved. Then, run the following command:

```
go run main.go -dns <domain-name>
```
Replace <domain-name> with the domain name you want to look up. The tool will display the IP addresses and CNAME for the domain.

Build by running the below command in the current folder

`CGO_ENABLED=0 go build -o dns-check`

Run by:

```
dns-lookup-go git:(main) ✗ ./dns-check -dns divyendra.com
DNS:   divyendra.com 
CName: divyendra.com. 
IpsV4: [104.21.29.202 172.67.149.199] 
IpsV6: [2606:4700:3037::6815:1dca 2606:4700:3037::ac43:95c7]
```

## Dependencies

This program has no external dependencies other than the Go standard library.