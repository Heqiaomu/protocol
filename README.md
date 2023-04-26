# README

本项目用于存放公共protobuf

- ui: 与前端显示相关的message定义

## 执行下面命令进行编译

```sh
./protosh
```

## 使用方法

通过拷贝proto文件然后进行远程调用

以使用ui/input.proto为例

1. 创建`third_party/proto/ui`目录
2. 在`gomod.sh`中添加依赖以及下载最新input.proto
3. 目标proto（以target.proto为例）中引用input.proto
4. 目标proto（以target.proto为例）中使用input.proto定义的类型
5. `proto.sh`中引用third_party/proto/ui, 注意添加引用路径

```
# 1. 创建目录
mkdir third_party/proto/ui

# 2. 下载最新input.proto
go get -insecure github.com/Heqiaomu/protocol

echo "====> copy protocol/ui/input.proto"
F=$(grep protocol  go.mod | sed "s/l v/l@v/g" | awk '{print $1}')
cp -f $GOPATH"/pkg/mod/"$F"/ui/input.proto" ./third_party/proto/ui/

# 3. 在target.proto中引用input.proto
import "ui/input.proto";

# 4. 在target.proto中使用input.proto中定义的类型
 ui.Input CustomInput = 1;

 # 5. 引用，执行
protoc -I. -I/${WORKPATH}/third_party/proto/ --go_out=plugins=grpc:. target.proto
```
