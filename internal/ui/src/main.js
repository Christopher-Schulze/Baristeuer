import { mount } from 'svelte'
import './app.css'
import App from './App.svelte'
import { Backend } from './wailsjs/go/service/DataService'

const app = mount(App, {
  target: document.getElementById('app'),
})

export default app
export { Backend }
