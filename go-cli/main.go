package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const version = "1.0.0"

type SystemInfo struct {
	OS           string    `json:"os"`
	Architecture string    `json:"architecture"`
	GoVersion    string    `json:"go_version"`
	NumCPU       int       `json:"num_cpu"`
	Timestamp    time.Time `json:"timestamp"`
}

func main() {
	var (
		showVersion = flag.Bool("version", false, "Show version information")
		showSystem  = flag.Bool("system", false, "Show system information")
		jsonOutput  = flag.Bool("json", false, "Output in JSON format")
		greet       = flag.String("greet", "", "Greet someone")
		uppercase   = flag.Bool("upper", false, "Convert greeting to uppercase")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "OSS Release Demo CLI Tool v%s\n\n", version)
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s --version\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --system\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --greet World\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --greet World --upper\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --system --json\n", os.Args[0])
	}

	flag.Parse()

	if *showVersion {
		if *jsonOutput {
			output := map[string]string{
				"version": version,
				"name":    "oss-release-demo-cli",
			}
			jsonBytes, _ := json.MarshalIndent(output, "", "  ")
			fmt.Println(string(jsonBytes))
		} else {
			fmt.Printf("oss-release-demo-cli version %s\n", version)
			fmt.Printf("Built with %s\n", runtime.Version())
		}
		return
	}

	if *showSystem {
		info := SystemInfo{
			OS:           runtime.GOOS,
			Architecture: runtime.GOARCH,
			GoVersion:    runtime.Version(),
			NumCPU:       runtime.NumCPU(),
			Timestamp:    time.Now(),
		}

		if *jsonOutput {
			jsonBytes, _ := json.MarshalIndent(info, "", "  ")
			fmt.Println(string(jsonBytes))
		} else {
			fmt.Println("System Information:")
			fmt.Printf("  OS: %s\n", info.OS)
			fmt.Printf("  Architecture: %s\n", info.Architecture)
			fmt.Printf("  Go Version: %s\n", info.GoVersion)
			fmt.Printf("  CPU Cores: %d\n", info.NumCPU)
			fmt.Printf("  Timestamp: %s\n", info.Timestamp.Format(time.RFC3339))
		}
		return
	}

	if *greet != "" {
		greeting := fmt.Sprintf("Hello, %s!", *greet)
		if *uppercase {
			greeting = strings.ToUpper(greeting)
		}
		
		if *jsonOutput {
			output := map[string]string{
				"greeting": greeting,
			}
			jsonBytes, _ := json.MarshalIndent(output, "", "  ")
			fmt.Println(string(jsonBytes))
		} else {
			fmt.Println(greeting)
		}
		return
	}

	// If no flags provided, show usage
	flag.Usage()
}