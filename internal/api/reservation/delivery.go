package reservation

import "github.com/labstack/echo"

type Handler interface{ Create() echo.HandlerFunc }
