/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"

export enum GitProjectProjectStatus {
  IGNORE_PROJECT_STATUS = "IGNORE_PROJECT_STATUS",
  ACTIVE = "ACTIVE",
  INACTIVE = "INACTIVE",
  ARCHIVE = "ARCHIVE",
}

export enum FileStatusStatusCode {
  STATUS_CODE_UNSPECIFIED = "STATUS_CODE_UNSPECIFIED",
  MODIFIED = "MODIFIED",
  ADDED = "ADDED",
  DELETED = "DELETED",
  RENAMED = "RENAMED",
  COPIED = "COPIED",
  TYPE_CHANGED = "TYPE_CHANGED",
  UNMERGED = "UNMERGED",
}

export type GitProject = {
  uuid?: string
  name?: string
  path?: string
  description?: string
  org?: string
  projecturl?: string
  remoteurl?: string
  user?: string
  status?: GitProjectProjectStatus
}

export type GitProjectStatus = {
  branchName?: string
  upstreamBranch?: string
  aheadCount?: number
  behindCount?: number
  isClean?: boolean
  isPrivate?: boolean
  visibility?: string
  numModified?: string
  numUnstaged?: string
  numUntracked?: string
  numIgnored?: string
  numStashes?: string
  numPrOpen?: string
  numIssuesOpen?: string
  staged?: FileStatus[]
  unstaged?: FileStatus[]
  untracked?: string[]
  ignored?: string[]
  commonAncestor?: string
}

export type FileStatus = {
  path?: string
  oldPath?: string
  status?: FileStatusStatusCode
}

export type GitProjectRow = {
  project?: GitProject
  status?: GitProjectStatus
  createdAt?: GoogleProtobufTimestamp.Timestamp
}