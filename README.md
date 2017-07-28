db-sample-go
---

## 概要

サーバーサイドはgolang、クライアントサイドはHTML, Vue.jsで実装されています。

## 事前(参考)設定
GOROOT、GOPATH、PATHが設定されていること。以下は`~/go`がGOPATHとして設定されている前提での参考設定

```
export GOROOT=/usr/local/go
export GOPATH=~/go
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
```

## set up
依存関係の解決、必要なパッケージなどを取得します。

```
$ make install
```

## run
httpをPORT8080でlistenします。

```
$ make watch
```

