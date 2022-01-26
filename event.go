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

import (
	"container/list"
)

type Consumer interface {
	Execute(value string)
}

type Producer interface {
	Push(topic, value string)
	Register(topic string, c Consumer)
	Deregister(topic string, c Consumer)
	Notify(topic string)
}

//implements Producer
type simpleProducer struct {
	db DataBase
}

func NewSimpleProducer(db DataBase) *simpleProducer {
	s := new(simpleProducer)
	s.db = db
	return s
}

func (s *simpleProducer) Register(topic string, consumer Consumer) {
	s.db.SaveObv(topic, consumer)
}

func (s *simpleProducer) Deregister(topic string, consumer Consumer) {
	s.db.DeleteObv(topic, consumer)
}

func (s *simpleProducer) Push(topic, value string) {
	if nil == s.db.AllMsg() {
		dbValues := &map[string]*list.List{}
		s.db.InitMsg(dbValues)
	}
	s.db.SaveMsg(topic, Message{
		Value: value,
	})
	s.Notify(topic)
}

var SimpleProducer *simpleProducer

func init() {
	if nil == SimpleProducer {
		db := simpleLocalDb{}
		mg := make(map[string]*list.List)
		ob := make(map[string]*list.List)
		db.InitMsg(&mg)
		db.InitObv(&ob)
		SimpleProducer = NewSimpleProducer(&db)
	}
}

func (s *simpleProducer) Notify(topic string) {
	if nil == s.db.AllMsg() || len(*s.db.AllMsg()) == 0 {
		return
	}
	if nil == s.db.AllObv() || len(*s.db.AllObv()) == 0 {
		return
	}
	values := *s.db.AllMsg()
	vs := values[topic]
	allObs := *s.db.AllObv()
	obs := allObs[topic]
	if nil == obs {
		return
	}
	for v := vs.Front(); v != nil; v = v.Next() {
		msg, _ := v.Value.(Message)
		for ob := obs.Front(); ob != nil; ob = ob.Next() {
			ob.Value.(Consumer).Execute(msg.Value)
		}
		s.db.DeleteMsg(topic, v)
	}
}
