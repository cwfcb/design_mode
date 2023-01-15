package builder

type ConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
	}
}

func (c *ConfigBuilder) SetName(name string) *ConfigBuilder {
	if name == "" {
		panic("")
	}
	c.name = name
	return c
}

func (c *ConfigBuilder) SetMaxTotal(maxTotal int) *ConfigBuilder {
	if maxTotal <= 0 {
		panic("")
	}
	c.maxTotal = maxTotal
	return c
}

func (c *ConfigBuilder) SetMaxIdle(maxIdle int) *ConfigBuilder {
	if maxIdle <= 0 {
		panic("")
	}
	c.maxIdle = maxIdle
	return c
}

func (c *ConfigBuilder) SetMinIdle(minIdle int) *ConfigBuilder {
	if minIdle < 0 {
		panic("")
	}
	c.minIdle = minIdle
	return c
}

func (c *ConfigBuilder) Build() *ConfigBuilder {
	// 校验逻辑放到这里来做，包括必填项校验、依赖关系校验、约束条件校验等
	if c.name == "" {
		panic("")
	}
	if c.maxIdle > c.maxTotal {
		panic("")
	}
	if c.minIdle > c.maxIdle || c.minIdle > c.maxTotal {
		panic("")
	}
	return c
}

// optional program
type option func(c *ConfigBuilder)

func WithName(name string) option {
	return func(c *ConfigBuilder) {
		if name == "" {
			panic("")
		}
		c.name = name
	}
}

func WithMaxTotal(maxTotal int) option {
	return func(c *ConfigBuilder) {
		if maxTotal <= 0 {
			panic("")
		}
		c.maxTotal = maxTotal
	}
}

func WithMaxIdle(maxIdle int) option {
	return func(c *ConfigBuilder) {
		if maxIdle <= 0 {
			panic("")
		}
		c.maxIdle = maxIdle
	}
}

func WithMinIdle(minIdle int) option {
	return func(c *ConfigBuilder) {
		if minIdle < 0 {
			panic("")
		}
		c.minIdle = minIdle
	}
}

func NewBuilder(opts ...option) *ConfigBuilder {
	builder := NewConfigBuilder()
	for _, opt := range opts {
		opt(builder)
	}
	return builder.Build()
}
