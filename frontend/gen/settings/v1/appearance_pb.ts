// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file settings/v1/appearance.proto (package appearance.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message appearance.v1.AppearanceServiceGetRequest
 */
export class AppearanceServiceGetRequest extends Message<AppearanceServiceGetRequest> {
  constructor(data?: PartialMessage<AppearanceServiceGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "appearance.v1.AppearanceServiceGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppearanceServiceGetRequest {
    return new AppearanceServiceGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppearanceServiceGetRequest {
    return new AppearanceServiceGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppearanceServiceGetRequest {
    return new AppearanceServiceGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AppearanceServiceGetRequest | PlainMessage<AppearanceServiceGetRequest> | undefined, b: AppearanceServiceGetRequest | PlainMessage<AppearanceServiceGetRequest> | undefined): boolean {
    return proto3.util.equals(AppearanceServiceGetRequest, a, b);
  }
}

/**
 * @generated from message appearance.v1.AppearanceServiceGetResponse
 */
export class AppearanceServiceGetResponse extends Message<AppearanceServiceGetResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  constructor(data?: PartialMessage<AppearanceServiceGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "appearance.v1.AppearanceServiceGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppearanceServiceGetResponse {
    return new AppearanceServiceGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppearanceServiceGetResponse {
    return new AppearanceServiceGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppearanceServiceGetResponse {
    return new AppearanceServiceGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: AppearanceServiceGetResponse | PlainMessage<AppearanceServiceGetResponse> | undefined, b: AppearanceServiceGetResponse | PlainMessage<AppearanceServiceGetResponse> | undefined): boolean {
    return proto3.util.equals(AppearanceServiceGetResponse, a, b);
  }
}
