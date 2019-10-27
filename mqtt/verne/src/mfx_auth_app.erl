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
    PoolSize = case os:getenv("MF_MQTT_VERNEMQ_GRPC_POOL_SIZE") of
        false ->
            10;
        PoolSizeEnv ->
            {PoolSizeInt, _PoolSizeRest} = string:to_integer(PoolSizeEnv),
            PoolSizeInt
    end,
    
    ets:insert(mfx_cfg, [
        {nats_url, NatsUrl},
        {redis_url, RedisUrl},
        {instance_id, InstanceId},
        {pool_size, PoolSize}
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
    CL = create_channel_list(PoolSize, GrpcHost, GrpcPort, []),
    application:set_env(grpcbox, client, #{channels => CL}),
    application:ensure_all_started(grpcbox),

    % Start the MFX Auth process
    mfx_auth_sup:start_link(PoolSize).

create_channel_list(0, _GrpcHost, _GrpcPort, CL) ->
    CL;
create_channel_list(PoolSize, GrpcHost, GrpcPort, CL) ->
    DC = list_to_atom("channel_" ++ integer_to_list(PoolSize)),
    CL2 = CL ++ [{DC, [{http, GrpcHost, GrpcPort, []}], #{}}],
    create_channel_list(PoolSize-1, GrpcHost, GrpcPort, CL2).

stop(_State) ->
    ok.


