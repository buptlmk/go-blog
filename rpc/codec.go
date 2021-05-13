package rpc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/vmihailenco/msgpack/v5"
	"io"
	"net"
	"net/rpc"
	"reflect"
	"strings"
)

// 请求的包装
type MsgPackReq struct {
	rpc.Request             // head
	Arg         interface{} //body
}

// 回复的包装
type MsgPackResp struct {
	rpc.Response
	Rely interface{}
}

// 定义codec

type MsgServerCodec struct {
	rwc    io.ReadWriteCloser // 用于读写数据，实际上是一个网络链接
	req    MsgPackReq         // 用于缓存解析到的请求
	closed bool               // 标识codec是否关闭
}

type MsgClientCodec struct {
	rwc    io.ReadWriteCloser
	res    MsgPackResp
	closed bool
}

func NewServerCodec(conn net.Conn) *MsgServerCodec {
	return &MsgServerCodec{conn, MsgPackReq{}, false}
}
func NewClientCodec(conn net.Conn) *MsgClientCodec {
	return &MsgClientCodec{conn, MsgPackResp{}, false}
}

// 发送
func (c *MsgClientCodec) WriteRequest(r *rpc.Request, arg interface{}) error {

	if c.closed {
		return nil
	}
	request := &MsgPackReq{*r, arg}
	reqData, err := msgpack.Marshal(request)
	if err != nil {
		panic(err)
		return err
	}
	// 先发送数据长度
	head := make([]byte, 4)
	binary.BigEndian.PutUint32(head, uint32(len(reqData)))
	_, err = c.rwc.Write(head)

	// 再发送数据
	_, err = c.rwc.Write(reqData)
	return err

}

func (c *MsgClientCodec) ReadResponseHeader(r *rpc.Response) error {
	if c.closed {
		return nil
	}

	data, err := readData(c.rwc)
	if err != nil {
		// client 初始化后开始轮询数据，所以处理链接close的情况
		if strings.Contains(err.Error(), "use of closed network connection") {
			return nil
		}
		panic(err)
	}

	// 反序列化
	var res MsgPackResp
	err = msgpack.Unmarshal(data, &res)

	if err != nil {
		panic(err)
	}
	r.ServiceMethod = res.ServiceMethod
	r.Seq = res.Seq

	c.res = res
	return nil
}

func (c *MsgClientCodec) ReadResponseBody(reply interface{}) error {

	if c.res.Error != "" {
		return errors.New(c.res.Error)
	}
	if reply != nil {
		reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(c.res.Rely))
	}
	return nil
}

func (c *MsgClientCodec) Close() error {
	c.closed = true
	if c.rwc != nil {
		return c.rwc.Close()
	}
	return nil
}

// 处理TCP粘包
func readData(conn io.ReadWriteCloser) (data []byte, err error) {
	const HeadSize = 4 // 设定长度为4个字节
	headBuf := bytes.NewBuffer(make([]byte, 0, HeadSize))
	headData := make([]byte, HeadSize)
	for {
		readSize, err := conn.Read(headData)
		if err != nil {
			return nil, err
		}
		headBuf.Write(headData[0:readSize])
		if headBuf.Len() == HeadSize {
			break
		} else {
			headData = make([]byte, HeadSize-readSize)
		}
	}

	bodyLen := int(binary.BigEndian.Uint32(headBuf.Bytes()))
	bodyBuf := bytes.NewBuffer(make([]byte, 0, bodyLen))

	bodyData := make([]byte, bodyLen)
	for {
		readSize, err := conn.Read(bodyData)
		if err != nil {
			return nil, err
		}
		bodyBuf.Write(bodyData[0:readSize])
		if bodyBuf.Len() == bodyLen {
			break
		} else {
			bodyData = make([]byte, bodyLen-readSize)
		}
	}

	data = bodyBuf.Bytes()

	return data, nil

}

func (c *MsgServerCodec) WriteResponse(r *rpc.Response, reply interface{}) error {

	if c.closed {
		return nil
	}
	// 封包
	response := &MsgPackResp{*r, reply}
	// 编码
	resData, err := msgpack.Marshal(response)
	if err != nil {
		panic(err)
		return err
	}
	// fatou
	head := make([]byte, 4)
	binary.BigEndian.PutUint32(head, uint32(len(resData)))

	_, err = c.rwc.Write(head)

	_, err = c.rwc.Write(resData)
	return err
}

func (c *MsgServerCodec) ReadRequestHeader(r *rpc.Request) error {

	if c.closed {
		return nil
	}

	data, err := readData(c.rwc)

	if err != nil {
		if err == io.EOF {
			return err
		}
		if strings.Contains(err.Error(), "connection reset by peer") {
			return err
		}
		panic(err)
	}

	//
	var req MsgPackReq
	err = msgpack.Unmarshal(data, &req)

	if err != nil {
		panic(err)
	}

	r.ServiceMethod = req.ServiceMethod
	r.Seq = req.Seq
	c.req = req
	return err
}

func (c *MsgServerCodec) ReadRequestBody(arg interface{}) error {
	if arg != nil {
		//参数不为nil，通过反射将结果设置到arg变量
		reflect.ValueOf(arg).Elem().Set(reflect.ValueOf(c.req.Arg))
	}
	return nil
}

func (c *MsgServerCodec) Close() error {
	c.closed = true
	if c.rwc != nil {
		return c.rwc.Close()
	}
	return nil
}
