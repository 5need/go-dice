go build -ldflags "-s -w" -o ./bin/go-dice ./
# TODO: make a proper build script
# rsync -uvrP --info=progress2 ./ root@go-dice://root/go-dice/
# ssh root@go-dice "systemctl restart go-dice"
