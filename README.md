# txtrep
txtrep is uses marks to replace text.

## Install
```
$ go get -u github.com/nissy/txtrep/cmd/txtrep
```

## Usage
```
$ echo $DB_PASSWORD
aaaaaa

$ cat database.yaml
database:
  password: <env:DB_PASSWORD>

$ txtrep database.yaml
database:
  password: aaaaaa
```