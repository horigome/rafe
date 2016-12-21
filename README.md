# Rafe commandline service with Go

## Overview

コマンドラインを実行するリモートサービスです。 JSONでコマンドを指定できます。

## How to build

    > make vendor_update
    > make this

## require

* golang.org/x/text/encoding/japanese
* golang.org/x/text/transform

## REST API

#### (POST) /host:8080/command

 Request body  

    {
      "commands": [
          {
            "name":   "command name",
            "option": "command options"
          },

      ]
    }

Response

    stdout text


ex.
    > curl -XPOST localhost:8080/command -d '{
      "commands": [
          {"name": "ls", "option": "-la"}
      ]    
    }'


## License

MIT

---

2016  M.Horigome
