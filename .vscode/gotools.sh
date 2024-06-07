#!/usr/bin/env sh

GO_TOOLS="\
    golang.org/x/tools/gopls@latest \
    golang.org/x/lint/golint@latest \
    github.com/go-delve/delve/cmd/dlv@latest"

(echo "${GO_TOOLS}" | xargs -n 1 go install -v )2>&1 | tee -a /$GOPATH/gotools.log

# GO_TOOLS="\
#     golang.org/x/tools/gopls@latest \
#     honnef.co/go/tools/cmd/staticcheck@latest \
#     golang.org/x/lint/golint@latest \
#     github.com/mgechev/revive@latest \
#     github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
#     github.com/ramya-rao-a/go-outline@latest \
#     github.com/go-delve/delve/cmd/dlv@latest \
#     github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
