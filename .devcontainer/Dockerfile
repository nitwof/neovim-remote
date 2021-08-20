ARG goVersion="1.17"
FROM golang:${goVersion}

ENV WORKSPACE_PATH=/workspace \
    DATA_PATH=/data \
    HIST_PATH=/data/history
ENV \
    LANG=C.UTF-8 \
    EDITOR=vim \
    HISTFILE=${HIST_PATH}/.bash_history

RUN \
    export DEBIAN_FRONTEND="noninteractive" && \
    apt-get update -qq && \
    apt-get -y dist-upgrade && \
    # Install packages
    apt-get update -qq && \
    apt-get install -y --no-install-recommends \
        bash-completion \
        build-essential \
        ca-certificates \
        curl \
        file \
        git \
        make \
        pkg-config \
        sudo \
        tzdata \
        unzip \
        vim \
        wget \
    && \
    # Setup data dir
    mkdir -p ${WORKSPACE_PATH} ${HIST_PATH} && \
    # Clean up
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    truncate -s 0 /var/log/*log
##################################################

ARG golangcilintVersion="1.42.0"
ARG goreleaserVersion="0.175.0"

# Install dev tools
##################################################
RUN \
    # Install goreleaser
    curl -sL "https://github.com/goreleaser/goreleaser/releases/download/v${goreleaserVersion}/goreleaser_Linux_x86_64.tar.gz" -o /tmp/goreleaser.tar.gz && \
    mkdir /tmp/goreleaser && \
    tar -xvf /tmp/goreleaser.tar.gz -C /tmp/goreleaser && \
    cp /tmp/goreleaser/goreleaser /usr/local/bin/goreleaser && \
    cp /tmp/goreleaser/completions/goreleaser.bash /etc/bash_completion.d/goreleaser && \
    # Install golangci-lint
    curl -sL "https://github.com/golangci/golangci-lint/releases/download/v${golangcilintVersion}/golangci-lint-${golangcilintVersion}-linux-amd64.tar.gz" -o /tmp/golangci-lint.tar.gz && \
    mkdir /tmp/golangci-lint && \
    tar -xvf /tmp/golangci-lint.tar.gz -C /tmp && \
    cp /tmp/golangci-lint-1.42.0-linux-amd64/golangci-lint /usr/local/bin/golangci-lint && \
    # Clean up
    rm -rf /tmp/*
##################################################

# Setup user and permissions
##################################################
ARG user="docker"
ARG userUID="1000"
ARG userGID="1000"

ENV VSCODE_PATH=/home/${user}/.vscode-server

# Setup user and update permissions
##################################################
RUN \
    groupadd --gid=${userGID} ${user} && \
    useradd -m --uid=${userUID} --gid=${userGID} ${user} && \
    mkdir -p ${VSCODE_PATH} && \
    echo "${user} ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/${user} && \
    chown -R ${userUID}:${userGID} ${WORKSPACE_PATH} ${DATA_PATH} ${VSCODE_PATH}

USER ${userUID}:${userGID}
##################################################

ENV GOPATH=${DATA_PATH}/go
ENV GOBIN=${GOPATH}/bin
ENV PATH=${GOBIN}:${PATH}

WORKDIR ${WORKSPACE_PATH}
VOLUME ["${WORKSPACE_PATH}", "${DATA_PATH}", "${VSCODE_PATH}"]
