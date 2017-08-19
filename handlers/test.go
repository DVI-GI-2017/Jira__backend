package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, body []byte, _ map[string]string) {
	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusForbidden)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	action, err := pool.NewAction("Insert and Find")
	if err != nil {
		log.Printf("%v", err)
	}

	pool.Queue <- &pool.Job{
		ModelType: user,
		Action:    action,
	}

	result := <-pool.Results

	tools.JsonResponse(result, w)
}

// test: for i in {1..15}; do echo '{"email": "test", "password": "password"}' | curl -d @- http://localhost:3000/api/v1/test; done
