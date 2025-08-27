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
	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/wsdl"
	"github.com/way-platform/soap-go/xsd"
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
	g := codegen.NewFile(markdownFilename)

	// Generate markdown documentation
	if err := generateMarkdown(g, doc); err != nil {
		return err
	}

	// Write the markdown file
	content, err := g.Content()
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

// generateMarkdown generates markdown documentation for a WSDL file
func generateMarkdown(g *codegen.File, doc *wsdl.Definitions) error {
	// Title
	g.P("# ", doc.Name)
	if doc.TargetNamespace != "" {
		g.P()
		g.P("**Namespace:** `", doc.TargetNamespace, "`")
	}
	g.P()

	// Build schema map for element lookups
	schemaMap := buildSchemaMap(doc)

	// Generate documentation for each service
	for _, service := range doc.Service {
		if err := generateServiceDoc(g, &service, doc, schemaMap); err != nil {
			return err
		}
	}

	return nil
}

// buildSchemaMap creates a map of element names to their definitions for easy lookup
func buildSchemaMap(doc *wsdl.Definitions) map[string]*xsd.Element {
	schemaMap := make(map[string]*xsd.Element)

	if doc.Types != nil {
		for _, schema := range doc.Types.Schemas {
			for i := range schema.Elements {
				element := &schema.Elements[i]
				schemaMap[element.Name] = element
			}
		}
	}

	return schemaMap
}

// normalizeDocumentation normalizes whitespace in documentation strings
// by trimming leading/trailing whitespace and replacing any sequence of
// whitespace characters (including newlines) with a single space
func normalizeDocumentation(doc string) string {
	return strings.TrimSpace(strings.Join(strings.Fields(doc), " "))
}

// generateServiceDoc generates documentation for a single service
func generateServiceDoc(g *codegen.File, service *wsdl.Service, doc *wsdl.Definitions, schemaMap map[string]*xsd.Element) error {
	g.P("## ", service.Name)
	g.P()

	// Add service description if available
	if service.Documentation != "" {
		g.P(normalizeDocumentation(service.Documentation))
		g.P()
	}

	// Find the corresponding PortType for this service
	var portType *wsdl.PortType
	for _, binding := range doc.Binding {
		for _, port := range service.Ports {
			if strings.Contains(port.Binding, binding.Name) {
				// Find the PortType referenced by this binding
				typeName := strings.TrimPrefix(binding.Type, "tns:")
				for i := range doc.PortType {
					if doc.PortType[i].Name == typeName {
						portType = &doc.PortType[i]
						break
					}
				}
				break
			}
		}
		if portType != nil {
			break
		}
	}

	if portType == nil {
		g.P("*No operations found for this service.*")
		g.P()
		return nil
	}

	// Generate documentation for each operation
	for _, operation := range portType.Operations {
		if err := generateOperationDoc(g, &operation, doc, schemaMap); err != nil {
			return err
		}
	}

	return nil
}

// generateOperationDoc generates documentation for a single operation
func generateOperationDoc(g *codegen.File, operation *wsdl.Operation, doc *wsdl.Definitions, schemaMap map[string]*xsd.Element) error {
	g.P("### ", operation.Name)
	g.P()

	// Add operation description if available
	if operation.Documentation != "" {
		g.P(normalizeDocumentation(operation.Documentation))
		g.P()
	}

	// Generate request documentation
	if operation.Input != nil {
		if err := generateMessageDoc(g, "Request", operation.Input.Message, doc, schemaMap); err != nil {
			return err
		}
	}

	// Generate response documentation
	if operation.Output != nil {
		if err := generateMessageDoc(g, "Response", operation.Output.Message, doc, schemaMap); err != nil {
			return err
		}
	}

	g.P()
	return nil
}

// generateMessageDoc generates documentation for a request or response message
func generateMessageDoc(g *codegen.File, messageType, messageName string, doc *wsdl.Definitions, schemaMap map[string]*xsd.Element) error {
	g.P("#### ", messageType)
	g.P()

	// Find the message definition
	var message *wsdl.Message
	cleanMessageName := strings.TrimPrefix(messageName, "tns:")
	for i := range doc.Messages {
		if doc.Messages[i].Name == cleanMessageName {
			message = &doc.Messages[i]
			break
		}
	}

	if message == nil {
		g.P("*Message definition not found.*")
		g.P()
		return nil
	}

	// Generate field documentation for each part
	for _, part := range message.Parts {
		if part.Element != "" {
			elementName := strings.TrimPrefix(part.Element, "tns:")
			element := schemaMap[elementName]
			if element != nil {
				generateElementFields(g, element, 0)
			} else {
				g.P("- ", part.Name, " (element: ", part.Element, ")")
			}
		} else if part.Type != "" {
			g.P("- ", part.Name, " (type: ", part.Type, ")")
		}
	}

	g.P()
	return nil
}

// generateElementFields generates hierarchical bullet list for element fields
func generateElementFields(g *codegen.File, element *xsd.Element, depth int) {
	indent := strings.Repeat("  ", depth)

	// Generate the field name and type
	fieldName := element.Name
	fieldType := element.Type
	if fieldType == "" && element.ComplexType != nil {
		fieldType = "complex"
	}

	// Add occurrence information
	occurrenceInfo := ""
	if element.MinOccurs != "" || element.MaxOccurs != "" {
		min := element.MinOccurs
		max := element.MaxOccurs
		if min == "" {
			min = "1"
		}
		if max == "" {
			max = "1"
		}
		if min != "1" || max != "1" {
			occurrenceInfo = fmt.Sprintf(" (%s..%s)", min, max)
		}
	}

	if fieldType != "" {
		g.P(indent, "- **", fieldName, "** (", fieldType, ")", occurrenceInfo)
	} else {
		g.P(indent, "- **", fieldName, "**", occurrenceInfo)
	}

	// If this element has a complex type, recursively generate its fields
	if element.ComplexType != nil {
		generateComplexTypeFields(g, element.ComplexType, depth+1)
	}
}

// generateComplexTypeFields generates fields for a complex type
func generateComplexTypeFields(g *codegen.File, complexType *xsd.ComplexType, depth int) {
	if complexType.Sequence != nil {
		generateSequenceFields(g, complexType.Sequence, depth)
	}
	if complexType.Choice != nil {
		generateChoiceFields(g, complexType.Choice, depth)
	}
	if complexType.All != nil {
		generateAllFields(g, complexType.All, depth)
	}

	// Generate attributes
	for _, attr := range complexType.Attributes {
		indent := strings.Repeat("  ", depth)
		required := ""
		if attr.Use == "required" {
			required = " (required)"
		}
		g.P(indent, "- **@", attr.Name, "** (", attr.Type, ")", required, " *[attribute]*")
	}
}

// generateSequenceFields generates fields for a sequence
func generateSequenceFields(g *codegen.File, sequence *xsd.Sequence, depth int) {
	for i := range sequence.Elements {
		generateElementFields(g, &sequence.Elements[i], depth)
	}
	for i := range sequence.Sequences {
		generateSequenceFields(g, &sequence.Sequences[i], depth)
	}
	for i := range sequence.Choices {
		generateChoiceFields(g, &sequence.Choices[i], depth)
	}
}

// generateChoiceFields generates fields for a choice
func generateChoiceFields(g *codegen.File, choice *xsd.Choice, depth int) {
	indent := strings.Repeat("  ", depth)
	g.P(indent, "- **Choice of:**")

	for i := range choice.Elements {
		generateElementFields(g, &choice.Elements[i], depth+1)
	}
	for i := range choice.Sequences {
		generateSequenceFields(g, &choice.Sequences[i], depth+1)
	}
	for i := range choice.Choices {
		generateChoiceFields(g, &choice.Choices[i], depth+1)
	}
}

// generateAllFields generates fields for an all group
func generateAllFields(g *codegen.File, all *xsd.All, depth int) {
	for i := range all.Elements {
		generateElementFields(g, &all.Elements[i], depth)
	}
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
