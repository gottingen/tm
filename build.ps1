# For `--version`
$PD_PKG = "github.com/gottingen/tm"
$GO_LDFLAGS = "-X `"$PD_PKG/server.PDReleaseVersion=$(git describe --tags --dirty --always)`""
$GO_LDFLAGS += " -X `"$PD_PKG/server.PDBuildTS=$(date -u '+%Y-%m-%d_%I:%M:%S')`""
$GO_LDFLAGS += " -X `"$PD_PKG/server.PDGitHash=$(git rev-parse HEAD)`""
$GO_LDFLAGS += " -X `"$PD_PKG/server.PDGitBranch=$(git rev-parse --abbrev-ref HEAD)`""

# Download Dashboard UI
powershell.exe -File ./scripts/embed-dashboard-ui.ps1

# Output binaries
go build -ldflags $GO_LDFLAGS -o bin/tm-server.exe cmd/tm-server/main.go
echo "bin/tm-server.exe"
go build -ldflags $GO_LDFLAGS -o bin/tm-ctl.exe tools/tm-ctl/main.go
echo "bin/tm-ctl.exe"
go build -o bin/tm-tso-bench.exe tools/tm-tso-bench/main.go
echo "bin/tm-tso-bench.exe"
go build -o bin/tm-recover.exe tools/tm-recover/main.go
echo "bin/tm-recover.exe"
