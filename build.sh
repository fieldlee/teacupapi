#! /bin/bash
export logfile=./logs/teacupapi.log
export binaryName=teacupapi
export imageName=teacupapi:latest
export containerName=teacupapi
export containerVolumn=/data/teacupapi/logs
export nodeVolumn=/home/ubuntu/logs
export outdir=/home/ubuntu/data
export indir=/data/teacupapi/data

function reload() {
  pid=$(ps -ef | grep $binaryName | grep -v grep | awk '{print $2}')
  kill -HUP $pid
  start
  sleep 1
  newpid=$(ps -ef | grep $binaryName | grep -v grep | awk '{print $2}')
  echo "reload..., pid=$newpid"
}

function start() {
  nohup ./$binaryName &
}

function stop() {
  pid=$(ps -ef | grep $binaryName | grep -v grep | awk '{print $2}')
  echo $pid
  kill -9 $pid
  echo "apigateway stop"
}

function restart() {
  stop
  sleep 1
  start
}

function gitPull() {
  sudo git pull
  if [ $? -ne 0 ]; then
    sudo git stash
    sudo git pull
    sudo git stash clean
  fi
  echo "git pull current branch success"
}

function build() {
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -ldflags "-s -w" -o $binaryName ./cmd/api 
  echo "build success"
}

function iosbuild() {
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags=jsoniter -ldflags "-s -w" -o $binaryName ./cmd/api 
  echo "ios build success"
}

function dockerDevBuild() {
  docker build --build-arg config=./config/app.dev.ini -t $imageName .
  echo "dev docker build success"
}

function dockerPreBuild() {
  docker build --build-arg config=./config/app.pre.ini -t $imageName .
  echo "pre docker build success"
}

function dockerProdBuild() {
  docker build --build-arg config=./config/app.prod.ini -t $imageName .
  echo "prod docker build success"
}

function dockerStop() {
  docker stop $containerName && docker rm -f $containerName
  echo "docker stop success"
}

function dockerRun() {
  docker run --restart=always --network=host --name $containerName -idt  -v $outdir:$indir -v $nodeVolumn:$containerVolumn $imageName
  echo "docker run success"
}

function dockerClean() {
  docker images | grep none | awk '{print $3 }' | xargs docker rmi
}

function dockerDevDeploy() {
  gitPull
  build
  dockerDevBuild
  dockerStop
  dockerRun
}

function dockerPreDeploy() {
  build
  dockerPreBuild
  dockerStop
  dockerRun
}

function dockerProdDeploy() {
  build
  dockerProdBuild
  dockerStop
  dockerRun
}

function tailf() {
  tail -f $logfile
}

function help() {
  echo "$0 start|stop|restart"
}

if [ "$1" == "" ]; then
  help
elif [ "$1" == "start" ]; then
  start
elif [ "$1" == "stop" ]; then
  stop
elif [ "$1" == "restart" ]; then
  restart
elif [ "$1" == "reload" ]; then
  reload
elif [ "$1" == "build" ]; then
  build
elif [ "$1" == "iosbuild" ]; then
  iosbuild
elif [ "$1" == "dockerRun" ]; then
  dockerRun
elif [ "$1" == "dockerStop" ]; then
  dockerStop
elif [ "$1" == "dockerDevDeploy" ]; then
  dockerDevDeploy
elif [ "$1" == "dockerPreDeploy" ];then
    dockerPreDeploy
elif [ "$1" == "dockerProdDeploy" ];then
    dockerProdDeploy
elif [ "$1" == "dockerClean" ]; then
  dockerClean
elif [ "$1" == "tail" ]; then
  tailf
else
  help
fi
