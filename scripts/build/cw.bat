echo on
setlocal

set GOOS=windows
set GOARCH=amd64

go build -o ../../../bin/windows/auth_microservice.exe ../../cmd/auth_microservice/main.go

if %errorlevel% == 0 (
    copy ../../configs/auth_microservice.conf ../../../bin/windows/auth_microservice.conf
    echo "Build successful!"
) else (
    echo "Build failed!"
)

endlocal