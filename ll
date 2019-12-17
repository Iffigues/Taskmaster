#! /bin/sh -
echo Starting
/usr/local/bin/unshare -fmp -- sh -c '
  umount /proc
  mount -nt proc p /proc
  exec bash <&2' &
ifconfig lo 127.1/8
exec socat tcp-listen:1234,fork,reuseaddr system:"ps -efH; echo still running"
