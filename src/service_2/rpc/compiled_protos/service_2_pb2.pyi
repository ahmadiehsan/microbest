"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""

import builtins
import google.protobuf.descriptor
import google.protobuf.message
import typing

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing.final
class EchoRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    MESSAGE_FIELD_NUMBER: builtins.int
    message: builtins.str
    def __init__(
        self,
        *,
        message: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing.Literal["message", b"message"]) -> None: ...

global___EchoRequest = EchoRequest

@typing.final
class EchoResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    MESSAGE_FIELD_NUMBER: builtins.int
    message: builtins.str
    def __init__(
        self,
        *,
        message: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing.Literal["message", b"message"]) -> None: ...

global___EchoResponse = EchoResponse
