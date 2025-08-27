package doc

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/way-platform/soap-go/internal/docgen"
	"github.com/way-platform/soap-go/wsdl"
)

// NewCommand creates a new [cobra.Command] for the doc command.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "doc",
		Short:   "Display documentation for a SOAP API",
		GroupID: "doc",
	}
	inputFile := cmd.Flags().StringP("input", "i", "", "input WSDL file (required)")
	_ = cmd.MarkFlagRequired("input")
	_ = cmd.MarkFlagFilename("input", "wsdl")
	outputFile := cmd.Flags().StringP("output", "o", "-", "output file (required)")
	pager := cmd.Flags().BoolP("pager", "p", true, "use interactive pager")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context(), *inputFile, *outputFile, *pager)
	}
	return cmd
}

func run(ctx context.Context, inputFile, outputFile string, usePager bool) error {
	doc, err := wsdl.ParseFromFile(inputFile)
	if err != nil {
		return err
	}
	markdownFilename := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile)) + ".md"

	// Generate markdown documentation using docgen
	generator := docgen.NewGenerator(markdownFilename, doc)
	if err := generator.Generate(); err != nil {
		return err
	}

	// Write the markdown file
	content, err := generator.File().Content()
	if err != nil {
		return err
	}
	if outputFile == "-" {
		if usePager {
			// Use interactive pager
			return showInPager(string(content))
		} else {
			// Render markdown with glamour for terminal output
			// Use "auto" to automatically detect light/dark theme
			out, err := glamour.Render(string(content), "auto")
			if err != nil {
				// Fallback to plain text if glamour fails
				fmt.Fprintf(os.Stdout, "%s", string(content))
			} else {
				fmt.Print(out)
			}
		}
	} else {
		return os.WriteFile(outputFile, content, 0o644)
	}
	return nil
}

// TUI Components for pager

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

type pagerModel struct {
	content  string
	ready    bool
	viewport viewport.Model
}

func (m pagerModel) Init() tea.Cmd {
	return nil
}

func (m pagerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m pagerModel) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m pagerModel) headerView() string {
	title := titleStyle.Render("SOAP API Documentation")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m pagerModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%% • q/esc to quit", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// showInPager displays the markdown content in an interactive pager
func showInPager(markdown string) error {
	// Render markdown with glamour for the pager
	content, err := glamour.Render(markdown, "auto")
	if err != nil {
		// Fallback to plain text if glamour fails
		content = markdown
	}

	p := tea.NewProgram(
		pagerModel{content: content},
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	_, err = p.Run()
	return err
}
