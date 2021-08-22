package ibofos

import (	
	"bytes"
	"fmt"
	"errors"
	"io"
	"net"
	log "github.com/sirupsen/logrus"
)

func  (rq Requester)  ConnectToIBoFOS() (net.Conn, error) {
	var err error = nil
	uri := fmt.Sprintf("%s:%d",rq.IbofServerIP, rq.IbofServerPort)

	conn, err := net.Dial("tcp", uri)
	if err != nil {		
		log.Errorf("ConnectToIBoFOS : [%v]", err)
		return nil, err		
	}
	return conn, err
}

func (rq Requester) DisconnectToIBoFOS(conn net.Conn) error {
	var err error = nil

	if conn != nil {
		log.Info("Connection Cloase : ", conn.LocalAddr().String())
		err = conn.Close()
	}

	return err
}

func (rq Requester) ReadFromIBoFSocket(conn net.Conn) (bytes.Buffer, error) {
	var err error
	var buf bytes.Buffer

	log.Info("readFromIBoFSocket Start")

	if conn == nil {
		log.Info("readFromIBoFSocket : Conn is nil")
	} else {

		_, err := io.Copy(&buf, conn)

		if err != nil || err == io.EOF {
			log.Info("readFromIBoFSocket : Message Receive Fail :", err)
		} else {
			log.Info("readFromIBoFSocket : Message Receive Success")
		}
	}
	return buf, err
}

func (rq Requester) WriteToIBoFSocket(conn net.Conn, marshaled []byte) error {
	var err error = nil
	if conn == nil {
		err = errors.New("WriteToIBoFSocket : Conn is nil")
		log.Error(err)
	} else {
		_, err = conn.Write(marshaled)
		if err != nil {
			//conn.Close()
			//conn = nil
			log.Infof("WriteToIBoFSocket : Writre Fail - %s\n", err)
			log.Infof("WriteToIBoFSocket : Conn closed\n")
		} else {
			log.Infof("WriteToIBoFSocket : Write Success\n")
		}
	}
	return err
}