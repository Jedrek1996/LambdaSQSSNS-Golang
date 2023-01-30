Steps to convert into zip and deploy onto AWS Lambda

//Change env to linux
$ENV:GOOS = "linux"     

//Build as LAMBDA only reads binary
go build main.go
