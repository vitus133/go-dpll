FROM registry.redhat.io/rhel9/go-toolset:1.25 AS builder
COPY . .
RUN make build

FROM registry.redhat.io/rhel9/support-tools:latest
WORKDIR /
USER root

RUN --mount=type=secret,id=rh_user --mount=type=secret,id=rh_pass \
    user="$(cat /run/secrets/rh_user)" && \
    pass="$(base64 -d /run/secrets/rh_pass)" && \
    subscription-manager register --username "${user}" --password "${pass}" && \
    /bin/dnf install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-9.noarch.rpm \
    python3-pip lshw pciutils git ethtool jq linuxptp gcc make hwdata synce4l gpsd-minimal gpsd-minimal-clients i2c-tools && dnf clean all && \
    ln -s /usr/bin/gpspipe /usr/local/bin/gpspipe && ln -s /usr/sbin/gpsd /usr/local/sbin/gpsd && ln -s /usr/bin/ubxtool /usr/local/bin/ubxtool &&\
    subscription-manager unregister && subscription-manager clean &&\
    git clone --depth=1 https://git.kernel.org/pub/scm/linux/kernel/git/netdev/net-next.git &&\ 
    rm -rf $(find /net-next -mindepth 1 -maxdepth 1 \( -name 'tools' -o -name 'Documentation' \) -prune -o -print |xargs) &&\
    pip install -r /net-next/tools/net/ynl/requirements.txt 

COPY --from=builder /opt/app-root/src/dpll-cli /usr/local/bin/
WORKDIR /net-next/tools/net/ynl/pyynl


