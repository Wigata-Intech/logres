package assets

import "embed"

//go:embed css/* fonts/* images/* js/*
var Assets embed.FS
