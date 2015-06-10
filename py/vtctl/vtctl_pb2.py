# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vtctl.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='vtctl.proto',
  package='vtctl',
  serialized_pb=_b('\n\x0bvtctl.proto\x12\x05vtctl\",\n\x04Time\x12\x0f\n\x07seconds\x18\x01 \x01(\x03\x12\x13\n\x0bnanoseconds\x18\x02 \x01(\x03\"b\n\x0bLoggerEvent\x12\x19\n\x04time\x18\x01 \x01(\x0b\x32\x0b.vtctl.Time\x12\r\n\x05level\x18\x02 \x01(\x03\x12\x0c\n\x04\x66ile\x18\x03 \x01(\t\x12\x0c\n\x04line\x18\x04 \x01(\x03\x12\r\n\x05value\x18\x05 \x01(\t\"X\n\x1a\x45xecuteVtctlCommandRequest\x12\x0c\n\x04\x61rgs\x18\x01 \x03(\t\x12\x16\n\x0e\x61\x63tion_timeout\x18\x02 \x01(\x03\x12\x14\n\x0clock_timeout\x18\x03 \x01(\x03\"@\n\x1b\x45xecuteVtctlCommandResponse\x12!\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x12.vtctl.LoggerEventb\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_TIME = _descriptor.Descriptor(
  name='Time',
  full_name='vtctl.Time',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='seconds', full_name='vtctl.Time.seconds', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='nanoseconds', full_name='vtctl.Time.nanoseconds', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=22,
  serialized_end=66,
)


_LOGGEREVENT = _descriptor.Descriptor(
  name='LoggerEvent',
  full_name='vtctl.LoggerEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='vtctl.LoggerEvent.time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='level', full_name='vtctl.LoggerEvent.level', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='file', full_name='vtctl.LoggerEvent.file', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='line', full_name='vtctl.LoggerEvent.line', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='vtctl.LoggerEvent.value', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=68,
  serialized_end=166,
)


_EXECUTEVTCTLCOMMANDREQUEST = _descriptor.Descriptor(
  name='ExecuteVtctlCommandRequest',
  full_name='vtctl.ExecuteVtctlCommandRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='args', full_name='vtctl.ExecuteVtctlCommandRequest.args', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='action_timeout', full_name='vtctl.ExecuteVtctlCommandRequest.action_timeout', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='lock_timeout', full_name='vtctl.ExecuteVtctlCommandRequest.lock_timeout', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=168,
  serialized_end=256,
)


_EXECUTEVTCTLCOMMANDRESPONSE = _descriptor.Descriptor(
  name='ExecuteVtctlCommandResponse',
  full_name='vtctl.ExecuteVtctlCommandResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='event', full_name='vtctl.ExecuteVtctlCommandResponse.event', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=258,
  serialized_end=322,
)

_LOGGEREVENT.fields_by_name['time'].message_type = _TIME
_EXECUTEVTCTLCOMMANDRESPONSE.fields_by_name['event'].message_type = _LOGGEREVENT
DESCRIPTOR.message_types_by_name['Time'] = _TIME
DESCRIPTOR.message_types_by_name['LoggerEvent'] = _LOGGEREVENT
DESCRIPTOR.message_types_by_name['ExecuteVtctlCommandRequest'] = _EXECUTEVTCTLCOMMANDREQUEST
DESCRIPTOR.message_types_by_name['ExecuteVtctlCommandResponse'] = _EXECUTEVTCTLCOMMANDRESPONSE

Time = _reflection.GeneratedProtocolMessageType('Time', (_message.Message,), dict(
  DESCRIPTOR = _TIME,
  __module__ = 'vtctl_pb2'
  # @@protoc_insertion_point(class_scope:vtctl.Time)
  ))
_sym_db.RegisterMessage(Time)

LoggerEvent = _reflection.GeneratedProtocolMessageType('LoggerEvent', (_message.Message,), dict(
  DESCRIPTOR = _LOGGEREVENT,
  __module__ = 'vtctl_pb2'
  # @@protoc_insertion_point(class_scope:vtctl.LoggerEvent)
  ))
_sym_db.RegisterMessage(LoggerEvent)

ExecuteVtctlCommandRequest = _reflection.GeneratedProtocolMessageType('ExecuteVtctlCommandRequest', (_message.Message,), dict(
  DESCRIPTOR = _EXECUTEVTCTLCOMMANDREQUEST,
  __module__ = 'vtctl_pb2'
  # @@protoc_insertion_point(class_scope:vtctl.ExecuteVtctlCommandRequest)
  ))
_sym_db.RegisterMessage(ExecuteVtctlCommandRequest)

ExecuteVtctlCommandResponse = _reflection.GeneratedProtocolMessageType('ExecuteVtctlCommandResponse', (_message.Message,), dict(
  DESCRIPTOR = _EXECUTEVTCTLCOMMANDRESPONSE,
  __module__ = 'vtctl_pb2'
  # @@protoc_insertion_point(class_scope:vtctl.ExecuteVtctlCommandResponse)
  ))
_sym_db.RegisterMessage(ExecuteVtctlCommandResponse)


import abc
from grpc.early_adopter import implementations
from grpc.framework.alpha import utilities
# @@protoc_insertion_point(module_scope)
