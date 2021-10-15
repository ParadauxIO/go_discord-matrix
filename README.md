# go_discord-matrix
Discord - Matrix bridge written in golang


# TODO

- Dockerfile

## Setup

``` 
git clone https://github.com/ParadauxIO/go_discord-matrix.git
mv config.json.example config.json
``` 

## Compiling
``` 
go build 
```  

## Running

The final version will be intended to be run as a docker container, however for now you can run the
project by simply executing the compiled go_discord-matrix binary.

```
chmod +x go_discord-matrix # Linux only, marks it as executable
./go_discord-matrix
```