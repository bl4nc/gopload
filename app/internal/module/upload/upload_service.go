package uploadmodule

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func NormalizePath(path string) (string, error) {
	if path == "" {
		return "", nil
	}

	// Remove espaços, acentos e caracteres especiais
	path = removeAccents(path)
	// Garantir que o caminho comece com uma única barra
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = "/" + strings.TrimLeft(path, "/")
	// Remover múltiplas barras consecutivas
	path = strings.ReplaceAll(path, "//", "/")

	// Validar o caminho normalizado
	validPath, err := isValidPath(path)
	if !validPath || err != nil {
		return "", fmt.Errorf("invalid path: %v", err)
	}

	return path, nil
}

func removeAccents(s string) string {
	t := transformString(s, func(r rune) rune {
		if unicode.IsSpace(r) {
			return '_'
		}
		return r
	})
	return t
}

func transformString(s string, f func(rune) rune) string {
	t := norm.NFD.String(s)
	var b strings.Builder
	for _, r := range t {
		if unicode.IsMark(r) {
			continue
		}
		b.WriteRune(f(r))
	}
	return b.String()
}

func isValidPath(path string) (bool, error) {
	validPathPattern := regexp.MustCompile(`^[a-zA-Z0-9/_-]+$`)
	if !validPathPattern.MatchString(path) {
		slog.Error("path contains invalid characters")
		return false, fmt.Errorf("path contains invalid characters")
	}
	return true, nil
}
