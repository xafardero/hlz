# hlz - My set of development tools for hlz

## 📜 Table of contents

- [🏗️ Installation](#-installation)
- [🛠️ Configuration](#️-configuration)

---

## 🏗️ Installation

### From binary

```bash
sudo wget -O /usr/local/bin/hlz https://github.com/xafardero/hlz/releases/latest/download/hlz
```
```bash
sudo chmod +x /usr/local/bin/hlz
```

### From source code

```bash
wget https://github.com/xafardero/hlz/releases/latest/download/hlz.zip
```
```bash
go build -o hlz *.go
```

```bash
sudo mv hlz /usr/local/bin/hlz
```
---
## 🛠️ Configuration
Config your hlz command in ~/hlz.yaml

Example:
```
code_path: /home/xaf/Code/holaluz
github_key_path: /home/xaf/.ssh/id_rsa_github
```

### Fast "change directory" to your projects
Add this function to your bash to be able to do `h core`

```
function h {
    dir="$(hlz cd $1)" && cd $dir
}
```