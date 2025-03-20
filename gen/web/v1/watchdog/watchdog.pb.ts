/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"

export enum CertificateStatusEnumCertStatus {
  EnumCertStatusIgnore = "EnumCertStatusIgnore",
  Valid = "Valid",
  Expiring = "Expiring",
  WrongCertificate = "WrongCertificate",
}

export type DomainItem = {
  uuid?: string
  name?: string
  endpoints?: string[]
}

export type DomainWatch = {
  name?: string
  refreshInterval?: string
  domains?: DomainItem[]
}

export type EndpointStatus = {
  endpoint?: string
  statusCode?: string
}

export type CertificateStatus = {
  status?: CertificateStatusEnumCertStatus
  certValid?: boolean
  certExpiry?: string
}

export type DomainSummary = {
  domain?: DomainItem
  reachable?: boolean
  resolvable?: boolean
  certsStatus?: CertificateStatus[]
  whoIsMutated?: boolean
  whoIsMutatedAt?: string
  createdAt?: GoogleProtobufTimestamp.Timestamp
  httpsRedirect?: boolean
}

export type DomainInfo = {
  ipAddresses?: string[]
  certificates?: Uint8Array[]
  endpointStatuses?: EndpointStatus[]
  whois?: string
  createdAt?: GoogleProtobufTimestamp.Timestamp
}

export type DomainRow = {
  domain?: DomainItem
  info?: DomainInfo
  summary?: DomainSummary
  createdAt?: GoogleProtobufTimestamp.Timestamp
}