import {EditorView, highlightActiveLine, lineNumbers} from '@codemirror/view'
import {xml} from '@codemirror/lang-xml'
import {autocompletion} from '@codemirror/autocomplete'
import {aura} from '../theme/aura'
import {dracula} from '../theme/dracula'
import {githubDark} from '../theme/github-dark'
import {githubLight} from '../theme/github-light'
import {materialDark} from '../theme/material-dark'
import {materialLight} from '../theme/material-light'
import {solarizedDark} from '../theme/solarized-dark'
import {solarizedLight} from '../theme/solarized-light'
import {tokyoNight} from '../theme/tokyo-night'
import {tokyoNightStorm} from '../theme/tokyo-night-storm'
import {tokyoNightDay} from '../theme/tokyo-night-day'
import {zenburn} from '../theme/zenburn'

const parent = document.getElementById('examples')
const doc = document.getElementById('code').textContent.trim()

const createEditor = (theme, name) => {
  console.log(theme)
  const view = new EditorView({
    doc: '// Theme: ' + name + '\n\n'+ doc,
    extensions: [
      lineNumbers(),
      highlightActiveLine(),
      xml(),
      autocompletion(),
      theme,
    ],
    parent,
  })
}

createEditor(aura, 'aura')
createEditor(dracula, 'dracula')
createEditor(githubDark, 'githubDark')
createEditor(githubLight,'githubLight')
createEditor(materialDark,'materialDark')
createEditor(materialLight,'materialLight')
createEditor(solarizedDark,'solarizedDark')
createEditor(solarizedLight,'solarizedLight')
createEditor(tokyoNight,'tokyoNight')
createEditor(tokyoNightStorm,'tokyoNightStorm')
createEditor(tokyoNightDay,'tokyoNightDay')
createEditor(zenburn,'zenburn')
