locals {
  namespace   = "eg"
  tenant      = "plat"
  environment = "ue2"
  stage       = "dev"
  name        = "test"
  delimiter   = "-"

  result1 = yamldecode(data.atmos_label.example1.output)
}

data "atmos_label" "example1" {
  namespace   = local.namespace
  tenant      = local.tenant
  environment = local.environment
  stage       = local.stage
  name        = local.name
  delimiter   = local.delimiter
}
