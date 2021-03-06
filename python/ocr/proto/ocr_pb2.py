# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/ocr.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/ocr.proto',
  package='ocr_main',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0fproto/ocr.proto\x12\x08ocr_main\"\x91\x01\n\x13ReqNewDetectedImage\x12\x10\n\x08ImageURL\x18\x01 \x01(\t\x12\r\n\x05\x43\x61mID\x18\x02 \x01(\t\x12\x0f\n\x07\x43\x61mName\x18\x07 \x01(\t\x12\x13\n\x0b\x43\x61ptureTime\x18\x03 \x01(\t\x12\x10\n\x08\x44\x61taType\x18\x05 \x01(\t\x12\x12\n\nTrackingID\x18\x06 \x01(\t\x12\r\n\x05Score\x18\x04 \x01(\t\"\n\n\x08ResEmpty2N\n\x03OCR\x12G\n\x10newDetectedImage\x12\x1d.ocr_main.ReqNewDetectedImage\x1a\x12.ocr_main.ResEmpty(\x01\x62\x06proto3'
)




_REQNEWDETECTEDIMAGE = _descriptor.Descriptor(
  name='ReqNewDetectedImage',
  full_name='ocr_main.ReqNewDetectedImage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='ImageURL', full_name='ocr_main.ReqNewDetectedImage.ImageURL', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CamID', full_name='ocr_main.ReqNewDetectedImage.CamID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CamName', full_name='ocr_main.ReqNewDetectedImage.CamName', index=2,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CaptureTime', full_name='ocr_main.ReqNewDetectedImage.CaptureTime', index=3,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='DataType', full_name='ocr_main.ReqNewDetectedImage.DataType', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='TrackingID', full_name='ocr_main.ReqNewDetectedImage.TrackingID', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Score', full_name='ocr_main.ReqNewDetectedImage.Score', index=6,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=30,
  serialized_end=175,
)


_RESEMPTY = _descriptor.Descriptor(
  name='ResEmpty',
  full_name='ocr_main.ResEmpty',
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
  serialized_start=177,
  serialized_end=187,
)

DESCRIPTOR.message_types_by_name['ReqNewDetectedImage'] = _REQNEWDETECTEDIMAGE
DESCRIPTOR.message_types_by_name['ResEmpty'] = _RESEMPTY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ReqNewDetectedImage = _reflection.GeneratedProtocolMessageType('ReqNewDetectedImage', (_message.Message,), {
  'DESCRIPTOR' : _REQNEWDETECTEDIMAGE,
  '__module__' : 'proto.ocr_pb2'
  # @@protoc_insertion_point(class_scope:ocr_main.ReqNewDetectedImage)
  })
_sym_db.RegisterMessage(ReqNewDetectedImage)

ResEmpty = _reflection.GeneratedProtocolMessageType('ResEmpty', (_message.Message,), {
  'DESCRIPTOR' : _RESEMPTY,
  '__module__' : 'proto.ocr_pb2'
  # @@protoc_insertion_point(class_scope:ocr_main.ResEmpty)
  })
_sym_db.RegisterMessage(ResEmpty)



_OCR = _descriptor.ServiceDescriptor(
  name='OCR',
  full_name='ocr_main.OCR',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=189,
  serialized_end=267,
  methods=[
  _descriptor.MethodDescriptor(
    name='newDetectedImage',
    full_name='ocr_main.OCR.newDetectedImage',
    index=0,
    containing_service=None,
    input_type=_REQNEWDETECTEDIMAGE,
    output_type=_RESEMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_OCR)

DESCRIPTOR.services_by_name['OCR'] = _OCR

# @@protoc_insertion_point(module_scope)
