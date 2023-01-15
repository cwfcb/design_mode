package factory_method

import (
	"errors"
	"mine/design_pattern/creational_pattern/factory_method/simple_factory"
)

/*
	工厂方法
*/

type IConfigParserFactory interface {
	CreateParser() simple_factory.RuleConfigParser
}

type (
	JsonRuleConfigParserFactory struct {
	}

	XmlRuleConfigParserFactory struct {
	}

	YamlRuleConfigParserFactory struct {
	}

	PropertiesRuleConfigParserFactory struct {
	}
)

func (j JsonRuleConfigParserFactory) CreateParser() simple_factory.RuleConfigParser {
	return simple_factory.JsonRuleConfigParser{}
}

func (x XmlRuleConfigParserFactory) CreateParser() simple_factory.RuleConfigParser {
	return simple_factory.XmlRuleConfigParser{}
}

func (y YamlRuleConfigParserFactory) CreateParser() simple_factory.RuleConfigParser {
	return simple_factory.YamlRuleConfigParser{}
}

func (p PropertiesRuleConfigParserFactory) CreateParser() simple_factory.RuleConfigParser {
	return simple_factory.PropertiesRuleConfigParser{}
}

// ConfigParserSource 是创建工厂对象的工厂类
type ConfigParserSource struct {
}

func (c *ConfigParserSource) CreateConfigParser(fileExtension string) (simple_factory.RuleConfigParser, error) {
	var factory IConfigParserFactory
	if fileExtension == "json" {
		factory = JsonRuleConfigParserFactory{}
	} else if fileExtension == "xml" {
		factory = XmlRuleConfigParserFactory{}
	} else if fileExtension == "yaml" {
		factory = YamlRuleConfigParserFactory{}
	} else if fileExtension == "properties" {
		factory = PropertiesRuleConfigParserFactory{}
	} else {
		return nil, errors.New("unsupported file extension")
	}

	parser := factory.CreateParser()
	return parser, nil
}

var cachedFactories map[string]IConfigParserFactory

func init() {
	cachedFactories = make(map[string]IConfigParserFactory)
	cachedFactories["json"] = JsonRuleConfigParserFactory{}
	cachedFactories["xml"] = XmlRuleConfigParserFactory{}
	cachedFactories["yaml"] = YamlRuleConfigParserFactory{}
	cachedFactories["properties"] = PropertiesRuleConfigParserFactory{}
}

func (c *ConfigParserSource) CreateConfigParserWithCache(fileExtension string) (simple_factory.RuleConfigParser, error) {
	factory, ok := cachedFactories[fileExtension]
	if !ok {
		return nil, errors.New("unsupported file extension")
	}
	parser := factory.CreateParser()
	return parser, nil
}
