env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o aws-ss-macos_x64
env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o aws-ss-linux_x64
zip aws-ss.zip aws-ss-macos_x64 aws-ss-linux_x64