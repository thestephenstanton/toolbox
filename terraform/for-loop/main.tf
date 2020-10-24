// 0.11

variable "fubars" {
  type    = "list"
  default = ["foo", "bar"]
}

resource "aws_instance" "example" {
  count = "${length(var.fubars)}"

  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"

  tags = {
    Name = "${element(var.fubars, count.index)}"
  }
}
