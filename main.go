package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/shirou/gopsutil/v3/cpu"

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
	if version == "gotip" {
		// Download gotip
		if _, err := os.Stat(filepath.Join(GOBIN, "gotip")); err != nil {
			// go install golang.org/dl/gotip@latest
			cmd := exec.Command("go", "install", "golang.org/dl/gotip@latest")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				panic(err)
			}
		}

		// Bootstrap gotip
		cmd := exec.Command(filepath.Join(GOBIN, "gotip"), "download")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
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

func SetupBenchstat() {
	if _, err := os.Stat(filepath.Join(GOBIN, "benchstat")); err == nil {
		return
	}

	cmd := exec.Command("go", "install", "golang.org/x/perf/cmd/benchstat@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func HashFileName(name string) string {
	h := sha256.Sum256([]byte(name))
	return name + "__" + hex.EncodeToString(h[:5])
}

type BenchmarkResult struct {
	Name       string               `json:"name"`
	BuildTimes map[string][]float64 `json:"build_times"`
	RunTimes   map[string][]float64 `json:"run_times"`
}

func main() {
	cpuinfo, err := cpu.Info()
	if err != nil {
		panic(err)
	}
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
	SetupBenchstat()

	if stat, err := os.Stat(BENCHMARK_DIR); os.IsNotExist(err) {
		panic("Benchmark directory does not exist.")
	} else if !stat.IsDir() {
		panic("Benchmark directory is not a directory.")
	}
	lists, err := os.ReadDir(BENCHMARK_DIR)
	if err != nil {
		panic(err)
	}

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

	results := map[string]BenchmarkResult{}

	for _, list := range lists {
		if !list.IsDir() {
			continue
		}

		if len(os.Args) > 1 {
			requestedBenchmark := os.Args[1:]
			found := false
			for _, requested := range requestedBenchmark {
				if list.Name() == requested {
					found = true
					break
				}
			}
			if !found {
				continue
			}
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
		var resultFiles []string
		for _, version := range config.Versions {
			if version < minver && version != "gotip" {
				continue
			}
			var buildTimes []time.Duration
			var program string
			// Build the program.
			for i := 0; i < config.Itercount; i++ {
				t := time.Now()
				Cleanup()
				if version == "gotip" {
					cmd := exec.Command(filepath.Join(GOBIN, "gotip"), "build", "std")
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					if err != nil {
						panic(err)
					}
					t = time.Now()
					program = BuildGoProgram(listName, "gotip")
				} else {
					cmd := exec.Command(filepath.Join(GOBIN, "go"+version), "build", "std")
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					if err != nil {
						panic(err)
					}
					t = time.Now()
					program = BuildGoProgram(listName, "go"+version)
				}
				buildTime := time.Since(t)
				buildTimes = append(buildTimes, buildTime)
			}

			fname := filepath.Join(RESULTS_DIR, HashFileName(benchmarkName)+"_"+version+".txt")
			resultFiles = append(resultFiles, HashFileName(benchmarkName)+"_"+version)
			resultsFile, err := os.Create(fname)
			if err != nil {
				panic(err)
			}
			rw := bufio.NewWriter(resultsFile)

			// Run the program.
			var runTimes []time.Duration
			for i := 0; i < config.Itercount; i++ {
				fmt.Fprintf(rw, "Benchmark%s\t", benchmarkName)
				t := time.Now()
				cmd := exec.Command(program)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					panic(err)
				}
				runTime := time.Since(t)
				sysT := cmd.ProcessState.SystemTime()
				userT := cmd.ProcessState.UserTime()
				runTimes = append(runTimes, runTime)
				fmt.Fprintf(rw, "1\t%d ns/op", runTime.Nanoseconds())
				fmt.Fprintf(rw, "\t%d user-ns/op", userT.Nanoseconds())
				fmt.Fprintf(rw, "\t%d sys-ns/op\n", sysT.Nanoseconds())
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
					BuildTimes: map[string][]float64{},
					RunTimes:   map[string][]float64{},
				}
			}
			var buildTimesFloat []float64
			var runTimesFloat []float64
			for _, buildTime := range buildTimes {
				buildTimesFloat = append(buildTimesFloat, float64(buildTime.Nanoseconds()))
			}
			for _, runTime := range runTimes {
				runTimesFloat = append(runTimesFloat, float64(runTime.Nanoseconds()))
			}
			results[benchmarkName].BuildTimes[version] = buildTimesFloat
			results[benchmarkName].RunTimes[version] = runTimesFloat
			rw.Flush()
			resultsFile.Close()
		}

		// Run Benchstat.
		path := filepath.Join(RESULTS_DIR, "benchstat"+"_"+benchmarkName)
		err = os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
		var b bytes.Buffer
		for i := 0; i < len(resultFiles); i++ {
			for j := i + 1; j < len(resultFiles); j++ {
				b.Reset()
				cmd := exec.Command(filepath.Join(GOBIN, "benchstat"),
					filepath.Join(RESULTS_DIR, resultFiles[i]+".txt"),
					filepath.Join(RESULTS_DIR, resultFiles[j]+".txt"),
				)
				cmd.Stdout = &b
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					panic(err)
				}

				// Write the results to a file.
				f, err := os.Create(filepath.Join(path, fmt.Sprintf("%s_%s_%s.txt", resultFiles[i], resultFiles[j], benchmarkName)))
				if err != nil {
					panic(err)
				}
				_, err = f.Write(b.Bytes())
				if err != nil {
					f.Close()
					panic(err)
				}
				f.Close()
			}
		}
	}

	// ns => ms
	for _, result := range results {
		for _, buildTimes := range result.BuildTimes {
			for i, buildTime := range buildTimes {
				buildTimes[i] = buildTime / 1000000.0
			}
		}
		for _, runTimes := range result.RunTimes {
			for i, runTime := range runTimes {
				runTimes[i] = runTime / 1000000.0
			}
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

	// Make a plot of the results.
	var benchnameStrings []string = make([]string, 0, len(results))
	for key := range results {
		benchnameStrings = append(benchnameStrings, key)
	}
	sort.Strings(benchnameStrings)

	for _, benchname := range benchnameStrings {
		result := results[benchname]
		c := chart.BarChart{}
		c.Title = benchname
		c.YAxis.Name = "Time (ns)"

		c.Bars = make([]chart.Value, len(result.BuildTimes))
		var sortedKeys []string = make([]string, 0, len(result.BuildTimes))
		for k := range result.BuildTimes {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Strings(sortedKeys)

		if len(sortedKeys) <= 0 {
			continue
		}

		var maxs []float64
		var mins []float64

		for i, k := range sortedKeys {
			data := stats.LoadRawData(result.RunTimes[k])
			median, err := data.Median()
			if err != nil {
				panic(err)
			}
			c.Bars[i].Value = median
			c.Bars[i].Label = k
			max, err := data.Max()
			if err != nil {
				panic(err)
			}
			min, err := data.Min()
			if err != nil {
				panic(err)
			}
			maxs = append(maxs, max)
			mins = append(mins, min)
		}

		maxsData := stats.LoadRawData(maxs)
		minsData := stats.LoadRawData(mins)
		max, err := maxsData.Max()
		if err != nil {
			panic(err)
		}
		min, err := minsData.Min()
		if err != nil {
			panic(err)
		}

		_ = max

		c.YAxis.Range = &chart.ContinuousRange{Min: min / 2, Max: max * 1.3}

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

		f, err = os.Create(filepath.Join(RESULTS_DIR, HashFileName(benchname)+".svg"))
		if err != nil {
			log.Println(err)
			continue
		}
		defer f.Close()
		err = c.Render(chart.SVG, f)
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
	// CPU Info
	fmt.Fprintf(f, "## Environment\n\n")
	fmt.Fprintf(f, "NumCPU: %d\n\n", runtime.NumCPU())
	fmt.Fprintf(f, "Arch: %s\n\n", runtime.GOARCH)
	fmt.Fprintf(f, "OS: %s\n\n", runtime.GOOS)
	fmt.Fprintf(f, "Version: %s\n\n", runtime.Version())
	fmt.Fprintf(f, "Itercount: %d\n\n", config.Itercount)
	for i := range cpuinfo {
		fmt.Fprintf(f, "### CPU %d\n\n", i)
		fmt.Fprintf(f, "Model: %s\n\n", cpuinfo[i].ModelName)
		fmt.Fprintf(f, "Cores: %d\n\n", cpuinfo[i].Cores)
		fmt.Fprintf(f, "Mhz: %f\n\n", cpuinfo[i].Mhz)
		fmt.Fprintf(f, "CacheSize: %d\n\n", cpuinfo[i].CacheSize)
		fmt.Fprintf(f, "Microcode: %s\n\n", cpuinfo[i].Microcode)
	}
	for _, name := range benchnameStrings {
		result := results[name]
		fmt.Fprintf(f, "## %s\n\n", result.Name)
		fmt.Fprintf(f, "| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |\n")
		fmt.Fprintf(f, "| ------ | ------ | ------ | ------ | ------ |\n")
		var sortedKeys []string = make([]string, 0, len(result.BuildTimes))
		for k := range result.BuildTimes {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Strings(sortedKeys)
		for _, version := range sortedKeys {
			BuildTimeData := stats.LoadRawData(result.BuildTimes[version])
			BuildTimeMedian, err := BuildTimeData.Median()
			if err != nil {
				panic(err)
			}
			BuildTimeStdDev, err := BuildTimeData.StandardDeviation()
			if err != nil {
				panic(err)
			}
			RunTimeData := stats.LoadRawData(result.RunTimes[version])
			RunTimeMedian, err := RunTimeData.Median()
			if err != nil {
				panic(err)
			}
			RunTimeStdDev, err := RunTimeData.StandardDeviation()
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(f, "| %s | %f | %f | %f | %f |\n", version, BuildTimeMedian, BuildTimeStdDev, RunTimeMedian, RunTimeStdDev)
		}
		fmt.Fprintf(f, "\n")

		// Include the plot.
		fmt.Fprintf(f, "![%s](./%s.png)\n\n", result.Name, HashFileName(result.Name))
	}
}
