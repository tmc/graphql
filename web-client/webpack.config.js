module.exports = {
  entry: './index.js',
  output: { path: './dist/', filename: 'app.js' },
  module: { loaders: [
    { test: /\.css$/, loader: 'style!css' },
    { test: /\.scss$/, loader: 'style!sass' },
    { test: /\.jsx?$/, loader: 'babel-loader?optional[]=runtime&stage=0', exclude: /node_modules/ }
  ] },
  externals: { 'react': 'React' }
};
