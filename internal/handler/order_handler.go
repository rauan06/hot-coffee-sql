package handler

import "net/http"

type OrderHandler struct{}

func (o *OrderHandler) CreateOrder(http.ResponseWriter, *http.Request)           {}
func (o *OrderHandler) RetrieveAllOrders(http.ResponseWriter, *http.Request)     {}
func (o *OrderHandler) RetrieveSepcificOrder(http.ResponseWriter, *http.Request) {}
func (o *OrderHandler) UpdateOrder(http.ResponseWriter, *http.Request)           {}
func (o *OrderHandler) DeleteOrder(http.ResponseWriter, *http.Request)           {}
func (o *OrderHandler) CloseOrder(http.ResponseWriter, *http.Request)            {}
