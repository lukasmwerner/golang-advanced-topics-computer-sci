#!/bin/bash
cd /root/noVNC-1.2.0
tigervncserver -xstartup /usr/bin/startlxde
/root/noVNC-1.2.0/utils/launch.sh --listen 8080 --vnc localhost:5901