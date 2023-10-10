package interfaces

import "github.com/gin-gonic/gin"

type GET func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

type POST func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

type PUT func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

type DELETE func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

type Middleware func(handlers ...gin.HandlerFunc)
