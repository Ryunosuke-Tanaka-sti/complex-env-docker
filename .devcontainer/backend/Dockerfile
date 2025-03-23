ARG GO_VER
FROM golang:${GO_VER}

RUN apt update \
  && apt install -y vim

# 非rootユーザーを作成
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

ENV GO111MODULE on
    
# 所有ディレクトリ
WORKDIR /workspace

# 所有権を変更
RUN chown -R $USERNAME:$USERNAME /workspace

# ユーザーを切り替え
USER $USERNAME

COPY --chown=$USERNAME:$USERNAME ./go-app .

