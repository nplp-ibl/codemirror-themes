import {EditorView} from '@codemirror/view'
import {Extension} from '@codemirror/state'
import {HighlightStyle, tags as t} from '@codemirror/highlight'

export const solarizedLightTheme = EditorView.theme({
  $: {
    color: '#586e75',
    backgroundColor: '#fdf6e3',
    '& ::selection': {backgroundColor: '#eee8d5'},
    caretColor: '#657b83',
  },

  '$$focused $cursor': {borderLeftColor: '#657b83'},
  '$$focused $selectionBackground': {backgroundColor: '#eee8d5'},

  $panels: {backgroundColor: '#eee8d5', color: '#586e75'},
  '$panels.top': {borderBottom: '2px solid black'},
  '$panels.bottom': {borderTop: '2px solid black'},

  $searchMatch: {
    backgroundColor: '#eee8d5',
    outline: `1px solid #d3af86`
  },
  '$searchMatch.selected': {
    backgroundColor: '#eee8d5'
  },

  $activeLine: {backgroundColor: '#eee8d5'},
  $selectionMatch: {backgroundColor: '#eee8d5'},

  '$matchingBracket, $nonmatchingBracket': {
    backgroundColor: '',
    outline: 'none'
  },
  $gutters: {
    backgroundColor: '#fdf6e3',
    color: '#586e75',
    border: 'none'
  },
  '$gutterElement.lineNumber': {color: 'inherit'},

  $foldPlaceholder: {
    backgroundColor: 'transparent',
    border: 'none',
    color: '#586e75'
  },

  $tooltip: {
    border: '1px solid #d3af86',
    backgroundColor: '#eee8d5'
  },
  '$tooltip.autocomplete': {
    '& > ul > li[aria-selected]': {
      backgroundColor: '#eee8d5',
      color: '#586e75'
    }
  }
}, {dark: true})

export const solarizedLightHighlightStyle = HighlightStyle.define(
  {tag: t.keyword, color: '#859900'},
  {tag: [t.name, t.deleted, t.character, t.propertyName, t.macroName], color: '#268BD2'},
  {tag: [t.processingInstruction, t.string, t.inserted], color: '#2AA198'},
  {tag: [t.function(t.variableName), t.labelName], color: '#268BD2'},
  {tag: [t.color, t.constant(t.name), t.standard(t.name)], color: '#CB4B16'},
  {tag: [t.definition(t.name), t.separator], color: '#268BD2'},
  {tag: [t.typeName, t.className, t.number, t.changed, t.annotation, t.modifier, t.self, t.namespace], color: '#268BD2'},
  {tag: [t.operator, t.operatorKeyword, t.url, t.escape, t.regexp, t.link, t.special(t.string)], color: '#D30102'},
  {tag: [t.meta, t.comment], color: '#93A1A1'},
  {tag: t.strong, fontWeight: 'bold'},
  {tag: t.emphasis, fontStyle: 'italic'},
  {tag: t.link, color: '#859900', textDecoration: 'underline'},
  {tag: t.heading, fontWeight: 'bold', color: '#268BD2'},
  {tag: [t.atom, t.bool, t.special(t.variableName)], color: ''},
  {tag: t.invalid, color: ''},
)

export const solarizedLight: Extension = [
  solarizedLightTheme,
  solarizedLightHighlightStyle,
]
