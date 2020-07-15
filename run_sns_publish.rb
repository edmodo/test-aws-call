#!/usr/bin/env ruby

require 'aws-sdk-v1'

sns = AWS::SNS::Client.new(
  :region => ENV['AWS_REGION'] || "me-south-1",
  :access_key_id => ENV['AWS_ACCESS_KEY_ID'],
  :secret_access_key => ENV['AWS_SECRET_ACCESS_KEY']
)

topic_arn = sns.create_topic(name: 'test-aws-call')[:topic_arn]
p sns.publish(topic_arn: topic_arn, message: "random message")
sns.delete_topic(topic_arn: topic_arn)
