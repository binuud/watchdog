/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type ReloadRequest = {
}

export type ReloadResponse = {
}

export class WatchDog {
  static Reload(req: ReloadRequest, initReq?: fm.InitReq): Promise<ReloadResponse> {
    return fm.fetchReq<ReloadRequest, ReloadResponse>(`/v1/watchdog/reload`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}