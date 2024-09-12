package streamx

var ClientTpl = `// Code generated by Kitex {{.Version}}. DO NOT EDIT.

package {{ToLower .ServiceName}}

import (
	{{- range $path, $aliases := .Imports}}
		{{- if not $aliases}}
			"{{$path}}"
		{{- else}}
			{{- range $alias, $is := $aliases}}
				{{$alias}} "{{$path}}"
			{{- end}}
		{{- end}}
	{{- end}}
)
{{- $protocol := .Protocol | getStreamxRef}}

type Client interface {
{{- range .AllMethods}}
{{- $unary := and (not .ServerStreaming) (not .ClientStreaming)}}
{{- $clientSide := and .ClientStreaming (not .ServerStreaming)}}
{{- $serverSide := and (not .ClientStreaming) .ServerStreaming}}
{{- $bidiSide := and .ClientStreaming .ServerStreaming}}
{{- $arg := index .Args 0}}
    {{.Name}}{{- if $unary}}(ctx context.Context, req {{$arg.Type}}, callOptions ...streamxcallopt.CallOption) (r {{.Resp.Type}}, err error)
             {{- else if $clientSide}}(ctx context.Context, callOptions ...streamxcallopt.CallOption) (stream streamx.ClientStreamingClient[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr $arg.Type}}, {{NotPtr .Resp.Type}}], err error)
             {{- else if $serverSide}}(ctx context.Context, req {{$arg.Type}}, callOptions ...streamxcallopt.CallOption) (stream streamx.ServerStreamingClient[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr .Resp.Type}}], err error)
             {{- else if $bidiSide}}(ctx context.Context, callOptions ...streamxcallopt.CallOption) (stream streamx.BidiStreamingClient[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr $arg.Type}}, {{NotPtr .Resp.Type}}], err error)
             {{- end}}
{{- end}}
}

func NewClient(destService string, opts ...streamxclient.Option) (Client, error) {
	var options []streamxclient.Option
	options = append(options, streamxclient.WithDestService(destService))
	options = append(options, opts...)
	cp, err := {{$protocol}}.NewClientProvider(svcInfo)
	if err != nil {
		return nil, err
	}
	options = append(options, streamxclient.WithProvider(cp))
	cli, err := streamxclient.NewClient(svcInfo, options...)
	if err != nil {
		return nil, err
	}
	kc := &kClient{streamer: cli, caller: cli.(client.Client)}
	return kc, nil
}

var _ Client = (*kClient)(nil)

type kClient struct {
    caller client.Client
    streamer streamxclient.Client
}

{{- range .AllMethods}}
{{- $unary := and (not .ServerStreaming) (not .ClientStreaming)}}
{{- $clientSide := and .ClientStreaming (not .ServerStreaming)}}
{{- $serverSide := and (not .ClientStreaming) .ServerStreaming}}
{{- $bidiSide := and .ClientStreaming .ServerStreaming}}
{{- $mode := ""}}
	{{- if $bidiSide -}} {{- $mode = "serviceinfo.StreamingBidirectional" }}
	{{- else if $serverSide -}} {{- $mode = "serviceinfo.StreamingServer" }}
	{{- else if $clientSide -}} {{- $mode = "serviceinfo.StreamingClient" }}
	{{- else if $unary -}} {{- $mode = "serviceinfo.StreamingUnary" }}
    {{- end}}
{{- $arg := index .Args 0}}
func (c *kClient) {{.Name}}{{- if $unary}}(ctx context.Context, req {{$arg.Type}}, callOptions ...streamxcallopt.CallOption) ({{.Resp.Type}}, error) {
    res := new({{NotPtr .Resp.Type}})
    _, err := streamxclient.InvokeStream[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr $arg.Type}}, {{NotPtr .Resp.Type}}](
        ctx, c.streamer, {{$mode}}, "{{.RawName}}", req, res, callOptions...)
    if err != nil {
        return nil, err
    }
    return res, nil
{{- else if $clientSide}}(ctx context.Context, callOptions ...streamxcallopt.CallOption) (stream streamx.ClientStreamingClient[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr $arg.Type}}, {{NotPtr .Resp.Type}}], err error) {
    return streamxclient.InvokeStream[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr $arg.Type}}, {{NotPtr .Resp.Type}}](
        ctx, c.streamer, {{$mode}}, "{{.RawName}}", nil, nil, callOptions...)
{{- else if $serverSide}}(ctx context.Context, req {{$arg.Type}}, callOptions ...streamxcallopt.CallOption) (stream streamx.ServerStreamingClient[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr .Resp.Type}}], err error) {
    return streamxclient.InvokeStream[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr $arg.Type}}, {{NotPtr .Resp.Type}}](
        ctx, c.streamer, {{$mode}}, "{{.RawName}}", req, nil, callOptions...)
{{- else if $bidiSide}}(ctx context.Context, callOptions ...streamxcallopt.CallOption) (stream streamx.BidiStreamingClient[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr $arg.Type}}, {{NotPtr .Resp.Type}}], err error) {
    return streamxclient.InvokeStream[{{$protocol}}.Header, {{$protocol}}.Trailer, {{NotPtr $arg.Type}}, {{NotPtr .Resp.Type}}](
        ctx, c.streamer, {{$mode}}, "{{.RawName}}", nil, nil, callOptions...)
{{- end}}
}
{{- end}}
`
