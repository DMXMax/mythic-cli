package game

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "text/template"
    "time"

    "github.com/DMXMax/mge/chart"
    "github.com/DMXMax/mythic-cli/util/db"
    gdb "github.com/DMXMax/mythic-cli/util/game"
    "github.com/DMXMax/mythic-cli/util/input"
    "github.com/spf13/cobra"
)

const defaultTemplatePath = "data/templates/game.md.tmpl"

var (
    exportTemplatePath string
    exportOutPath      string
    exportForce        bool
)

// exportCmd renders a game to a Markdown file using a template file
var exportCmd = &cobra.Command{
    Use:   "export [name]",
    Short: "export a game to Markdown",
    Long:  "Export a game to a Markdown file using a Go text/template file. If no name is provided, the current game is exported.",
    RunE: func(cmd *cobra.Command, args []string) error {
        // Determine target game name
        var name string
        if len(args) > 0 {
            name = strings.TrimSpace(args[0])
        } else if gdb.Current != nil {
            name = gdb.Current.Name
        }
        if name == "" {
            return fmt.Errorf("no game name specified and no current game selected")
        }

        // Load game with logs
        var game gdb.Game
        if err := db.GamesDB.Preload("Log").Where("name = ?", name).First(&game).Error; err != nil {
            return fmt.Errorf("failed to load game '%s': %w", name, err)
        }

        // Resolve output path
        outPath := exportOutPath
        if strings.TrimSpace(outPath) == "" {
            safe := sanitizeFilename(game.Name)
            outPath = safe + ".md"
        }
        if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil && filepath.Dir(outPath) != "." {
            return fmt.Errorf("failed to create output directory: %w", err)
        }

        // Parse template
        tplPath := exportTemplatePath
        if strings.TrimSpace(tplPath) == "" {
            tplPath = defaultTemplatePath
        }
        funcMap := template.FuncMap{
            "formatTime": func(t time.Time, layout string) string { return t.Format(layout) },
            "oddsName":   func(v int8) string { if v < 0 || int(v) >= len(chart.OddsStrList) { return fmt.Sprintf("%d", v) }; return chart.OddsStrList[v] },
        }
        tpl, err := template.New(filepath.Base(tplPath)).Funcs(funcMap).ParseFiles(tplPath)
        if err != nil {
            return fmt.Errorf("failed to parse template '%s': %w", tplPath, err)
        }

        // If the output file exists, prompt to overwrite
        if info, err := os.Stat(outPath); err == nil && !info.IsDir() && !exportForce {
            ans, err := input.Ask(fmt.Sprintf("File '%s' already exists. Overwrite? [y/N]: ", outPath))
            if err != nil {
                return fmt.Errorf("failed to read confirmation: %w", err)
            }
            a := strings.TrimSpace(strings.ToLower(ans))
            if a != "y" && a != "yes" {
                return fmt.Errorf("export canceled; file exists: %s", outPath)
            }
        }

        // Create output file
        f, err := os.Create(outPath)
        if err != nil {
            return fmt.Errorf("failed to create output file '%s': %w", outPath, err)
        }
        defer f.Close()

        // Execute template with game as root
        if err := tpl.Execute(f, game); err != nil {
            return fmt.Errorf("failed to render template: %w", err)
        }

        cmd.Printf("Exported game '%s' to %s\n", game.Name, outPath)
        return nil
    },
}

func init() {
    exportCmd.Flags().StringVarP(&exportTemplatePath, "template", "t", defaultTemplatePath, "path to the Markdown template file")
    exportCmd.Flags().StringVarP(&exportOutPath, "out", "o", "", "output Markdown file path (default: <game>.md)")
    exportCmd.Flags().BoolVarP(&exportForce, "force", "f", false, "overwrite output file without prompting")
}

func sanitizeFilename(s string) string {
    s = strings.TrimSpace(s)
    // Replace path separators and problematic characters
    repl := []string{"/", "-", "\\", "-", ":", "-", "*", "-", "?", "-", "\"", "-", "<", "-", ">", "-", "|", "-"}
    r := strings.NewReplacer(repl...)
    s = r.Replace(s)
    // Collapse whitespace
    s = strings.Join(strings.Fields(s), " ")
    return s
}
