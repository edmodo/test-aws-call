#!/usr/bin/env ruby

require 'fog'
require 'pathname'

Fog.credentials = {
  aws_access_key_id: ENV['AWS_ACCESS_KEY_ID'],
  aws_secret_access_key: ENV['AWS_SECRET_ACCESS_KEY'],
  region: ENV['AWS_REGION'] || "me-south-1"
}

storage = Fog::Storage.new(provider: 'AWS')
fog_directory = storage.directories.new(key: 'test-aws-call')

Pathname.new("./files").each_child do |file|
  p "uploading #{file.to_s}"
  fog_directory.files.create(body: file.binread, key: file.basename)
end

