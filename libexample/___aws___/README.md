aws
==================

# アクセスユーザーを作成
https://console.aws.amazon.com/iam/home#home

# コード

## 認証
アクセスキーとシークレットで認証。

```go
creds := aws.Creds(accessKey, secretKey, "")
```

## S3アクセス

クライアント生成。

```go
c := s3.New(creds, "ap-northeast-1", nil)
```

ファイルを取得

```go
req := s3.GetObjectRequest{}
req.Bucket = aws.String("kyokomi-foo")
req.Key = aws.String("bar/media.json")

res, err := c.GetObject(&req)
if err != nil {
    log.Fatalln(err)
}
defer res.Body.Close()
```

普通のhttp.Responseっぽく扱える

```go
data, err := ioutil.ReadAll(res.Body)
if err != nil {
    log.Fatalln(err)
}

fmt.Println(string(data))
```
