-module(version).

-export([init/2]).

init(Req, Opts) ->
    _ = lager:warning("~s GET /version~n", [?MODULE_STRING]),
    Req2 = cowboy_req:reply(200,
        #{<<"content-type">> => <<"application/json">>},
        jsone:encode([{<<"status">>,<<"running">>}]),
        Req),
    {ok, Req2, Opts}.