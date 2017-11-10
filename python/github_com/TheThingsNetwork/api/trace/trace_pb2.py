# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: github_com/TheThingsNetwork/api/trace/trace.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github_com.gogo.protobuf.gogoproto import gogo_pb2 as github__com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='github_com/TheThingsNetwork/api/trace/trace.proto',
  package='trace',
  syntax='proto3',
  serialized_pb=_b('\n1github_com/TheThingsNetwork/api/trace/trace.proto\x12\x05trace\x1a-github_com/gogo/protobuf/gogoproto/gogo.proto\"\xef\x01\n\x05Trace\x12\x12\n\x02id\x18\x01 \x01(\tB\x06\xe2\xde\x1f\x02ID\x12\x0c\n\x04time\x18\x02 \x01(\x03\x12!\n\nservice_id\x18\x03 \x01(\tB\r\xe2\xde\x1f\tServiceID\x12\x14\n\x0cservice_name\x18\x04 \x01(\t\x12\r\n\x05\x65vent\x18\x05 \x01(\t\x12,\n\x08metadata\x18\x06 \x03(\x0b\x32\x1a.trace.Trace.MetadataEntry\x12\x1d\n\x07parents\x18\x0b \x03(\x0b\x32\x0c.trace.Trace\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42r\n\x1eorg.thethingsnetwork.api.traceB\nTraceProtoP\x01Z%github_com/TheThingsNetwork/api/trace\xaa\x02\x1aTheThingsNetwork.API.Traceb\x06proto3')
  ,
  dependencies=[github__com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2.DESCRIPTOR,])




_TRACE_METADATAENTRY = _descriptor.Descriptor(
  name='MetadataEntry',
  full_name='trace.Trace.MetadataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='trace.Trace.MetadataEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='trace.Trace.MetadataEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=300,
  serialized_end=347,
)

_TRACE = _descriptor.Descriptor(
  name='Trace',
  full_name='trace.Trace',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='trace.Trace.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\002ID'))),
    _descriptor.FieldDescriptor(
      name='time', full_name='trace.Trace.time', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='service_id', full_name='trace.Trace.service_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\tServiceID'))),
    _descriptor.FieldDescriptor(
      name='service_name', full_name='trace.Trace.service_name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='event', full_name='trace.Trace.event', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='trace.Trace.metadata', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='parents', full_name='trace.Trace.parents', index=6,
      number=11, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_TRACE_METADATAENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=108,
  serialized_end=347,
)

_TRACE_METADATAENTRY.containing_type = _TRACE
_TRACE.fields_by_name['metadata'].message_type = _TRACE_METADATAENTRY
_TRACE.fields_by_name['parents'].message_type = _TRACE
DESCRIPTOR.message_types_by_name['Trace'] = _TRACE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Trace = _reflection.GeneratedProtocolMessageType('Trace', (_message.Message,), dict(

  MetadataEntry = _reflection.GeneratedProtocolMessageType('MetadataEntry', (_message.Message,), dict(
    DESCRIPTOR = _TRACE_METADATAENTRY,
    __module__ = 'github_com.TheThingsNetwork.api.trace.trace_pb2'
    # @@protoc_insertion_point(class_scope:trace.Trace.MetadataEntry)
    ))
  ,
  DESCRIPTOR = _TRACE,
  __module__ = 'github_com.TheThingsNetwork.api.trace.trace_pb2'
  # @@protoc_insertion_point(class_scope:trace.Trace)
  ))
_sym_db.RegisterMessage(Trace)
_sym_db.RegisterMessage(Trace.MetadataEntry)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\036org.thethingsnetwork.api.traceB\nTraceProtoP\001Z%github_com/TheThingsNetwork/api/trace\252\002\032TheThingsNetwork.API.Trace'))
_TRACE_METADATAENTRY.has_options = True
_TRACE_METADATAENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_TRACE.fields_by_name['id'].has_options = True
_TRACE.fields_by_name['id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\002ID'))
_TRACE.fields_by_name['service_id'].has_options = True
_TRACE.fields_by_name['service_id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\tServiceID'))
try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)