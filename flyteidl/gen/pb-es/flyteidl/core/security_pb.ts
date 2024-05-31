// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file flyteidl/core/security.proto (package flyteidl.core, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * Secret encapsulates information about the secret a task needs to proceed. An environment variable
 * FLYTE_SECRETS_ENV_PREFIX will be passed to indicate the prefix of the environment variables that will be present if
 * secrets are passed through environment variables.
 * FLYTE_SECRETS_DEFAULT_DIR will be passed to indicate the prefix of the path where secrets will be mounted if secrets
 * are passed through file mounts.
 *
 * @generated from message flyteidl.core.Secret
 */
export class Secret extends Message<Secret> {
  /**
   * The name of the secret group where to find the key referenced below. For K8s secrets, this should be the name of
   * the v1/secret object. For Confidant, this should be the Credential name. For Vault, this should be the secret name.
   * For AWS Secret Manager, this should be the name of the secret.
   * +required
   *
   * @generated from field: string group = 1;
   */
  group = "";

  /**
   * The group version to fetch. This is not supported in all secret management systems. It'll be ignored for the ones
   * that do not support it.
   * +optional
   *
   * @generated from field: string group_version = 2;
   */
  groupVersion = "";

  /**
   * The name of the secret to mount. This has to match an existing secret in the system. It's up to the implementation
   * of the secret management system to require case sensitivity. For K8s secrets, Confidant and Vault, this should
   * match one of the keys inside the secret. For AWS Secret Manager, it's ignored.
   * +optional
   *
   * @generated from field: string key = 3;
   */
  key = "";

  /**
   * mount_requirement is optional. Indicates where the secret has to be mounted. If provided, the execution will fail
   * if the underlying key management system cannot satisfy that requirement. If not provided, the default location
   * will depend on the key management system.
   * +optional
   *
   * @generated from field: flyteidl.core.Secret.MountType mount_requirement = 4;
   */
  mountRequirement = Secret_MountType.ANY;

  constructor(data?: PartialMessage<Secret>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Secret";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "group_version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "mount_requirement", kind: "enum", T: proto3.getEnumType(Secret_MountType) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Secret {
    return new Secret().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Secret {
    return new Secret().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Secret {
    return new Secret().fromJsonString(jsonString, options);
  }

  static equals(a: Secret | PlainMessage<Secret> | undefined, b: Secret | PlainMessage<Secret> | undefined): boolean {
    return proto3.util.equals(Secret, a, b);
  }
}

/**
 * @generated from enum flyteidl.core.Secret.MountType
 */
export enum Secret_MountType {
  /**
   * Default case, indicates the client can tolerate either mounting options.
   *
   * @generated from enum value: ANY = 0;
   */
  ANY = 0,

  /**
   * ENV_VAR indicates the secret needs to be mounted as an environment variable.
   *
   * @generated from enum value: ENV_VAR = 1;
   */
  ENV_VAR = 1,

  /**
   * FILE indicates the secret needs to be mounted as a file.
   *
   * @generated from enum value: FILE = 2;
   */
  FILE = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Secret_MountType)
proto3.util.setEnumType(Secret_MountType, "flyteidl.core.Secret.MountType", [
  { no: 0, name: "ANY" },
  { no: 1, name: "ENV_VAR" },
  { no: 2, name: "FILE" },
]);

/**
 * @generated from message flyteidl.core.Connection
 */
export class Connection extends Message<Connection> {
  /**
   * The task type that the connection is used for.
   *
   * @generated from field: string task_type = 1;
   */
  taskType = "";

  /**
   * The credentials to use for the connection, such as API keys, OAuth2 tokens, etc.
   * The key is the name of the secret, and it's defined in the flytekit.
   * flytekit uses the key to locate the desired secret within the map.
   *
   * @generated from field: map<string, string> secrets = 2;
   */
  secrets: { [key: string]: string } = {};

  /**
   * The configuration to use for the connection, such as the endpoint, account name, etc.
   * The key is the name of the config, and it's defined in the flytekit.
   *
   * @generated from field: map<string, string> configs = 3;
   */
  configs: { [key: string]: string } = {};

  constructor(data?: PartialMessage<Connection>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Connection";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "secrets", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 3, name: "configs", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Connection {
    return new Connection().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Connection {
    return new Connection().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Connection {
    return new Connection().fromJsonString(jsonString, options);
  }

  static equals(a: Connection | PlainMessage<Connection> | undefined, b: Connection | PlainMessage<Connection> | undefined): boolean {
    return proto3.util.equals(Connection, a, b);
  }
}

/**
 * OAuth2Client encapsulates OAuth2 Client Credentials to be used when making calls on behalf of that task.
 *
 * @generated from message flyteidl.core.OAuth2Client
 */
export class OAuth2Client extends Message<OAuth2Client> {
  /**
   * client_id is the public id for the client to use. The system will not perform any pre-auth validation that the
   * secret requested matches the client_id indicated here.
   * +required
   *
   * @generated from field: string client_id = 1;
   */
  clientId = "";

  /**
   * client_secret is a reference to the secret used to authenticate the OAuth2 client.
   * +required
   *
   * @generated from field: flyteidl.core.Secret client_secret = 2;
   */
  clientSecret?: Secret;

  constructor(data?: PartialMessage<OAuth2Client>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.OAuth2Client";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "client_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "client_secret", kind: "message", T: Secret },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OAuth2Client {
    return new OAuth2Client().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OAuth2Client {
    return new OAuth2Client().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OAuth2Client {
    return new OAuth2Client().fromJsonString(jsonString, options);
  }

  static equals(a: OAuth2Client | PlainMessage<OAuth2Client> | undefined, b: OAuth2Client | PlainMessage<OAuth2Client> | undefined): boolean {
    return proto3.util.equals(OAuth2Client, a, b);
  }
}

/**
 * Identity encapsulates the various security identities a task can run as. It's up to the underlying plugin to pick the
 * right identity for the execution environment.
 *
 * @generated from message flyteidl.core.Identity
 */
export class Identity extends Message<Identity> {
  /**
   * iam_role references the fully qualified name of Identity & Access Management role to impersonate.
   *
   * @generated from field: string iam_role = 1;
   */
  iamRole = "";

  /**
   * k8s_service_account references a kubernetes service account to impersonate.
   *
   * @generated from field: string k8s_service_account = 2;
   */
  k8sServiceAccount = "";

  /**
   * oauth2_client references an oauth2 client. Backend plugins can use this information to impersonate the client when
   * making external calls.
   *
   * @generated from field: flyteidl.core.OAuth2Client oauth2_client = 3;
   */
  oauth2Client?: OAuth2Client;

  /**
   * execution_identity references the subject who makes the execution
   *
   * @generated from field: string execution_identity = 4;
   */
  executionIdentity = "";

  constructor(data?: PartialMessage<Identity>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Identity";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "iam_role", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "k8s_service_account", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "oauth2_client", kind: "message", T: OAuth2Client },
    { no: 4, name: "execution_identity", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Identity {
    return new Identity().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Identity {
    return new Identity().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Identity {
    return new Identity().fromJsonString(jsonString, options);
  }

  static equals(a: Identity | PlainMessage<Identity> | undefined, b: Identity | PlainMessage<Identity> | undefined): boolean {
    return proto3.util.equals(Identity, a, b);
  }
}

/**
 * OAuth2TokenRequest encapsulates information needed to request an OAuth2 token.
 * FLYTE_TOKENS_ENV_PREFIX will be passed to indicate the prefix of the environment variables that will be present if
 * tokens are passed through environment variables.
 * FLYTE_TOKENS_PATH_PREFIX will be passed to indicate the prefix of the path where secrets will be mounted if tokens
 * are passed through file mounts.
 *
 * @generated from message flyteidl.core.OAuth2TokenRequest
 */
export class OAuth2TokenRequest extends Message<OAuth2TokenRequest> {
  /**
   * name indicates a unique id for the token request within this task token requests. It'll be used as a suffix for
   * environment variables and as a filename for mounting tokens as files.
   * +required
   *
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * type indicates the type of the request to make. Defaults to CLIENT_CREDENTIALS.
   * +required
   *
   * @generated from field: flyteidl.core.OAuth2TokenRequest.Type type = 2;
   */
  type = OAuth2TokenRequest_Type.CLIENT_CREDENTIALS;

  /**
   * client references the client_id/secret to use to request the OAuth2 token.
   * +required
   *
   * @generated from field: flyteidl.core.OAuth2Client client = 3;
   */
  client?: OAuth2Client;

  /**
   * idp_discovery_endpoint references the discovery endpoint used to retrieve token endpoint and other related
   * information.
   * +optional
   *
   * @generated from field: string idp_discovery_endpoint = 4;
   */
  idpDiscoveryEndpoint = "";

  /**
   * token_endpoint references the token issuance endpoint. If idp_discovery_endpoint is not provided, this parameter is
   * mandatory.
   * +optional
   *
   * @generated from field: string token_endpoint = 5;
   */
  tokenEndpoint = "";

  constructor(data?: PartialMessage<OAuth2TokenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.OAuth2TokenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(OAuth2TokenRequest_Type) },
    { no: 3, name: "client", kind: "message", T: OAuth2Client },
    { no: 4, name: "idp_discovery_endpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "token_endpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OAuth2TokenRequest {
    return new OAuth2TokenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OAuth2TokenRequest {
    return new OAuth2TokenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OAuth2TokenRequest {
    return new OAuth2TokenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: OAuth2TokenRequest | PlainMessage<OAuth2TokenRequest> | undefined, b: OAuth2TokenRequest | PlainMessage<OAuth2TokenRequest> | undefined): boolean {
    return proto3.util.equals(OAuth2TokenRequest, a, b);
  }
}

/**
 * Type of the token requested.
 *
 * @generated from enum flyteidl.core.OAuth2TokenRequest.Type
 */
export enum OAuth2TokenRequest_Type {
  /**
   * CLIENT_CREDENTIALS indicates a 2-legged OAuth token requested using client credentials.
   *
   * @generated from enum value: CLIENT_CREDENTIALS = 0;
   */
  CLIENT_CREDENTIALS = 0,
}
// Retrieve enum metadata with: proto3.getEnumType(OAuth2TokenRequest_Type)
proto3.util.setEnumType(OAuth2TokenRequest_Type, "flyteidl.core.OAuth2TokenRequest.Type", [
  { no: 0, name: "CLIENT_CREDENTIALS" },
]);

/**
 * SecurityContext holds security attributes that apply to tasks.
 *
 * @generated from message flyteidl.core.SecurityContext
 */
export class SecurityContext extends Message<SecurityContext> {
  /**
   * run_as encapsulates the identity a pod should run as. If the task fills in multiple fields here, it'll be up to the
   * backend plugin to choose the appropriate identity for the execution engine the task will run on.
   *
   * @generated from field: flyteidl.core.Identity run_as = 1;
   */
  runAs?: Identity;

  /**
   * secrets indicate the list of secrets the task needs in order to proceed. Secrets will be mounted/passed to the
   * pod as it starts. If the plugin responsible for kicking of the task will not run it on a flyte cluster (e.g. AWS
   * Batch), it's the responsibility of the plugin to fetch the secret (which means propeller identity will need access
   * to the secret) and to pass it to the remote execution engine.
   *
   * @generated from field: repeated flyteidl.core.Secret secrets = 2;
   */
  secrets: Secret[] = [];

  /**
   * tokens indicate the list of token requests the task needs in order to proceed. Tokens will be mounted/passed to the
   * pod as it starts. If the plugin responsible for kicking of the task will not run it on a flyte cluster (e.g. AWS
   * Batch), it's the responsibility of the plugin to fetch the secret (which means propeller identity will need access
   * to the secret) and to pass it to the remote execution engine.
   *
   * @generated from field: repeated flyteidl.core.OAuth2TokenRequest tokens = 3;
   */
  tokens: OAuth2TokenRequest[] = [];

  /**
   * The name of the connection.
   * The connection is defined in the externalResourceAttributes or flyteadmin configmap.
   * The connection config take precedence in the following order:
   * 1. connection in the externalResourceAttributes in the project-domain settings.
   * 2. connection in the externalResourceAttributes in the project settings.
   * 3. connection in the flyteadmin configmap.
   * +optional
   *
   * @generated from field: string connection_ref = 4;
   */
  connectionRef = "";

  constructor(data?: PartialMessage<SecurityContext>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.SecurityContext";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "run_as", kind: "message", T: Identity },
    { no: 2, name: "secrets", kind: "message", T: Secret, repeated: true },
    { no: 3, name: "tokens", kind: "message", T: OAuth2TokenRequest, repeated: true },
    { no: 4, name: "connection_ref", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SecurityContext {
    return new SecurityContext().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SecurityContext {
    return new SecurityContext().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SecurityContext {
    return new SecurityContext().fromJsonString(jsonString, options);
  }

  static equals(a: SecurityContext | PlainMessage<SecurityContext> | undefined, b: SecurityContext | PlainMessage<SecurityContext> | undefined): boolean {
    return proto3.util.equals(SecurityContext, a, b);
  }
}

