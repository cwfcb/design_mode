package builder

import "testing"

func TestNewResourcePoolConfig(t *testing.T) {
	config := NewResourcePoolConfig("test")
	config.SetMaxTotal(10)
	config.SetMaxIdle(8)
	config.SetMinIdle(2)
	t.Log(config)
}

func TestNewConfigBuilder(t *testing.T) {
	builder := NewConfigBuilder().
		SetName("test").
		SetMaxTotal(10).
		SetMaxIdle(8).
		SetMinIdle(2).
		Build()

	config := NewResourcePoolConfigByBuilder(builder)
	t.Log(config)
}

func TestNewConfigWithOption(t *testing.T) {
	opts := []option{
		WithName("test"),
		WithMaxTotal(10),
		WithMaxIdle(8),
		WithMinIdle(2),
	}
	builder := NewBuilder(opts...)
	config := NewResourcePoolConfigByBuilder(builder)
	t.Log(config)
}
