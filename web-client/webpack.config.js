module.exports = {
  entry: './index.js',
  output: { path: './dist/', filename: 'app.js' },
  module: { loaders: [ {
    test: /\.js$/,
    loader: 'babel-loader?stage=0'
  } ] }
};
