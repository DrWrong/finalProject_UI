#the UI Server of an Service composition system
It is built under the framework *Macaron* and *Thrift*
As an Web Service it invoke the internal `weatherInfoService` and `securityService` to provide an combined service

## Directory structure

+ conf/   the system's configure file, The connection info of `weatherInfoService` and `securityService` is configured there.
+ controller/ following the MVC pattern it server as the web service's controller, it process the http request from Browser.
+ public/   static file like `css` and `js` file
+ utils/  some commonly used codes like get Configure , get Service Client and so on.
+ views/ template views. it use pongo2 template render engine to render.
