// @flow
/*:: type Api = {
  cache: boolean => void,
  env: () => string,
}
*/

// Cache in the module. This can get called from multiple places and env vars can get lost
let isElectron = process.env.BABEL_ENV === 'Electron'
module.exports = function(api /*: Api */) {
  api.cache(true)

  console.error('\nbabel.config.js config for', isElectron ? 'Electron' : 'React Native')

  if (isElectron) {
    console.error('Babel for Electron')
    return {
      plugins: [
        '@babel/plugin-proposal-object-rest-spread',
        '@babel/transform-flow-strip-types',
        '@babel/plugin-proposal-class-properties',
      ],
      presets: ['@babel/preset-env', '@babel/preset-react'],
    }
  } else {
    console.error('Babel for RN')
    return {}
  }
}
