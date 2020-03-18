# hlz - My set of development tools for hlz

## 📜 Table of contents

- [🚀 Usage](#️-usage)
- [🏗️ Installation](#-installation)
- [🛠️ Configuration](#️-configuration)

---

## 🚀 Usage

Clone you repositories easily with:

```
hlz clone core
```

Generate random cups from your cli:

```
hlz cups
```

Enter to the aws instances:

```
hlz ssh core
```

---

## 🏗️ Installation

### From binary (Linux)

```bash
sudo wget -O /usr/local/bin/hlz https://github.com/xafardero/hlz/releases/latest/download/hlz
```
```bash
sudo chmod +x /usr/local/bin/hlz
```

### From binary (Mac)

```bash
sudo wget -O /usr/local/bin/hlz https://github.com/xafardero/hlz/releases/latest/download/ hlz_darwin
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

There you can config the path where are going to be cloned your repositories or your github ssh key.

Example:
```
code_path: /home/xaf/Code/holaluz
github_key_path: /home/xaf/.ssh/id_rsa
```

### Fast "change directory" to your projects
Add this function to your zsh config (~/.zshrc) to be able to do `h core`

```
function h {
    dir="$(hlz cd $1)" && cd $dir
}
```
