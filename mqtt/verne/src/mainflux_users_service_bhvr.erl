%%%-------------------------------------------------------------------
%% @doc Behaviour to implement for grpc service mainflux.UsersService.
%% @end
%%%-------------------------------------------------------------------

%% this module was generated on 2019-10-25T22:19:30+00:00 and should not be modified manually

-module(mainflux_users_service_bhvr).

%% @doc Unary RPC
-callback identify(ctx:ctx(), internal_pb:token()) ->
    {ok, internal_pb:user_id(), ctx:ctx()} | grpcbox_stream:grpc_error_response().

