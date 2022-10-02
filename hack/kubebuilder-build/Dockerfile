FROM golang:1.19.1-buster

RUN curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH) && \
    chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
