#!/usr/bin/env ruby

require 'pathname'
require 'aws-sdk-v1'

s3 = AWS::S3.new(
  :region => ENV['AWS_REGION'] || "me-south-1",
  :access_key_id => ENV['AWS_ACCESS_KEY_ID'],
  :secret_access_key => ENV['AWS_SECRET_ACCESS_KEY']
)

bucket = s3.buckets['test-aws-call']
if !bucket.exists?
  s3.buckets.create('test-aws-call')
end

Pathname.new("./files").each_child do |file|
  p "uploading #{file.to_s}"
  bucket.objects[file.basename].write(file)
end
