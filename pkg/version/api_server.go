package version

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"fmt"

	pb "github.com/bhojpur/data/pkg/api/v1/version"
	"github.com/bhojpur/data/pkg/internal/errors"

	"github.com/gogo/protobuf/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type apiServer struct {
	version *pb.Version
	options APIServerOptions
}

func newAPIServer(version *pb.Version, options APIServerOptions) *apiServer {
	return &apiServer{version, options}
}

func (a *apiServer) GetVersion(ctx context.Context, request *types.Empty) (response *pb.Version, err error) {
	return a.version, nil
}

// APIServerOptions are options when creating a new APIServer.
type APIServerOptions struct {
	DisableLogging bool
}

// NewAPIServer creates a new APIServer for the given Version.
func NewAPIServer(version *pb.Version, options APIServerOptions) pb.APIServer {
	return newAPIServer(version, options)
}

// GetServerVersion gets the server *Version given the *grpc.ClientConn.
func GetServerVersion(clientConn *grpc.ClientConn) (*pb.Version, error) {
	res, err := pb.NewAPIClient(clientConn).GetVersion(
		context.Background(),
		&types.Empty{},
	)
	return res, errors.EnsureStack(err)
}

// String returns a string representation of the Version.
func String(v *pb.Version) string {
	return fmt.Sprintf("%d.%d.%d%s", v.Major, v.Minor, v.Micro, v.Additional)
}
