-module(mfx_grpc).
-behaviour(gen_server).

-export([
    start_link/0,
    start_link/1,
    init/1,
    handle_call/3,
    handle_cast/2,
    handle_info/2,
    terminate/2
]).

-record(state, {pool_size}).

init(_Args) ->
    error_logger:info_msg("mfx_grpc genserver has started (~w)~n", [self()]),
    [{_, PoolSize}] = ets:lookup(mfx_cfg, pool_size),
    {ok, #state{pool_size = PoolSize}}.

start_link() ->
    gen_server:start_link({local, ?MODULE}, ?MODULE, [], []).

start_link(Args) ->
    gen_server:start_link(?MODULE, Args, []).

handle_call({identify, Message}, _From, #state{pool_size = PoolSize} = State) ->
    error_logger:info_msg("mfx_grpc message: ~p", [Message]),
    Opts = #{channel => list_to_atom("channel_" ++ integer_to_list(rand:uniform(PoolSize)))},
    {ok, Resp, HeadersAndTrailers} = mainflux_things_service_client:identify(Message, Opts),
    case maps:get(<<":status">>, maps:get(headers, HeadersAndTrailers)) of
        <<"200">> ->
            {reply, {ok, maps:get(value, Resp)}, State};
        ErrorStatus ->
            {reply, {error, ErrorStatus}, State}
    end;

handle_call({can_access_by_id, Message}, _From, #state{pool_size = PoolSize} = State) ->
    error_logger:info_msg("mfx_grpc message: ~p", [Message]),
    Opts = #{channel => list_to_atom("channel_" ++ integer_to_list(rand:uniform(PoolSize)))},
    {ok, _, HeadersAndTrailers} = mainflux_things_service_client:can_access_by_id(Message, Opts),
    error_logger:info_msg("mfx_grpc can_access_by_id() HeadersAndTrailers: ~p", [HeadersAndTrailers]),
    case maps:get(<<":status">>, maps:get(headers, HeadersAndTrailers)) of
        <<"200">> ->
            {reply, ok, State};
        ErrorStatus ->
            {reply, {error, ErrorStatus}, State}
    end.

handle_cast(_Request, State) ->
    {noreply, State}.

handle_info(_Info, State) ->
    {noreply, State}.

terminate(_Reason, _State) ->
    [].

