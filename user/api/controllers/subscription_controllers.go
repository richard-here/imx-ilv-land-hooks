package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/richard-here/imx-ilv-land-hooks/user/api/auth"
	"github.com/richard-here/imx-ilv-land-hooks/user/api/responses"
	"github.com/richard-here/imx-ilv-land-hooks/user/api/utils/formaterror"
	"github.com/richard-here/imx-ilv-land-hooks/user/models"
)

func (server *Server) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	subscription := models.Subscription{}
	err = json.Unmarshal(body, &subscription)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	subscription.Prepare()
	subscription.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if uid != subscription.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	subscriptionCreated, err := subscription.Subscribe(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, subscriptionCreated.ID))
	responses.JSON(w, http.StatusCreated, subscriptionCreated)
}

func (server *Server) GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	subscription := models.Subscription{}

	subscriptions, err := subscription.FindAllSubscriptions(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, subscriptions)
}

func (server *Server) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	subscription := models.Subscription{}
	err = server.DB.Debug().Model(models.Subscription{}).Where("id = ?", sid).Take(&subscription).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	if uid != subscription.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	_, err = subscription.Unsubscribe(server.DB, sid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", sid))
	responses.JSON(w, http.StatusNoContent, "")
}
