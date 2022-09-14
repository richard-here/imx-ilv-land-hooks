package imx

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/richard-here/imx-ilv-land-hooks/orders/internal/imx/models"
)

func FindById(id string) error {
	requestUrl := fmt.Sprintf("https://api.ropsten.x.immutable.com/v1/orders/%s", id)
	res, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	var order models.Order
	err = json.Unmarshal(bodyBytes, &order)
	fmt.Printf("%+v", order)
	return err
}
