# go-email


# 1 time setup
```
$ source init-env-golang-1.18.sh
$ go mod init weetech.ch/email
$ go mod tidy
```

# release
```
$ git tag v0.0.1
$ git push origin v0.0.1
Total 0 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:weetech-software/go-email.git
 * [new tag]         v0.0.1 -> v0.0.1
```

# verify after release
```
$ go list -m github.com/weetech-software/go-email@v0.0.1
```

# get the module
```
$ go get github.com/weetech-software/go-email@v0.0.1
```
