package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v3"
)

var (
	colorHeader  = color.New(color.FgCyan, color.Bold)
	colorSuccess = color.New(color.FgGreen)
	colorWarning = color.New(color.FgYellow)
	colorDanger  = color.New(color.FgRed)
	colorInfo    = color.New(color.FgCyan)
)

func printHeader(msg string) {
	colorHeader.Println("\n" + msg)
	fmt.Println(strings.Repeat("─", 50))
}

func noResult(msg string) {
	colorInfo.Println("  ✔ " + msg)
}

func FileDetectCommand() *cli.Command {
	return &cli.Command{
		Name:  "fdetect",
		Usage: "Detect and scan files with various filters",
		Flags: []cli.Flag{

			&cli.BoolFlag{Name: "emt", Usage: "Find empty files"},
			&cli.IntFlag{Name: "rec", Usage: "Find files modified in last N days"},
			&cli.Float64Flag{Name: "lrg", Usage: "Find files larger than N GB"},

			&cli.BoolFlag{Name: "edir", Usage: "Find empty directories"},
			&cli.BoolFlag{Name: "ext", Usage: "Show file extension summary"},
			&cli.BoolFlag{Name: "count", Usage: "Count total files and directories"},
			&cli.IntFlag{Name: "top", Usage: "Show top N largest files"},
			&cli.BoolFlag{Name: "dup", Usage: "Find duplicate files by name"},
			&cli.IntFlag{Name: "old", Usage: "Find files not modified in last N days"},
			&cli.BoolFlag{Name: "hidden", Usage: "Find hidden files (dot files)"},
			&cli.StringFlag{Name: "perm", Usage: "Find files with specific permission (e.g. 777)"},
			&cli.BoolFlag{Name: "new", Usage: "Find the newest file in directory"},
			&cli.BoolFlag{Name: "log", Usage: "Find all log files (.log)"},
			&cli.BoolFlag{Name: "tmp", Usage: "Find temp files (.tmp .cache .swp)"},
			&cli.IntFlag{Name: "depth", Value: -1, Usage: "Limit scan depth (-1 = unlimited)"},

			&cli.StringFlag{Name: "p", Value: ".", Usage: "Directory path to scan"},
		},
		Action: runDetect,
	}
}

func runDetect(ctx context.Context, c *cli.Command) error {

	start := time.Now()

	path := c.String("p")
	maxDepth := c.Int("depth")

	colorInfo.Println("📁 Scanning path:", path)

	if !c.Bool("emt") && !c.Bool("edir") && !c.Bool("ext") &&
		!c.Bool("count") && !c.Bool("dup") && !c.Bool("hidden") &&
		!c.Bool("new") && !c.Bool("log") && !c.Bool("tmp") &&
		c.Int("rec") == 0 && c.Float64("lrg") == 0 &&
		c.Int("top") == 0 && c.Int("old") == 0 &&
		c.String("perm") == "" {
		colorWarning.Println("⚠  No flag provided. Use --help to see all available flags.")
		return nil
	}

	withinDepth := func(p string) bool {
		if maxDepth == -1 {
			return true
		}
		rel, err := filepath.Rel(filepath.Clean(path), p)
		if err != nil {
			return false
		}
		return len(strings.Split(rel, string(os.PathSeparator))) <= maxDepth
	}

	// ── EMPTY FILES ──────────────────────────────────────────────

	if c.Bool("emt") {
		printHeader("🔍 Empty Files")
		foundAny := false
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			info, err := d.Info()
			if err != nil {
				return nil
			}
			if info.Size() == 0 {
				colorDanger.Printf("  ✖ %s\n", p)
				foundAny = true
			}
			return nil
		})
		if !foundAny {
			noResult("No empty files found.")
		}
	}

	// ── EMPTY DIRECTORIES ────────────────────────────────────────

	if c.Bool("edir") {
		printHeader("📂 Empty Directories")
		foundAny := false
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || !d.IsDir() || p == path || !withinDepth(p) {
				return nil
			}
			entries, err := os.ReadDir(p)
			if err != nil {
				return nil
			}
			if len(entries) == 0 {
				colorDanger.Printf("  ✖ %s\n", p)
				foundAny = true
			}
			return nil
		})
		if !foundAny {
			noResult("No empty directories found.")
		}
	}

	// ── RECENT FILES ─────────────────────────────────────────────

	if c.Int("rec") > 0 {
		days := c.Int("rec")
		cutoff := time.Now().AddDate(0, 0, -days)
		printHeader(fmt.Sprintf("🕐 Files Modified in Last %d Day(s)", days))
		foundAny := false
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			info, err := d.Info()
			if err != nil {
				return nil
			}
			if info.ModTime().After(cutoff) {
				colorSuccess.Printf("  ✔ %s\n", p)
				colorInfo.Printf("     Modified: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
				foundAny = true
			}
			return nil
		})
		if !foundAny {
			noResult(fmt.Sprintf("No files modified in last %d day(s).", days))
		}
	}

	// ── OLD FILES ────────────────────────────────────────────────

	if c.Int("old") > 0 {
		days := c.Int("old")
		cutoff := time.Now().AddDate(0, 0, -days)
		printHeader(fmt.Sprintf("🗓  Files Not Modified in Last %d Day(s)", days))
		foundAny := false
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			info, err := d.Info()
			if err != nil {
				return nil
			}
			if info.ModTime().Before(cutoff) {
				colorWarning.Printf("  ⚠  %s\n", p)
				colorInfo.Printf("     Last modified: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
				foundAny = true
			}
			return nil
		})
		if !foundAny {
			noResult(fmt.Sprintf("No files older than %d day(s).", days))
		}
	}

	// ── LARGE FILES ──────────────────────────────────────────────

	if c.Float64("lrg") > 0 {
		sizeGB := c.Float64("lrg")
		minBytes := int64(sizeGB * 1024 * 1024 * 1024)
		printHeader(fmt.Sprintf("💾 Files Larger Than %.2f GB", sizeGB))
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"FILE", "SIZE (GB)"})
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			info, err := d.Info()
			if err != nil {
				return nil
			}
			if info.Size() > minBytes {
				sizeGBActual := float64(info.Size()) / (1024 * 1024 * 1024)
				table.Append([]string{p, fmt.Sprintf("%.2f", sizeGBActual)})
			}
			return nil
		})
		if table.NumLines() == 0 {
			noResult(fmt.Sprintf("No files larger than %.2f GB.", sizeGB))
		} else {
			table.Render()
		}
	}

	// ── TOP N LARGEST FILES ──────────────────────────────────────

	if c.Int("top") > 0 {
		n := c.Int("top")
		printHeader(fmt.Sprintf("🏆 Top %d Largest Files", n))

		type fileEntry struct {
			path string
			size int64
		}
		var files []fileEntry

		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			info, err := d.Info()
			if err != nil {
				return nil
			}
			files = append(files, fileEntry{p, info.Size()})
			return nil
		})

		sort.Slice(files, func(i, j int) bool {
			return files[i].size > files[j].size
		})

		if len(files) == 0 {
			noResult("No files found.")
		} else {
			if n > len(files) {
				n = len(files)
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"RANK", "FILE", "SIZE (MB)"})
			for i := 0; i < n; i++ {
				sizeMB := float64(files[i].size) / (1024 * 1024)
				table.Append([]string{
					fmt.Sprintf("#%d", i+1),
					files[i].path,
					fmt.Sprintf("%.2f", sizeMB),
				})
			}
			table.Render()
		}
	}

	// ── COUNT FILES & DIRS ───────────────────────────────────────

	if c.Bool("count") {
		printHeader("🔢 File & Directory Count")
		fileCount := 0
		dirCount := 0
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || p == path || !withinDepth(p) {
				return nil
			}
			if d.IsDir() {
				dirCount++
			} else {
				fileCount++
			}
			return nil
		})
		colorSuccess.Printf("  Files       : %d\n", fileCount)
		colorInfo.Printf("  Directories : %d\n", dirCount)
		colorHeader.Printf("  Total       : %d\n", fileCount+dirCount)
	}

	// ── EXTENSION SUMMARY ────────────────────────────────────────

	if c.Bool("ext") {
		printHeader("📋 File Extension Summary")
		extMap := make(map[string]int)
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			extension := strings.ToLower(filepath.Ext(p))
			if extension == "" {
				extension = "(no extension)"
			}
			extMap[extension]++
			return nil
		})
		if len(extMap) == 0 {
			noResult("No files found.")
		} else {
			var keys []string
			for k := range extMap {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"EXTENSION", "COUNT"})
			for _, k := range keys {
				table.Append([]string{k, fmt.Sprintf("%d", extMap[k])})
			}
			table.Render()
		}
	}

	// ── DUPLICATE FILES BY NAME ──────────────────────────────────

	if c.Bool("dup") {
		printHeader("♊ Duplicate File Names")
		nameMap := make(map[string][]string)
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			nameMap[d.Name()] = append(nameMap[d.Name()], p)
			return nil
		})
		foundAny := false
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"FILE NAME", "PATHS"})
		for name, paths := range nameMap {
			if len(paths) > 1 {
				table.Append([]string{name, strings.Join(paths, "\n")})
				foundAny = true
			}
		}
		if !foundAny {
			noResult("No duplicate file names found.")
		} else {
			table.Render()
		}
	}

	// ── HIDDEN FILES ─────────────────────────────────────────────

	if c.Bool("hidden") {
		printHeader("👻 Hidden Files")
		foundAny := false
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || p == path || !withinDepth(p) {
				return nil
			}
			if strings.HasPrefix(d.Name(), ".") {
				colorWarning.Printf("  ⚠  %s\n", p)
				foundAny = true
			}
			return nil
		})
		if !foundAny {
			noResult("No hidden files found.")
		}
	}

	// ── PERMISSION FILES ─────────────────────────────────────────

	if c.String("perm") != "" {
		perm := c.String("perm")
		printHeader(fmt.Sprintf("🔐 Files With Permission %s", perm))
		foundAny := false
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			info, err := d.Info()
			if err != nil {
				return nil
			}
			if fmt.Sprintf("%o", info.Mode().Perm()) == perm {
				colorWarning.Printf("  ⚠  %s\n", p)
				foundAny = true
			}
			return nil
		})
		if !foundAny {
			noResult(fmt.Sprintf("No files with permission %s found.", perm))
		}
	}

	// ── NEWEST FILE ──────────────────────────────────────────────

	if c.Bool("new") {
		printHeader("🆕 Newest File")
		var newestPath string
		var newestTime time.Time
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			info, err := d.Info()
			if err != nil {
				return nil
			}
			if info.ModTime().After(newestTime) {
				newestTime = info.ModTime()
				newestPath = p
			}
			return nil
		})
		if newestPath == "" {
			noResult("No files found.")
		} else {
			colorSuccess.Printf("  ✔ %s\n", newestPath)
			colorInfo.Printf("     Modified: %s\n", newestTime.Format("2006-01-02 15:04:05"))
		}
	}

	// ── LOG FILES ────────────────────────────────────────────────

	if c.Bool("log") {
		printHeader("📄 Log Files")
		foundAny := false
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			if strings.HasSuffix(strings.ToLower(d.Name()), ".log") {
				colorWarning.Printf("  ⚠  %s\n", p)
				foundAny = true
			}
			return nil
		})
		if !foundAny {
			noResult("No log files found.")
		}
	}

	// ── TEMP FILES ───────────────────────────────────────────────

	if c.Bool("tmp") {
		printHeader("🗑  Temp Files")
		tmpExts := []string{".tmp", ".cache", ".swp", ".bak", ".temp"}
		foundAny := false
		filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !withinDepth(p) {
				return nil
			}
			ext := strings.ToLower(filepath.Ext(d.Name()))
			for _, t := range tmpExts {
				if ext == t {
					colorWarning.Printf("  ⚠  %s\n", p)
					foundAny = true
					break
				}
			}
			return nil
		})
		if !foundAny {
			noResult("No temp files found.")
		}
	}

	fmt.Println()
	colorSuccess.Println("⏱ Scan completed in:", time.Since(start))

	return nil
}
