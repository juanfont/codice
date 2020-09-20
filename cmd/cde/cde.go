package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	codice "github.com/juanfont/codice"
	"github.com/spf13/cobra"
)

const version = "0.1"

// The version command prints this service.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Long:  "The version of the tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

var codiceCmd = &cobra.Command{
	Use:   "codice",
	Short: "codice - a CLI tool to work with CODICE files from contrataciondelestado.es",
	Long: fmt.Sprintf(`
	codice is a CLI tool to work with contrataciondelestado.es files.
v%s 
Juan Font Alonso <juanfontalonso@gmail.com> - 2020
https://github.com/juanfont/codice`, version),
}

var xmlFileCmd = &cobra.Command{
	Use:   "xml file.xml output_base",
	Short: "Parses a single XML from the filesystem",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires three args")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		c, err := codice.NewCodiceApp()
		if err != nil {
			log.Fatalln(err)
		}
		entries, err := c.LoadXMLFromFs(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
		err = c.FlattenToCsv(entries, os.Args[3])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var zipCmd = &cobra.Command{
	Use:   "zip file.zip output_prefix",
	Short: "Parses a ZIP",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires three args")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		c, err := codice.NewCodiceApp()
		if err != nil {
			log.Fatalln(err)
		}
		entries, err := c.LoadWebZip(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
		err = c.FlattenToCsv(entries, os.Args[3])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func main() {
	addCommands()

	if err := codiceCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func addCommands() {
	codiceCmd.AddCommand(versionCmd)
	codiceCmd.AddCommand(xmlFileCmd)
	codiceCmd.AddCommand(zipCmd)
}
