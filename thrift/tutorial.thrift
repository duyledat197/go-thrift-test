include "test.thrift"

namespace go tutorial

exception InvalidOperation {
  1: i32 code,
  2: string message
}

struct AddRequest {
    1: i32 a,
    2: i32 b
}

struct AddResponse {
    1: i32 result
}

/**
 * Ahh, now onto the cool part, defining a service. Services just need a name
 * and can optionally inherit from another service using the extends keyword.
 */
service Calculator extends test.SharedService {
   AddResponse add(AddRequest req) throws (1:InvalidOperation ouch),
}