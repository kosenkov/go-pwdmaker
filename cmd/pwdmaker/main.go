package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/kosenkov/go-pwdmaker/internal/pwdgenerator"
)

func main() {

	var opts struct {
		Resource string `short:"r" long:"resource" required:"true" description:"The name of the resource for which the password is being created"`

		Login string `short:"l" long:"login" required:"true" description:"User login for the specified resource"`

		Number int `short:"n" long:"number" required:"true" description:"Required number of characters"`
	}

	args := os.Args

	_, err := flags.ParseArgs(&opts, args)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Resource: %s\n", opts.Resource)
	fmt.Printf("Login: %s\n", opts.Login)
	fmt.Printf("Generated pwd %s:\n", pwdgenerator.Generate(opts.Resource+opts.Login, opts.Number, true))
}
