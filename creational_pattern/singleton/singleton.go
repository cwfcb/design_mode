package singleton

import (
	"fmt"
	"sync"
)

/*
	单例模式的三种实现方式
*/

type SingleInstance struct{}

var (
	instance *SingleInstance
	mutex    sync.Mutex
)

func init() {
	instance = &SingleInstance{}
}

// GetInstance 饥汉式
func GetInstance() *SingleInstance {
	return instance
}

// GetInstanceLazily 懒汉式
func GetInstanceLazily() *SingleInstance {
	if instance == nil {
		instance = &SingleInstance{}
	}
	return instance
}

// GetInstanceSafety 双重检测式
func GetInstanceSafety() *SingleInstance {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &SingleInstance{}
		}
		mutex.Unlock()
	}
	fmt.Println("Single instance already created.")
	return instance
}
