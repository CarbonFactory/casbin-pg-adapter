package model

import (
	"strings"
)

const (
	policyLinePrefix = ", "
)

// CasbinRule is the model for casbin rule
type CasbinRule struct {
	TableName struct{} `sql:"x_policy" pg:",discard_unknown_columns" `
	PType     string   `sql:",pType" db:"p_type" `
	V0        string   `sql:",v0" db:"v0" `
	V1        string   `sql:",v1" db:"v1" `
	V2        string   `sql:",v2" db:"v2" `
	V3        string   `sql:",v3" db:"v3" `
	V4        string   `sql:",v4" db:"v4" `
	V5        string   `sql:",v5" db:"v5" `
}

// NewCasbinRuleFromPTypeAndRule returns a CasbinRule from pType and rule
func NewCasbinRuleFromPTypeAndRule(pType string, rule []string) CasbinRule {
	casbinRule := CasbinRule{
		PType: pType,
	}
	ruleLength := len(rule)
	if ruleLength > 0 {
		casbinRule.V0 = rule[0]
	}
	if ruleLength > 1 {
		casbinRule.V1 = rule[1]
	}
	if ruleLength > 2 {
		casbinRule.V2 = rule[2]
	}
	if ruleLength > 3 {
		casbinRule.V3 = rule[3]
	}
	if ruleLength > 4 {
		casbinRule.V4 = rule[4]
	}
	if ruleLength > 5 {
		casbinRule.V5 = rule[5]
	}
	return casbinRule
}

// NewCasbinRuleFromPTypeAndFilter returns a CasbinRule from pType and filter
func NewCasbinRuleFromPTypeAndFilter(pType string, fieldIndex int, fieldValues ...string) CasbinRule {
	casbinRule := CasbinRule{
		PType: pType,
	}

	idx := fieldIndex + len(fieldValues)
	if fieldIndex <= 0 && idx > 0 {
		casbinRule.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && idx > 1 {
		casbinRule.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && idx > 2 {
		casbinRule.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && idx > 3 {
		casbinRule.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && idx > 4 {
		casbinRule.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && idx > 5 {
		casbinRule.V5 = fieldValues[5-fieldIndex]
	}

	return casbinRule
}

// ToPolicyLine map casbinRule to a policy line used in casbin
func (casbinRule CasbinRule) ToPolicyLine() string {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(casbinRule.PType)
	if len(casbinRule.V0) > 0 {
		stringBuilder.WriteString(policyLinePrefix)
		stringBuilder.WriteString(casbinRule.V0)
	}
	if len(casbinRule.V1) > 0 {
		stringBuilder.WriteString(policyLinePrefix)
		stringBuilder.WriteString(casbinRule.V1)
	}
	if len(casbinRule.V2) > 0 {
		stringBuilder.WriteString(policyLinePrefix)
		stringBuilder.WriteString(casbinRule.V2)
	}
	if len(casbinRule.V3) > 0 {
		stringBuilder.WriteString(policyLinePrefix)
		stringBuilder.WriteString(casbinRule.V3)
	}
	if len(casbinRule.V4) > 0 {
		stringBuilder.WriteString(policyLinePrefix)
		stringBuilder.WriteString(casbinRule.V4)
	}
	if len(casbinRule.V5) > 0 {
		stringBuilder.WriteString(policyLinePrefix)
		stringBuilder.WriteString(casbinRule.V5)
	}

	return stringBuilder.String()
}
