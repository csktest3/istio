// Copyright 2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
)

// ClientHandshake return the client handshake info.
func (s *handler) ClientHandshake(_ context.Context, _ string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	info := CredInfo{Err: ErrInvalidConnection}
	return conn, info, nil
}

// ServerHandshake return the server handshake info.
func (s *handler) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	var creds CredInfo

	if s.creds == nil {
		creds = CredInfo{Err: ErrNoCredentials}
	} else {
		creds = *s.creds
	}
	return conn, creds, nil
}

// Info returns the proto info.
func (s *handler) Info() credentials.ProtocolInfo {
	return credentials.ProtocolInfo{
		SecurityProtocol: authType,
		SecurityVersion:  "0.1",
		ServerName:       "workloadhandler",
	}
}

// Clone returns the clone info.
func (s *handler) Clone() credentials.TransportCredentials {
	return &(*s)
}

// OverrideServerName overrides server name.
func (s *handler) OverrideServerName(_ string) error {
	return nil
}

// GetCred gets the cred.
func (s *handler) GetCred() credentials.TransportCredentials {
	return s
}
