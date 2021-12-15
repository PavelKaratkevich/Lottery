package main

import (
	"Lottery/lotterypb"
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection to localhost:8000 failed: %v", err)
	}
	cs := lotterypb.NewLotteryServiceClient(conn)
	g := gin.Default()
	g.POST("/buy/:first_name/:last_name/:id_number/", func(ctx *gin.Context) {
		a, b, c := ctx.Param("first_name"), ctx.Param("last_name"), ctx.Param("id_number")
		req := &lotterypb.Request{FirstName: a, LastName: b, IdNumber: c}
		response, err := cs.BuyLotteryTicket(context.Background(), req)
		if err == nil {
			ctx.JSON(http.StatusOK, fmt.Sprintf("ID of your ticket: %v", response.TicketId))
		} else {
			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.AlreadyExists:
					ctx.JSON(http.StatusForbidden, e.Message())
				case codes.Internal:
					ctx.JSON(http.StatusInternalServerError, e.Message())
				case codes.Aborted:
					ctx.JSON(http.StatusInternalServerError, e.Message())
				case codes.ResourceExhausted:
					ctx.JSON(http.StatusGone, e.Message())
				}
			}
		}
	})
	if err1 := g.Run(":8080"); err1 != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
