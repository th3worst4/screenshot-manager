package main

import (
    "os"
    "io"
    "bufio"
    "fmt"
)

const HISTORYSIZE uint8 = 10

var Home_Var string = os.Getenv("HOME")
var Documents string = Home_Var + "/Documents"
var Downloads string = Home_Var + "/Downloads"
var Music string = Home_Var + "/Music"
var Pictures string = Home_Var + "/Pictures"
var Videos string = Home_Var + "/Videos"

// Doubly linked list implementation
// This is for future history feature (maybe)

type node struct {
    nextNode *node
	previousNode *node
    value string
    index uint8
}

func CreateNode(value string) *node {
    newNode := new(node)
    newNode.value = value
    newNode.nextNode = nil
	newNode.previousNode = nil
    newNode.index = 0
    
    return newNode
}

func InsertNewEntry(head **node, newEntry *node) *node {
    (*newEntry).nextNode = *head
	(**head).previousNode = newEntry
    if (**head).index == HISTORYSIZE{ 
        (*newEntry).index = HISTORYSIZE
        RemoveTail(newEntry)
    }else{
        (*newEntry).index = (**head).index + 1
    }
    
    *head = newEntry

    return newEntry 
}

func RemoveTail(head *node){
    tmp := (*head).nextNode
    for i := uint8(0); i < HISTORYSIZE - uint8(2); i++ {
        tmp = (*tmp).nextNode
		(*tmp).index--
    }
    tmp.nextNode = nil
}

func PrintNodes(head *node){
    tmp := head
    for tmp != nil{
        fmt.Printf("%v %v %v %v- ", (*tmp).value, (*tmp).index, (*tmp).nextNode, (*tmp).previousNode)
        tmp = (*tmp).nextNode
    }
    fmt.Printf("\n")
}

func PathExists(path string) (bool, error){
    _, err := os.Stat(path) 
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

func SaveFile(file *os.File, dest string){
	r := bufio.NewReader(file)
	output, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer output.Close()
	
	w := bufio.NewWriter(output)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		w.Write(buf[:n])
		w.Flush()
	}
}
