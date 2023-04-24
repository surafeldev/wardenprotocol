// TEMPORARY AUTOGENERATED FILE: tinyjson stub code to make the package
// compilable during generation.

package  src

import (
  "github.com/CosmWasm/tinyjson/jwriter"
  "github.com/CosmWasm/tinyjson/jlexer"
)

func ( ExecuteMsg ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* ExecuteMsg ) UnmarshalJSON([]byte) error { return nil }
func ( ExecuteMsg ) MarshalTinyJSON(w *jwriter.Writer) {}
func (* ExecuteMsg ) UnmarshalTinyJSON(l *jlexer.Lexer) {}

type TinyJSON_exporter_ExecuteMsg *ExecuteMsg

func ( InstantiateMsg ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* InstantiateMsg ) UnmarshalJSON([]byte) error { return nil }
func ( InstantiateMsg ) MarshalTinyJSON(w *jwriter.Writer) {}
func (* InstantiateMsg ) UnmarshalTinyJSON(l *jlexer.Lexer) {}

type TinyJSON_exporter_InstantiateMsg *InstantiateMsg

func ( Test ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* Test ) UnmarshalJSON([]byte) error { return nil }
func ( Test ) MarshalTinyJSON(w *jwriter.Writer) {}
func (* Test ) UnmarshalTinyJSON(l *jlexer.Lexer) {}

type TinyJSON_exporter_Test *Test

func ( Verify ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* Verify ) UnmarshalJSON([]byte) error { return nil }
func ( Verify ) MarshalTinyJSON(w *jwriter.Writer) {}
func (* Verify ) UnmarshalTinyJSON(l *jlexer.Lexer) {}

type TinyJSON_exporter_Verify *Verify

func ( VerifyResponse ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* VerifyResponse ) UnmarshalJSON([]byte) error { return nil }
func ( VerifyResponse ) MarshalTinyJSON(w *jwriter.Writer) {}
func (* VerifyResponse ) UnmarshalTinyJSON(l *jlexer.Lexer) {}

type TinyJSON_exporter_VerifyResponse *VerifyResponse
