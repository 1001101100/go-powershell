# run fmt vet and lint

go fmt github.com/1001101100/go-powershell/pkg/powershell
go fmt github.com/1001101100/go-powershell/pkg/logger
go fmt github.com/1001101100/go-powershell/examples/simple
go fmt github.com/1001101100/go-powershell/examples/cmd

go vet github.com/1001101100/go-powershell/pkg/powershell
go vet github.com/1001101100/go-powershell/pkg/logger
go vet github.com/1001101100/go-powershell/examples/simple
go vet github.com/1001101100/go-powershell/examples/cmd

golint github.com/1001101100/go-powershell/pkg/powershell
golint github.com/1001101100/go-powershell/pkg/logger
golint github.com/1001101100/go-powershell/examples/simple
golint github.com/1001101100/go-powershell/examples/cmd