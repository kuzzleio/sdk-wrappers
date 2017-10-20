package main

// map which stores instances to keep references in case the gc passes
var instances map[interface{}]interface{}

// register new instance to the instances map
func registerInstance(instance interface{}) {
	instances[instance] = nil
}