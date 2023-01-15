package abstract_factory

import (
	"errors"
	"mine/design_pattern/creational_pattern/factory_method/simple_factory"
)

/*
	抽象工厂
*/

type IConfigParserFactory interface {
	CreateRuleParser() simple_factory.RuleConfigParser
	CreateSystemParser() ISystemConfigParser
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

func (j JsonRuleConfigParserFactory) CreateRuleParser() simple_factory.RuleConfigParser {
	return simple_factory.JsonRuleConfigParser{}
}

func (j JsonRuleConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return JsonSystemConfigParser{}
}

func (x XmlRuleConfigParserFactory) CreateRuleParser() simple_factory.RuleConfigParser {
	return simple_factory.XmlRuleConfigParser{}
}

func (x XmlRuleConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return XmlSystemConfigParser{}
}

func (y YamlRuleConfigParserFactory) CreateRuleParser() simple_factory.RuleConfigParser {
	return simple_factory.YamlRuleConfigParser{}
}

func (y YamlRuleConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return YamlSystemConfigParser{}
}

func (p PropertiesRuleConfigParserFactory) CreateRuleParser() simple_factory.RuleConfigParser {
	return simple_factory.PropertiesRuleConfigParser{}
}

func (p PropertiesRuleConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return PropertiesSystemConfigParser{}
}

// ConfigParserSource 是创建工厂对象的工厂类
type ConfigParserSource struct {
}

func (c *ConfigParserSource) CreateConfigParser(fileExtension, configType string) (interface{}, error) {
	if configType == "rule" {
		return c.createRuleConfigParser(fileExtension)
	}
	if configType == "system" {
		return c.createSystemConfigParser(fileExtension)
	}
	return nil, errors.New("unsupported config type")
}

func (c *ConfigParserSource) createRuleConfigParser(fileExtension string) (simple_factory.RuleConfigParser, error) {
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

	return factory.CreateRuleParser(), nil
}

func (c *ConfigParserSource) createSystemConfigParser(fileExtension string) (ISystemConfigParser, error) {
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

	return factory.CreateSystemParser(), nil
}

var cachedFactories map[string]IConfigParserFactory

func init() {
	cachedFactories = make(map[string]IConfigParserFactory)
	cachedFactories["json"] = JsonRuleConfigParserFactory{}
	cachedFactories["xml"] = XmlRuleConfigParserFactory{}
	cachedFactories["yaml"] = YamlRuleConfigParserFactory{}
	cachedFactories["properties"] = PropertiesRuleConfigParserFactory{}
}

func (c *ConfigParserSource) createRuleConfigParserWithCache(fileExtension string) (simple_factory.RuleConfigParser, error) {
	factory, ok := cachedFactories[fileExtension]
	if !ok {
		return nil, errors.New("unsupported file extension")
	}
	parser := factory.CreateRuleParser()
	return parser, nil
}

func (c *ConfigParserSource) createSystemConfigParserWithCache(fileExtension string) (ISystemConfigParser, error) {
	factory, ok := cachedFactories[fileExtension]
	if !ok {
		return nil, errors.New("unsupported file extension")
	}
	parser := factory.CreateSystemParser()
	return parser, nil
}
