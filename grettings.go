package grettings 
import "fmt"

func Hello(name string) string {
    message: fmt.Sprint("Hello %v Welocme!", name)
    return message
}
