package conparams

import (
	"flag"
	"os"
	"testing"
)

// resetFlags сбрасывает флаги для тестирования
func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestNewParams_DefaultValues(t *testing.T) {
	resetFlags()
	
	// Симулируем запуск с минимальными параметрами
	os.Args = []string{"program", "test_pattern"}
	
	params := NewParams()
	
	if params.Pattern != "test_pattern" {
		t.Errorf("Pattern = %v, want %v", params.Pattern, "test_pattern")
	}
	
	if params.A != 0 {
		t.Errorf("A = %v, want %v", params.A, 0)
	}
	
	if params.B != 0 {
		t.Errorf("B = %v, want %v", params.B, 0)
	}
	
	if params.C != 0 {
		t.Errorf("C = %v, want %v", params.C, 0)
	}
	
	if params.Count != false {
		t.Errorf("Count = %v, want %v", params.Count, false)
	}
	
	if params.I != false {
		t.Errorf("I = %v, want %v", params.I, false)
	}
	
	if params.V != false {
		t.Errorf("V = %v, want %v", params.V, false)
	}
	
	if params.F != false {
		t.Errorf("F = %v, want %v", params.F, false)
	}
	
	if params.N != false {
		t.Errorf("N = %v, want %v", params.N, false)
	}
}

func TestNewParams_AllFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected Params
	}{
		{
			name: "After context flag",
			args: []string{"program", "-A", "3", "pattern"},
			expected: Params{
				A: 3, B: 0, C: 0, Count: false, I: false, V: false, F: false, N: false,
				Pattern: "pattern",
			},
		},
		{
			name: "Before context flag",
			args: []string{"program", "-B", "2", "pattern"},
			expected: Params{
				A: 0, B: 2, C: 0, Count: false, I: false, V: false, F: false, N: false,
				Pattern: "pattern",
			},
		},
		{
			name: "Around context flag",
			args: []string{"program", "-C", "5", "pattern"},
			expected: Params{
				A: 0, B: 0, C: 5, Count: false, I: false, V: false, F: false, N: false,
				Pattern: "pattern",
			},
		},
		{
			name: "Count flag",
			args: []string{"program", "-c", "pattern"},
			expected: Params{
				A: 0, B: 0, C: 0, Count: true, I: false, V: false, F: false, N: false,
				Pattern: "pattern",
			},
		},
		{
			name: "Case insensitive flag",
			args: []string{"program", "-i", "pattern"},
			expected: Params{
				A: 0, B: 0, C: 0, Count: false, I: true, V: false, F: false, N: false,
				Pattern: "pattern",
			},
		},
		{
			name: "Invert flag",
			args: []string{"program", "-v", "pattern"},
			expected: Params{
				A: 0, B: 0, C: 0, Count: false, I: false, V: true, F: false, N: false,
				Pattern: "pattern",
			},
		},
		{
			name: "Fixed string flag",
			args: []string{"program", "-f", "pattern"},
			expected: Params{
				A: 0, B: 0, C: 0, Count: false, I: false, V: false, F: true, N: false,
				Pattern: "pattern",
			},
		},
		{
			name: "Line numbers flag",
			args: []string{"program", "-n", "pattern"},
			expected: Params{
				A: 0, B: 0, C: 0, Count: false, I: false, V: false, F: false, N: true,
				Pattern: "pattern",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			os.Args = tt.args
			
			params := NewParams()
			
			if params.A != tt.expected.A {
				t.Errorf("A = %v, want %v", params.A, tt.expected.A)
			}
			if params.B != tt.expected.B {
				t.Errorf("B = %v, want %v", params.B, tt.expected.B)
			}
			if params.C != tt.expected.C {
				t.Errorf("C = %v, want %v", params.C, tt.expected.C)
			}
			if params.Count != tt.expected.Count {
				t.Errorf("Count = %v, want %v", params.Count, tt.expected.Count)
			}
			if params.I != tt.expected.I {
				t.Errorf("I = %v, want %v", params.I, tt.expected.I)
			}
			if params.V != tt.expected.V {
				t.Errorf("V = %v, want %v", params.V, tt.expected.V)
			}
			if params.F != tt.expected.F {
				t.Errorf("F = %v, want %v", params.F, tt.expected.F)
			}
			if params.N != tt.expected.N {
				t.Errorf("N = %v, want %v", params.N, tt.expected.N)
			}
			if params.Pattern != tt.expected.Pattern {
				t.Errorf("Pattern = %v, want %v", params.Pattern, tt.expected.Pattern)
			}
		})
	}
}

func TestNewParams_CombinedFlags(t *testing.T) {
	tests := []struct {
		name string
		args []string
		test func(*testing.T, *Params)
	}{
		{
			name: "Multiple flags combined",
			args: []string{"program", "-i", "-v", "-n", "-A", "2", "test"},
			test: func(t *testing.T, p *Params) {
				if !p.I || !p.V || !p.N || p.A != 2 {
					t.Errorf("Combined flags not set correctly: I=%v, V=%v, N=%v, A=%v", 
						p.I, p.V, p.N, p.A)
				}
			},
		},
		{
			name: "Context flags combined",
			args: []string{"program", "-A", "1", "-B", "2", "-C", "3", "test"},
			test: func(t *testing.T, p *Params) {
				if p.A != 1 || p.B != 2 || p.C != 3 {
					t.Errorf("Context flags not set correctly: A=%v, B=%v, C=%v", 
						p.A, p.B, p.C)
				}
			},
		},
		{
			name: "All boolean flags",
			args: []string{"program", "-c", "-i", "-v", "-f", "-n", "test"},
			test: func(t *testing.T, p *Params) {
				if !p.Count || !p.I || !p.V || !p.F || !p.N {
					t.Errorf("Boolean flags not set correctly: Count=%v, I=%v, V=%v, F=%v, N=%v", 
						p.Count, p.I, p.V, p.F, p.N)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			os.Args = tt.args
			
			params := NewParams()
			tt.test(t, params)
		})
	}
}

func TestNewParams_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		args []string
		test func(*testing.T, *Params)
	}{
		{
			name: "No pattern provided",
			args: []string{"program"},
			test: func(t *testing.T, p *Params) {
				if p.Pattern != "" {
					t.Errorf("Pattern should be empty when not provided, got %v", p.Pattern)
				}
			},
		},
		{
			name: "Zero context values",
			args: []string{"program", "-A", "0", "-B", "0", "-C", "0", "test"},
			test: func(t *testing.T, p *Params) {
				if p.A != 0 || p.B != 0 || p.C != 0 {
					t.Errorf("Zero context values not handled correctly: A=%v, B=%v, C=%v", 
						p.A, p.B, p.C)
				}
			},
		},
		{
			name: "Large context values",
			args: []string{"program", "-A", "100", "-B", "200", "-C", "300", "test"},
			test: func(t *testing.T, p *Params) {
				if p.A != 100 || p.B != 200 || p.C != 300 {
					t.Errorf("Large context values not handled correctly: A=%v, B=%v, C=%v", 
						p.A, p.B, p.C)
				}
			},
		},
		{
			name: "Pattern with special characters",
			args: []string{"program", "test@#$%^&*()"},
			test: func(t *testing.T, p *Params) {
				expected := "test@#$%^&*()"
				if p.Pattern != expected {
					t.Errorf("Special character pattern not handled correctly: got %v, want %v", 
						p.Pattern, expected)
				}
			},
		},
		{
			name: "Empty pattern",
			args: []string{"program", ""},
			test: func(t *testing.T, p *Params) {
				if p.Pattern != "" {
					t.Errorf("Empty pattern not handled correctly: got %v, want empty string", 
						p.Pattern)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			os.Args = tt.args
			
			params := NewParams()
			tt.test(t, params)
		})
	}
}

// Тест на валидацию типов параметров
func TestParams_StructValidation(t *testing.T) {
	params := &Params{
		A:       5,
		B:       3,
		C:       7,
		Count:   true,
		I:       true,
		V:       false,
		F:       true,
		N:       false,
		Pattern: "test_pattern",
	}

	// Проверяем, что все поля имеют правильные типы
	if params.A != 5 {
		t.Errorf("A field type validation failed")
	}
	if params.B != 3 {
		t.Errorf("B field type validation failed")
	}
	if params.C != 7 {
		t.Errorf("C field type validation failed")
	}
	if params.Count != true {
		t.Errorf("Count field type validation failed")
	}
	if params.I != true {
		t.Errorf("I field type validation failed")
	}
	if params.V != false {
		t.Errorf("V field type validation failed")
	}
	if params.F != true {
		t.Errorf("F field type validation failed")
	}
	if params.N != false {
		t.Errorf("N field type validation failed")
	}
	if params.Pattern != "test_pattern" {
		t.Errorf("Pattern field type validation failed")
	}
}

// Бенчмарк для создания параметров
func BenchmarkNewParams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resetFlags()
		os.Args = []string{"program", "-i", "-v", "-A", "2", "benchmark_pattern"}
		NewParams()
	}
}