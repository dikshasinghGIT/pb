# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: addsvc.proto

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
  name='addsvc.proto',
  package='addsvc',
  syntax='proto3',
  serialized_pb=_b('\n\x0c\x61\x64\x64svc.proto\x12\x06\x61\x64\x64svc\"\"\n\nSumRequest\x12\t\n\x01\x61\x18\x01 \x01(\x03\x12\t\n\x01\x62\x18\x02 \x01(\x03\"\"\n\x08SumReply\x12\t\n\x01v\x18\x01 \x01(\x03\x12\x0b\n\x03\x65rr\x18\x02 \x01(\t\"%\n\rConcatRequest\x12\t\n\x01\x61\x18\x01 \x01(\t\x12\t\n\x01\x62\x18\x02 \x01(\t\"%\n\x0b\x43oncatReply\x12\t\n\x01v\x18\x01 \x01(\t\x12\x0b\n\x03\x65rr\x18\x02 \x01(\t2l\n\x03\x41\x64\x64\x12-\n\x03Sum\x12\x12.addsvc.SumRequest\x1a\x10.addsvc.SumReply\"\x00\x12\x36\n\x06\x43oncat\x12\x15.addsvc.ConcatRequest\x1a\x13.addsvc.ConcatReply\"\x00\x62\x06proto3')
)




_SUMREQUEST = _descriptor.Descriptor(
  name='SumRequest',
  full_name='addsvc.SumRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='a', full_name='addsvc.SumRequest.a', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='b', full_name='addsvc.SumRequest.b', index=1,
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
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=58,
)


_SUMREPLY = _descriptor.Descriptor(
  name='SumReply',
  full_name='addsvc.SumReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='v', full_name='addsvc.SumReply.v', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='err', full_name='addsvc.SumReply.err', index=1,
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
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=60,
  serialized_end=94,
)


_CONCATREQUEST = _descriptor.Descriptor(
  name='ConcatRequest',
  full_name='addsvc.ConcatRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='a', full_name='addsvc.ConcatRequest.a', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='b', full_name='addsvc.ConcatRequest.b', index=1,
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
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=96,
  serialized_end=133,
)


_CONCATREPLY = _descriptor.Descriptor(
  name='ConcatReply',
  full_name='addsvc.ConcatReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='v', full_name='addsvc.ConcatReply.v', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='err', full_name='addsvc.ConcatReply.err', index=1,
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
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=135,
  serialized_end=172,
)

DESCRIPTOR.message_types_by_name['SumRequest'] = _SUMREQUEST
DESCRIPTOR.message_types_by_name['SumReply'] = _SUMREPLY
DESCRIPTOR.message_types_by_name['ConcatRequest'] = _CONCATREQUEST
DESCRIPTOR.message_types_by_name['ConcatReply'] = _CONCATREPLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SumRequest = _reflection.GeneratedProtocolMessageType('SumRequest', (_message.Message,), dict(
  DESCRIPTOR = _SUMREQUEST,
  __module__ = 'addsvc_pb2'
  # @@protoc_insertion_point(class_scope:addsvc.SumRequest)
  ))
_sym_db.RegisterMessage(SumRequest)

SumReply = _reflection.GeneratedProtocolMessageType('SumReply', (_message.Message,), dict(
  DESCRIPTOR = _SUMREPLY,
  __module__ = 'addsvc_pb2'
  # @@protoc_insertion_point(class_scope:addsvc.SumReply)
  ))
_sym_db.RegisterMessage(SumReply)

ConcatRequest = _reflection.GeneratedProtocolMessageType('ConcatRequest', (_message.Message,), dict(
  DESCRIPTOR = _CONCATREQUEST,
  __module__ = 'addsvc_pb2'
  # @@protoc_insertion_point(class_scope:addsvc.ConcatRequest)
  ))
_sym_db.RegisterMessage(ConcatRequest)

ConcatReply = _reflection.GeneratedProtocolMessageType('ConcatReply', (_message.Message,), dict(
  DESCRIPTOR = _CONCATREPLY,
  __module__ = 'addsvc_pb2'
  # @@protoc_insertion_point(class_scope:addsvc.ConcatReply)
  ))
_sym_db.RegisterMessage(ConcatReply)



_ADD = _descriptor.ServiceDescriptor(
  name='Add',
  full_name='addsvc.Add',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=174,
  serialized_end=282,
  methods=[
  _descriptor.MethodDescriptor(
    name='Sum',
    full_name='addsvc.Add.Sum',
    index=0,
    containing_service=None,
    input_type=_SUMREQUEST,
    output_type=_SUMREPLY,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Concat',
    full_name='addsvc.Add.Concat',
    index=1,
    containing_service=None,
    input_type=_CONCATREQUEST,
    output_type=_CONCATREPLY,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_ADD)

DESCRIPTOR.services_by_name['Add'] = _ADD

try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities


  class AddStub(object):
    """The Add service definition.
    """

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.Sum = channel.unary_unary(
          '/addsvc.Add/Sum',
          request_serializer=SumRequest.SerializeToString,
          response_deserializer=SumReply.FromString,
          )
      self.Concat = channel.unary_unary(
          '/addsvc.Add/Concat',
          request_serializer=ConcatRequest.SerializeToString,
          response_deserializer=ConcatReply.FromString,
          )


  class AddServicer(object):
    """The Add service definition.
    """

    def Sum(self, request, context):
      """Sums two integers.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def Concat(self, request, context):
      """Concatenates two strings
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_AddServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'Sum': grpc.unary_unary_rpc_method_handler(
            servicer.Sum,
            request_deserializer=SumRequest.FromString,
            response_serializer=SumReply.SerializeToString,
        ),
        'Concat': grpc.unary_unary_rpc_method_handler(
            servicer.Concat,
            request_deserializer=ConcatRequest.FromString,
            response_serializer=ConcatReply.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'addsvc.Add', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaAddServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    """The Add service definition.
    """
    def Sum(self, request, context):
      """Sums two integers.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def Concat(self, request, context):
      """Concatenates two strings
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaAddStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    """The Add service definition.
    """
    def Sum(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Sums two integers.
      """
      raise NotImplementedError()
    Sum.future = None
    def Concat(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Concatenates two strings
      """
      raise NotImplementedError()
    Concat.future = None


  def beta_create_Add_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('addsvc.Add', 'Concat'): ConcatRequest.FromString,
      ('addsvc.Add', 'Sum'): SumRequest.FromString,
    }
    response_serializers = {
      ('addsvc.Add', 'Concat'): ConcatReply.SerializeToString,
      ('addsvc.Add', 'Sum'): SumReply.SerializeToString,
    }
    method_implementations = {
      ('addsvc.Add', 'Concat'): face_utilities.unary_unary_inline(servicer.Concat),
      ('addsvc.Add', 'Sum'): face_utilities.unary_unary_inline(servicer.Sum),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_Add_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('addsvc.Add', 'Concat'): ConcatRequest.SerializeToString,
      ('addsvc.Add', 'Sum'): SumRequest.SerializeToString,
    }
    response_deserializers = {
      ('addsvc.Add', 'Concat'): ConcatReply.FromString,
      ('addsvc.Add', 'Sum'): SumReply.FromString,
    }
    cardinalities = {
      'Concat': cardinality.Cardinality.UNARY_UNARY,
      'Sum': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'addsvc.Add', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
