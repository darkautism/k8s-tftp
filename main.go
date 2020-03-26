package main

import (
	"pack.ag/tftp"
	"log"
	"os"
	"io"
)

func main() {
	s, err := tftp.NewServer(":69", tftp.ServerSinglePort(true))
	if err != nil {
		panic(err)
	}
	readHandler := tftp.ReadHandlerFunc(proxyTFTP)
	s.ReadHandler(readHandler)
	s. ListenAndServe()
	select{}

}

func proxyTFTP(w tftp.ReadRequest) {
	log.Printf("[%s] GET %s\n", w.Addr().IP.String(), w.Name() )
	file, err := os.Open("/tftpboot/" + w.Name()) // For read access.
	if err != nil {
		log.Println(err)
		w.WriteError(tftp.ErrCodeFileNotFound, err.Error())
		return
	}
	defer file.Close()

	if _, err := io.Copy(w, file); err != nil {
		log.Println(err)
	}
}
