package main

import (
	customrules "example/customerules"
	"flag"

	"github.com/yoheimuta/protolint/linter/rule"
	"github.com/yoheimuta/protolint/plugin"
)

var ()

func main() {
	flag.Parse()

	plugin.RegisterCustomRules(
		customrules.NewEnumNamesLowerSnakeCaseRule(),

		// Wrapping with RuleGen allows referring to command-line flags.
		plugin.RuleGen(func(
			verbose bool,
			fixMode bool,
		) rule.Rule {
			return customrules.NewSimpleRule(verbose, fixMode, rule.SeverityError)
		}),
	)
}
