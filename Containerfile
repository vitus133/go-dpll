FROM registry.redhat.io/rhel9/go-toolset:1.21 AS builder
COPY . .
RUN make build

FROM registry.redhat.io/rhel9/support-tools:latest
WORKDIR /
USER root
RUN dnf install -y git python3 python3-pip lshw && \
    git clone --depth=1 https://github.com/torvalds/linux.git && \
    pip install -r /linux/tools/net/ynl/requirements.txt && \
    dnf remove -y git

WORKDIR /linux/tools/net/ynl/pyynl

# Uncomment this line if you want cli to block while waiting on netlink notifications
# RUN sed -i 's/, socket.MSG_DONTWAIT//g' lib/ynl.py 
COPY --from=builder /opt/app-root/src/dpll-cli /usr/local/bin/

