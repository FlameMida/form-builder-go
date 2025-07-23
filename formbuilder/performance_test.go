package formbuilder

import (
	"fmt"
	"testing"

	"github.com/FlameMida/form-builder-go/components"
	"github.com/FlameMida/form-builder-go/ui/elm"
)

// BenchmarkFormBuilder_SmallForm tests performance with small forms (10 components)
func BenchmarkFormBuilder_SmallForm(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(10)
	config := map[string]interface{}{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		form, err := NewForm(bootstrap, "/test", rules, config)
		if err != nil {
			b.Fatal(err)
		}
		_ = form.FormRule()
	}
}

// BenchmarkFormBuilder_MediumForm tests performance with medium forms (100 components)
func BenchmarkFormBuilder_MediumForm(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(100)
	config := map[string]interface{}{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		form, err := NewForm(bootstrap, "/test", rules, config)
		if err != nil {
			b.Fatal(err)
		}
		_ = form.FormRule()
	}
}

// BenchmarkFormBuilder_LargeForm tests performance with large forms (1000 components)
func BenchmarkFormBuilder_LargeForm(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(1000)
	config := map[string]interface{}{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		form, err := NewForm(bootstrap, "/test", rules, config)
		if err != nil {
			b.Fatal(err)
		}
		_ = form.FormRule()
	}
}

// BenchmarkFormBuilder_JSONSerialization tests JSON serialization performance
func BenchmarkFormBuilder_JSONSerialization(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(100)
	config := map[string]interface{}{}
	
	form, err := NewForm(bootstrap, "/test", rules, config)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := form.ParseFormRule()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFormBuilder_JSONSerialization_WithCache tests cached JSON serialization performance
func BenchmarkFormBuilder_JSONSerialization_WithCache(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(100)
	config := map[string]interface{}{}
	
	form, err := NewForm(bootstrap, "/test", rules, config)
	if err != nil {
		b.Fatal(err)
	}
	
	// Warm up cache
	_, err = form.ParseFormRule()
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := form.ParseFormRule()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFormBuilder_Append tests append operation performance
func BenchmarkFormBuilder_Append(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	form, err := NewForm(bootstrap, "/test", []interface{}{}, map[string]interface{}{})
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		component := components.NewInput(fmt.Sprintf("field_%d", i), "Test Field")
		_, err := form.Append(component)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFormBuilder_FieldUniqueCheck tests field uniqueness check performance
func BenchmarkFormBuilder_FieldUniqueCheck(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := NewForm(bootstrap, "/test", rules, map[string]interface{}{})
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFormBuilder_ConcurrentAccess tests concurrent access performance
func BenchmarkFormBuilder_ConcurrentAccess(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(100)
	form, err := NewForm(bootstrap, "/test", rules, map[string]interface{}{})
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = form.FormRule()
		}
	})
}

// BenchmarkFormBuilder_FormDataSet tests form data setting performance
func BenchmarkFormBuilder_FormDataSet(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(100)
	form, err := NewForm(bootstrap, "/test", rules, map[string]interface{}{})
	if err != nil {
		b.Fatal(err)
	}
	
	formData := make(map[string]interface{})
	for i := 0; i < 100; i++ {
		formData[fmt.Sprintf("field_%d", i)] = fmt.Sprintf("value_%d", i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		form.FormData(formData)
	}
}

// createTestComponents creates test components for benchmarking
func createTestComponents(count int) []interface{} {
	rules := make([]interface{}, count)
	for i := 0; i < count; i++ {
		switch i % 5 {
		case 0:
			rules[i] = components.NewInput(fmt.Sprintf("field_%d", i), fmt.Sprintf("Input %d", i))
		case 1:
			rules[i] = components.NewTextarea(fmt.Sprintf("field_%d", i), fmt.Sprintf("Textarea %d", i))
		case 2:
			rules[i] = components.NewSwitch(fmt.Sprintf("field_%d", i), fmt.Sprintf("Switch %d", i))
		case 3:
			input := components.NewInput(fmt.Sprintf("field_%d", i), fmt.Sprintf("Required Input %d", i))
			input.Required()
			rules[i] = input
		case 4:
			textarea := components.NewTextarea(fmt.Sprintf("field_%d", i), fmt.Sprintf("Large Textarea %d", i))
			textarea.Rows(10)
			rules[i] = textarea
		}
	}
	return rules
}

// BenchmarkMemoryAllocation tests memory allocation patterns
func BenchmarkMemoryAllocation_FormCreation(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(100)
	config := map[string]interface{}{}
	
	b.ReportAllocs()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		form, err := NewForm(bootstrap, "/test", rules, config)
		if err != nil {
			b.Fatal(err)
		}
		_ = form
	}
}

// BenchmarkMemoryAllocation_FormRule tests memory allocation for FormRule
func BenchmarkMemoryAllocation_FormRule(b *testing.B) {
	bootstrap := elm.NewBootstrap()
	rules := createTestComponents(100)
	form, err := NewForm(bootstrap, "/test", rules, map[string]interface{}{})
	if err != nil {
		b.Fatal(err)
	}
	
	b.ReportAllocs()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = form.FormRule()
	}
}