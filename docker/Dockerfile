FROM progrium/busybox
MAINTAINER leetchang@gmail.com

# Copy toolkit binary files
ADD kubeme-toolkit /kubeme-toolkit

EXPOSE 8001/tcp

ENTRYPOINT [ "/kubeme-toolkit" ]
