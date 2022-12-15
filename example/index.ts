import {aura} from '../theme/aura'
import {dracula} from '../theme/dracula'
import {githubDark} from '../theme/github-dark'
import {githubLight} from '../theme/github-light'
import {materialDark} from '../theme/material-dark'
import {materialLight} from '../theme/material-light'
import {mirage} from '../theme/mirage'
import {oneDarkPro} from '../theme/one-dark-pro'
import {panda} from '../theme/panda'
import {pastel} from '../theme/pastel'
import {shadesOfPurple} from '../theme/shades-of-purple'
import {solarizedDark} from '../theme/solarized-dark'
import {solarizedLight} from '../theme/solarized-light'
import {tokyoNightStorm} from '../theme/tokyo-night-storm'
import {tokyoNight} from '../theme/tokyo-night'
import {zenburn} from '../theme/zenburn'

import {EditorView, highlightActiveLine, lineNumbers} from '@codemirror/view'
import {xml} from '@codemirror/lang-xml'
import {autocompletion} from '@codemirror/autocomplete'

const parent = document.getElementById('examples')
const doc = document.getElementById('code').textContent.trim()

const createEditor = (theme, name) => {
  console.log(theme)
  const view = new EditorView({
    doc: '<!-- Theme: ' + name + ' -->\n\n'+ doc,
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
createEditor(mirage,'mirage')
createEditor(oneDarkPro, 'oneDarkPro')
createEditor(panda,'panda')
createEditor(pastel,'pastel')
createEditor(shadesOfPurple,'shadesOfPurple')
createEditor(solarizedDark,'solarizedDark')
createEditor(solarizedLight,'solarizedLight')
createEditor(tokyoNightStorm,'tokyoNightStorm')
createEditor(tokyoNight,'tokyoNight')
createEditor(zenburn,'zenburn')
