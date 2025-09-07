package mygrep

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"L2.12/conparams"
)

// createTestFile создает временный тестовый файл с указанным содержимым
func createTestFile(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	assetsDir := filepath.Join(tmpDir, "assets")
	err := os.MkdirAll(assetsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create assets directory: %v", err)
	}

	filePath := filepath.Join(assetsDir, "file.txt")
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Меняем рабочую директорию на временную
	oldWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	t.Cleanup(func() { os.Chdir(oldWd) })

	return filePath
}

// Тестовые данные
const testContent = `Hello world
Go is awesome
GREPed utility test
Another line with Go
No match here
Test line for grep
Go again!
UPPERCASE TEXT
lowercase text
Mixed CaSe TeXt`

func TestMygrep_buildFlags(t *testing.T) {
	tests := []struct {
		name     string
		iFlag    bool
		expected string
	}{
		{
			name:     "No flags",
			iFlag:    false,
			expected: "(?)",
		},
		{
			name:     "Case insensitive flag",
			iFlag:    true,
			expected: "(?i)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := &conparams.Params{I: tt.iFlag}
			g := mygrep{params: params}
			result := g.buildFlags()
			if result != tt.expected {
				t.Errorf("buildFlags() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMygrep_buildPattern(t *testing.T) {
	tests := []struct {
		name     string
		flags    string
		pattern  string
		fFlag    bool
		expected string
	}{
		{
			name:     "Basic pattern",
			flags:    "(?)",
			pattern:  "test",
			fFlag:    false,
			expected: "(?)test",
		},
		{
			name:     "Fixed string pattern",
			flags:    "(?)",
			pattern:  "test",
			fFlag:    true,
			expected: "(?)\\btest\\b",
		},
		{
			name:     "Case insensitive fixed pattern",
			flags:    "(?i)",
			pattern:  "Go",
			fFlag:    true,
			expected: "(?i)\\bGo\\b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := &conparams.Params{Pattern: tt.pattern, F: tt.fFlag}
			g := mygrep{params: params}
			result := g.buildPattern(tt.flags)
			if result != tt.expected {
				t.Errorf("buildPattern() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMygrep_checkStrings_BasicSearch(t *testing.T) {
	createTestFile(t, testContent)
	lines := strings.Split(testContent, "\n")

	tests := []struct {
		name     string
		pattern  string
		expected []string
	}{
		{
			name:    "Search for 'Go'",
			pattern: "Go",
			expected: []string{
				"Go is awesome",
				"Another line with Go",
				"Go again!",
			},
		},
		{
			name:    "Search for 'test'",
			pattern: "test",
			expected: []string{
				"GREPed utility test",
				"Test line for grep",
			},
		},
		{
			name:     "No matches",
			pattern:  "xyz123",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := &conparams.Params{Pattern: tt.pattern}
			g := &mygrep{params: params}
			result := g.checkStrings(lines)
			
			if len(result) != len(tt.expected) {
				t.Errorf("checkStrings() returned %d results, want %d", len(result), len(tt.expected))
			}
			
			for i, expected := range tt.expected {
				if i >= len(result) || result[i] != expected {
					t.Errorf("checkStrings()[%d] = %v, want %v", i, result[i], expected)
				}
			}
		})
	}
}

func TestMygrep_checkStrings_CaseInsensitive(t *testing.T) {
	createTestFile(t, testContent)
	lines := strings.Split(testContent, "\n")

	params := &conparams.Params{Pattern: "go", I: true}
	g := &mygrep{params: params}
	result := g.checkStrings(lines)

	expected := []string{
		"Go is awesome",
		"Another line with Go",
		"Go again!",
	}

	if len(result) != len(expected) {
		t.Errorf("Case insensitive search returned %d results, want %d", len(result), len(expected))
	}
}

func TestMygrep_checkStrings_InvertMatch(t *testing.T) {
	createTestFile(t, testContent)
	lines := strings.Split(testContent, "\n")

	params := &conparams.Params{Pattern: "Go", V: true}
	g := &mygrep{params: params}
	result := g.checkStrings(lines)

	// Должны получить все строки, которые НЕ содержат "Go"
	expected := []string{
		"Hello world",
		"GREPed utility test",
		"No match here",
		"Test line for grep",
		"UPPERCASE TEXT",
		"lowercase text",
		"Mixed CaSe TeXt",
	}

	if len(result) != len(expected) {
		t.Errorf("Inverted search returned %d results, want %d", len(result), len(expected))
	}
}

func TestMygrep_checkStrings_FixedString(t *testing.T) {
	createTestFile(t, testContent)
	lines := strings.Split(testContent, "\n")

	params := &conparams.Params{Pattern: "Go", F: true}
	g := &mygrep{params: params}
	result := g.checkStrings(lines)

	// С флагом F должны найти только точные совпадения слова "Go"
	expected := []string{
		"Go is awesome",
		"Another line with Go",
		"Go again!",
	}

	if len(result) != len(expected) {
		t.Errorf("Fixed string search returned %d results, want %d", len(result), len(expected))
	}
}

func TestMygrep_checkStrings_ContextAfter(t *testing.T) {
	createTestFile(t, testContent)
	lines := strings.Split(testContent, "\n")

	params := &conparams.Params{Pattern: "awesome", A: 2}
	g := &mygrep{params: params}
	result := g.checkStrings(lines)

	// Должны получить найденную строку + 2 строки после
	expected := []string{
		"Go is awesome",
		"GREPed utility test",
		"Another line with Go",
	}

	if len(result) != len(expected) {
		t.Errorf("Context after search returned %d results, want %d", len(result), len(expected))
	}
}

func TestMygrep_checkStrings_ContextBefore(t *testing.T) {
	createTestFile(t, testContent)
	lines := strings.Split(testContent, "\n")

	params := &conparams.Params{Pattern: "awesome", B: 1}
	g := &mygrep{params: params}
	result := g.checkStrings(lines)

	// Должны получить 1 строку до + найденную строку
	expected := []string{
		"Hello world",
		"Go is awesome",
	}

	if len(result) != len(expected) {
		t.Errorf("Context before search returned %d results, want %d", len(result), len(expected))
	}
}

func TestMygrep_checkStrings_ContextAround(t *testing.T) {
	createTestFile(t, testContent)
	lines := strings.Split(testContent, "\n")

	params := &conparams.Params{Pattern: "awesome", C: 2}
	g := &mygrep{params: params}
	result := g.checkStrings(lines)

	// C=2 означает 1 строку до и 1 строку после ((2+1)/2 = 1)
	expected := []string{
		"Hello world",
		"Go is awesome",
		"GREPed utility test",
	}

	if len(result) != len(expected) {
		t.Errorf("Context around search returned %d results, want %d", len(result), len(expected))
	}
}

func TestMygrep_checkStrings_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		content string
		pattern string
		params  conparams.Params
		wantLen int
	}{
		{
			name:    "Empty file",
			content: "",
			pattern: "test",
			params:  conparams.Params{Pattern: "test"},
			wantLen: 0,
		},
		{
			name:    "Single line file",
			content: "single line",
			pattern: "line",
			params:  conparams.Params{Pattern: "line"},
			wantLen: 1,
		},
		{
			name:    "Context beyond boundaries",
			content: "line1\nline2\nline3",
			pattern: "line2",
			params:  conparams.Params{Pattern: "line2", A: 10, B: 10},
			wantLen: 3, // Все строки
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTestFile(t, tt.content)
			lines := strings.Split(tt.content, "\n")
			if tt.content == "" {
				lines = []string{}
			}

			g := &mygrep{params: &tt.params}
			result := g.checkStrings(lines)

			if len(result) != tt.wantLen {
				t.Errorf("checkStrings() returned %d results, want %d", len(result), tt.wantLen)
			}
		})
	}
}

func TestMygrep_RegexPatterns(t *testing.T) {
	testContentRegex := `test123
abc456
789def
special@chars
line.with.dots
line with spaces`
	
	createTestFile(t, testContentRegex)
	lines := strings.Split(testContentRegex, "\n")

	tests := []struct {
		name     string
		pattern  string
		expected int
	}{
		{
			name:     "Digits pattern",
			pattern:  "\\d+",
			expected: 3, // test123, abc456, 789def
		},
		{
			name:     "Word pattern",
			pattern:  "\\w+",
			expected: 6, // Все строки содержат word characters
		},
		{
			name:     "Dot pattern",
			pattern:  "\\.",
			expected: 1, // line.with.dots
		},
		{
			name:     "Special chars",
			pattern:  "@",
			expected: 1, // special@chars
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := &conparams.Params{Pattern: tt.pattern}
			g := &mygrep{params: params}
			result := g.checkStrings(lines)

			if len(result) != tt.expected {
				t.Errorf("Regex pattern '%s' returned %d results, want %d", tt.pattern, len(result), tt.expected)
			}
		})
	}
}

func TestMygrep_CombinedFlags(t *testing.T) {
	createTestFile(t, testContent)
	lines := strings.Split(testContent, "\n")

	tests := []struct {
		name   string
		params conparams.Params
		minLen int
		maxLen int
	}{
		{
			name: "Case insensitive + invert",
			params: conparams.Params{
				Pattern: "go",
				I:       true,
				V:       true,
			},
			minLen: 5,
			maxLen: 10,
		},
		{
			name: "Fixed string + case insensitive",
			params: conparams.Params{
				Pattern: "test",
				F:       true,
				I:       true,
			},
			minLen: 1,
			maxLen: 3,
		},
		{
			name: "Context + invert",
			params: conparams.Params{
				Pattern: "xyz",
				V:       true,
				A:       1,
			},
			minLen: 10, // Все строки + контекст
			maxLen: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &mygrep{params: &tt.params}
			result := g.checkStrings(lines)

			if len(result) < tt.minLen || len(result) > tt.maxLen {
				t.Errorf("Combined flags test '%s' returned %d results, want between %d and %d", 
					tt.name, len(result), tt.minLen, tt.maxLen)
			}
		})
	}
}

// Бенчмарк тесты
func BenchmarkMygrep_checkStrings_SimplePattern(b *testing.B) {
	lines := strings.Split(testContent, "\n")
	params := &conparams.Params{Pattern: "Go"}
	g := &mygrep{params: params}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.result = nil // Очищаем результаты
		g.checkStrings(lines)
	}
}

func BenchmarkMygrep_checkStrings_RegexPattern(b *testing.B) {
	lines := strings.Split(testContent, "\n")
	params := &conparams.Params{Pattern: "\\w+"}
	g := &mygrep{params: params}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.result = nil
		g.checkStrings(lines)
	}
}

func BenchmarkMygrep_checkStrings_WithContext(b *testing.B) {
	lines := strings.Split(testContent, "\n")
	params := &conparams.Params{Pattern: "Go", A: 2, B: 2}
	g := &mygrep{params: params}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.result = nil
		g.checkStrings(lines)
	}
}