package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Proceso struct {
	Bandera    bool
	Id_proceso uint64
	Iter       uint64
}
type ListaProcesos struct {
	Procesos []*Proceso
}

func (lp *ListaProcesos) Agregar(p *Proceso) {
	lp.Procesos = append(lp.Procesos, p)
}
func (p *Proceso) CorrerProceso() {
	p.Iter = 0
	p.Bandera = true
	for {
		p.Iter++
		time.Sleep(time.Millisecond * 500)
		if !p.Bandera {
			break
		}
	}

}

func (lp *ListaProcesos) Mostrar(bandera chan bool) {

	for {
		select {
		case <-bandera:
			return
		default:
			for _, v := range lp.Procesos {
				fmt.Println("Id: ", v.Id_proceso, " :", v.Iter)
				time.Sleep(time.Millisecond * 500)

			}

		}
	}
}

func (lp *ListaProcesos) Eliminar(id uint64) {
	var procesoAuxiliar []*Proceso
	for _, v := range lp.Procesos {
		if v.Id_proceso == id {
			fmt.Println("Proceso : ", v.Id_proceso, "Finalizado")
			v.Bandera = false

		} else {
			procesoAuxiliar = append(procesoAuxiliar, v)
		}
	}
	lp.Procesos = procesoAuxiliar
}

func main() {
	bandera := make(chan bool)
	var stopPrint = bufio.NewScanner(os.Stdin)
	var respuesta int
	var id uint64
	var aProceso *Proceso
	var idProcesos []*Proceso
	var listaProcesos = &ListaProcesos{
		Procesos: idProcesos,
	}
	//var flag bool
	//c := make(chan string)
	for ok := true; ok; ok = (respuesta != 0) {

		fmt.Println(" 1. Agregar Proceso \n 2.Mostrar Procesos \n 3.Eliminar Procesos \n 0.Salir")
		fmt.Scanln(&respuesta)
		switch respuesta {
		case 1:

			fmt.Println("Agregar Proceso")
			fmt.Scanln(&id)
			aProceso = &Proceso{Id_proceso: id}
			listaProcesos.Agregar(aProceso)
			go aProceso.CorrerProceso()
			break
		case 2:
			fmt.Println("Mostrando...")
			go listaProcesos.Mostrar(bandera)
			stopPrint.Scan()
			bandera <- true
			break
		case 3:
			fmt.Println("Ingresa el ID a eliminar")
			var idDel uint64
			fmt.Scanln(&idDel)
			go listaProcesos.Eliminar(idDel)

			break
		}

	}
}
