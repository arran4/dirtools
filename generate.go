package dirtools

//go:generate sh -c "(command -v gosubc >/dev/null 2>&1 && gosubc generate || go run github.com/arran4/go-subcommand/cmd/gosubc generate) && sed -i '/\"github.com\/arran4\/dirtools\"/d' cmd/dirquery/root.go cmd/extdirisolate/root.go"
