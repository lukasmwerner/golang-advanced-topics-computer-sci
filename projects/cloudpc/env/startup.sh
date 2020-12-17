#!/bin/bash
cd /root/noVNC-1.2.0
tigervncserver -xstartup /usr/bin/startlxde -SecurityTypes None
/root/noVNC-1.2.0/utils/launch.sh --listen 10000 --vnc localhost:5901