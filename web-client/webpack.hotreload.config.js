var path = require('path');
var webpack = require('webpack');

module.exports = {
  entry: [
    'webpack-dev-server/client?http://localhost:3000',
    'webpack/hot/only-dev-server',
    './index.js'
  ],
  output: { path: path.join(__dirname, 'dist'), filename: 'app.js' },
  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoErrorsPlugin()
  ],
  module: { loaders: [
    { test: /\.css$/, loader: 'style!css' },
    { test: /\.scss$/, loader: 'style!sass' },
    { test: /\.jsx?$/, loaders: ['react-hot', 'babel-loader?stage=0'], exclude: /node_modules/ }
  ] }
};
