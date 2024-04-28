package frontend

import (
	"embed"
	"io/fs"
	"log/slog"
	"os"
)

//go:embed dist
var EMBED_UI embed.FS

func GetUiFs() fs.FS {
	embedRoot, err := fs.Sub(EMBED_UI, "build")
	if err != nil {
		slog.Error("Unable to get root for web ui", slog.String("error", err.Error()))
		os.Exit(1)
	}
	return embedRoot
}
