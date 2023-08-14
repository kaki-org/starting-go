# starting-go

## requirements

- `go version go1.21.0 darwin/arm64`


## environments

`~/.zshrc` などに以下を追加

```zsh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## install

[asdf-vm](https://asdf-vm.com/)を利用してインストールを実施

## TODO

asdf利用時に以下のメッセージが出力される

```
Notice: Behaving like ASDF_GOLANG_MOD_VERSION_ENABLED=true
        In the future this will have to be set to continue
        reading from the go.mod and go.work files
```

[こちら](https://blog.yukii.work/posts/2023-06-29-asdf-golang-mod-version-enabled/)を参考に対応する。以下を設定しておけばメッセージ自体は表示されなくなる模様。

```zsh
export ASDF_GOLANG_MOD_VERSION_ENABLED=true
```
