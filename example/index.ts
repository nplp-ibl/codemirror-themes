import {aura} from '../theme/aura'
import {ayuDark} from '../theme/ayu-dark'
import {ayuMirage} from '../theme/ayu-mirage'
import {ayu} from '../theme/ayu'
import {cobalt} from '../theme/cobalt'
import {cyberpunk} from '../theme/cyberpunk'
import {dracula} from '../theme/dracula'
import {githubDark} from '../theme/github-dark'
import {githubLight} from '../theme/github-light'
import {gruvboxDark} from '../theme/gruvbox-dark'
import {gruvboxLight} from '../theme/gruvbox-light'
import {materialDark} from '../theme/material-dark'
import {materialLight} from '../theme/material-light'
import {mirage} from '../theme/mirage'
import {monokai} from '../theme/monokai'
import {noctisAzureus} from '../theme/noctis-azureus'
import {noctisBordo} from '../theme/noctis-bordo'
import {noctisHibernus} from '../theme/noctis-hibernus'
import {noctisLilac} from '../theme/noctis-lilac'
import {noctisLux} from '../theme/noctis-lux'
import {noctisMinimus} from '../theme/noctis-minimus'
import {noctisUva} from '../theme/noctis-uva'
import {noctisViola} from '../theme/noctis-viola'
import {noctis} from '../theme/noctis'
import {oneDarkPro} from '../theme/one-dark-pro'
import {panda} from '../theme/panda'
import {pastel} from '../theme/pastel'
import {poimandres} from '../theme/poimandres'
import {raDark} from '../theme/ra-dark'
import {rubyBlue} from '../theme/ruby-blue'
import {setiBlack} from '../theme/seti-black'
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
createEditor(ayuDark, 'ayuDark')
createEditor(ayuMirage, 'ayuMirage')
createEditor(ayu, 'ayu')
createEditor(cobalt, 'cobalt')
createEditor(cyberpunk, 'cyberpunk')
createEditor(dracula, 'dracula')
createEditor(githubDark, 'githubDark')
createEditor(githubLight,'githubLight')
createEditor(gruvboxDark,'gruvboxDark')
createEditor(gruvboxLight,'gruvboxLight')
createEditor(materialDark,'materialDark')
createEditor(materialLight,'materialLight')
createEditor(mirage,'mirage')
createEditor(monokai,'monokai')
createEditor(noctisAzureus,'noctisAzureus')
createEditor(noctisBordo,'noctisBordo')
createEditor(noctisHibernus,'noctisHibernus')
createEditor(noctisLilac,'noctisLilac')
createEditor(noctisLux,'noctisLux')
createEditor(noctisMinimus,'noctisMinimus')
createEditor(noctisUva,'noctisUva')
createEditor(noctisViola,'noctisViola')
createEditor(noctis,'noctis')
createEditor(oneDarkPro, 'oneDarkPro')
createEditor(panda,'panda')
createEditor(pastel,'pastel')
createEditor(poimandres,'poimandres')
createEditor(raDark,'raDark')
createEditor(rubyBlue,'rubyBlue')
createEditor(setiBlack,'setiBlack')
createEditor(shadesOfPurple,'shadesOfPurple')
createEditor(solarizedDark,'solarizedDark')
createEditor(solarizedLight,'solarizedLight')
createEditor(tokyoNightStorm,'tokyoNightStorm')
createEditor(tokyoNight,'tokyoNight')
createEditor(zenburn,'zenburn')
