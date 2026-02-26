FROM registry.redhat.io/rhel9/go-toolset:latest AS builder
COPY . .
RUN make build

FROM registry.redhat.io/rhel9/support-tools:latest
WORKDIR /
USER root

RUN --mount=type=secret,id=rh_user --mount=type=secret,id=rh_pass \
    user="$(cat /run/secrets/rh_user)" && \
    pass="$(base64 -d /run/secrets/rh_pass)" && \
    subscription-manager register --username "${user}" --password "${pass}" && \
    subscription-manager repos --enable=rhel-9-for-x86_64-baseos-rpms --enable=rhel-9-for-x86_64-appstream-rpms && \
    /bin/dnf install -y man-db coreutils-common https://dl.fedoraproject.org/pub/epel/epel-release-latest-9.noarch.rpm \
    gcc make bison flex libmnl-devel libcap-devel elfutils-libelf-devel ncurses-devel libmnl-devel sysfsutils perf  \
    python3-pip lshw pciutils sysstat git ethtool jq linuxptp hwdata synce4l gpsd-minimal gpsd-minimal-clients i2c-tools acpica-tools && dnf clean all && \
    ln -s /usr/bin/gpspipe /usr/local/bin/gpspipe && ln -s /usr/sbin/gpsd /usr/local/sbin/gpsd && ln -s /usr/bin/ubxtool /usr/local/bin/ubxtool && \
    subscription-manager unregister && subscription-manager clean && \
    git clone --depth=1 https://git.kernel.org/pub/scm/linux/kernel/git/netdev/net-next.git && \ 
    rm -rf $(find /net-next -mindepth 1 -maxdepth 1 \( -name 'tools' -o -name 'Documentation' \) -prune -o -print |xargs) && \
    pip install -r /net-next/tools/net/ynl/requirements.txt  && \
    git clone  https://git.kernel.org/pub/scm/network/iproute2/iproute2.git && cd iproute2 && \ 
    ./configure && make && make install && cd .. && rm -rf iproute2


COPY --from=builder /opt/app-root/src/dpll-cli /usr/local/bin/
WORKDIR /net-next/tools/net/ynl/pyynl


