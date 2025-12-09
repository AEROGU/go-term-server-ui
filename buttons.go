package main

type Button struct {
	Label    string
	Handler  func()
	Disabled bool
}
