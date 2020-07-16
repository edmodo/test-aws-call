#!/usr/bin/env ruby

require 'aws-sdk-cognitoidentityprovider'

client = Aws::CognitoIdentityProvider::Client.new(
  :region => "me-south-1",
  :access_key_id => ENV['AWS_ACCESS_KEY_ID'],
  :secret_access_key => ENV['AWS_SECRET_ACCESS_KEY']
)

pools = client.list_user_pools(
  next_token: "PaginationKeyType",
  max_results: 1
)

p pools
