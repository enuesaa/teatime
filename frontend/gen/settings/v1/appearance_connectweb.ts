// @generated by protoc-gen-connect-web v0.7.0 with parameter "target=ts"
// @generated from file settings/v1/appearance.proto (package appearance.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AppearanceServiceGetRequest, AppearanceServiceGetResponse } from "./appearance_pb.ts";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service appearance.v1.AppearanceService
 */
export const AppearanceService = {
  typeName: "appearance.v1.AppearanceService",
  methods: {
    /**
     * @generated from rpc appearance.v1.AppearanceService.Get
     */
    get: {
      name: "Get",
      I: AppearanceServiceGetRequest,
      O: AppearanceServiceGetResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

