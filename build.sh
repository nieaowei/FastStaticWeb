CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags v1.0 -o bin/faststaticweb-drawin ./main
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags v1.0 -o bin/faststaticweb-linux ./main
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags v1.0 -o bin/faststaticweb-windows ./main