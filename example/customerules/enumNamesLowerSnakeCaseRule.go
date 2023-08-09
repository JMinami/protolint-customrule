package customrules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"

	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/rule"
	"github.com/yoheimuta/protolint/linter/strs"
	"github.com/yoheimuta/protolint/linter/visitor"
)

// EnumNamesLowerSnakeCaseRule verifies that all enum names are LowerSnakeCase.
type EnumNamesLowerSnakeCaseRule struct {
}

// NewEnumNamesLowerSnakeCaseRule creates a new EnumNamesLowerSnakeCaseRule.
func NewEnumNamesLowerSnakeCaseRule() EnumNamesLowerSnakeCaseRule {
	return EnumNamesLowerSnakeCaseRule{}
}

// ID returns the ID of this rule.
func (r EnumNamesLowerSnakeCaseRule) ID() string {
	return "ENUM_NAMES_LOWER_SNAKE_CASE"
}

// Purpose returns the purpose of this rule.
func (r EnumNamesLowerSnakeCaseRule) Purpose() string {
	return "Verifies that all enum names are LowerSnakeCase."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r EnumNamesLowerSnakeCaseRule) IsOfficial() bool {
	return true
}

// Severity gets the severity of the rule
func (r EnumNamesLowerSnakeCaseRule) Severity() rule.Severity {
	return rule.SeverityWarning
}

// Apply applies the rule to the proto.
func (r EnumNamesLowerSnakeCaseRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &enumNamesLowerSnakeCaseVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID(), string(r.Severity())),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type enumNamesLowerSnakeCaseVisitor struct {
	*visitor.BaseAddVisitor
}

// VisitEnum checks the enum field.
func (v *enumNamesLowerSnakeCaseVisitor) VisitEnum(e *parser.Enum) bool {
	if !strs.IsLowerSnakeCase(e.EnumName) {
		v.AddFailuref(e.Meta.Pos, "Enum name %q must be underscore_separated_names", e.EnumName)
	}
	return false
}

func (v *enumNamesLowerSnakeCaseVisitor) VisitExtend(e *parser.Extend) bool {
	v.AddFailuref(e.Meta.LastPos, "Message Type %s", e.MessageType)
	v.AddFailuref(e.Meta.LastPos, "InlinceCommentBefineLeftCurly %s", e.InlineCommentBehindLeftCurly)
	return false
}

func (v *enumNamesLowerSnakeCaseVisitor) VisitField(f *parser.Field) bool {
	v.AddFailuref(f.Meta.LastPos, "type %s, %s", f.Type, f.FieldName)
	for _, c := range f.Comments {
		v.AddFailuref(f.Meta.LastPos, "comment, %s", c.Raw)
	}
	// v.AddFailuref(f.Meta.LastPos, "inline comment, %s", f.InlineComment.Raw)
	if f.InlineComment != nil {
		v.AddFailuref(f.Meta.LastPos, "inline comment, %s", f.InlineComment.Raw)
	}
	return false
}
