package main

import (
	"log"

	"github.com/go-kratos/kratos/cmd/kratos/v2/internal/project"
	"github.com/go-kratos/kratos/cmd/kratos/v2/internal/proto"
	"github.com/go-kratos/kratos/cmd/kratos/v2/internal/upgrade"
	"github.com/spf13/cobra"
)

var (
	version string = "v2.0.0-beta2"

	rootCmd = &cobra.Command{
		Use:     "kratos",
		Short:   "Kratos: An elegant toolkit for Go microservices.",
		Long:    `Kratos: An elegant toolkit for Go microservices.`,
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
