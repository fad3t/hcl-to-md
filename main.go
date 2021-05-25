package main

import (
  "fmt"
  "log"
  "github.com/hashicorp/hcl/v2/hclsimple"
)

type Variable struct {
  Name        string `hcl:"name,label"`
  Description string `hcl:"description,optional"`
  Default     string `hcl:"default,optional"`
  Type        string `hcl:"type,optional"`
}

type Config struct {
  Variables []Variable `hcl:"variable,block"`
}

func main() {
  var config Config
  err := hclsimple.DecodeFile("variables.hcl", nil, &config)
  if err != nil {
    log.Fatalf("Failed to load configuration: %s", err)
  }
  fmt.Printf("| Name | Type | Description | Default |\n")
  fmt.Printf("|------|------|-------------|---------|\n")
  for _, s := range config.Variables {
    fmt.Printf("| %v | %v | %v | %v |\n", s.Name, s.Type, s.Description, s.Default)
  }
}
