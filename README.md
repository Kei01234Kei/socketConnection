# socketConnection

研究室の導入課題用リポジトリです。
Go言語でソケット通信の実装を行いました。

## テスト方法

### サーバ起動

```bash
cd server
go run main.go
```

### クライアントからサーバにデータを送る

```bash
cd client
go run main.go helloWorld # helloWorldという文字列をサーバに送る
```
