# Command spoofer

Provides commands alternatives for unsupported commands, emulate outputs

## Definition

Every command uses JSON file with match arguments to outputs

### lsusb

`lsusb` provides alternative outputs based on file-flag `/tmp/lsusb_alternative`. if the file exists
output from lsusb2.json will be shown

## Generation

`json` can be generated by using `generate-json`

```
Usage of dist/generate-json:
  -append
    	append to existing file (default true)
  -args string
    	arguments ',' as separator
  -exec string
    	executable (default "lspci")
  -output string
    	result file output (default "res.json")
```

Usage Examples:
```
generate-json --exec=lsusb --args="-h,--help,-t,--tree,-V,--version," --append=false --output lsusb.json
```

# Installation

1. Get the latest release link from the Releases page and download with `wget` and extract to  `~/bin`
    ```bash
    wget https://github.com/codio-content/command-spoofing/releases/latest/download/commands-linux-amd64.tgz -O - | sudo tar -zx -C /bin
    ```
