# go-bullet
Bullet-hell shooter with Go's coroutines

## Live reloading
Install makiuchi-d/arelo:

```
go install github.com/makiuchi-d/arelo@latest
```

Start live reloading:

```
arelo -p '**/*.go' -i '**/.*' -i '**/*_test.go' -- ./run-wsl.sh
```
