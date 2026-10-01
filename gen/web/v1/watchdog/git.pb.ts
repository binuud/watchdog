/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

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
  repo?: string
  user?: string
  status?: GitProjectProjectStatus
}

export type GitStatusResponse = {
  branchName?: string
  upstreamBranch?: string
  aheadCount?: number
  behindCount?: number
  isClean?: boolean
  numStaged?: string
  numUnstaged?: string
  numUntracked?: string
  numIgnored?: string
  staged?: FileStatus[]
  unstaged?: FileStatus[]
  untracked?: string[]
  ignored?: string[]
}

export type FileStatus = {
  path?: string
  oldPath?: string
  status?: FileStatusStatusCode
}