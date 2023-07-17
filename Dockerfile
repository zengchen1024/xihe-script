FROM golang:latest as BUILDER

# build binary
RUN mkdir -p /go/src/mindspore/xihe-script

COPY . /go/src/mindspore/xihe-script

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN cd /go/src/mindspore/xihe-script && go mod tidy && CGO_ENABLED=1 go build -v -o ./xihe-script .

# copy binary config and utils
FROM openeuler/openeuler:22.03

RUN yum update -y && yum install -y python3 && yum install -y python3-pip

RUN mkdir -p /opt/app/xihe-script/py/data

COPY ./py /opt/app/xihe-script/py

RUN chmod 755 -R /opt/app/xihe-script/py

ENV EVALUATE /opt/app/xihe-script/py/evaluate.py
ENV CALCULATE /opt/app/xihe-script/py/calculate_fid.py
ENV UPLOAD /opt/app/xihe-script/py/data/

RUN pip install -r /opt/app/xihe-script/py/requirements.txt

COPY --from=BUILDER /go/src/mindspore/xihe-script/xihe-script /opt/app/xihe-script
WORKDIR /opt/app/xihe-script/

ENTRYPOINT ["/opt/app/xihe-script/xihe-script"]