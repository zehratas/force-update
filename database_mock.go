package main

type mockDatabase struct {
	storage map[string]string
}

func (d *mockDatabase) Get(key string) string {
	return d.storage[key]
}
