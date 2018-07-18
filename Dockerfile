From ubuntu:18.04 

MAINTAINER yusuf <mhyusufibrahim@gmail.com>

RUN apt-get -y update && apt-get -y install build-essential vim wget git &&\
    wget -q --show-progress https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz &&\
    tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz &&\
    echo 'export PATH=$PATH:/usr/local/go/bin\nexport GOPATH=/root/go\nexport PATH=$PATH:/root/go/bin' >> /root/.bashrc 

WORKDIR /root/go/src/app/Foreign-Currency-BE-Exercise

ENV GOROOT="/usr/local/go"
ENV GOPATH="/root/go"

ENV  PATH="/usr/local/go/bin:/root/go/bin:${PATH}"

RUN  /usr/local/go/bin/go get github.com/mattn/go-sqlite3 &&\
     /usr/local/go/bin/go get github.com/go-sql-driver/mysql &&\	
     /usr/local/go/bin/go get github.com/lib/pq &&\
     /usr/local/go/bin/go get github.com/jinzhu/gorm &&\
     /usr/local/go/bin/go get github.com/revel/revel &&\
     /usr/local/go/bin/go get github.com/revel/cmd/revel

CMD bash -c "/root/go/bin/revel run"
