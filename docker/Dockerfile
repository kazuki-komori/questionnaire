FROM golang:1.15.2-alpine
ENV GO111MODULE=off
#ARG MYSQL_PASS
# アップデートとgitのインストール！！
RUN apk update && apk add git
# ワーキングディレクトリの設定
WORKDIR /go/src/github.com/kazuki-komori/questionnaire_server
COPY . .
# ホットリロード
RUN go get -u github.com/oxequa/realize
#CMD ["realize", "start"]
#RUN  go get -u github.com/cosmtrek/air
#CMD ["air"]
# go-mod
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/kazuki-komori/questionnaire_server
#RUN go get github.com/pilu/fresh
#CMD ["fresh"]
