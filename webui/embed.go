package webui

import (
	"embed"
)

//go:embed templates/*
var EmbedTemplates embed.FS

//go:embed assets/*
var EmbedAssets embed.FS
