
Add this function to your bash to be able to do `h core`

```
function h {
    dir="$(hlz cd $1)" && cd $dir
}
```

~/hlz.yaml

Example:
```
code_path: /home/xaf/Code/holaluz
github_key_path: /home/xaf/.ssh/id_rsa_github
```

Building
```
go build -o hlz *.go
```

```
sudo mv hlz /usr/local/bin/hlz
```

