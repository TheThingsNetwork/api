# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: trace/trace.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "trace.Trace" do
    optional :id, :string, 1
    optional :time, :int64, 2
    optional :service_id, :string, 3
    optional :service_name, :string, 4
    optional :event, :string, 5
    map :metadata, :string, :string, 6
    repeated :parents, :message, 11, "trace.Trace"
  end
end

module Trace
  Trace = Google::Protobuf::DescriptorPool.generated_pool.lookup("trace.Trace").msgclass
end
