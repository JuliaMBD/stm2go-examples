// This file was generated by a program.
// Please do not edit this file directly.
package testpackage

import (
// package names to be imported
)

type TimerStartStm2State uint8
const (
TimerStartStm2Count TimerStartStm2State = iota
TimerStartStm2StopPush
)

var TimerStartStm2Eod Eod
var TimerStartStm2CurrentState TimerStartStm2State
var TimerStartStm2NextState TimerStartStm2State

func init() {
TimerStartStm2Initialize()
}

func TimerStartStm2Initialize() {
TimerStartStm2Eod = Entry
TimerStartStm2CurrentState = TimerStartStm2Count
TimerStartStm2NextState = TimerStartStm2Count
}

func TimerStartStm2Task() {
switch TimerStartStm2CurrentState {
case TimerStartStm2Count:
if TimerStartStm2Eod == Entry {
TimerStartStm2CountEntry()
TimerStartStm2Eod = Do
}
if TimerStartStm2Eod == Do {
TimerStartStm2CountDo()
if TimerStartStm2CountOneSecCond() {
TimerStartStm2CountOneSecAction()
TimerStartStm2NextState = TimerStartStm2Count
TimerStartStm2Eod = Exit
}
if TimerStartStm2CountPushCond() {
TimerStartStm2CountPushAction()
TimerStartStm2NextState = TimerStartStm2StopPush
TimerStartStm2Eod = Exit
}
}
if TimerStartStm2Eod == Exit {
TimerStartStm2CountExit()
TimerStartStm2Eod = Entry
}
case TimerStartStm2StopPush:
if TimerStartStm2Eod == Entry {
TimerStartStm2StopPushEntry()
TimerStartStm2Eod = Do
}
if TimerStartStm2Eod == Do {
TimerStartStm2StopPushDo()
}
if TimerStartStm2Eod == Exit {
TimerStartStm2StopPushExit()
TimerStartStm2Eod = Entry
}
}
}

func TimerStartStm2Update() {
switch TimerStartStm2CurrentState {
case TimerStartStm2Count:
case TimerStartStm2StopPush:
}
TimerStartStm2CurrentState = TimerStartStm2NextState
}

