locals {
  component   = "test/test-component-override"
  stack       = "tenant1-ue2-dev"
  tenant      = "tenant1"
  environment = "ue2"
  stage       = "dev"

  result1 = yamldecode(data.atmos_component_config.example1.output)
  result2 = yamldecode(data.atmos_component_config.example2.output)
}

data "atmos_component_config" "example1" {
  component     = local.component
  stack         = local.stack
  ignore_errors = false
}

data "atmos_component_config" "example2" {
  component     = local.component
  tenant        = local.tenant
  environment   = local.environment
  stage         = local.stage
  ignore_errors = false
}
