import {EditorView} from '@codemirror/view'
import {Extension} from '@codemirror/state'
import {HighlightStyle, syntaxHighlighting} from '@codemirror/language'
import {tags as t} from '@lezer/highlight'

export const config = {
  name: 'poimandres',
  dark: true,
  background: '#1b1e28',
  foreground: '#a6accd',
  selection: '#717cb425',
  cursor: '#a6accd',
  dropdownBackground: '#1b1e28',
  dropdownBorder: '#ffffff10',
  activeLine: '#717cb425',
  matchingBracket: '#00000000',
  keyword: '#a6accd',
  storage: '#a6accd',
  variable: '#e4f0fb',
  parameter: '#e4f0fb',
  function: '#ADD7FF',
  string: '#5DE4c7',
  constant: '#5DE4c7',
  type: '#a6accdC0',
  class: '#91B4D5',
  number: '#5DE4c7',
  comment: '#767c9dB0',
  heading: '#5DE4c7',
  invalid: '#d0679d',
  regexp: '#5fb3a1',
}

export const poimandresTheme = EditorView.theme({
  '&': {
    color: config.foreground,
    backgroundColor: config.background,
  },

  '.cm-content': {caretColor: config.cursor},

  '&.cm-focused .cm-cursor': {borderLeftColor: config.cursor},
  '&.cm-focused .cm-selectionBackground, .cm-selectionBackground, & ::selection': {backgroundColor: config.selection},

  '.cm-panels': {backgroundColor: config.dropdownBackground, color: config.foreground},
  '.cm-panels.cm-panels-top': {borderBottom: '2px solid black'},
  '.cm-panels.cm-panels-bottom': {borderTop: '2px solid black'},

  '.cm-searchMatch': {
    backgroundColor: config.dropdownBackground,
    outline: `1px solid ${config.dropdownBorder}`
  },
  '.cm-searchMatch.cm-searchMatch-selected': {
    backgroundColor: config.selection
  },

  '.cm-activeLine': {backgroundColor: config.activeLine},
  '.cm-selectionMatch': {backgroundColor: config.selection},

  '&.cm-focused .cm-matchingBracket, &.cm-focused .cm-nonmatchingBracket': {
    backgroundColor: config.matchingBracket,
    outline: 'none'
  },

  '.cm-gutters': {
    backgroundColor: config.background,
    color: config.foreground,
    border: 'none'
  },
  '.cm-activeLineGutter': {backgroundColor: config.background},

  '.cm-foldPlaceholder': {
    backgroundColor: 'transparent',
    border: 'none',
    color: config.foreground
  },

  '.cm-tooltip': {
    border: `1px solid ${config.dropdownBorder}`,
    backgroundColor: config.dropdownBackground,
    color: config.foreground
  },
  '.cm-tooltip .cm-tooltip-arrow:before': {
    borderTopColor: 'transparent',
    borderBottomColor: 'transparent'
  },
  '.cm-tooltip .cm-tooltip-arrow:after': {
    borderTopColor: config.foreground,
    borderBottomColor: config.foreground,
  },
  '.cm-tooltip.cm-tooltip-autocomplete': {
    '& > ul > li[aria-selected]': {
      background: config.selection,
      color: config.foreground
    }
  },
}, {dark: config.dark})

export const poimandresHighlightStyle = HighlightStyle.define([
  {tag: t.keyword, color: config.keyword},
  {tag: [t.name, t.deleted, t.character, t.macroName], color: config.variable},
  {tag: [t.propertyName], color: config.function},
  {tag: [t.processingInstruction, t.string, t.inserted, t.special(t.string)], color: config.string},
  {tag: [t.function(t.variableName), t.labelName], color: config.function},
  {tag: [t.color, t.constant(t.name), t.standard(t.name)], color: config.constant},
  {tag: [t.definition(t.name), t.separator], color: config.variable},
  {tag: [t.className], color: config.class},
  {tag: [t.number, t.changed, t.annotation, t.modifier, t.self, t.namespace], color: config.number},
  {tag: [t.typeName], color: config.type, fontStyle: config.type},
  {tag: [t.operator, t.operatorKeyword], color: config.keyword},
  {tag: [t.url, t.escape, t.regexp, t.link], color: config.regexp},
  {tag: [t.meta, t.comment], color: config.comment},
  {tag: t.strong, fontWeight: 'bold'},
  {tag: t.emphasis, fontStyle: 'italic'},
  {tag: t.link, textDecoration: 'underline'},
  {tag: t.heading, fontWeight: 'bold', color: config.heading},
  {tag: [t.atom, t.bool, t.special(t.variableName)], color: config.variable},
  {tag: t.invalid, color: config.invalid},
  {tag: t.strikethrough, textDecoration: 'line-through'},
])

export const poimandres: Extension = [
  poimandresTheme,
  syntaxHighlighting(poimandresHighlightStyle),
]
