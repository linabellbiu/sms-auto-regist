package main

import (
	"auto-reg/register"
	"auto-reg/www_yunjiema_top"
)

func main() {

	go register.Register()

	www_yunjiema_top.On()
}
