#!/usr/bin/env bash
#
# Function: 通用脚本底层组件
# Called by: 其他脚本

SYSLOG_ERROR="user.error"
SYSLOG_INFO="user.info"
SYSLOG_DEBUG="user.debug"

__green() {
  printf '\33[1;32m%b\33[0m' "$1"
}

__red() {
  printf '\33[1;31m%b\33[0m' "$1"
}

_printargs() {
  _exitstatus="$?"
  printf -- "%s" "[$(date)] "
  if [ -z "$2" ]; then
    printf -- "%s" "$1"
  else
    printf -- "%s" "$1='$2'"
  fi
  printf "\n"
  # return the saved exit status
  return "$_exitstatus"
}

_contains() {
  _str="$1"
  _sub="$2"
  echo "$_str" | grep -- "$_sub" >/dev/null 2>&1
}

#class
_syslog() {
  _exitstatus="$?"
  _logclass="$1"
  shift
  if [ -z "$__logger_i" ]; then
    if _contains "$(logger --help 2>&1)" "-i"; then
      __logger_i="logger -i"
    else
      __logger_i="logger"
    fi
  fi
  $__logger_i -t "$PROJECT_NAME" -p "$_logclass" "$(_printargs "$@")" >/dev/null 2>&1
  return "$_exitstatus"
}

_log() {
  [ -z "$LOG_FILE" ] && return
  _printargs "$@" >>"$LOG_FILE"
}

_info() {
  _log "$@"
  _syslog "$SYSLOG_INFO" "$@"
  _printargs "$@"
}

_err() {
  _syslog "$SYSLOG_ERROR" "$@"
  _log "$@"
  printf -- "%s" "[$(date)] " >&2
  if [ -z "$2" ]; then
    __red "$1" >&2
  else
    __red "$1='$2'" >&2
  fi
  printf "\n" >&2
  return 1
}
