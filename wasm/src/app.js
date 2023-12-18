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

      const result = '<p>' + this.add(v1, v2) + '</p>'
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
  }
}

export default App 
