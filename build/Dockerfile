FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/metal3-operator \
    USER_UID=1001 \
    USER_NAME=metal3-operator

# install templates
COPY templates /usr/local/templates

# install operator binary
COPY build/_output/bin/metal3-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
