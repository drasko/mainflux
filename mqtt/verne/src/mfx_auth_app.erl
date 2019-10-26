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

    NatsUrl = case os:getenv("MF_NATS_URL") of
        false -> "nats://localhost:4222";
        NatsEnv -> NatsEnv
    end,
    RedisUrl = case os:getenv("MF_MQTT_ADAPTER_ES_URL") of
        false -> "tcp://localhost:6379";
        RedisEnv -> RedisEnv
    end,
    InstanceId = case os:getenv("MF_MQTT_INSTANCE_ID") of
        false -> "";
        InstanceEnv -> InstanceEnv
    end,
    
    ets:insert(mfx_cfg, [
        {nats_url, NatsUrl},
        {redis_url, RedisUrl},
        {instance_id, InstanceId}
    ]),

    % Also, init one ETS table for keeping the #{ClientId => Username} mapping
    ets:new(mfx_client_map, [set, named_table, public]),

    % Start Grpcbox
    GrpcUrl = case os:getenv("MF_THINGS_AUTH_GRPC_URL") of
        false -> "tcp://localhost:8183";
        GrpcEnv -> GrpcEnv
    end,
    {ok, {_, _, GrpcHost, GrpcPort, _, _}} = http_uri:parse(GrpcUrl),
    error_logger:info_msg("gRPC host: ~p,  port: ~p", [GrpcHost, GrpcPort]),

    application:load(grpcbox),
    application:set_env(grpcbox, client, #{channels => [{default_channel, [{http, GrpcHost, GrpcPort, []}], #{}}]}),
    application:ensure_all_started(grpcbox),

    % Start the MFX Auth process
    mfx_auth_sup:start_link().

stop(_State) ->
    ok.


