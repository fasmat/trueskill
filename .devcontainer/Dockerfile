# See here for image contents: https://github.com/devcontainers/images/blob/v0.4.15/src/base-ubuntu/.devcontainer/Dockerfile

# [Choice] Ubuntu version: ubuntu-24.04, ubuntu-22.04, ubuntu-20.04
ARG VARIANT=ubuntu-24.04
FROM mcr.microsoft.com/devcontainers/base:${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    bash-completion \
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/*

RUN curl https://raw.githubusercontent.com/docker/docker-ce/master/components/cli/contrib/completion/bash/docker \
    -o /etc/bash_completion.d/docker.sh
