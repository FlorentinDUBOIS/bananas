# bananas

A slack bot who love bananas!

## Installation

Firstly, clone the repository:

```
$ git clone git@github.com:FlorentinDUBOIS/bananas.git $GOPATH/src/github.com/FlorentinDUBOIS/bananas
```

Secondly, install dependencies with Go [dep](https://github.com/golang/dep):

```
$ dep ensure
```

Thirdly, compile it:

```
$ go build
```

Now, you are able to use it! So, to begin just do:

```
$ ./bananas --help

A slack bot who love bananas!

Usage:
  bananas [flags]

Flags:
  -h, --help           help for bananas
  -p, --port int32     set the port to listen (default 8080)
      --token string   set the bearer token to use
  -v, --verbose        enable verbose output
```

## Environment variables

| Name  | Description                           | Required |
| ----- | ------------------------------------- | -------- |
| PORT  | Set the port to listen                | No       |
| TOKEN | Set the token to set on each response | Yes      |

## Availables routes

| Verbs  | Path        | Description           |
| ------ | ----------- | --------------------- |
| `POST` | `/messages` | Get a random sentence |
