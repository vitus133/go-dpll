FROM registry.ci.openshift.org/ocp/builder:rhel-9-golang-1.20-openshift-4.14 AS builder
RUN ls -al
RUN pwd
COPY . .
RUN make build

FROM registry.access.redhat.com/ubi9-micro
WORKDIR /
USER root
COPY --from=builder /go/src/github.com/openshift/origin/dpll-cli /usr/local/bin/

