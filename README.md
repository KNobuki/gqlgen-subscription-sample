# gqlgen-subscription-sample

gqlgenによるsubscription実装のサンプル

## 起動方法

``` shell
docker-compose up -d
```

## GraphQL実行方法

### リンク

http://localhost:8080/

### マット情報登録 mutation

```graphql
mutation {
  createSmartMat(currentWeight: 100) {
    id,
    currentWeight
  }
}
```

### マット重量変更 mutation

```graphql
mutation {
    updateSmartMatWeight(id: 1, currentWeight:200) {
        id
        currentWeight
    }
}
```

### マット情報取得 query

```graphql
query {
	smartMats {
    id
    currentWeight
  }
}
```

### マット情報取得 subscription

```graphql
subscription {
  smartMatWeightUpdated(id: 1) {
    id
    currentWeight
  }
}
```
