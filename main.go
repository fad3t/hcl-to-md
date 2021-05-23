package main

import (
  "log"
  "github.com/hashicorp/hcl/hclsimple"
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
  err := hclsimple.DecodeFile("variables.tfvars", nil, &config)
  if err != nil {
    log.Fatalf("Failed to load configuration: %s", err)
  }
  log.Printf("Configuration is %#v", config)
}
