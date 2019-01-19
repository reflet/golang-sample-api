# golang-api

## 構成
- go1.11.4 linux/amd64
- dep v0.5.0
- goa

## ローカル環境構築

### 前提条件
- macOS
- Homebrew
- git ( `$ brew install git` )
- docker for mac ( `$ brew cask install docker` )

### サーバ起動
```bash
# ワークス−ペース作成 (任意フォルダ)
$ mkdir -p ~/workspace/golang-sample-api/
$ cd ~/workspace/golang-sample-api/

# コード配置
$ git clone git@github.com:reflet/golang-sample-api.git .

# dockerイメージを構築する
$ docker-compose build

# goパッケージをインストールする
$ docker-compose run go dep ensure -v -vendor-only=true

# アプリケーション構築
$ docker-compose run go go build -o server

# コンテナを起動する
$ docker-compose up -d
```

ブラウザを起動して動作確認する
```bash
# 起動に少し時間がかかるのでちょっと時間をおいて実行する
$ open http://192.168.99.100:8080/bottles/1
$ open http://192.168.99.100:8080/swagger.json
$ open http://192.168.99.100:8080/swagger.yaml
```
※ `192.168.99.100`は、環境によって異なることがあります。

### 日々の作業
```bash
# サーバの状態確認
$ docker-compose ps

# サーバのログ確認 (最新10件を継続して監視 - ctrl + c で確認終了)
$ docker-compose logs -f --tail 10 go 

# サーバ停止
$ docker-compose stop

# サーバ開始
$ docker-compose start

# サーバ破棄 (コンテナのみ)
$ docker-compose down

# サーバ破棄 (コンテナ＋ボリューム)
$ docker-compose down -v
```
