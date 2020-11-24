FROM ubuntu:bionic


RUN apt-get update && \
    apt-get --no-install-recommends -y install  software-properties-common && \
    add-apt-repository ppa:otto-kesselgulasch/gimp 

RUN apt-get update && \
    apt-get --no-install-recommends install -y  firefox kdenlive gimp


RUN apt-get install -y tigervnc-standalone-server 
#tigervnc-xorg-extension

ENV TERM=xterm
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get -y --no-install-recommends install wget git unzip python-numpy

RUN apt-get --no-install-recommends install -y xterm lxde openbox

ENV USER=user

WORKDIR /root

RUN wget https://github.com/novnc/noVNC/archive/v1.2.0.zip && unzip v1.2.0.zip 

WORKDIR /root/noVNC-1.2.0

COPY .vnc/ /root/.vnc/

COPY startup.sh startup.sh

#CMD bash
CMD /root/noVNC-1.2.0/startup.sh