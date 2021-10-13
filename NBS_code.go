package main
  
import (
    "fmt"
    "os"
    "net"
)
  
// Main function
func main() {
  
    // list all environment variables and their values
    for _, env := range os.Environ() {
        fmt.Println(env)
    }

    // Print details of network interface
   interfaces, _ := net.Interfaces()
   for _, inter := range interfaces {
      fmt.Println("{Index:",inter.Index,"MTU:",inter.MTU,"Name:",inter.Name,"HardWwareAddr:",inter.HardwareAddr,"Flags :",inter.Flags,"}")
      addrs,_ := inter.Addrs()
      for _, ipaddr := range addrs {
         fmt.Println("[",ipaddr,"]")
      }
      fmt.Println()     
   }
}