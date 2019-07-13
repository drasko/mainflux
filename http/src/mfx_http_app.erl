-module(mfx_auth_app).

-behaviour(application).

%% Application callbacks
-export([start/2, stop/1]).

%% ===================================================================
%% Application callbacks
%% ===================================================================

start(_StartType, _StartArgs) ->

    % Put ENV variables in ETS
    ets:new(mfx_cfg, [set, named_table, public]),

    AuthUrl = case os:getenv("MF_THINGS_AUTH_HTTP_URL") of
        false -> "http://localhost:8989";
        AuthEnv -> AuthEnv
    end,
    NatsUrl = case os:getenv("MF_NATS_URL") of
        false -> "nats://localhost:4222";
        NatsEnv -> NatsEnv
    end,

    ets:insert(mfx_cfg, [
        {auth_url, AuthUrl},
        {nats_url, NatsUrl}
    ]),

    % Start the process
    {ok, Pid} = 'mfx_http_sup':start_link(),
	Routes = [ {
        '_',
        [
            {"/messages", messages, []}
        ]
    } ],
    Dispatch = cowboy_router:compile(Routes),

    TransOpts = [ {ip, {0,0,0,0}}, {port, 8089} ],
	ProtoOpts = #{env => #{dispatch => Dispatch}},

	{ok, _} = cowboy:start_clear(esk_cowboy,
		TransOpts, ProtoOpts),

    {ok, Pid}.

stop(_State) ->
    ok.
