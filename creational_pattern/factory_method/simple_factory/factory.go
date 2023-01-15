package simple_factory

import "errors"

type ConfigParserFactory struct {
}

func (c ConfigParserFactory) CreateConfigParser(fileExtension string) (RuleConfigParser, error) {
	var parser RuleConfigParser
	if fileExtension == "json" {
		parser = JsonRuleConfigParser{}
	} else if fileExtension == "xml" {
		parser = XmlRuleConfigParser{}
	} else if fileExtension == "yaml" {
		parser = YamlRuleConfigParser{}
	} else if fileExtension == "properties" {
		parser = PropertiesRuleConfigParser{}
	} else {
		return nil, errors.New("unsupported file extension")
	}
	return parser, nil
}

var cachedParsers map[string]RuleConfigParser

func init() {
	cachedParsers = make(map[string]RuleConfigParser)
	cachedParsers["json"] = JsonRuleConfigParser{}
	cachedParsers["xml"] = XmlRuleConfigParser{}
	cachedParsers["yaml"] = YamlRuleConfigParser{}
	cachedParsers["properties"] = PropertiesRuleConfigParser{}
}

func (c ConfigParserFactory) CreateConfigParserWithCache(fileExtension string) (RuleConfigParser, error) {
	parser, ok := cachedParsers[fileExtension]
	if ok {
		return parser, nil
	}
	return nil, errors.New("unsupported file extension")
}
