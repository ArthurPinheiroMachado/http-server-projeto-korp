package projetokorp

import (
	"fmt"
	"http-server-projeto-korp/internal/util"
	"net/http"
	"time"
)

type AboutProjetoKorp struct {
	Name    string `json:"nome"`
	TimeNow string `json:"horario"`
}

func GetProjetoKorp(projectName string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		trace := util.CreateErrorContext("projeto_korp.GetProjetoKorp")
		w.Header().Set("Content-Type", "application/json")

		response := AboutProjetoKorp{
			Name:    projectName,
			TimeNow: time.Now().UTC().Format("02/01/2006 15:04:05"),
		}

		if err := util.JsonEncodeToWriter(w, response); err != nil {
			util.SendHttpError(w, http.StatusInternalServerError, trace.Apply(fmt.Errorf("erro ao codificar resposta final: %v", err)))
		}
	}
}
