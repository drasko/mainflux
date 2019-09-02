%% -*- coding: utf-8 -*-
%% Automatically generated, do not edit
%% Generated by gpb_compile version 4.9.0

-ifndef(message).
-define(message, true).

-define(message_gpb_version, "4.9.0").

-ifndef('MAINFLUX.RAWMESSAGE_PB_H').
-define('MAINFLUX.RAWMESSAGE_PB_H', true).
-record('mainflux.RawMessage',
        {channel = []           :: iodata() | undefined, % = 1
         subtopic = []          :: iodata() | undefined, % = 2
         publisher = []         :: iodata() | undefined, % = 3
         protocol = []          :: iodata() | undefined, % = 4
         contentType = []       :: iodata() | undefined, % = 5
         payload = <<>>         :: iodata() | undefined % = 6
        }).
-endif.

-ifndef('MAINFLUX.MESSAGE_PB_H').
-define('MAINFLUX.MESSAGE_PB_H', true).
-record('mainflux.Message',
        {channel = []           :: iodata() | undefined, % = 1
         subtopic = []          :: iodata() | undefined, % = 2
         publisher = []         :: iodata() | undefined, % = 3
         protocol = []          :: iodata() | undefined, % = 4
         name = []              :: iodata() | undefined, % = 5
         unit = []              :: iodata() | undefined, % = 6
         value                  :: {floatValue, float() | integer() | infinity | '-infinity' | nan} | {stringValue, iodata()} | {boolValue, boolean() | 0 | 1} | {dataValue, iodata()} | undefined, % oneof
         valueSum = undefined   :: message:'mainflux.SumValue'() | undefined, % = 11
         time = 0.0             :: float() | integer() | infinity | '-infinity' | nan | undefined, % = 12
         updateTime = 0.0       :: float() | integer() | infinity | '-infinity' | nan | undefined, % = 13
         link = []              :: iodata() | undefined % = 14
        }).
-endif.

-ifndef('MAINFLUX.SUMVALUE_PB_H').
-define('MAINFLUX.SUMVALUE_PB_H', true).
-record('mainflux.SumValue',
        {value = 0.0            :: float() | integer() | infinity | '-infinity' | nan | undefined % = 1
        }).
-endif.

-endif.
