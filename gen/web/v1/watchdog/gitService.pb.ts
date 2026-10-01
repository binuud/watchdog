/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"

export enum GitHealthResponseHealthStatus {
  HealthStatusIgnore = "HealthStatusIgnore",
  Active = "Active",
  Error = "Error",
}

export type GitHealthRequest = {
}

export type GitHealthResponse = {
  createdAt?: string
  numProjects?: string
  status?: GitHealthResponseHealthStatus
}

export class Git {
  static Health(req: GitHealthRequest, initReq?: fm.InitReq): Promise<GitHealthResponse> {
    return fm.fetchReq<GitHealthRequest, GitHealthResponse>(`/v1/git/health?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}