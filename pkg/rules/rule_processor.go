package rules

type RuleProcessor interface {
	Name() string
}

func GetRuleProcessor(name string) RuleProcessor {
	switch name {
	case "github.com/spicehq/labels":
		return &spicehqLabels{}
	}

	return nil
}

type spicehqLabels struct {
}

func (s *spicehqLabels) Name() string {
	return "github.com/spicehq/labels"
}
