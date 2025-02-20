// Copyright © 2018 Alex Kolbasov
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

/*
Package hmsclient provides methods for accessing Hive Metastore over
Thrift protocol. Only unsecured Thrift connections are supported.

Example usage:

  import(
    "log"
    "github.com/SananGuliyev/gometastore/hmsclient"
  )

  func printDatabases() {
    client, err := hmsclient.Open("localhost", 9083)
    if err != nil {
      log.Fatal(err)
    }
    defer client.Close()

    databases, err := client.GetAllDatabases()
    if err != nil {
      log.Fatal(err)
	}
    for _, d := range databases {
      fmt.Println(d)
    }
  }
*/
package hmsclient
