package main
/**
内嵌interface
 */

type LearnByHear interface {
	hear()
}

type LearnByLokk interface {
	look()
}

type Learn interface {
	LearnByHear
	LearnByLokk
}

