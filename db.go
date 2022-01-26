/*
 * Copyright (c) 2022 Govent
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 */
package govent

import "container/list"

type DataBase interface {
	SaveMsg(topic string, msg Message)
	DeleteMsg(topic string, e *list.Element)
	AllMsg() *map[string]*list.List
	InitMsg(values *map[string]*list.List) *map[string]*list.List
	SaveObv(topic string, consumer Consumer)
	DeleteObv(topic string, consumer Consumer)
	AllObv() *map[string]*list.List
	InitObv(values *map[string]*list.List) *map[string]*list.List
}

type Message struct {
	Value string
}

var localMessages *map[string]*list.List
var localObservers *map[string]*list.List

type simpleLocalDb struct {
}

func (db *simpleLocalDb) DeleteMsg(topic string, e *list.Element) {
	values := *localMessages
	v := values[topic]
	if nil == v {
		return
	}
	v.Remove(e)
	values[topic] = v
	localMessages = &values
}

func (db *simpleLocalDb) SaveMsg(topic string, msg Message) {
	values := *localMessages
	v := values[topic]
	if nil == v {
		v = list.New()
	}
	v.PushBack(msg)
	values[topic] = v
	localMessages = &values
}

func (db *simpleLocalDb) AllMsg() *map[string]*list.List {
	return localMessages
}

func (db *simpleLocalDb) InitMsg(values *map[string]*list.List) *map[string]*list.List {
	localMessages = values
	return localMessages
}

func (db *simpleLocalDb) DeleteObv(topic string, consumer Consumer) {
	obs := *localObservers
	v := obs[topic]
	if nil == v {
		return
	}
	for ob := v.Front(); ob != nil; ob = ob.Next() {
		if ob.Value.(*Consumer) == &consumer {
			v.Remove(ob)
			break
		}
	}
	obs[topic] = v
	localObservers = &obs
}

func (db *simpleLocalDb) SaveObv(topic string, consumer Consumer) {
	values := *localObservers
	obs := values[topic]
	if nil == obs {
		obs = list.New()
	}
	obs.PushBack(consumer)
	values[topic] = obs
	localObservers = &values
}

func (db *simpleLocalDb) AllObv() *map[string]*list.List {
	return localObservers
}

func (db *simpleLocalDb) InitObv(values *map[string]*list.List) *map[string]*list.List {
	localObservers = values
	return localObservers
}
