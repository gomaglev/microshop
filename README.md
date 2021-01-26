#[WIP]

# microshop

A scaffolding using gRPC, Wire and Gorm.

## Install and Development

Requires Go 1.14+, Protobuf 3.0+, Wire

Copy env file
```
cp .env.example .env
```

install modules
```
go get .
```

install modules
```
go get ./cmd/microshop
```

run application
```
make start
```

debug application
use dlv, vscode settings.json
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug with dlv",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "host": "127.0.0.1",
            "port": 10088,
            "program": "${workspaceFolder}/cmd/microshop/main.go",
            "cwd": "${workspaceFolder}",
            "env": {},
            "args": ["start"],
            "showLog": true
        },
    ]
}
```

### Test

Use moke injectors to run unit tests.  
internal/injector/mock  

## Postgres cert sample
https://github.com/CrunchyData/crunchy-containers/tree/master/examples

## Project layout
Project layout follows the standard below:
https://github.com/golang-standards/project-layout