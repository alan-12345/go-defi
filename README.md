# Generate Go file from ABI
`abigen --abi PATH/TO/ABI/Abi.abi --pkg PACKAGE_NAME --type NAME --out PATH/TO/EXPORT/file.go`

# Clean up unused packages
`go mod tidy`

# Build .exe
`go build PATH/TO/MAIN/*.go`