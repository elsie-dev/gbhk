provider "aws"{
  region= "us-east-1"
  version= "~>4.61.0"
}

resource "aws_vpc" "my_vpc" {
  cidr_block = "172.16.0.0/16"

  tags = {
    Name = "twalavpc"
  }
}

resource "aws_subnet" "my_subnet" {
  vpc_id            = aws_vpc.my_vpc.id
  cidr_block        = "172.16.10.0/24"
  availability_zone = "us-east-1"

  tags = {
    Name = "twalasubnet"
  }
}

resource "aws_network_interface" "mockserver" {
  subnet_id   = aws_subnet.my_subnet.id
  private_ips = ["172.16.10.100"]

  tags = {
    Name = "primary_network_interface"
  }
}

resource "aws_instance" "foo" {
  ami           = "ami-007855ac798b5175e" # us-east-1
  instance_type = "t3.medium"

  network_interface {
    network_interface_id = aws_network_interface.mockserver.id
    device_index         = 0
  }

  credit_specification {
    cpu_credits = "unlimited"
  }
}
