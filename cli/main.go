package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	services "ai-service"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "krouly",
	Short: "CLI tool for managing data extraction workflows",
}

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new data extraction client",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Creating data extraction workflow app: %s\n", name)

		clientDir := filepath.Join("..", "client", name)
		if err := os.MkdirAll(clientDir, 0755); err != nil {
			fmt.Println(err)
			return
		}

		// Install Preact and Preact Router in the client directory
		if err := installPreact(clientDir); err != nil {
			fmt.Println(err)
			return
		}

		// Generate Preact files in the client directory
		if err := generateMainJS(clientDir); err != nil {
			fmt.Println(err)
			return
		}
		if err := generateIndexHTML(clientDir, name); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Preact setup completed successfully.")
	},
}

var runCmd = &cobra.Command{
	Use:   "run [name]",
	Short: "Run the Preact app",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Running the Preact app for: %s\n", name)

		clientDir := filepath.Join("..", "client", name) // Adjust app name and directory as needed
		webappDir := filepath.Join(clientDir, "webapp")

		runCommand := exec.Command("npm", "start")
		runCommand.Dir = webappDir
		runCommand.Stdout = os.Stdout
		runCommand.Stderr = os.Stderr
		if err := runCommand.Run(); err != nil {
			fmt.Println(err)
			return
		}
	},
}

var runAICmd = &cobra.Command{
	Use:   "run-ai",
	Short: "Run the AI inference on extracted data",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := "../storage/cryptodata.json"
		predictions, err := services.RunAIInference(filePath)
		if err != nil {
			fmt.Println("Error running AI inference:", err)
			return
		}

		fmt.Println("Predictions:", predictions)
	},
}

func installPreact(targetDir string) error {
	cmd := exec.Command("npm", "install", "preact", "preact-router")
	cmd.Dir = targetDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to install Preact in %s: %v\n%s", targetDir, err, output)
	}
	return nil
}

func generateMainJS(targetDir string) error {
	mainJS := `import { h, render } from 'preact';
import App from './components/App';

render(<App />, document.getElementById('root'));`
	err := os.MkdirAll(filepath.Join(targetDir, "webapp", "static", "js"), 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	err = os.WriteFile(filepath.Join(targetDir, "webapp", "static", "js", "main.js"), []byte(mainJS), 0644)
	if err != nil {
		return fmt.Errorf("failed to generate main.js: %v", err)
	}
	return nil
}

func generateIndexHTML(targetDir, name string) error {
	indexHTML := fmt.Sprintf(`<html>
<head>
	<title>%s - Powered by Preact</title>
	<script src="/static/js/preact.min.js"></script>
	<script src="/static/js/preact-router.min.js"></script>
</head>
<body>
	<div id="root">Hi friend! My name is %s.</div>
	<script src="/static/js/main.js"></script>
</body>
</html>`, name, name)
	err := os.MkdirAll(filepath.Join(targetDir, "webapp", "views"), 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	err = os.WriteFile(filepath.Join(targetDir, "webapp", "views", "index.html"), []byte(indexHTML), 0644)
	if err != nil {
		return fmt.Errorf("failed to generate index.html: %v", err)
	}
	return nil
}

func main() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(runAICmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
