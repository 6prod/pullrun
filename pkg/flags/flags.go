package flags

import (
	"flag"
	"fmt"
	"os"
)

const (
	FlagDefaultDir     = "/var/local/pullrun"
	FlagDefaultDirPerm = 0750
)

const (
	FlagConfig       = "config"
	FlagDeleteConfig = "delete-config"
	FlagDir          = "dir"
	FlagDirPerm      = "dir-perm"
)

var Config = flag.String(FlagConfig, "", "configuration file or json configuration between single quotes")
var ConfigDelete = flag.Bool(FlagDeleteConfig, false, "delete the configuration once loaded")
var Dir = flag.String(FlagDir, FlagDefaultDir, "download dir, created if missing")
var DirPerm = flag.Uint(FlagDirPerm, FlagDefaultDirPerm, "download dir permission")

func Parse() {
	flag.Parse()
}

func Lookup(name string) *flag.Flag {
	return flag.Lookup(name)
}

func Usage() {
	cmdOut := flag.CommandLine.Output()
	fmt.Fprintf(cmdOut, "Usage: %s [flags] command\n", os.Args[0])
	fmt.Fprintf(cmdOut, "\nPull then run commands\n")
	fmt.Fprintf(cmdOut, "\nFlags:\n")
	flag.PrintDefaults()
}
