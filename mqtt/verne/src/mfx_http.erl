-module(mfx_http).
-behaviour(gen_server).
-behaviour(poolboy_worker).

-export([
    start_link/0,
    start_link/1,
    init/1,
    send/2,
    handle_call/3,
    handle_cast/2,
    handle_info/2,
    terminate/2
]).

-record(state, {url}).

init(_Args) ->
    error_logger:info_msg("mfx_http genserver has started (~w)~n", [self()]),

    [{_, AuthUrl}] = ets:lookup(mfx_cfg, auth_url),
    error_logger:info_msg("Auth URL: ~p", [AuthUrl]),
    {ok, #state{url = AuthUrl}}.

start_link() ->
    gen_server:start_link({local, ?MODULE}, ?MODULE, [], []).

start_link(Args) ->
    gen_server:start_link(?MODULE, Args, []).

send(Method, Message) ->
    gen_server:call(?MODULE, {send, Method, Message}).

handle_call({identify, Password}, _From, #state{url = AuthUrl} = State) ->
  error_logger:info_msg("mfx_http message: ~p", [Password]),
  URL = [list_to_binary(AuthUrl), <<"/identify">>],
  ReqBody = jsone:encode(#{<<"token">> => Password}),
  ReqHeaders = [{<<"Content-Type">>, <<"application/json">>}],
  Options = [{pool, default}],
  case hackney:request(post, URL, ReqHeaders, ReqBody, Options) of
      {ok, Status, _, Ref} ->
          case Status of
              200 ->
                  case hackney:body(Ref) of
                      {ok, RespBody} ->
                          {[{<<"id">>, ThingId}]} = jsone:decode(RespBody, [{object_format, tuple}]),
                          error_logger:info_msg("identify: ~p", [URL]),
                          {reply, {ok, ThingId}, State};
                      _ ->
                          {reply, {error, hackney_error}, State}
                  end;
              403 ->
                  {reply, {error, invalid_credentials}, State};
              _ ->
                  {reply, {error, auth_error}, State}
          end;
      {error, checkout_timeout} ->
          {reply, {error, checkout_timeout}, State};
      _ ->
          {reply, {error, auth_req_error}, erStateror}
  end;
handle_call({can_access_by_id, UserName, ChannelId}, _From, #state{url = AuthUrl} = State) ->
  error_logger:info_msg("mfx_http UserName: ~p, ChannelId: ~p", [UserName, ChannelId]),
  URL = [list_to_binary(AuthUrl), <<"/channels/">>, ChannelId, <<"/access-by-id">>],
  error_logger:info_msg("URL: ~p", [URL]),
  ReqBody = jsone:encode(#{<<"thing_id">> => UserName}),
  ReqHeaders = [{<<"Content-Type">>, <<"application/json">>}],
  Options = [{pool, default}],
  case hackney:request(post, URL, ReqHeaders, ReqBody, Options) of
      {ok, Status, _RespHeaders, _ClientRef} ->
          case Status of
              200 ->
                  {reply, ok, State};
              403 ->
                  {reply, {error, forbidden}, State};
              _ ->
                  {reply, {error, auth_error}, State}
          end;
      {error, checkout_timeout} ->
          {reply, {error, checkout_timeout}, State};
      _ ->
          {reply, {error,auth_req_error}, State}
  end.

handle_cast(_Request, State) ->
    {noreply, State}.

handle_info(_Info, State) ->
    {noreply, State}.

terminate(_Reason, _State) ->
    ok.
