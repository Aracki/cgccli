# cgccli
The Cancer Genomics Cloud - CLI tool

[![Go Report Card](https://goreportcard.com/badge/github.com/aracki/cgccli)](https://goreportcard.com/report/github.com/aracki/cgccli)
[![GoDoc](https://godoc.org/github.com/Aracki/cgccli?status.svg)](https://godoc.org/github.com/Aracki/cgccli)
![GitHub All Releases](https://img.shields.io/github/downloads/aracki/cgccli/total.svg)

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

![Alt Text](https://imgur.com/YODUYuv)

### CLI tool supports following operations

* List projects (​https://docs.cancergenomicscloud.org/docs/list-all-your-projects​)
* List files in project (​https://docs.cancergenomicscloud.org/docs/list-files-in-a-project​)
* Get file details (​https://docs.cancergenomicscloud.org/docs/get-file-details​)
* Update file details (​https://docs.cancergenomicscloud.org/docs/update-file-details​)
* Download file (​https://docs.cancergenomicscloud.org/docs/get-download-url-for-a-file​)
