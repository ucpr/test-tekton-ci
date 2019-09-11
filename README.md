See below for Tekton installation.
https://github.com/tektoncd/pipeline/blob/master/docs/install.md

説明
----
- `go test`と`docker image`のbuild, pushまで可能

usage
-----
- `tekton/`にあるyamlを`kubectl apply -f`で.
- docker hubにpushするために以下の設定が必要
  - https://github.com/tektoncd/pipeline/blob/master/docs/auth.md
