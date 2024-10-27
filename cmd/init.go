package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	"sync-it/common"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		configPath := path.Join(cmd.Flag(common.CMDInitFlagDir).Value.String(), common.ConfigFilename)
		if _, err := os.Stat(configPath); err == nil {
			print("Configuration file already exists at", configPath)
			return
		}
		err := writeConfigToFile(common.Configuration{}, configPath)
		if err != nil {
			fmt.Printf("Error saving configuration: %v \n at %s", err, configPath)
			return
		}
		createDirectory(common.DefaultLogPath)
		println("Configuration initialized successfully!")
		println("Configuration file saved at:", configPath)
		println("Use 'sync add' to set up specific sync tasks or 'sync run' to start syncing.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")
	initCmd.PersistentFlags().String(common.CMDInitFlagDir, common.DefaultConfigDir, "set directory")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initCmd.Flags().BoolP("test", "t", false, "turn on/off test mode")
}

func writeConfigToFile(cfg common.Configuration, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	defer file.Close()

	if err = json.NewEncoder(file).Encode(cfg); err != nil {
		return err
	}
	return err
}

func createDirectory(path string) {
	_ = os.WriteFile(path, []byte{}, 0644)
}
