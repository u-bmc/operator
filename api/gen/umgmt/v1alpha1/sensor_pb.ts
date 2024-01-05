// @generated by protoc-gen-es v1.6.0 with parameter "target=ts"
// @generated from file umgmt/v1alpha1/sensor.proto (package umgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum umgmt.v1alpha1.SensorType
 */
export enum SensorType {
  /**
   * @generated from enum value: SENSOR_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SENSOR_TYPE_TEMPERATURE = 1;
   */
  TEMPERATURE = 1,

  /**
   * @generated from enum value: SENSOR_TYPE_FAN = 2;
   */
  FAN = 2,

  /**
   * @generated from enum value: SENSOR_TYPE_VOLTAGE = 3;
   */
  VOLTAGE = 3,

  /**
   * @generated from enum value: SENSOR_TYPE_CURRENT = 4;
   */
  CURRENT = 4,

  /**
   * @generated from enum value: SENSOR_TYPE_POWER = 5;
   */
  POWER = 5,

  /**
   * @generated from enum value: SENSOR_TYPE_ENERGY = 6;
   */
  ENERGY = 6,

  /**
   * @generated from enum value: SENSOR_TYPE_HUMIDITY = 7;
   */
  HUMIDITY = 7,

  /**
   * @generated from enum value: SENSOR_TYPE_AIRFLOW = 8;
   */
  AIRFLOW = 8,

  /**
   * @generated from enum value: SENSOR_TYPE_PERCENTAGE = 9;
   */
  PERCENTAGE = 9,

  /**
   * @generated from enum value: SENSOR_TYPE_CAPACITY = 10;
   */
  CAPACITY = 10,

  /**
   * @generated from enum value: SENSOR_TYPE_FREQUENCY = 11;
   */
  FREQUENCY = 11,

  /**
   * @generated from enum value: SENSOR_TYPE_BINARY = 12;
   */
  BINARY = 12,
}
// Retrieve enum metadata with: proto3.getEnumType(SensorType)
proto3.util.setEnumType(SensorType, "umgmt.v1alpha1.SensorType", [
  { no: 0, name: "SENSOR_TYPE_UNSPECIFIED" },
  { no: 1, name: "SENSOR_TYPE_TEMPERATURE" },
  { no: 2, name: "SENSOR_TYPE_FAN" },
  { no: 3, name: "SENSOR_TYPE_VOLTAGE" },
  { no: 4, name: "SENSOR_TYPE_CURRENT" },
  { no: 5, name: "SENSOR_TYPE_POWER" },
  { no: 6, name: "SENSOR_TYPE_ENERGY" },
  { no: 7, name: "SENSOR_TYPE_HUMIDITY" },
  { no: 8, name: "SENSOR_TYPE_AIRFLOW" },
  { no: 9, name: "SENSOR_TYPE_PERCENTAGE" },
  { no: 10, name: "SENSOR_TYPE_CAPACITY" },
  { no: 11, name: "SENSOR_TYPE_FREQUENCY" },
  { no: 12, name: "SENSOR_TYPE_BINARY" },
]);

/**
 * @generated from enum umgmt.v1alpha1.MeasurementUnit
 */
export enum MeasurementUnit {
  /**
   * @generated from enum value: MEASUREMENT_UNIT_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_CELSIUS = 1;
   */
  CELSIUS = 1,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_RPM = 2;
   */
  RPM = 2,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_VOLTS = 3;
   */
  VOLTS = 3,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_AMPS = 4;
   */
  AMPS = 4,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_WATTS = 5;
   */
  WATTS = 5,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_JOULES = 6;
   */
  JOULES = 6,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_PERCENT = 7;
   */
  PERCENT = 7,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_CMM = 8;
   */
  CMM = 8,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_HZ = 9;
   */
  HZ = 9,

  /**
   * @generated from enum value: MEASUREMENT_UNIT_RH = 10;
   */
  RH = 10,
}
// Retrieve enum metadata with: proto3.getEnumType(MeasurementUnit)
proto3.util.setEnumType(MeasurementUnit, "umgmt.v1alpha1.MeasurementUnit", [
  { no: 0, name: "MEASUREMENT_UNIT_UNSPECIFIED" },
  { no: 1, name: "MEASUREMENT_UNIT_CELSIUS" },
  { no: 2, name: "MEASUREMENT_UNIT_RPM" },
  { no: 3, name: "MEASUREMENT_UNIT_VOLTS" },
  { no: 4, name: "MEASUREMENT_UNIT_AMPS" },
  { no: 5, name: "MEASUREMENT_UNIT_WATTS" },
  { no: 6, name: "MEASUREMENT_UNIT_JOULES" },
  { no: 7, name: "MEASUREMENT_UNIT_PERCENT" },
  { no: 8, name: "MEASUREMENT_UNIT_CMM" },
  { no: 9, name: "MEASUREMENT_UNIT_HZ" },
  { no: 10, name: "MEASUREMENT_UNIT_RH" },
]);

/**
 * @generated from enum umgmt.v1alpha1.FanProfile
 */
export enum FanProfile {
  /**
   * @generated from enum value: FAN_PROFILE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: FAN_PROFILE_SILENT = 1;
   */
  SILENT = 1,

  /**
   * @generated from enum value: FAN_PROFILE_BALANCED = 2;
   */
  BALANCED = 2,

  /**
   * @generated from enum value: FAN_PROFILE_PERFORMANCE = 3;
   */
  PERFORMANCE = 3,

  /**
   * @generated from enum value: FAN_PROFILE_FULL_SPEED = 4;
   */
  FULL_SPEED = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(FanProfile)
proto3.util.setEnumType(FanProfile, "umgmt.v1alpha1.FanProfile", [
  { no: 0, name: "FAN_PROFILE_UNSPECIFIED" },
  { no: 1, name: "FAN_PROFILE_SILENT" },
  { no: 2, name: "FAN_PROFILE_BALANCED" },
  { no: 3, name: "FAN_PROFILE_PERFORMANCE" },
  { no: 4, name: "FAN_PROFILE_FULL_SPEED" },
]);

/**
 * @generated from message umgmt.v1alpha1.Threshold
 */
export class Threshold extends Message<Threshold> {
  /**
   * @generated from field: int32 lower_critical = 1;
   */
  lowerCritical = 0;

  /**
   * @generated from field: int32 lower_non_critical = 2;
   */
  lowerNonCritical = 0;

  /**
   * @generated from field: int32 upper_non_critical = 3;
   */
  upperNonCritical = 0;

  /**
   * @generated from field: int32 upper_critical = 4;
   */
  upperCritical = 0;

  constructor(data?: PartialMessage<Threshold>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.Threshold";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "lower_critical", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "lower_non_critical", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "upper_non_critical", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "upper_critical", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Threshold {
    return new Threshold().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Threshold {
    return new Threshold().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Threshold {
    return new Threshold().fromJsonString(jsonString, options);
  }

  static equals(a: Threshold | PlainMessage<Threshold> | undefined, b: Threshold | PlainMessage<Threshold> | undefined): boolean {
    return proto3.util.equals(Threshold, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.Sensor
 */
export class Sensor extends Message<Sensor> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: umgmt.v1alpha1.SensorType type = 2;
   */
  type = SensorType.UNSPECIFIED;

  /**
   * @generated from field: umgmt.v1alpha1.MeasurementUnit unit = 3;
   */
  unit = MeasurementUnit.UNSPECIFIED;

  /**
   * @generated from field: umgmt.v1alpha1.Threshold threshold = 4;
   */
  threshold?: Threshold;

  /**
   * @generated from field: int32 value = 5;
   */
  value = 0;

  /**
   * @generated from field: string description = 6;
   */
  description = "";

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 7;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Sensor>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.Sensor";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(SensorType) },
    { no: 3, name: "unit", kind: "enum", T: proto3.getEnumType(MeasurementUnit) },
    { no: 4, name: "threshold", kind: "message", T: Threshold },
    { no: 5, name: "value", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Sensor {
    return new Sensor().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Sensor {
    return new Sensor().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Sensor {
    return new Sensor().fromJsonString(jsonString, options);
  }

  static equals(a: Sensor | PlainMessage<Sensor> | undefined, b: Sensor | PlainMessage<Sensor> | undefined): boolean {
    return proto3.util.equals(Sensor, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.ChangeFanProfilesRequest
 */
export class ChangeFanProfilesRequest extends Message<ChangeFanProfilesRequest> {
  /**
   * @generated from field: umgmt.v1alpha1.FanProfile fan_profile = 1;
   */
  fanProfile = FanProfile.UNSPECIFIED;

  constructor(data?: PartialMessage<ChangeFanProfilesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.ChangeFanProfilesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "fan_profile", kind: "enum", T: proto3.getEnumType(FanProfile) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ChangeFanProfilesRequest {
    return new ChangeFanProfilesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ChangeFanProfilesRequest {
    return new ChangeFanProfilesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ChangeFanProfilesRequest {
    return new ChangeFanProfilesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ChangeFanProfilesRequest | PlainMessage<ChangeFanProfilesRequest> | undefined, b: ChangeFanProfilesRequest | PlainMessage<ChangeFanProfilesRequest> | undefined): boolean {
    return proto3.util.equals(ChangeFanProfilesRequest, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.ChangeFanProfilesResponse
 */
export class ChangeFanProfilesResponse extends Message<ChangeFanProfilesResponse> {
  constructor(data?: PartialMessage<ChangeFanProfilesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.ChangeFanProfilesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ChangeFanProfilesResponse {
    return new ChangeFanProfilesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ChangeFanProfilesResponse {
    return new ChangeFanProfilesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ChangeFanProfilesResponse {
    return new ChangeFanProfilesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ChangeFanProfilesResponse | PlainMessage<ChangeFanProfilesResponse> | undefined, b: ChangeFanProfilesResponse | PlainMessage<ChangeFanProfilesResponse> | undefined): boolean {
    return proto3.util.equals(ChangeFanProfilesResponse, a, b);
  }
}

