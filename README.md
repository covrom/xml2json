# xml2json
Unmarshal xml and marshal it to json universally.

## Example of translation

This xml
```xml
<definitions name = "HelloService"
   targetNamespace = "http://www.examples.com/wsdl/HelloService.wsdl"
   xmlns = "http://schemas.xmlsoap.org/wsdl/"
   xmlns:soap = "http://schemas.xmlsoap.org/wsdl/soap/"
   xmlns:tns = "http://www.examples.com/wsdl/HelloService.wsdl"
   xmlns:xsd = "http://www.w3.org/2001/XMLSchema">
 
   <message name = "SayHelloRequest">
      <part name = "firstName" type = "xsd:string"/>
   </message>
	
   <message name = "SayHelloResponse">
      <part name = "greeting" type = "xsd:string"/>
   </message>

   <portType name = "Hello_PortType">
      <operation name = "sayHello">
         <input message = "tns:SayHelloRequest"/>
         <output message = "tns:SayHelloResponse"/>
      </operation>
   </portType>

   <binding name = "Hello_Binding" type = "tns:Hello_PortType">
      <soap:binding style = "rpc"
         transport = "http://schemas.xmlsoap.org/soap/http"/>
      <operation name = "sayHello">
         <soap:operation soapAction = "sayHello"/>
         <input>
            <soap:body
               encodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
               namespace = "urn:examples:helloservice"
               use = "encoded"/>
         </input>
		
         <output>
            <soap:body
               encodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
               namespace = "urn:examples:helloservice"
               use = "encoded"/>
         </output>
      </operation>
   </binding>

   <service name = "Hello_Service">
      <documentation>WSDL File for HelloService</documentation>
      <port binding = "tns:Hello_Binding" name = "Hello_Port">
         <soap:address
            location = "http://www.examples.com/SayHello/" />
      </port>
   </service>
</definitions>
```
produces this json
```json
{
  "definitions": {
    "binding": {
      "binding": {
        "style": "rpc",
        "transport": "http://schemas.xmlsoap.org/soap/http"
      },
      "name": "Hello_Binding",
      "operation": {
        "input": {
          "body": {
            "encodingStyle": "http://schemas.xmlsoap.org/soap/encoding/",
            "namespace": "urn:examples:helloservice",
            "use": "encoded"
          }
        },
        "name": "sayHello",
        "operation": {
          "soapAction": "sayHello"
        },
        "output": {
          "body": {
            "encodingStyle": "http://schemas.xmlsoap.org/soap/encoding/",
            "namespace": "urn:examples:helloservice",
            "use": "encoded"
          }
        }
      },
      "type": "tns:Hello_PortType"
    },
    "message": [
      {
        "name": "SayHelloRequest",
        "part": {
          "name": "firstName",
          "type": "xsd:string"
        }
      },
      {
        "name": "SayHelloResponse",
        "part": {
          "name": "greeting",
          "type": "xsd:string"
        }
      }
    ],
    "name": "HelloService",
    "portType": {
      "name": "Hello_PortType",
      "operation": {
        "input": {
          "message": "tns:SayHelloRequest"
        },
        "name": "sayHello",
        "output": {
          "message": "tns:SayHelloResponse"
        }
      }
    },
    "service": {
      "documentation": "WSDL File for HelloService",
      "name": "Hello_Service",
      "port": {
        "address": {
          "location": "http://www.examples.com/SayHello/"
        },
        "binding": "tns:Hello_Binding",
        "name": "Hello_Port"
      }
    },
    "soap": "http://schemas.xmlsoap.org/wsdl/soap/",
    "targetNamespace": "http://www.examples.com/wsdl/HelloService.wsdl",
    "tns": "http://www.examples.com/wsdl/HelloService.wsdl",
    "xmlns": "http://schemas.xmlsoap.org/wsdl/",
    "xsd": "http://www.w3.org/2001/XMLSchema"
  }
}
```
Benchamrk for this sample:
```
BenchmarkXMLNode_UnmarshalXML-4            10000            135491 ns/op
BenchmarkXMLNode_MarshalJSON-4             20000             85595 ns/op
```