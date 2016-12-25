# Rafe REST API Servie
Command exec Over the Web API

## Overview

コマンドラインを実行するREST API Service です。 JSON形式で実行するコマンドを指定できます。

## How to build

require golang. and GNU make

    > make vendor_update
    > make this

## require

* golang.org/x/text/encoding/japanese
* golang.org/x/text/transform

## REST API

#### (POST) /host:8080/command

コマンドをサービスで実行し、結果をJSONで返却します。

Request header

    application/json

body    

    {
      "commands": [
          {
            "name":   "command name",
            "option": "command options"
          },

      ]
    }

Response
* 200 : OK
* 400 : Bad Request
* 500 : Internal error

Body  (stdout text)

    stdout text


ex.

    > curl -XPOST localhost:8080/command -d '{
      "commands": [
          {"name": "ls", "option": "-la"}
      ]    
    }'


#### (GET) /host:8080/version

サービスのバージョンを取得  

 response
* 200 : OK
* 400 : Bad Request

body  

    {
      "version": "1.0.0.0"
      "description": "rafe service"
      ]
    }

ex.

    > curl -XGET localhost:8080/version


---

## License

MIT


2016  M.Horigome
