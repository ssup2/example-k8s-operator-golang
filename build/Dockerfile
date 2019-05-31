FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/example-k8s-operator-memcached \
    USER_UID=1001 \
    USER_NAME=example-k8s-operator-memcached

# install operator binary
COPY build/_output/bin/example-k8s-operator-memcached ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
