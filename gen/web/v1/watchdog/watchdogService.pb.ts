/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as WatchdogWatchdog from "./watchdog.pb"

export enum HealthResponseHealthStatus {
  HealthStatusIgnore = "HealthStatusIgnore",
  Active = "Active",
  Error = "Error",
}

export type ReloadRequest = {
}

export type ReloadResponse = {
}

export type GetRequest = {
  uuid?: string
  name?: string
}

export type GetResponse = {
  summary?: WatchdogWatchdog.DomainSummary
}

export type GetDetailsRequest = {
  uuid?: string
  name?: string
}

export type GetDetailsResponse = {
  summary?: WatchdogWatchdog.DomainSummary
}

export type ListSummariesRequest = {
  page?: string
  perPage?: string
}

export type ListSummariesResponse = {
  page?: string
  perPage?: string
  summaries?: WatchdogWatchdog.DomainSummary[]
}

export type HealthRequest = {
}

export type HealthResponse = {
  createdAt?: string
  numDomains?: string
  status?: HealthResponseHealthStatus
}

export class WatchDog {
  static Reload(req: ReloadRequest, initReq?: fm.InitReq): Promise<ReloadResponse> {
    return fm.fetchReq<ReloadRequest, ReloadResponse>(`/v1/watchdog/reload`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Get(req: GetRequest, initReq?: fm.InitReq): Promise<GetResponse> {
    return fm.fetchReq<GetRequest, GetResponse>(`/v1/watchdog/get?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetDetails(req: GetDetailsRequest, initReq?: fm.InitReq): Promise<GetDetailsResponse> {
    return fm.fetchReq<GetDetailsRequest, GetDetailsResponse>(`/v1/watchdog/getDetails?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ListSummaries(req: ListSummariesRequest, initReq?: fm.InitReq): Promise<ListSummariesResponse> {
    return fm.fetchReq<ListSummariesRequest, ListSummariesResponse>(`/v1/watchdog/getAll?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static Health(req: HealthRequest, initReq?: fm.InitReq): Promise<HealthResponse> {
    return fm.fetchReq<HealthRequest, HealthResponse>(`/v1/watchdog/health?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}