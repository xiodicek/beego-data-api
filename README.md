# beego-data-api

## 導入
beego入手
```
go get -u github.com/beego/bee
go get -u github.com/astaxie/beego
```

git clone
```
cd $GOPATH/src
git clone https://github.com/xiodicek/beego-data-api.git
```
※ `$GOPATH/src`  直下に置かないと動きません

起動
```
cd $GOPATH/src/beego-data-api
bee run
```

## 基本的な構成
ディレクトリ構成は下記のようになっています
```
beego-data-api
├── conf
├── controllers
│   └── export.go
├── lastupdate.tmp
├── main.go
├── models
│   ├── appsflyerReport.go
│   └── googleReport.go
├── routers
├── swagger
└── tests
```

### コア部分
- `controllers/export.go` でパラメータ受け付け、各モデルへ処理を振り分け
- `models/` 配下にテーブルごとにモデルを作成し、データ抽出処理を定義

### その他
- `conf/` 設定ファイル
- `routers/` ルーティング定義
- `swagger/` swaggerドキュメントファイル（ `-downdoc=true -gendoc=true` で自動生成されます）

### APIドキュメント
[https://xiodicek.github.io/beego-data-api/swagger/](https://xiodicek.github.io/beego-data-api/swagger/) 

## リンク
- [公式ドキュメント](https://beego.me/docs/intro/) 
-  [Go製フレームワーク（BEEGO）の使い方 - Qiita](https://qiita.com/ken0renaissance/items/89a64ea5ae955892d9f2) 
