resource "aws_kinesis_stream" "stryker" {
  name             = "stryker"
  shard_count      = 1
  

  shard_level_metrics = [
    "IncomingRecords",
  ]
}
