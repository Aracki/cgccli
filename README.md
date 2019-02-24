# cgccli
The CLI tool written in Go for accessing Cancer Genomics Cloud Public API 

[![Go Report Card](https://goreportcard.com/badge/github.com/aracki/cgccli)](https://goreportcard.com/report/github.com/aracki/cgccli)
[![GoDoc](https://godoc.org/github.com/Aracki/cgccli?status.svg)](https://godoc.org/github.com/Aracki/cgccli)
![GitHub All Releases](https://img.shields.io/github/downloads/aracki/cgccli/total.svg)
![https://img.shields.io/badge/CGC-API%20Reference-blue.svg](https://img.shields.io/github/downloads/aracki/cgccli/total.svg))
<a href="https://docs.cancergenomicscloud.org/docs/the-cgc-api)"><img src="https://img.shields.io/badge/CGC-API%20Reference-blue.svg"></a>

### Installation
If you don't have Go installed you can download the appropriate binary for your system from the [releases page](https://github.com/Aracki/cgccli/releases) and put it in your path.

If you do have Go:

```
go get -u github.com/aracki/cgccli
```

### Sample usage
```
cgccli --token {token} projects list
cgccli --token {token} files list --project {project_id}
cgccli --token {token} files stat --file {file_id} 
cgccli --token {token} files update --file {file_id} name={name}
cgccli --token {token} files update --file {file_id} metadata.{key}={value}
cgccli --token {token} files download --file {file_id} --dest {file_destination}
```

![Alt Text](https://imgur.com/YODUYuv.gif)
