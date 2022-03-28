# Kemomimi-OJS Worker

## About
Dockerコンテナ内で自動的にコンパイルとコードの実行を行い,  
テストケースと照合し結果を返します

## Usage 
```shell
$ go build main.go
$ ./main <CompilerType> <TestCaseID>
```
※ 同じ階層に`test-cace/<ID>/(ans|test)_{ID}.txt` が存在しないとエラーになります
