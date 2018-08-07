FROM alpine:3.6

RUN adduser -D redeploy-operator
USER redeploy-operator

ADD tmp/_output/bin/redeploy-operator /usr/local/bin/redeploy-operator
