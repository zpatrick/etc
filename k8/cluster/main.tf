provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "${var.aws_region}"
}

data "aws_availability_zones" "available" {}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  name   = "${var.name}"
  cidr   = "10.0.0.0/16"

  azs             = "${data.aws_availability_zones.available.names}"
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway           = true
  single_nat_gateway           = true
  create_database_subnet_group = false

  tags = {
    Name = "${var.name}"
  }
}

module "cluster" {
  source        = "/Users/zpatrick/go/src/github.com/quintilesims/terraform-k8-cluster"
  name          = "${var.name}"
  azs           = "${data.aws_availability_zones.available.names}"
  master_ami_id = "todo"
  node_ami_id   = "todo"
  public_subnet_ids = "${module.vpc.public_subnets}"
}
