#!/bin/bash

action=${1}

pid_exists(){
   test -d /proc/"$1"
}

# get project dir
baseDir=$(cd `dirname $0`;cd ..;pwd)
cd ${baseDir}

port=5601
env="dev"
grpcServerName="usrinfo_grpc_server"
grpcBuildDir="${baseDir}/src/grpc_server"
httpServerName="usrinfo_http_server"
httpBuildDir="${baseDir}/src/http_server"
binDir="${baseDir}/bin"
configDir="${baseDir}/config"

daemon_is_running(){
   serverName=$1

   [ ! -e "${baseDir}/run/${serverName}.pid" ] && return 1

   pid=$(cat "${baseDir}/run/${serverName}.pid")

   test -z ${pid} && return 1
   test -n "$pid" && pid_exists "$pid"
} >/dev/null 2>&1

buildServer(){
   serverName=$1
   buildDir=$2
   [ -x ${binDir}${serverName} ] && rm -rf ${binDir}/${serverName}
   cd "${buildDir}" && go build -o ${binDir}/${serverName}
   echo "Compile ${serverName} successfully."
   cd "${baseDir}"
}

startServer(){
   serverName=$1
   cd ${baseDir}
   [ -d "${baseDir}/run" ] || mkdir -p "${baseDir}/run"
   [ -d "${baseDir}/log" ] || mkdir -p "${baseDir}/log"

   if daemon_is_running "${serverName}"; then
      echo "Start ${serverName} FAILED.${serverName} is already running."
      return
   fi

   if [ ! -f "${configDir}/seelog.xml" ];then
      echo "Start ${serverName} FAILED.File ${configDir}/seelog.xml not exist."
      return
   fi
   cd ${baseDir} && ./bin/${serverName} > ${baseDir}/log/nohup.out 2>&1 &
   pid=$!
   sleep 1
   if kill -0 ${pid}; then
      echo "${pid}" > "${baseDir}/run/${serverName}.pid"
      echo "Start ${serverName} OK."
   else
      wait ${pid}; daemonexit=$?
      echo "Start ${serverName} FAILED."
      return 1
   fi
   cd "${baseDir}"
}

stopServer(){
   serverName=$1
   if [ ! -e "${baseDir}/run/${serverName}.pid" ];then
      echo "${serverName} is not running."
      return
   fi
   pid=$(cat "${baseDir}/run/${serverName}.pid")
   if [ X"${pid}" != X"" ] && [ X"${pid}" != X"1" ]
   then
      kill -9 ${pid}
   fi
   rm -f ${baseDir}/run/${serverName}.pid
   echo "Stop ${serverName} OK."
}

build(){
   buildServer ${grpcServerName} ${grpcBuildDir}
   buildServer ${httpServerName} ${httpBuildDir}
}

start(){
   startServer ${grpcServerName}
   startServer ${httpServerName}
}

stop(){
   stopServer ${grpcServerName}
   stopServer ${httpServerName}
}

case ${action} in
   start)
      start
   ;;
   stop)
      stop
   ;;
   restart)
      stop
      sleep 1
      start
   ;;
   build)
      build
   ;;
   *)
      echo "Usage:sh ${0} start|stop|restart|build"
   ;;
esac