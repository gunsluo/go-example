import tpltext from './tpl/add.html'
import $ from "jquery";

class App {
  constructor() {
    this.name = 'an wasm example';
    this.init()
  }

  async render() {
    $('#app').html(tpltext)

    $('#addBtn').on('click', event => {
      const v1 = parseInt($('#v1').val())
      const v2 = parseInt($('#v2').val())

      const result = '<p>' + add(v1, v2) + '</p>'
      $('#result').html(result)
    })

  }

  add(i, j) {
    if (Number.isNaN(i)) {
      i = 0
    }

    if (Number.isNaN(j)) {
      j = 0
    }

    return i + j
  }

  init() {
    this.loadGoWasm()
  }

  async loadGoWasm() {
    const go = new Go();
    let mod, inst;
    WebAssembly.instantiateStreaming(fetch("add.wasm"), go.importObject).then(async (result) => {
      mod = result.module;
      inst = result.instance;
      await go.run(inst);
    });
  }
}

export default App 
