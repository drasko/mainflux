%% -*- coding: utf-8 -*-
%% Automatically generated, do not edit
%% Generated by gpb_compile version 4.9.0

-ifndef(internal).
-define(internal, true).

-define(internal_gpb_version, "4.9.0").

-ifndef('ACCESSREQ_PB_H').
-define('ACCESSREQ_PB_H', true).
-record('AccessReq',
        {token = []             :: iodata() | undefined, % = 1
         chanID = []            :: iodata() | undefined % = 2
        }).
-endif.

-ifndef('THINGID_PB_H').
-define('THINGID_PB_H', true).
-record('ThingID',
        {value = []             :: iodata() | undefined % = 1
        }).
-endif.

-ifndef('TOKEN_PB_H').
-define('TOKEN_PB_H', true).
-record('Token',
        {value = []             :: iodata() | undefined % = 1
        }).
-endif.

-ifndef('USERID_PB_H').
-define('USERID_PB_H', true).
-record('UserID',
        {value = []             :: iodata() | undefined % = 1
        }).
-endif.

-endif.
