package abstract_factory

type ISystemConfigParser interface {
	Parse(content string) ([]byte, error)
}

type (
	JsonSystemConfigParser struct {
	}

	XmlSystemConfigParser struct {
	}

	YamlSystemConfigParser struct {
	}

	PropertiesSystemConfigParser struct {
	}
)

func (j JsonSystemConfigParser) Parse(content string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (x XmlSystemConfigParser) Parse(content string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (y YamlSystemConfigParser) Parse(content string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (p PropertiesSystemConfigParser) Parse(content string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}
