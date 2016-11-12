# chech that go version is >= 1.7
go version

mkdir workspace
cd workspace
export GOPATH=`pwd`
go get github.com/estarter/mvnc
