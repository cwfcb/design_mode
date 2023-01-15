package builder

const (
	// default value
	defaultMaxTotal = 8
	defaultMaxIdle  = 8
	defaultMinIdle  = 2
)

type ResourcePoolConfig struct {
	name     string // required
	maxTotal int    // optional
	maxIdle  int    // optional
	minIdle  int    // optional
}

func NewResourcePoolConfig(name string) ResourcePoolConfig {
	if name == "" {
		panic("empty name")
	}
	return ResourcePoolConfig{
		name:     name,
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
	}
}

func (r *ResourcePoolConfig) SetMaxTotal(maxTotal int) {
	if maxTotal <= 0 {
		panic("")
	}
	r.maxTotal = maxTotal
}

func (r *ResourcePoolConfig) SetMaxIdle(maxIdle int) {
	if maxIdle <= 0 {
		panic("")
	}
	r.maxIdle = maxIdle
}

func (r *ResourcePoolConfig) SetMinIdle(minIdle int) {
	if minIdle < 0 {
		panic("")
	}
	r.minIdle = minIdle
}

func NewResourcePoolConfigByBuilder(builder *ConfigBuilder) ResourcePoolConfig {
	return ResourcePoolConfig{
		name:     builder.name,
		maxTotal: builder.maxTotal,
		maxIdle:  builder.maxIdle,
		minIdle:  builder.minIdle,
	}
}
