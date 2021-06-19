# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/service.proto',
  package='main',
  syntax='proto3',
  serialized_options=b'Z\t./;mygrpc',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x13proto/service.proto\x12\x04main\"D\n\tResStatus\x12\x12\n\nstatusCode\x18\x01 \x01(\x05\x12\x0f\n\x07\x65rrCode\x18\x02 \x01(\t\x12\x12\n\nerrMessage\x18\x03 \x01(\t\"+\n\x08ResEmpty\x12\x1f\n\x06status\x18\x01 \x01(\x0b\x32\x0f.main.ResStatus\"\n\n\x08ReqEmpty\"C\n\x0bReqMLResult\x12\x13\n\x0b\x43ontainerID\x18\x01 \x01(\t\x12\x10\n\x08ImageURL\x18\x02 \x01(\t\x12\r\n\x05Score\x18\x03 \x01(\x02\"v\n\x0bResMLResult\x12\x1f\n\x06status\x18\x01 \x01(\x0b\x32\x0f.main.ResStatus\x12\x10\n\x08\x43\x61\x63hedID\x18\x05 \x01(\x05\x12\x13\n\x0b\x43ontainerID\x18\x02 \x01(\t\x12\x10\n\x08ImageURL\x18\x03 \x01(\t\x12\r\n\x05Score\x18\x04 \x01(\x02\"P\n\x15ReqConfirmContainerID\x12\x15\n\x08\x43\x61\x63hedID\x18\x01 \x01(\x05H\x00\x88\x01\x01\x12\x13\n\x0b\x43ontainerID\x18\x02 \x01(\tB\x0b\n\t_CachedID\"Y\n\x11\x43ontainerTracking\x12\n\n\x02ID\x18\x01 \x01(\x05\x12\x13\n\x0b\x43ontainerID\x18\x02 \x01(\t\x12\x10\n\x08ImageURL\x18\x03 \x01(\t\x12\x11\n\tCreatedAt\x18\x04 \x01(\x05\"h\n\x19ResListContainerTrackings\x12\x1f\n\x06status\x18\x01 \x01(\x0b\x32\x0f.main.ResStatus\x12*\n\ttrackings\x18\x02 \x03(\x0b\x32\x17.main.ContainerTracking2\xfd\x01\n\x06MyGRPC\x12\x30\n\x0bnewMLResult\x12\x11.main.ReqMLResult\x1a\x0e.main.ResEmpty\x12\x33\n\x0cpullMLResult\x12\x0e.main.ReqEmpty\x1a\x11.main.ResMLResult0\x01\x12I\n\x16listContainerTrackings\x12\x0e.main.ReqEmpty\x1a\x1f.main.ResListContainerTrackings\x12\x41\n\x12\x63onfirmContainerID\x12\x1b.main.ReqConfirmContainerID\x1a\x0e.main.ResEmptyB\x0bZ\t./;mygrpcb\x06proto3'
)




_RESSTATUS = _descriptor.Descriptor(
  name='ResStatus',
  full_name='main.ResStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='statusCode', full_name='main.ResStatus.statusCode', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='errCode', full_name='main.ResStatus.errCode', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='errMessage', full_name='main.ResStatus.errMessage', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=29,
  serialized_end=97,
)


_RESEMPTY = _descriptor.Descriptor(
  name='ResEmpty',
  full_name='main.ResEmpty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='main.ResEmpty.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=99,
  serialized_end=142,
)


_REQEMPTY = _descriptor.Descriptor(
  name='ReqEmpty',
  full_name='main.ReqEmpty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=144,
  serialized_end=154,
)


_REQMLRESULT = _descriptor.Descriptor(
  name='ReqMLResult',
  full_name='main.ReqMLResult',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='ContainerID', full_name='main.ReqMLResult.ContainerID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ImageURL', full_name='main.ReqMLResult.ImageURL', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Score', full_name='main.ReqMLResult.Score', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=156,
  serialized_end=223,
)


_RESMLRESULT = _descriptor.Descriptor(
  name='ResMLResult',
  full_name='main.ResMLResult',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='main.ResMLResult.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CachedID', full_name='main.ResMLResult.CachedID', index=1,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ContainerID', full_name='main.ResMLResult.ContainerID', index=2,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ImageURL', full_name='main.ResMLResult.ImageURL', index=3,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Score', full_name='main.ResMLResult.Score', index=4,
      number=4, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=225,
  serialized_end=343,
)


_REQCONFIRMCONTAINERID = _descriptor.Descriptor(
  name='ReqConfirmContainerID',
  full_name='main.ReqConfirmContainerID',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='CachedID', full_name='main.ReqConfirmContainerID.CachedID', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ContainerID', full_name='main.ReqConfirmContainerID.ContainerID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='_CachedID', full_name='main.ReqConfirmContainerID._CachedID',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=345,
  serialized_end=425,
)


_CONTAINERTRACKING = _descriptor.Descriptor(
  name='ContainerTracking',
  full_name='main.ContainerTracking',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='ID', full_name='main.ContainerTracking.ID', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ContainerID', full_name='main.ContainerTracking.ContainerID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ImageURL', full_name='main.ContainerTracking.ImageURL', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CreatedAt', full_name='main.ContainerTracking.CreatedAt', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=427,
  serialized_end=516,
)


_RESLISTCONTAINERTRACKINGS = _descriptor.Descriptor(
  name='ResListContainerTrackings',
  full_name='main.ResListContainerTrackings',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='main.ResListContainerTrackings.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='trackings', full_name='main.ResListContainerTrackings.trackings', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=518,
  serialized_end=622,
)

_RESEMPTY.fields_by_name['status'].message_type = _RESSTATUS
_RESMLRESULT.fields_by_name['status'].message_type = _RESSTATUS
_REQCONFIRMCONTAINERID.oneofs_by_name['_CachedID'].fields.append(
  _REQCONFIRMCONTAINERID.fields_by_name['CachedID'])
_REQCONFIRMCONTAINERID.fields_by_name['CachedID'].containing_oneof = _REQCONFIRMCONTAINERID.oneofs_by_name['_CachedID']
_RESLISTCONTAINERTRACKINGS.fields_by_name['status'].message_type = _RESSTATUS
_RESLISTCONTAINERTRACKINGS.fields_by_name['trackings'].message_type = _CONTAINERTRACKING
DESCRIPTOR.message_types_by_name['ResStatus'] = _RESSTATUS
DESCRIPTOR.message_types_by_name['ResEmpty'] = _RESEMPTY
DESCRIPTOR.message_types_by_name['ReqEmpty'] = _REQEMPTY
DESCRIPTOR.message_types_by_name['ReqMLResult'] = _REQMLRESULT
DESCRIPTOR.message_types_by_name['ResMLResult'] = _RESMLRESULT
DESCRIPTOR.message_types_by_name['ReqConfirmContainerID'] = _REQCONFIRMCONTAINERID
DESCRIPTOR.message_types_by_name['ContainerTracking'] = _CONTAINERTRACKING
DESCRIPTOR.message_types_by_name['ResListContainerTrackings'] = _RESLISTCONTAINERTRACKINGS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ResStatus = _reflection.GeneratedProtocolMessageType('ResStatus', (_message.Message,), {
  'DESCRIPTOR' : _RESSTATUS,
  '__module__' : 'proto.service_pb2'
  # @@protoc_insertion_point(class_scope:main.ResStatus)
  })
_sym_db.RegisterMessage(ResStatus)

ResEmpty = _reflection.GeneratedProtocolMessageType('ResEmpty', (_message.Message,), {
  'DESCRIPTOR' : _RESEMPTY,
  '__module__' : 'proto.service_pb2'
  # @@protoc_insertion_point(class_scope:main.ResEmpty)
  })
_sym_db.RegisterMessage(ResEmpty)

ReqEmpty = _reflection.GeneratedProtocolMessageType('ReqEmpty', (_message.Message,), {
  'DESCRIPTOR' : _REQEMPTY,
  '__module__' : 'proto.service_pb2'
  # @@protoc_insertion_point(class_scope:main.ReqEmpty)
  })
_sym_db.RegisterMessage(ReqEmpty)

ReqMLResult = _reflection.GeneratedProtocolMessageType('ReqMLResult', (_message.Message,), {
  'DESCRIPTOR' : _REQMLRESULT,
  '__module__' : 'proto.service_pb2'
  # @@protoc_insertion_point(class_scope:main.ReqMLResult)
  })
_sym_db.RegisterMessage(ReqMLResult)

ResMLResult = _reflection.GeneratedProtocolMessageType('ResMLResult', (_message.Message,), {
  'DESCRIPTOR' : _RESMLRESULT,
  '__module__' : 'proto.service_pb2'
  # @@protoc_insertion_point(class_scope:main.ResMLResult)
  })
_sym_db.RegisterMessage(ResMLResult)

ReqConfirmContainerID = _reflection.GeneratedProtocolMessageType('ReqConfirmContainerID', (_message.Message,), {
  'DESCRIPTOR' : _REQCONFIRMCONTAINERID,
  '__module__' : 'proto.service_pb2'
  # @@protoc_insertion_point(class_scope:main.ReqConfirmContainerID)
  })
_sym_db.RegisterMessage(ReqConfirmContainerID)

ContainerTracking = _reflection.GeneratedProtocolMessageType('ContainerTracking', (_message.Message,), {
  'DESCRIPTOR' : _CONTAINERTRACKING,
  '__module__' : 'proto.service_pb2'
  # @@protoc_insertion_point(class_scope:main.ContainerTracking)
  })
_sym_db.RegisterMessage(ContainerTracking)

ResListContainerTrackings = _reflection.GeneratedProtocolMessageType('ResListContainerTrackings', (_message.Message,), {
  'DESCRIPTOR' : _RESLISTCONTAINERTRACKINGS,
  '__module__' : 'proto.service_pb2'
  # @@protoc_insertion_point(class_scope:main.ResListContainerTrackings)
  })
_sym_db.RegisterMessage(ResListContainerTrackings)


DESCRIPTOR._options = None

_MYGRPC = _descriptor.ServiceDescriptor(
  name='MyGRPC',
  full_name='main.MyGRPC',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=625,
  serialized_end=878,
  methods=[
  _descriptor.MethodDescriptor(
    name='newMLResult',
    full_name='main.MyGRPC.newMLResult',
    index=0,
    containing_service=None,
    input_type=_REQMLRESULT,
    output_type=_RESEMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='pullMLResult',
    full_name='main.MyGRPC.pullMLResult',
    index=1,
    containing_service=None,
    input_type=_REQEMPTY,
    output_type=_RESMLRESULT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='listContainerTrackings',
    full_name='main.MyGRPC.listContainerTrackings',
    index=2,
    containing_service=None,
    input_type=_REQEMPTY,
    output_type=_RESLISTCONTAINERTRACKINGS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='confirmContainerID',
    full_name='main.MyGRPC.confirmContainerID',
    index=3,
    containing_service=None,
    input_type=_REQCONFIRMCONTAINERID,
    output_type=_RESEMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_MYGRPC)

DESCRIPTOR.services_by_name['MyGRPC'] = _MYGRPC

# @@protoc_insertion_point(module_scope)
