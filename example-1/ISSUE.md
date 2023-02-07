# Argument on Interace fields

## Fails:
```graphql
query {
  user(id: "1") {
      sayHi(name: "someone")
  }
}
```

## Works:
```graphql
query {
  user(id: "1") {
    ... on GitHubUser {
      sayHi(name: "someone")
    }
  }
}
```
