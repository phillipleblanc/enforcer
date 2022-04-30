package rules

type Rules interface {
	Name() string
}

func GetRule(name string) Rules {
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
