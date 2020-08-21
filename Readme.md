# livedl
A tool which could record niconico's new HTML5 live streaming. It could also record live streaming from Twitcasting and YouTube with `streamlink` and `youtube-dl`.

## Build in Linux, e.g. Ubuntu

### Installation of ffmpeg
```
sudo add-apt-repository ppa:jonathonf/ffmpeg-4
sudo apt update
sudo apt install ffmpeg
```

### Installation of streamlink
```
sudo add-apt-repository ppa:nilarimogard/webupd8
sudo apt update
sudo apt install streamlink
```

### Installation of youtube-dl
```
sudo curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl
sudo chmod a+rx /usr/local/bin/youtube-dl
```

### Installation of go
AS Go 1.15 is required otherwise the installation of gin would go wrong.
```
wget https://golang.org/dl/go1.15.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.15.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```
If necessary, add `PATH` in bashrc.

### Installation of git
```
sudo apt-get install git
```

### Installation of compiler tools like gcc
```
sudo apt-get install build-essential
```

### Installation of essential modules in go
```
go get github.com/gorilla/websocket
go get golang.org/x/crypto/sha3
go get github.com/mattn/go-sqlite3
go get github.com/gin-gonic/gin
```

### Get the source of livedl
```
git clone https://github.com/glloyd0314/livedl.git
```

### Compile livedl
```
cd livedl
```

#### To select the version of livedl (Optional)
```
$ git tag
20180513.6
20180514.7
...
20181008.33
20181107.34
20181215.35
v2alpha
$ git checkout 20181107.34 (The version you select)
```

#### To select the latest version of livedl (Optional)
```
git checkout master
```

Now build it.
```
go build src/livedl.go
```

Check the version of livedl
```
./livedl -h
livedl (20181215.35-linux-amd64)
```

### Way to use it
First to login nico and record 
`
./livedl -nico-login 'example@example.com,password' lv321234567
`
From 2nd time, record directly
`
./livedl lv321234567
`


For more please check the website of https://himananiito.hatenablog.jp/entry/livedl (Japanese)


## Windows(32bit及び64bit上での32bit向け)コンパイル方法

### gccのインストール

gcc には必ず以下を使用すること。

http://tdm-gcc.tdragon.net/download

環境変数で（例）`C:\TDM-GCC-64\bin`が他のgccより優先されるように設定すること。

### 必要なgoのモジュール

linuxの説明に倣ってインストールする。

### コンパイル

PowerSellで、`build-386.ps1` を実行する。または以下を実行する。

```
set-item env:GOARCH -value 386
set-item env:CGO_ENABLED -value 1
go build -o livedl.x86.exe src/livedl.go
```

### 32bit環境で`x509: certificate signed by unknown authority`が出る

動けばいいのであればオプションで以下を指定する。

`-http-skip-verify=on`

## Dockerでビルド

### livedlのソースを取得
```
git clone https://github.com/glloyd0314/livedl.git
cd livedl
git checkout master # Or another version that supports docker (contains Dockerfile)
```

### イメージ作成
```
docker build -t <your_image_tag> .
```

### イメージの使い方

- 出力フォルダを/livedlにマウント
- 通常のパラメーターに加えて`--no-chdir`を渡す

```
docker run -it --rm -v ~/livedl:/livedl <your_image_tag> livedl --no-chdir <other_parameters> ...
```
