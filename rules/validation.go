// Package rules provides validation rules for form components
package rules

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/FlameMida/form-builder-go/contracts"
)

// RequiredRule validates that a field is not empty
type RequiredRule struct {
	message string
}

// NewRequiredRule creates a new required validation rule
func NewRequiredRule(message string) *RequiredRule {
	if message == "" {
		message = "此字段为必填项"
	}
	return &RequiredRule{message: message}
}

// Type returns the rule type
func (r *RequiredRule) Type() string {
	return "required"
}

// Message returns the validation message
func (r *RequiredRule) Message() string {
	return r.message
}

// Validate validates that the value is not empty
func (r *RequiredRule) Validate(value interface{}) error {
	if value == nil {
		return errors.New(r.message)
	}

	switch v := value.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return errors.New(r.message)
		}
	case []interface{}:
		if len(v) == 0 {
			return errors.New(r.message)
		}
	default:
		if reflect.ValueOf(value).IsZero() {
			return errors.New(r.message)
		}
	}

	return nil
}

// ToMap returns the rule as a map
func (r *RequiredRule) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"required": true,
		"message":  r.message,
		"type":     r.Type(),
	}
}

// LengthRule validates string length
type LengthRule struct {
	Min     int
	Max     int
	message string
}

// NewLengthRule creates a new length validation rule
func NewLengthRule(min, max int, message string) *LengthRule {
	if message == "" {
		message = fmt.Sprintf("长度必须在 %d 到 %d 之间", min, max)
	}
	return &LengthRule{
		Min:     min,
		Max:     max,
		message: message,
	}
}

// Type returns the rule type
func (r *LengthRule) Type() string {
	return "length"
}

// Message returns the validation message
func (r *LengthRule) Message() string {
	return r.message
}

// Validate validates the string length
func (r *LengthRule) Validate(value interface{}) error {
	if value == nil {
		return nil // Length validation only applies to non-nil values
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("长度验证只适用于字符串类型")
	}

	length := len([]rune(str)) // Count Unicode characters properly
	if length < r.Min || length > r.Max {
		return errors.New(r.message)
	}

	return nil
}

// ToMap returns the rule as a map
func (r *LengthRule) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"min":     r.Min,
		"max":     r.Max,
		"message": r.message,
		"type":    r.Type(),
	}
}

// EmailRule validates email format
type EmailRule struct {
	message string
	pattern *regexp.Regexp
}

// NewEmailRule creates a new email validation rule
func NewEmailRule(message string) *EmailRule {
	if message == "" {
		message = "请输入有效的邮箱地址"
	}

	// Basic email regex pattern
	pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return &EmailRule{
		message: message,
		pattern: pattern,
	}
}

// Type returns the rule type
func (r *EmailRule) Type() string {
	return "email"
}

// Message returns the validation message
func (r *EmailRule) Message() string {
	return r.message
}

// Validate validates email format
func (r *EmailRule) Validate(value interface{}) error {
	if value == nil {
		return nil // Email validation only applies to non-nil values
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("邮箱验证只适用于字符串类型")
	}

	if strings.TrimSpace(str) == "" {
		return nil // Empty values are handled by required rule
	}

	if !r.pattern.MatchString(str) {
		return errors.New(r.message)
	}

	return nil
}

// ToMap returns the rule as a map
func (r *EmailRule) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":    r.Type(),
		"message": r.message,
	}
}

// NumberRule validates numeric values
type NumberRule struct {
	Min     *float64
	Max     *float64
	message string
}

// NewNumberRule creates a new number validation rule
func NewNumberRule(min, max *float64, message string) *NumberRule {
	if message == "" {
		if min != nil && max != nil {
			message = fmt.Sprintf("数值必须在 %.2f 到 %.2f 之间", *min, *max)
		} else if min != nil {
			message = fmt.Sprintf("数值必须大于等于 %.2f", *min)
		} else if max != nil {
			message = fmt.Sprintf("数值必须小于等于 %.2f", *max)
		} else {
			message = "请输入有效的数值"
		}
	}

	return &NumberRule{
		Min:     min,
		Max:     max,
		message: message,
	}
}

// Type returns the rule type
func (r *NumberRule) Type() string {
	return "number"
}

// Message returns the validation message
func (r *NumberRule) Message() string {
	return r.message
}

// Validate validates numeric values
func (r *NumberRule) Validate(value interface{}) error {
	if value == nil {
		return nil // Number validation only applies to non-nil values
	}

	var num float64
	var err error

	switch v := value.(type) {
	case float64:
		num = v
	case int:
		num = float64(v)
	case string:
		if strings.TrimSpace(v) == "" {
			return nil // Empty values are handled by required rule
		}
		num, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return errors.New("请输入有效的数值")
		}
	default:
		return errors.New("数值验证只适用于数字或字符串类型")
	}

	if r.Min != nil && num < *r.Min {
		return errors.New(r.message)
	}

	if r.Max != nil && num > *r.Max {
		return errors.New(r.message)
	}

	return nil
}

// ToMap returns the rule as a map
func (r *NumberRule) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"type":    r.Type(),
		"message": r.message,
	}

	if r.Min != nil {
		result["min"] = *r.Min
	}

	if r.Max != nil {
		result["max"] = *r.Max
	}

	return result
}

// PatternRule validates against a regular expression
type PatternRule struct {
	Pattern *regexp.Regexp
	message string
}

// NewPatternRule creates a new pattern validation rule
func NewPatternRule(pattern string, message string) (*PatternRule, error) {
	if message == "" {
		message = "格式不正确"
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("无效的正则表达式: %w", err)
	}

	return &PatternRule{
		Pattern: regex,
		message: message,
	}, nil
}

// Type returns the rule type
func (r *PatternRule) Type() string {
	return "pattern"
}

// Message returns the validation message
func (r *PatternRule) Message() string {
	return r.message
}

// Validate validates against the pattern
func (r *PatternRule) Validate(value interface{}) error {
	if value == nil {
		return nil // Pattern validation only applies to non-nil values
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("模式验证只适用于字符串类型")
	}

	if strings.TrimSpace(str) == "" {
		return nil // Empty values are handled by required rule
	}

	if !r.Pattern.MatchString(str) {
		return errors.New(r.message)
	}

	return nil
}

// ToMap returns the rule as a map
func (r *PatternRule) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"pattern": r.Pattern.String(),
		"message": r.message,
		"type":    r.Type(),
	}
}

// ValidatorChain allows chaining multiple validation rules
type ValidatorChain struct {
	rules []contracts.ValidateRule
}

// NewValidatorChain creates a new validator chain
func NewValidatorChain() *ValidatorChain {
	return &ValidatorChain{
		rules: make([]contracts.ValidateRule, 0),
	}
}

// Add adds a validation rule to the chain
func (v *ValidatorChain) Add(rule contracts.ValidateRule) *ValidatorChain {
	v.rules = append(v.rules, rule)
	return v
}

// Required adds a required validation rule
func (v *ValidatorChain) Required(message string) *ValidatorChain {
	return v.Add(NewRequiredRule(message))
}

// Length adds a length validation rule
func (v *ValidatorChain) Length(min, max int, message string) *ValidatorChain {
	return v.Add(NewLengthRule(min, max, message))
}

// Email adds an email validation rule
func (v *ValidatorChain) Email(message string) *ValidatorChain {
	return v.Add(NewEmailRule(message))
}

// Number adds a number validation rule
func (v *ValidatorChain) Number(min, max *float64, message string) *ValidatorChain {
	return v.Add(NewNumberRule(min, max, message))
}

// Pattern adds a pattern validation rule
func (v *ValidatorChain) Pattern(pattern, message string) *ValidatorChain {
	rule, err := NewPatternRule(pattern, message)
	if err != nil {
		panic(err) // In production, handle this more gracefully
	}
	return v.Add(rule)
}

// Validate validates a value against all rules in the chain
func (v *ValidatorChain) Validate(value interface{}) error {
	for _, rule := range v.rules {
		if err := rule.Validate(value); err != nil {
			return err
		}
	}
	return nil
}

// GetRules returns all rules in the chain
func (v *ValidatorChain) GetRules() []contracts.ValidateRule {
	return v.rules
}
