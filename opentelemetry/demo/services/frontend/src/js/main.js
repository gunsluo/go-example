
import { SimpleSpanProcessor } from '@opentelemetry/tracing';
import { WebTracerProvider } from '@opentelemetry/web';
import { ZoneContextManager } from '@opentelemetry/context-zone';
import { CollectorTraceExporter } from '@opentelemetry/exporter-collector';
import { B3Propagator } from '@opentelemetry/core';
import { SpanKind } from '@opentelemetry/api';


const tracerProvider = new WebTracerProvider();

const collectorOptions = {
  serviceName: 'frontend',
  url: 'http://localhost:55678/v1/trace', // url is optional and can be omitted - default is localhost:55678
  // url: 'localhost:55680',
};
tracerProvider.addSpanProcessor(new SimpleSpanProcessor(new CollectorTraceExporter(collectorOptions)));

tracerProvider.register({
  contextManager: new ZoneContextManager(),
  propagator: new B3Propagator(),
});

const tracer = tracerProvider.getTracer('frontend', '1.0.0');



const getData = (ctx, url, method) => new Promise((resolve, reject) => {
  let traceparent = '00-' + ctx.traceId + '-' + ctx.spanId + '-0' + ctx.traceFlags
  // let tracestate = ctx.traceState.serialize()
  console.log(traceparent)
  // console.log(tracestate)

  const req = new XMLHttpRequest();
  req.open(method, url, true);
  req.setRequestHeader('Content-Type', 'application/json');
  req.setRequestHeader('Accept', 'application/json');
  req.setRequestHeader('traceparent', traceparent);
  // req.setRequestHeader('tracestate', tracestate);
  req.onload = () => {
    resolve(req.responseText);
  };
  req.onerror = () => {
    reject();
  };
  req.send();
});


const prepareClickEvent = () => {
  const pUrl = 'http://localhost:8080/profile?id=123';
  const element = document.getElementById('btn');
  const div = document.getElementById('show');
  const method = 'POST'

  const onClick = () => {
    const span = tracer.startSpan(method + ' ' + pUrl, {
      parent: tracer.getCurrentSpan(),
      kind: SpanKind.CLIENT,
      attributes: {
        'http.method': method,
        'http.scheme': 'http',
      },
    });

    getData(span.context(), pUrl, method).then((data) => {
      div.innerHTML = data;
      span.setAttribute('http.status_code', 200)
      span.end();
    }, () => {
      span.end();
    })
  };
  element.addEventListener('click', onClick);
};

window.addEventListener('load', prepareClickEvent);
