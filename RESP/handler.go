package resp

import (
	"sync"
)

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"SET":  set,
	"GET":  get,
	"HSET": hset,
	"HGET": hget,
}
var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}
var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{Typ: "string", Str: "Pong"}
	}
	return Value{Typ: "string", Str: args[0].Bulk}
}
func set(args []Value) Value {
	if len(args) != 2 {
		return Value{Typ: "error", Str: "Err wrong number of arguments for Set command"}
	}
	key := args[0].Bulk
	value := args[1].Bulk
	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()
	return Value{Typ: "string", Str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{Typ: "error", Str: "Err wrong number of arguments for Get command"}
	}
	key := args[0].Bulk
	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()
	if !ok {
		return Value{Typ: "null"}
	}
	return Value{Typ: "string", Str: value}
}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{Typ: "error", Str: "Err wrong number of arguments for Set command"}
	}
	set := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk
	HSETsMu.Lock()
	if _, ok := HSETs[set]; !ok {
		HSETs[set] = map[string]string{}
	}
	HSETs[set][key] = value
	HSETsMu.Unlock()
	return Value{Typ: "string", Str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{Typ: "error", Str: "Err wrong number of arguments for Get command"}
	}
	set := args[0].Bulk
	key := args[1].Bulk
	HSETsMu.RLock()
	value, ok := HSETs[set][key]
	HSETsMu.RUnlock()
	if !ok {
		return Value{Typ: "null"}
	}
	return Value{Typ: "string", Str: value}
}
