/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

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

export type CertificateStatus = {
  status?: CertificateStatusEnumCertStatus
  certValid?: boolean
  certExpiry?: string
}

export type DomainSummary = {
  domain?: DomainItem
  reachable?: boolean
  resolvable?: boolean
  certStatus?: CertificateStatus
  whoIsMutated?: boolean
  whoIsMutatedAt?: string
}