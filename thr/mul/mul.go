// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package mul

import(
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type Multiple interface {
  // Parameters:
  //  - A
  //  - B
  Mul(ctx context.Context, a int32, b int32) (r int32, err error)
}

type MultipleClient struct {
  c thrift.TClient
}

func NewMultipleClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MultipleClient {
  return &MultipleClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewMultipleClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MultipleClient {
  return &MultipleClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewMultipleClient(c thrift.TClient) *MultipleClient {
  return &MultipleClient{
    c: c,
  }
}

func (p *MultipleClient) Client_() thrift.TClient {
  return p.c
}
// Parameters:
//  - A
//  - B
func (p *MultipleClient) Mul(ctx context.Context, a int32, b int32) (r int32, err error) {
  var _args0 MultipleMulArgs
  _args0.A = a
  _args0.B = b
  var _result1 MultipleMulResult
  if err = p.Client_().Call(ctx, "Mul", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type MultipleProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Multiple
}

func (p *MultipleProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *MultipleProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *MultipleProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewMultipleProcessor(handler Multiple) *MultipleProcessor {

  self2 := &MultipleProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["Mul"] = &multipleProcessorMul{handler:handler}
return self2
}

func (p *MultipleProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x3

}

type multipleProcessorMul struct {
  handler Multiple
}

func (p *multipleProcessorMul) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := MultipleMulArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Mul", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := MultipleMulResult{}
var retval int32
  var err2 error
  if retval, err2 = p.handler.Mul(ctx, args.A, args.B); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Mul: " + err2.Error())
    oprot.WriteMessageBegin("Mul", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Mul", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - A
//  - B
type MultipleMulArgs struct {
  A int32
  B int32
}

func NewMultipleMulArgs() *MultipleMulArgs {
  return &MultipleMulArgs{}
}


func (p *MultipleMulArgs) GetA() int32 {
  return p.A
}

func (p *MultipleMulArgs) GetB() int32 {
  return p.B
}
func (p *MultipleMulArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case -1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField_1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case -2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField_2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MultipleMulArgs)  ReadField_1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field -1: ", err)
} else {
  p.A = v
}
  return nil
}

func (p *MultipleMulArgs)  ReadField_2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field -2: ", err)
} else {
  p.B = v
}
  return nil
}

func (p *MultipleMulArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Mul_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField_2(oprot); err != nil { return err }
    if err := p.writeField_1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MultipleMulArgs) writeField_2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("b", thrift.I32, -2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error -2:b: ", p), err) }
  if err := oprot.WriteI32(int32(p.B)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.b (-2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error -2:b: ", p), err) }
  return err
}

func (p *MultipleMulArgs) writeField_1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("a", thrift.I32, -1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error -1:a: ", p), err) }
  if err := oprot.WriteI32(int32(p.A)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.a (-1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error -1:a: ", p), err) }
  return err
}

func (p *MultipleMulArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MultipleMulArgs(%+v)", *p)
}

// Attributes:
//  - Success
type MultipleMulResult struct {
  Success *int32 `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewMultipleMulResult() *MultipleMulResult {
  return &MultipleMulResult{}
}

var MultipleMulResult_Success_DEFAULT int32
func (p *MultipleMulResult) GetSuccess() int32 {
  if !p.IsSetSuccess() {
    return MultipleMulResult_Success_DEFAULT
  }
return *p.Success
}
func (p *MultipleMulResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *MultipleMulResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MultipleMulResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *MultipleMulResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Mul_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MultipleMulResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteI32(int32(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *MultipleMulResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MultipleMulResult(%+v)", *p)
}


