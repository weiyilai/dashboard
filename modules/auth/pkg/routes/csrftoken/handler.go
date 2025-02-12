// Copyright 2017 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/xsrftoken"

	"k8s.io/dashboard/auth/pkg/router"
	"k8s.io/dashboard/csrf"
)

func init() {
	router.V1().GET("/csrftoken/:action", handleLogin)
}

func handleLogin(c *gin.Context) {
	action := c.Param("action")
	token := xsrftoken.Generate(csrf.Key(), "none", action)
	c.JSON(http.StatusOK, csrf.Response{Token: token})
}
