# go-pwdmaker
A simple command line tool for generating passwords written with golang

### How to use

#### Get help
```sh
go run ./cmd/pwdmaker/main.go -h
```


#### Generate password
```sh
go run ./cmd/pwdmaker/main.go -r habr -l mylogin -n 24
```

Output:
```sh
Resource: Habr
Login: mylogin
Generated pwd: HE1NmmB~khV=7rz7_Jh]Z!V2
````
