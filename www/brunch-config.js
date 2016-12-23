// See http://brunch.io for documentation.
module.exports = {
  files: {
    javascripts: {joinTo: 'app.js'},
    stylesheets: {joinTo: 'app.css'},
    templates: {joinTo: 'app.js'}
  },
  paths: {
    watched: ['js','css','assets']
  },
  plugins: {
    fingerprint: {
      autoClearOldFiles: true,
      alwaysRun: true,
      srcBasePath: 'public/',
      destBasePath: 'public/'
    }
  }
}
