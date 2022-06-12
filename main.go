package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

const BUILD_CACHE_DIR = ".build_cache"
const BUILD_OUTPUT_DIR = ".build_output"
const GOBIN = ".gobin"
const CONFIG_FILE = "config.json"
const BENCHMARK_DIR = "benchmarks"
const RESULTS_FILE = "results.json"
const RESULTS_DIR = "results"

type Config struct {
	Name      string   `json:"name"`
	Itercount int      `json:"itercount"`
	Versions  []string `json:"versions"`
}

type Benchmark struct {
	Name   string `json:"name"`
	Minver string `json:"minver"`
}

func LoadConfig(cfg *Config) error {
	data, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, cfg)
}

func BuildGoProgram(benchmarkName string, gobin string) string {
	// Build the benchmark.
	cmd := exec.Command(filepath.Join(GOBIN, gobin), "build", "-o", BUILD_OUTPUT_DIR+"/"+benchmarkName+".exe", BENCHMARK_DIR+"/"+benchmarkName+"/"+"main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return BUILD_OUTPUT_DIR + "/" + benchmarkName + ".exe"
}

func Cleanup() {
	os.RemoveAll(BUILD_CACHE_DIR)
	os.RemoveAll(BUILD_OUTPUT_DIR)
	os.MkdirAll(BUILD_CACHE_DIR, 0755)
	os.MkdirAll(BUILD_OUTPUT_DIR, 0755)
}

func DownloadGo(version string) {
	if version == "master" {
		return
	}

	if _, err := os.Stat(filepath.Join(GOBIN, "go"+version)); err == nil {
		return
	}

	if _, err := os.Stat(filepath.Join(GOBIN, "go"+version+".exe")); err == nil {
		return
	}

	cmd := exec.Command("go", "install", "golang.org/dl/go"+version+"@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command(filepath.Join(GOBIN, "go"+version), "download")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func HashFileName(name string) string {
	h := sha256.Sum256([]byte(name))
	return hex.EncodeToString(h[:])
}

type BenchmarkResult struct {
	Name       string             `json:"name"`
	BuildTimes map[string]float64 `json:"build_times"`
	RunTimes   map[string]float64 `json:"run_times"`
}

func main() {
	CacheDir, err := filepath.Abs(BUILD_CACHE_DIR)
	if err != nil {
		panic(err)
	}

	// Set GOCACHE to the cache directory.
	os.Setenv("GOCACHE", CacheDir)

	GOBINPath, err := filepath.Abs(GOBIN)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(GOBINPath); err != nil {
		os.MkdirAll(GOBINPath, 0755)
	}
	// Set GOBIN
	os.Setenv("GOBIN", GOBINPath)

	Cleanup()

	// Load the config file.
	config := Config{}
	err = LoadConfig(&config)
	if err != nil {
		panic(err)
	}

	for _, version := range config.Versions {
		DownloadGo(version)
	}

	if stat, err := os.Stat(BENCHMARK_DIR); os.IsNotExist(err) {
		panic("Benchmark directory does not exist.")
	} else if !stat.IsDir() {
		panic("Benchmark directory is not a directory.")
	}
	lists, err := os.ReadDir(BENCHMARK_DIR)
	if err != nil {
		panic(err)
	}

	results := map[string]BenchmarkResult{}

	for _, list := range lists {
		if !list.IsDir() {
			continue
		}
		listName := list.Name()
		configPath := filepath.Join(BENCHMARK_DIR, listName, "meta.json")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			continue
		}
		benchmark := Benchmark{}
		data, err := os.ReadFile(configPath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &benchmark)
		if err != nil {
			panic(err)
		}
		benchmarkName := benchmark.Name
		if benchmarkName == "" {
			benchmarkName = listName
		}
		minver := benchmark.Minver
		if minver == "" {
			minver = "0.0"
		}
		for _, version := range config.Versions {
			if version < minver && version != "master" {
				continue
			}
			var buildTimes []time.Duration
			var program string
			// Build the program.
			for i := 0; i < config.Itercount; i++ {
				t := time.Now()
				Cleanup()
				program = BuildGoProgram(listName, "go"+version)
				buildTime := time.Since(t)
				buildTimes = append(buildTimes, buildTime)
			}

			// Run the program.
			var runTimes []time.Duration
			for i := 0; i < config.Itercount; i++ {
				t := time.Now()
				cmd := exec.Command(program)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					panic(err)
				}
				runTime := time.Since(t)
				runTimes = append(runTimes, runTime)
			}

			// Print the results.
			buildAvg := time.Duration(0)
			for _, buildTime := range buildTimes {
				buildAvg += buildTime
			}

			runAvg := time.Duration(0)
			for _, runTime := range runTimes {
				runAvg += runTime
			}

			buildAvg /= time.Duration(len(buildTimes))
			runAvg /= time.Duration(len(runTimes))
			fmt.Printf("%s %s Build: %s Run: %s\n", benchmarkName, version, buildAvg, runAvg)
			if _, ok := results[benchmarkName]; !ok {
				results[benchmarkName] = BenchmarkResult{
					Name:       benchmarkName,
					BuildTimes: map[string]float64{},
					RunTimes:   map[string]float64{},
				}
			}
			results[benchmarkName].BuildTimes[version] = float64(buildAvg.Nanoseconds())
			results[benchmarkName].RunTimes[version] = float64(runAvg.Nanoseconds())
		}
	}

	// Write the results to a file.
	data, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(RESULTS_FILE, data, 0664)
	if err != nil {
		panic(err)
	}

	Cleanup()

	// Cleanup the results directory.
	if _, err := os.Stat(RESULTS_DIR); err == nil {
		err = os.RemoveAll(RESULTS_DIR)
		if err != nil {
			panic(err)
		}
	}
	err = os.MkdirAll(RESULTS_DIR, 0755)
	if err != nil {
		panic(err)
	}

	// Make a plot of the results.
	for benchname, result := range results {
		c := chart.BarChart{}
		c.Title = benchname
		c.YAxis.Name = "Time (ns)"

		c.Bars = make([]chart.Value, len(result.BuildTimes))
		var sortedKeys []string = make([]string, 0, len(result.BuildTimes))
		for k := range result.BuildTimes {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Strings(sortedKeys)

		for i, k := range sortedKeys {
			c.Bars[i].Value = result.RunTimes[k] / 1000000
			c.Bars[i].Label = k
		}

		f, err := os.Create(filepath.Join(RESULTS_DIR, HashFileName(benchname)+".png"))
		if err != nil {
			log.Println(err)
			continue
		}
		defer f.Close()
		err = c.Render(chart.PNG, f)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	// Make Results.md
	f, err := os.Create(filepath.Join(RESULTS_DIR, "results.md"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "# Benchmarks\n\n")
	for _, result := range results {
		fmt.Fprintf(f, "## %s\n\n", result.Name)
		fmt.Fprintf(f, "| Version | Build Time (ms) | Run Time (ms) |\n")
		fmt.Fprintf(f, "| ------ | ------ | ------ |\n")
		var sortedKeys []string = make([]string, 0, len(result.BuildTimes))
		for k := range result.BuildTimes {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Strings(sortedKeys)
		for _, version := range sortedKeys {
			fmt.Fprintf(f, "| %s | %f | %f |\n", version, result.BuildTimes[version]/(1000000), result.RunTimes[version]/(1000000))
		}
		fmt.Fprintf(f, "\n")

		// Include the plot.
		fmt.Fprintf(f, "![%s](./%s.png)\n\n", result.Name, HashFileName(result.Name))
	}
}
