package scraper

import "fmt"

type NoSuchContentsError struct{
    url string
    tag string
}

func (e NoSuchContentsError) Error() string{
    return fmt.Sprintf("not find contents, target url: %s, find tag : %s\n", e.url, e.tag)
}


