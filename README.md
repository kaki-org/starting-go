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

### ASDF_GOLANG_MOD_VERSION_ENABLED

```
Notice: Behaving like ASDF_GOLANG_MOD_VERSION_ENABLED=true
        In the future this will have to be set to continue
        reading from the go.mod and go.work files
```

[こちら](https://blog.yukii.work/posts/2023-06-29-asdf-golang-mod-version-enabled/)を参考に対応する。以下を設定しておけばメッセージ自体は表示されなくなる模様。

```zsh
export ASDF_GOLANG_MOD_VERSION_ENABLED=true
```

### gocode

現在はあまり使われていない模様なのでインストールはしない事にしている。

### gopls was not able to find modules in your workspace

[VSCodeでProject Managerを使っている場合にgoplsが動かない現象を解消する](https://qiita.com/y_tochukaso/items/da426190a4563c1dbebd)を参考に対応する。

コーディング中に出力されたメッセージは以下

```
gopls was not able to find modules in your workspace.
When outside of GOPATH, gopls needs to know which modules you are working on.
You can fix this by opening your workspace to a folder inside a Go module, or
by using a go.work file to specify multiple modules.
See the documentation for more information on setting up your workspace:
https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.go
```

`mkdir go.mod` で対応した。
