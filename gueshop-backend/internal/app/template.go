package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
	"go.uber.org/fx"
)

// ProvideTempalteRenderer lo usas en fx.Provide()
var ProvideTemplateRenderer = fx.Provide(NewTemplateRenderer)

type TemplateRenderer struct {
	baseDir string
}

func NewTemplateRenderer() *TemplateRenderer {
	dir := os.Getenv("TEMPLATE_PATH")
	if dir == "" {
		dir = "templates"
	}

	return &TemplateRenderer{
		baseDir: dir,
	}
}

func (t *TemplateRenderer) Render(templateName string, ctx pongo2.Context) (string, error) {
	fullPath := filepath.Join(t.baseDir, templateName)

	tpl, err := pongo2.FromFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("error al cargar template %s: %w", templateName, err)
	}

	out, err := tpl.Execute(ctx)
	if err != nil {
		return "", fmt.Errorf("error al renderizar template %s: %w", templateName, err)
	}

	return out, nil
}
