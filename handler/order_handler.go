package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ibrahimker/latihan-register/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type OrderHandlerInterface interface {
	OrderHandler(w http.ResponseWriter, r *http.Request)
}

type OrderHandler struct {
	postgrespool *pgxpool.Pool
}

func NewOrderHandler(postgrespool *pgxpool.Pool) OrderHandlerInterface {
	return &OrderHandler{postgrespool: postgrespool}
}

func (h *OrderHandler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// id := params["id"]

	switch r.Method {
	case http.MethodGet:
		h.getOrdersHandler(w, r)
	case http.MethodPost:
		// createUsersHandler(w, r)
	case http.MethodPut:
		// updateUserHandler(w, r, id)
	case http.MethodDelete:
		// deleteUserHandler(w, r, id)
	}
}

func (h *OrderHandler) getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// cara 1

	// queryString := `select
	// 	o.id as order_id
	// 	,o.customer_name
	// 	,o.ordered_at
	// 	,json_agg(json_build_object(
	// 		'item_id',i.id
	// 		,'item_code',i.code
	// 		,'description',i.description
	// 		,'quantity',i.quantity
	// 		,'order_id',i.order_id
	// 	)) as items
	// from public.order o join item i
	// on o.id = i.order_id
	// group by o.id`
	// rows, err := h.postgrespool.Query(ctx, queryString)
	// if err != nil {
	// 	fmt.Println("query row error", err)
	// }
	// defer rows.Close()

	// var orders []*entity.Order
	// for rows.Next() {
	// 	var o entity.Order
	// 	var itemsStr string
	// 	if serr := rows.Scan(&o.ID, &o.CustomerName, &o.OrderedAt, &itemsStr); serr != nil {
	// 		fmt.Println("Scan error", serr)
	// 	}
	// 	var items []entity.Item
	// 	if err := json.Unmarshal([]byte(itemsStr), &items); err != nil {
	// 		fmt.Errorf("Error when parsing items")
	// 	} else {
	// 		o.Items = append(o.Items, items...)
	// 	}
	// 	orders = append(orders, &o)
	// }
	// end cara 1

	// cara 2
	queryString := `select
		o.id as order_id
		,o.customer_name
		,o.ordered_at
	from public.order o`
	rows, err := h.postgrespool.Query(ctx, queryString)
	if err != nil {
		fmt.Println("query row error", err)
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var o entity.Order
		if serr := rows.Scan(&o.ID, &o.CustomerName, &o.OrderedAt); serr != nil {
			fmt.Println("Scan error", serr)
		}
		orders = append(orders, &o)
	}

	var wg sync.WaitGroup
	wg.Add(len(orders))
	for i := 0; i < len(orders); i++ {
		go func(i int) {
			defer wg.Done()
			queryString := `select
			i.id
			,i.code
			,i.description
			,i.quantity
			,i.order_id
		from public.item i where i.order_id = $1`
			rows, err := h.postgrespool.Query(ctx, queryString, orders[i].ID)
			if err != nil {
				fmt.Println("query row error", err)
			}
			defer rows.Close()

			var items []entity.Item
			for rows.Next() {
				var i entity.Item
				if serr := rows.Scan(&i.ID, &i.Code, &i.Description, &i.Quantity, &i.OrderID); serr != nil {
					fmt.Println("Scan error", serr)
				}
				items = append(items, i)
			}
			orders[i].Items = items
		}(i)
	}
	wg.Wait()

	// end cara 2
	jsonData, _ := json.Marshal(&orders)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
