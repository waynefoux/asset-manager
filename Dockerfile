# I copied this from an old work project and wanted to see if I can mod it for my use case.

FROM golang:1.20
LABEL description="Asset Management"
LABEL maintainer="ReadyRunBreak"
LABEL version="1.0"

ENV APP_HOME /opt/app/asset-manager

# I love utils!
#RUN apk update && \
#    apk add --no-cache bash python ca-certificates wget curl
RUN apt-get update && \
	apt-get -y install bash python wget curl

ADD https://github.com/energyhub/secretly/releases/download/0.0.6/secretly-linux-amd64 /usr/sbin/secretly

RUN chmod a+x /usr/sbin/secretly

# Copy app to directory
WORKDIR ${APP_HOME}
RUN mkdir bin conf logs
COPY build/linux/asset-manager bin/
COPY tools/launch.sh bin/
RUN chmod +x bin/launch.sh


# Start app
CMD secretly env bash ${APP_HOME}/bin/launch.sh