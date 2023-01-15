package simple_factory

/*
	简单工厂
*/

type RuleConfigParser interface {
	Parse(content string) (interface{}, error)
}

type (
	JsonRuleConfigParser struct{}

	XmlRuleConfigParser struct{}

	YamlRuleConfigParser struct{}

	PropertiesRuleConfigParser struct{}
)

func (j JsonRuleConfigParser) Parse(content string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (x XmlRuleConfigParser) Parse(content string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (y YamlRuleConfigParser) Parse(content string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (p PropertiesRuleConfigParser) Parse(content string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
