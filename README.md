# envrep
envrep is uses marks to replace text.

## Install
```
$ go get -u github.com/nissy/envrep/cmd/envrep
```

## Usage
```
$ echo $DB_PASSWORD
aaaaaa

$ cat database.yaml
database:
  password: <env:DB_PASSWORD>

$ envrep database.yaml
database:
  password: aaaaaa
```