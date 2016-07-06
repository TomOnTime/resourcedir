#!/bin/sh
 
# PROVIDE: velma_CHANGE_TO_INSTANCE
# REQUIRE: LOGIN FILESYSTEMS
# KEYWORDS: start stop restart status poll enabled rcvar

 
. /etc/rc.subr
 
instance_name="CHANGE_TO_INSTANCE"
name="velma_$instance_name"
rcvar=`set_rcvar`
pidfile="/var/run/${name}.pid"
command=daemon
start_cmd="$command -c -u velmad -r -P $pidfile -- /usr/local/bin/$name -flagfile /etc/velma/$instance_name.conf >/var/log/${name}.log"

load_rc_config $name
run_rc_command "$1"
