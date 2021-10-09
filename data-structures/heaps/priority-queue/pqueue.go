package main

import (
	"container/heap"
	"fmt"
)

// example of hospital patients and their danger
type Patient struct {
	ticketNo int
	severity int // higher == more dangerous
	desc string
}

func (i Patient) String() string {
	return fmt.Sprintf("doctor is treating patient no %d with severity %d (%s)",
		i.ticketNo, i.severity, i.desc)
}

type PQueue []Patient

func (pq PQueue) Len() int           { return len(pq) }
func (pq PQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq PQueue) Less(i, j int) bool {
	// higher severity patients are treated first
	// but if same severity, then look at who got here first from ticket nu
	if pq[i].severity == pq[j].severity {
		return pq[i].ticketNo < pq[j].ticketNo
	}
	return pq[i].severity > pq[j].severity
}

func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Patient))
}
func (pq *PQueue) Pop() interface{} {
	tmp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return tmp
}

func main() {
	pq := PQueue{}
	heap.Init(&pq)

	heap.Push(&pq, Patient{ticketNo: 101, severity: 1, desc: "coughing"})
	heap.Push(&pq, Patient{ticketNo: 102, severity: 9, desc: "heart pain"})
	heap.Push(&pq, Patient{ticketNo: 103, severity: 5, desc: "brain damage"})
	heap.Push(&pq, Patient{ticketNo: 104, severity: 3, desc: "stomachache"})
	heap.Push(&pq, Patient{ticketNo: 105, severity: 1, desc: "bruises"})

	for range pq {
		fmt.Println(heap.Pop(&pq))
	}
}