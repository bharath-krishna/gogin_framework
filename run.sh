#/bin/sh

export PATH="$PATH:$GOPATH/bin"
reflex -r '\.go$' --inverse-regex="/docs\/*.go" --shutdown-timeout=500ms -s --  sh -c "swag init; go build -o app; ./app"